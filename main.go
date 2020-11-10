package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
)

type Block struct {
	Cmd   string
	Args  []string
	UpInt int
	UpSig int
	Pos   int
}

var (
	updateChan  = make(chan struct{})
	barBytesArr = make([][]byte, len(Blocks))
	sigChan     = make(chan os.Signal, 16)
	signalMap   = make(map[os.Signal]Block)
	signalArr   = []os.Signal{}
	x           *xgb.Conn
	root        xproto.Window
)

func updateBar() {
	var finalBytesBuffer bytes.Buffer

	for i := 0; i < len(Blocks); i++ {
		if barBytesArr[i] != nil {
			finalBytesBuffer.Write(Delim)
			finalBytesBuffer.Write(barBytesArr[i])
		}
	}

	finalBytes := bytes.TrimPrefix(finalBytesBuffer.Bytes(), Delim)
	xproto.ChangeProperty(x, xproto.PropModeReplace, root, xproto.AtomWmName, xproto.AtomString, 8, uint32(len(finalBytes)), finalBytes)
}

func runBlock(block Block) {
	outputBytes, err := exec.Command(block.Cmd, block.Args[:]...).Output()
	if err != nil {
		log.Println("Failed to update", block.Cmd, block.Args[:], " -- ", err)
	} else {
		barBytesArr[block.Pos] = bytes.TrimSpace(outputBytes)
		updateChan <- struct{}{}
	}
}

func main() {
	/* setup X */
	var err error

	x, err = xgb.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	defer x.Close()
	root = xproto.Setup(x).DefaultScreen(x).Root

	/* initialize blocks */
	for i := 0; i < len(Blocks); i++ {
		go func(i int) {
			Blocks[i].Pos = i

			if len(Blocks[i].Args) < 1 {
				pCmd := strings.Split(Blocks[i].Cmd, " ")
				Blocks[i].Cmd = pCmd[0]
				Blocks[i].Args = pCmd[1:]
			}

			runBlock(Blocks[i])
			if Blocks[i].UpInt != 0 {
				for {
					time.Sleep(time.Duration(Blocks[i].UpInt) * time.Second)
					runBlock(Blocks[i])
				}
			}
		}(i)
	}

	/* handle signals */
	for _, block := range Blocks {
		if block.UpSig != 0 {
			signal.Notify(sigChan, syscall.Signal(34+block.UpSig))
			signalMap[syscall.Signal(34+block.UpSig)] = block
		}
	}

	go func() {
		for sig := range sigChan {
			block, _ := signalMap[sig]
			runBlock(block)
		}
	}()

	/* check for updates */
	for _ = range updateChan {
		updateBar()
	}
}

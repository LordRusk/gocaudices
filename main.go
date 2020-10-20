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
	sigChan      = make(chan os.Signal)
	updateChan   = make(chan int)
	barStringArr = make([]string, len(Blocks))
	signalMap    = make(map[os.Signal]Block)
)

func updateBar(x *xgb.Conn, root xproto.Window, stringArr []string) {
	var finalString strings.Builder

	for i := 0; i < len(stringArr); i++ {
		if stringArr[i] != "" {
			finalString.WriteString(Delim)
			finalString.WriteString(stringArr[i])
		}
	}

	xproto.ChangeProperty(x, xproto.PropModeReplace, root, xproto.AtomWmName, xproto.AtomString, 8, uint32(finalString.Len()), []byte(finalString.String()))
}

func runBlock(block Block, updateChan chan<- int) {
	outputBytes, err := exec.Command(block.Cmd, block.Args[:]...).Output()
	if err != nil {
		log.Println("Failed to update", block.Cmd, block.Args[:], " -- ", err)
	} else {
		barStringArr[block.Pos] = string(bytes.TrimSpace(outputBytes))
		updateChan <- 0
	}
}

func main() {
	/* setup X */
	x, err := xgb.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	defer x.Close()
	root := xproto.Setup(x).DefaultScreen(x).Root

	/* initialize blocks */
	for i := 0; i < len(Blocks); i++ {
		go func(i int) {
			Blocks[i].Pos = i

			if len(Blocks[i].Args) < 1 {
				pCmd := strings.Split(Blocks[i].Cmd, " ")
				Blocks[i].Cmd = pCmd[0]
				Blocks[i].Args = pCmd[1:]
			}

			runBlock(Blocks[i], updateChan)
			if Blocks[i].UpInt != 0 {
				for {
					time.Sleep(time.Duration(Blocks[i].UpInt) * time.Second)
					runBlock(Blocks[i], updateChan)
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
			runBlock(block, updateChan)
		}
	}()

	/* set status on update */
	for _ = range updateChan {
		updateBar(x, root, barStringArr)
	}
}

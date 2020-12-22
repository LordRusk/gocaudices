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
	cmd   string
	inSh  bool
	args  []string // used internally
	upInt int
	upSig int
	pos   int
}

var (
	updateChan  = make(chan interface{}, 1)
	barBytesArr = make([][]byte, len(Blocks))
	sigChan     = make(chan os.Signal, 1024)
	signalMap   = make(map[os.Signal][]Block)
	x           *xgb.Conn     // global X connection
	root        xproto.Window // global root window
)

func runBlock(block Block) {
	var outputBytes []byte
	var err error
	if len(block.args) < 2 {
		outputBytes, err = exec.Command(block.cmd).Output()
	} else {
		outputBytes, err = exec.Command(block.args[0], block.args[1:]...).Output()
	}
	if err != nil {
		log.Printf("Failed to update `%v` | %v\n", block.cmd, err)
	} else {
		barBytesArr[block.pos] = bytes.TrimSpace(outputBytes)
		updateChan <- nil
	}
}

func main() {
	// setup X
	var err error
	x, err = xgb.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	defer x.Close()
	root = xproto.Setup(x).DefaultScreen(x).Root

	for i := 0; i < len(Blocks); i++ { // initialize blocks
		go func(i int) {
			Blocks[i].pos = i

			if Blocks[i].inSh {
				Blocks[i].args = []string{shell, "-c", Blocks[i].cmd}
			} else {
				Blocks[i].args = strings.Split(Blocks[i].cmd, " ")
			}

			runBlock(Blocks[i])
			if Blocks[i].upInt != 0 {
				for {
					time.Sleep(time.Duration(Blocks[i].upInt) * time.Second)
					runBlock(Blocks[i])
				}
			}
		}(i)
	}

	for _, block := range Blocks { // handle signals
		if block.upSig != 0 {
			signal.Notify(sigChan, syscall.Signal(34+block.upSig))
			signalMap[syscall.Signal(34+block.upSig)] = append(signalMap[syscall.Signal(34+block.upSig)], block)
		}
	}

	go func() {
		for sig := range sigChan {
			blocks, _ := signalMap[sig]
			for _, block := range blocks {
				runBlock(block)
			}
		}
	}()

	var finalBytesBuffer bytes.Buffer
	for range updateChan { // update bar on signal
		for i := 0; i < len(Blocks); i++ {
			if barBytesArr[i] != nil {
				finalBytesBuffer.Write(delim)
				finalBytesBuffer.Write(barBytesArr[i])
			}
		}

		finalBytes := bytes.TrimPrefix(finalBytesBuffer.Bytes(), delim)
		finalBytesBuffer.Reset()

		// set the root window name
		xproto.ChangeProperty(x, xproto.PropModeReplace, root, xproto.AtomWmName, xproto.AtomString, 8, uint32(len(finalBytes)), finalBytes)
	}
}

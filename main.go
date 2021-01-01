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

type block struct {
	cmd   string
	args  []string // used internally
	inSh  bool
	upInt int
	upSig int
	pos   int // used internally
}

var updateChan = make(chan interface{}, 1)
var barBytesArr = make([][]byte, len(blocks))
var sigChan = make(chan os.Signal, 1024)
var signalMap = make(map[os.Signal][]block)

func runBlock(b block) {
	var outputBytes []byte
	var err error
	if b.args == nil {
		outputBytes, err = exec.Command(b.cmd).Output()
	} else {
		outputBytes, err = exec.Command(b.args[0], b.args[1:]...).Output()
	}
	if err != nil {
		log.Printf("Failed to update `%v` | %v\n", b.cmd, err)
		return
	}

	barBytesArr[b.pos] = bytes.TrimSpace(outputBytes)
}

func main() {
	// setup X
	x, err := xgb.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	defer x.Close()
	root := xproto.Setup(x).DefaultScreen(x).Root

	for i := 0; i < len(blocks); i++ { // initialize blocks
		blocks[i].pos = i

		if blocks[i].inSh {
			blocks[i].args = []string{shell, cmdstropt, blocks[i].cmd}
		} else {
			if strings.Contains(blocks[i].cmd, " ") {
				blocks[i].args = strings.Split(blocks[i].cmd, " ")
			}
		}

		if blocks[i].upSig != 0 {
			signal.Notify(sigChan, syscall.Signal(34+blocks[i].upSig))
			signalMap[syscall.Signal(34+blocks[i].upSig)] = append(signalMap[syscall.Signal(34+blocks[i].upSig)], blocks[i])
		}

		runBlock(blocks[i])
		if blocks[i].upInt != 0 {
			go func(i int) {
				for {
					time.Sleep(time.Duration(blocks[i].upInt) * time.Second)
					runBlock(blocks[i])
					updateChan <- nil
				}
			}(i)
		}
	}
	updateChan <- nil

	go func() { // handle signals
		for sig := range sigChan {
			bs, _ := signalMap[sig]
			for _, b := range bs {
				runBlock(b)
			}
			updateChan <- nil
		}
	}()

	var finalBytesBuffer bytes.Buffer
	for range updateChan { // update bar on signal
		for _, b := range blocks {
			if barBytesArr[b.pos] != nil {
				finalBytesBuffer.Write(delim)
				finalBytesBuffer.Write(barBytesArr[b.pos])
			}
		}

		finalBytes := bytes.TrimPrefix(finalBytesBuffer.Bytes(), delim)
		xproto.ChangeProperty(x, xproto.PropModeReplace, root, xproto.AtomWmName, xproto.AtomString, 8, uint32(len(finalBytes)), finalBytes) // set the root window name
		finalBytesBuffer.Reset()
	}
}

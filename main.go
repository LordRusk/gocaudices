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
	inSh  bool
	upInt int
	upSig int

	args []string // used internally
	pos  int      // used internally
}

var updateChan = make(chan interface{})
var barBytesArr = make([][]byte, len(blocks))

func runBlock(b block) {
	outputBytes, err := exec.Command(b.args[0], b.args[1:]...).Output()
	if err != nil {
		log.Printf("Failed to update `%v` | %v\n", b.cmd, err)
		return
	}

	barBytesArr[b.pos] = bytes.TrimSpace(outputBytes)
	updateChan <- nil
}

func main() {
	x, err := xgb.NewConn() // connect to X
	if err != nil {
		log.Fatalf("Cannot connect to X! | %v\n", err)
	}
	defer x.Close()
	root := xproto.Setup(x).DefaultScreen(x).Root

	sigChan := make(chan os.Signal, 1024)
	signalMap := make(map[os.Signal][]block)

	for i := 0; i < len(blocks); i++ { // initialize blocks
		go func(i int) {
			blocks[i].pos = i

			if blocks[i].inSh {
				blocks[i].args = []string{shell, cmdstropt, blocks[i].cmd}
			} else {
				blocks[i].args = strings.Split(blocks[i].cmd, " ")
			}

			if blocks[i].upSig != 0 {
				signal.Notify(sigChan, syscall.Signal(34+blocks[i].upSig))
				signalMap[syscall.Signal(34+blocks[i].upSig)] = append(signalMap[syscall.Signal(34+blocks[i].upSig)], blocks[i])
			}

			runBlock(blocks[i])
			if blocks[i].upInt != 0 {
				for {
					time.Sleep(time.Duration(blocks[i].upInt) * time.Second)
					runBlock(blocks[i])
				}
			}
		}(i)
	}

	go func() { // update bar on signal
		var finalBytesBuffer bytes.Buffer
		for range updateChan {
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
	}()
	updateChan <- nil // initially update the bar

	for sig := range sigChan { // handle signals
		bs, _ := signalMap[sig]
		for _, b := range bs {
			runBlock(b)
		}
		updateChan <- nil
	}
}

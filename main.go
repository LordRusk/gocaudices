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

var barBytesArr = make([][]byte, len(blocks))

func runBlock(b block) {
	var outputBytes []byte
	var err error
	if len(b.args) == 1 {
		outputBytes, err = exec.Command(b.args[0]).Output()
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
	// connect to X
	x, err := xgb.NewConn()
	if err != nil {
		log.Fatalf("Cannot connect to X! | %v\n", err)
	}
	defer x.Close()
	root := xproto.Setup(x).DefaultScreen(x).Root

	updateChan := make(chan interface{}, 1)
	sigChan := make(chan os.Signal, 1024)
	signalMap := make(map[os.Signal][]block)

	for i := 0; i < len(blocks); i++ { // initialize blocks
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
			go func(i int) {
				for {
					time.Sleep(time.Duration(blocks[i].upInt) * time.Second)
					runBlock(blocks[i])
					updateChan <- nil
				}
			}(i)
		}
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

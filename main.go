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

	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
)

type block struct {
	Cmd      string
	Shell    bool
	Interval int
	Signal   int

	// Internal
	args []string
	pos  int
}

var updateChan = make(chan struct{})
var barBytesArr = make([][]byte, len(blocks))

func (b *block) run() {
	outputBytes, err := exec.Command(b.args[0], b.args[1:]...).Output()
	if err != nil {
		log.Printf("Failed to update `%s`: %s\n", b.Cmd, err)
		return
	}

	barBytesArr[b.pos] = bytes.TrimSpace(bytes.Split(outputBytes, []byte("\n"))[0])
	updateChan <- struct{}{}
}

func main() {
	x, err := xgb.NewConn() // connect to X
	if err != nil {
		log.Fatalf("Cannot connect to X: %s\n", err)
	}
	defer x.Close()
	root := xproto.Setup(x).DefaultScreen(x).Root

	sigChan := make(chan os.Signal, 1024)
	signalMap := make(map[os.Signal][]block)

	for i := 0; i < len(blocks); i++ { // initialize blocks
		go func(i int) {
			blocks[i].pos = i

			if blocks[i].Shell {
				blocks[i].args = []string{shell, cmdstropt, blocks[i].Cmd}
			} else {
				blocks[i].args = strings.Split(blocks[i].Cmd, " ")
			}

			if blocks[i].Signal != 0 {
				signal.Notify(sigChan, syscall.Signal(34+blocks[i].Signal))
				signalMap[syscall.Signal(34+blocks[i].Signal)] = append(signalMap[syscall.Signal(34+blocks[i].Signal)], blocks[i])
			}

			blocks[i].run() // initially build bar
			if blocks[i].Interval != 0 {
				for {
					time.Sleep(time.Duration(blocks[i].Interval) * time.Second)
					blocks[i].run()
				}
			}
		}(i)
	}

	go func() { // update bar on signal
		var finalBytesBuffer bytes.Buffer
		for range updateChan {
			for i := 0; i < len(blocks); i++ {
				if barBytesArr[i] != nil {
					finalBytesBuffer.WriteString(delim)
					finalBytesBuffer.Write(barBytesArr[i])
				}
			}

			finalBytes := finalBytesBuffer.Bytes()[len(delim):]
			xproto.ChangeProperty(x, xproto.PropModeReplace, root, xproto.AtomWmName, xproto.AtomString, 8, uint32(len(finalBytes)), finalBytes) // set the root window name
			finalBytesBuffer.Reset()
		}
	}()

	for sig := range sigChan { // handle signals
		go func(sig *os.Signal) {
			bs := signalMap[*sig]
			for _, b := range bs {
				go b.run()
			}
		}(&sig)
	}
}

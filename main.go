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
	Interval uint
	Signal   uint

	// Internal
	args []string
	pos  int
}

var updateChan = make(chan int)
var barBytesArr = make([][]byte, len(blocks))

func (b *block) run() {
	outputBytes, err := exec.Command(b.args[0], b.args[1:]...).Output()
	if err != nil {
		log.Printf("block %q update failed: %s\n", b.Cmd, err.Error())
		return
	}

	barBytesArr[b.pos] = bytes.TrimSpace(bytes.Split(outputBytes, []byte("\n"))[0])
	updateChan <- 1
}

func main() {
	x, err := xgb.NewConn() // connect to X
	if err != nil {
		log.Fatalf("X connection failed: %s\n", err.Error())
	}
	defer x.Close()
	root := xproto.Setup(x).DefaultScreen(x).Root

	sigChan := make(chan os.Signal, 1024)
	signalMap := make(map[os.Signal][]block)

	// initialize blocks
	for i := range blocks {
		go func(bl *block, i int) {
			bl.pos = i

			if bl.Shell {
				bl.args = []string{shell, cmdstropt, bl.Cmd}
			} else {
				bl.args = strings.Split(bl.Cmd, " ")
			}

			if bl.Signal != 0 {
				signal.Notify(sigChan, syscall.Signal(34+bl.Signal))
				signalMap[syscall.Signal(34+bl.Signal)] = append(signalMap[syscall.Signal(34+bl.Signal)], *bl)
			}

			bl.run() // initially build bar
			if bl.Interval != 0 {
				for {
					time.Sleep(time.Duration(bl.Interval) * time.Second)
					bl.run()
				}
			}
		}(&blocks[i], i)
	}

	go func() { // update bar on signal
		var finalBytesBuffer bytes.Buffer
		for range updateChan {
			for i := range blocks {
				if barBytesArr[i] != nil {
					finalBytesBuffer.WriteString(delim)
					finalBytesBuffer.Write(barBytesArr[i])
				}
			}

			finalBytes := bytes.TrimPrefix(finalBytesBuffer.Bytes(), []byte(delim))
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

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
	UpInt int
	UpSig int
	Pos   int
}

var (
	sigChan      = make(chan os.Signal)
	updateChan   = make(chan bool)
	barStringArr = make([]string, len(Blocks))
	signalMap    = make(map[os.Signal]Block)
)

func mergeFinalString(stringArr []string) string {
	var finalString strings.Builder

	for i := 0; i < len(stringArr); i++ {
		if stringArr[i] != "" {
			finalString.WriteString(Delim)
			finalString.WriteString(stringArr[i])
		}
	}

	return finalString.String()
}

func execBlock(command string) (string, error) {
	outputBytes, err := exec.Command(Shell, RunIn, command).Output()
	if err != nil {
		return "", err
	}

	outputBytes = bytes.TrimSpace(outputBytes)

	return string(outputBytes), err

}

func runBlock(block Block, updateChan chan<- bool) {
	newString, err := execBlock(block.Cmd)
	if err != nil {
		log.Println("Failed to update", block.Cmd, " -- ", newString, err)
	} else {
		barStringArr[block.Pos] = newString
		updateChan <- true
	}
}

func main() {
	x, err := xgb.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	defer x.Close()
	root := xproto.Setup(x).DefaultScreen(x).Root

	/* run blocks */
	for i := 0; i < len(Blocks); i++ {
		go func(i int) {
			Blocks[i].Pos = i
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
		signal.Notify(sigChan, syscall.Signal(34+block.UpSig))
		signalMap[syscall.Signal(34+block.UpSig)] = block
	}

	go func() {
		for sig := range sigChan {
			block, _ := signalMap[sig]
			runBlock(block, updateChan)
		}
	}()

	/* set status on update */
	for _ = range updateChan {
		statusText := mergeFinalString(barStringArr)
		xproto.ChangeProperty(x, xproto.PropModeReplace, root, xproto.AtomWmName, xproto.AtomString, 8, uint32(len(statusText)), []byte(statusText))
	}
}

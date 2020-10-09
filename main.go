package main

// #cgo LDFLAGS: -lX11
// #include <X11/Xlib.h>
import "C"

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type Block struct {
	Cmd   string
	UpInt int
	UpSig int
	Pos   int
}

var (
	sigChan      = make(chan os.Signal, 512)
	updateChan   = make(chan bool, 512)
	barStringArr = make([]string, len(Blocks))
	dpy          = C.XOpenDisplay(nil)
	screen       = C.XDefaultScreen(dpy)
	root         = C.XRootWindow(dpy, screen)
)

func setStatus(s *C.char) {
	C.XStoreName(dpy, root, s)
	C.XFlush(dpy)
}

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
	/* start all the blocks */
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
	signalMap := make(map[int]Block)
	for _, block := range Blocks {
		signal.Notify(sigChan, syscall.Signal(34+block.UpSig))
		signalMap[34+block.UpSig] = block
	}

	go func() {
		for sig := range sigChan {
			psig := strings.Split(sig.String(), " ")
			sigNum, _ := strconv.Atoi(psig[1])
			block, _ := signalMap[sigNum]
			runBlock(block, updateChan)
		}
	}()

	/* set status on update */
	for _ = range updateChan {
		setStatus(C.CString(mergeFinalString(barStringArr)))
	}
}

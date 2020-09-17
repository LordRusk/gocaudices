package main

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
	sigChan      = make(chan os.Signal, len(Blocks))
	updateChan   = make(chan bool, len(Blocks))
	barStringArr = make([]string, len(Blocks))
)

func mergeFinalString(stringArr []string) string {
	var finalString strings.Builder

	for i := 0; i < len(stringArr); i++ {
		if stringArr[i] != "" {
			finalString.WriteString(stringArr[i])
			finalString.WriteString(Delim)
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
		if barStringArr[block.Pos] != newString {
			barStringArr[block.Pos] = newString
			if len(updateChan) == 0 {
				updateChan <- true
			}
		}
	}
}

func main() {
	/* get all the blocks running */
	for i := 0; i < len(Blocks); i++ {
		go func(i int) {
			Blocks[i].Pos = i
			runBlock(Blocks[i], updateChan)
			time.Sleep(time.Duration(i) * (time.Millisecond * 2))
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
		for {
			sig := <-sigChan
			psig := strings.Split(sig.String(), " ")
			sigNum, err := strconv.ParseInt(psig[1], 0, 64)
			if err != nil {
				log.Println(err)
			}

			if block, ok := signalMap[int(sigNum)]; !ok {
				log.Println("Unkown update signal:", psig[1])
			} else {
				runBlock(block, updateChan)
			}
		}
	}()

	/* watch for updates */
	if Receivers < 1 {
		Receivers++
	}
	for i := 0; i < Receivers; i++ {
		go func() {
			for _ = range updateChan {
				exec.Command("xsetroot", "-name", mergeFinalString(barStringArr)).Output()
			}
		}()
	}

	select {}
}

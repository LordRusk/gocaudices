package main

import (
	"os"
	"log"
	"time"
	"regexp"
	"syscall"
	"strings"
	"strconv"
	"os/exec"
	"os/signal"
)

type updateBlock struct {
	blockPos int
	newString string
}

type Block struct {
	Cmd string
	UpInt int
	UpSig int
	Pos int
}

var (
	updateChan = make(chan updateBlock, len(Blocks))
	sigChan = make(chan os.Signal, len(Blocks))
)

func execBlock(command string) (string, error) {
	newStringBytes, err := exec.Command(Shell, RunIn, command).Output()
	if err != nil {
		return "", err
	}

	newString := string(newStringBytes)
	re := regexp.MustCompile(` \n`)
	newString = re.ReplaceAllString(string(newString), "")
	re = regexp.MustCompile(`\n`)
	newString = re.ReplaceAllString(string(newString), "")

	return newString, err
}

func runBlock(block Block, updateChan chan<- updateBlock) {
	updateBlock := updateBlock { blockPos: block.Pos }

	newString, err := execBlock(block.Cmd)
	if err != nil {
		log.Println("Failed to update", block.Cmd, " -- ", newString, err)
	}

	updateBlock.newString = newString
	updateChan <- updateBlock
}

func main() {
	/* get all the blocks running */
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
		for {
			sig := <- sigChan
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

	/* build the original []string */
	barStringArr := make([]string, len(Blocks))
	for i := 0; i < len(Blocks); i++ {
		updatedBlock := <- updateChan
		barStringArr[updatedBlock.blockPos] = updatedBlock.newString

		_, err := exec.Command(Shell, RunIn, string("xsetroot -name \"Building Blocks... "+strconv.FormatInt(int64(i+1), 10)+" of "+strconv.FormatInt(int64(len(barStringArr)), 10)+"\"")).Output()
		if err != nil {
			log.Println(err)
		}
	}

	/* watch for updates and update accordingly */
	for {
		updatedBlock := <- updateChan
		barStringArr[updatedBlock.blockPos] = updatedBlock.newString

		_, err := exec.Command(Shell, RunIn, string("xsetroot -name \""+strings.Join(barStringArr[:], Delim)+"\"")).Output()
		if err != nil {
			log.Println(err)
		}
	}
}

package main

var (
	delim = []byte(" ") // the delimiter that will be used
	shell = "dash"      // shell used

	Blocks = []Block{
		{cmd: "recicon", upSig: 9},
		{cmd: "music", upSig: 11},
		{cmd: "pacpackages", upSig: 8},
		{cmd: "news", upSig: 6},
		{cmd: "georona | cut -d' ' -f1,3", inSh: true, upInt: 18000, upSig: 19}, // example of command that is run in shell
		{cmd: "torrent", upInt: 10, upSig: 7},
		{cmd: "memory", upInt: 6, upSig: 14},
		{cmd: "cpu", upInt: 3, upSig: 13},
		{cmd: "cpubars", upInt: 1, upSig: 22},
		{cmd: "disk /home", upInt: 7, upSig: 15},
		{cmd: "disk", upInt: 7, upSig: 15},
		{cmd: "astrological", upInt: 18000, upSig: 18},
		{cmd: "weather", upInt: 18000, upSig: 5},
		{cmd: "mailbox", upSig: 12},
		{cmd: "nettraf", upInt: 1, upSig: 16},
		{cmd: "volume", upSig: 10},
		{cmd: "battery", upInt: 5, upSig: 3},
		{cmd: "clock", upSig: 1},
		{cmd: "internet", upInt: 5, upSig: 4},
	}
)

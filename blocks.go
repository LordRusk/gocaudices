package main

var (
	/* config */
	delim = []byte(" ") // the delimiter that will be used
	shell = "sh"

	Blocks = []Block{
		Block{cmd: "recicon", upSig: 9},
		Block{cmd: "music", upSig: 11},
		Block{cmd: "pacpackages", upSig: 8},
		Block{cmd: "news", upSig: 6},
		Block{cmd: "georona | cut -d' ' -f1,3", inSh: true, upInt: 18000, upSig: 19}, // example of command that is run in shell
		Block{cmd: "torrent", upInt: 10, upSig: 7},
		Block{cmd: "memory", upInt: 6, upSig: 14},
		Block{cmd: "cpu", upInt: 3, upSig: 13},
		Block{cmd: "cpubars", upInt: 1, upSig: 22},
		Block{cmd: "disk /home", upInt: 7, upSig: 15},
		Block{cmd: "disk", upInt: 7, upSig: 15},
		Block{cmd: "astrological", upInt: 18000, upSig: 18},
		Block{cmd: "weather", upInt: 18000, upSig: 5},
		Block{cmd: "mailbox", upSig: 12},
		Block{cmd: "nettraf", upInt: 1, upSig: 16},
		Block{cmd: "volume", upSig: 10},
		Block{cmd: "battery", upInt: 5, upSig: 3},
		Block{cmd: "clock", upSig: 1},
		Block{cmd: "sip", upInt: 10, upSig: 2},
		Block{cmd: "internet", upInt: 5, upSig: 4},
	}
)

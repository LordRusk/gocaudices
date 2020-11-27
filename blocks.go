package main

var (
	/* config */
	Delim = []byte(" ") /* the delimiter that will be used */

	Blocks = []Block{
		Block{Cmd: "recicon", UpSig: 9},
		Block{Cmd: "music", UpSig: 11},
		Block{Cmd: "pacpackages", UpSig: 8},
		Block{Cmd: "news", UpSig: 6},
		Block{Cmd: "dash", Args: []string{"-c", "georona | cut -d' ' -f1,3"}, UpInt: 18000, UpSig: 19}, /* example of command that is run in shell */
		// Block{Cmd: "crypto", UpInt: 18000, UpSig: 17},
		Block{Cmd: "torrent", UpInt: 10, UpSig: 7},
		Block{Cmd: "memory", UpInt: 6, UpSig: 14},
		Block{Cmd: "cpu", UpInt: 3, UpSig: 13},
		Block{Cmd: "cpubars", UpInt: 1, UpSig: 22},
		Block{Cmd: "disk /home", UpInt: 7, UpSig: 15},
		Block{Cmd: "disk", UpInt: 7, UpSig: 15},
		Block{Cmd: "astrological", UpInt: 18000, UpSig: 18},
		Block{Cmd: "weather", UpInt: 18000, UpSig: 5},
		Block{Cmd: "mailbox", UpSig: 12},
		Block{Cmd: "nettraf", UpInt: 1, UpSig: 16},
		Block{Cmd: "volume", UpSig: 10},
		Block{Cmd: "battery", UpInt: 5, UpSig: 3},
		Block{Cmd: "clock", UpSig: 1},
		Block{Cmd: "sip", UpInt: 10, UpSig: 2},
		// Block{Cmd: "vpnstat express", UpSig: 21},
		Block{Cmd: "internet", UpInt: 5, UpSig: 4},
		// Block{Cmd: "cord", UpSig: 23},
		Block{Cmd: "help-icon"},
	}
)

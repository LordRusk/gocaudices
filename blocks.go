package main

var (
	/* config */
	Delim     = " "    /* the delimiter that will be used */
	Shell     = "dash" /* the shell to run the scripts in */
	RunIn     = "-c"   /* arg to make Shell run command from stdin */
	Receivers = 6      /* they amount of modules waiting for block updates | if value is less then 1, it will revert to 1 */

	Blocks = []Block{
		/* command */ /* update interval */ /* update sig */
		// Block{Cmd: "cat /tmp/recordingicon 2>/dev/null", UpInt: 0, UpSig: 9},
		// Block{Cmd: "music", UpInt: 0, UpSig: 11},
		Block{Cmd: "pacpackages", UpInt: 0, UpSig: 8},
		Block{Cmd: "news", UpInt: 0, UpSig: 6},
		// Block{Cmd: "georona | cut -d' ' -f1,3", UpInt: 0, UpSig: 19},
		// Block{Cmd: "crypto", UpInt: 18000, UpSig: 17},
		Block{Cmd: "torrent", UpInt: 20, UpSig: 7},
		Block{Cmd: "memory", UpInt: 10, UpSig: 14},
		Block{Cmd: "cpu", UpInt: 1, UpSig: 13},
		Block{Cmd: "cpubars", UpInt: 1, UpSig: 22},
		Block{Cmd: "disk /home", UpInt: 10, UpSig: 15},
		Block{Cmd: "disk", UpInt: 10, UpSig: 15},
		Block{Cmd: "astrological", UpInt: 18000, UpSig: 18},
		Block{Cmd: "weather", UpInt: 18000, UpSig: 5},
		Block{Cmd: "mailbox", UpInt: 0, UpSig: 12},
		Block{Cmd: "nettraf", UpInt: 1, UpSig: 16},
		Block{Cmd: "volume", UpInt: 0, UpSig: 10},
		Block{Cmd: "battery", UpInt: 5, UpSig: 3},
		Block{Cmd: "clock", UpInt: 0, UpSig: 1},
		// Block{Cmd: "sip", UpInt: 10, UpSig: 2},
		// Block{Cmd: "vpnstat express", UpInt: 0, UpSig: 21},
		Block{Cmd: "internet", UpInt: 5, UpSig: 4},
		// Block{Cmd: "cord", UpInt: 0, UpSig: 23},
		// Block{Cmd: "help-icon", UpInt: 0, UpSig: 20},
	}
)

/* quick list of block update signals
 * 1 clock
 * 2 sip
 * 3 battery
 * 4 internet
 * 5 weather
 * 6 news
 * 7 torrent
 * 8 pacpackages
 * 9 recicon
 * 10 volume
 * 11 music
 * 12 mailbox
 * 13 cpu
 * 14 memory
 * 15 disk
 * 16 nettraf
 * 17 crypto
 * 18 astrological
 * 19 georona
 * 20 help-icon
 * 21 vpnstat
 * 22 cpubar
 * 23 cord
 * 22 cpubar
 */

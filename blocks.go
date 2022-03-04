package main

var (
	delim     = " "  // the delimiter that will be used
	shell     = "sh" // shell used
	cmdstropt = "-c" // command string opt for shell

	blocks = []block{
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

/* Quick list of modules Update Signals
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
 */

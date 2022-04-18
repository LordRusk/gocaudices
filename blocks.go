package main

var (
	delim     = " "  // the delimiter that will be used
	shell     = "sh" // shell used
	cmdstropt = "-c" // command string opt for shell

	blocks = []block{
		{Cmd: "recicon", Signal: 9},
		{Cmd: "music", Signal: 11},
		{Cmd: "pacpackages", Signal: 8},
		{Cmd: "news", Signal: 6},
		{Cmd: "georona | cut -d' ' -f1,3", Shell: true, Interval: 18000, Signal: 19}, // example of command that is run in shell
		{Cmd: "torrent", Interval: 10, Signal: 7},
		{Cmd: "memory", Interval: 6, Signal: 14},
		{Cmd: "cpu", Interval: 3, Signal: 13},
		{Cmd: "cpubars", Interval: 1, Signal: 22},
		{Cmd: "disk /home", Interval: 7, Signal: 15},
		{Cmd: "disk", Interval: 7, Signal: 15},
		{Cmd: "astrological", Interval: 18000, Signal: 18},
		{Cmd: "weather", Interval: 18000, Signal: 5},
		{Cmd: "mailbox", Signal: 12},
		{Cmd: "nettraf", Interval: 1, Signal: 16},
		{Cmd: "volume", Signal: 10},
		{Cmd: "battery", Interval: 5, Signal: 3},
		{Cmd: "clock", Signal: 1},
		{Cmd: "internet", Interval: 5, Signal: 4},
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

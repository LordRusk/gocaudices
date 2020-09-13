//Modify this file to change what commands output to your statusbar, and recompile using the make command.
static const Block blocks[] = {
	/*Icon*/	/*Command*/		/*Update Interval*/	/*Update Signal*/
	{"", "cat /tmp/recordingicon 2>/dev/null",	0,			9},
	//{"",		"music",			0,			11},
	{"",		"pacpackages",			0,			8},
	{"",		"news",				0,			6},
	//{"",		"georona | cut -d' ' -f1,3",	18000,			19},
	//{"",		"crypto",			18000,			17},
	{"",		"torrent",			20,			7},
	{"",		"memory",			10,			14},
	{"",		"cpu",				10,			13},
	{"",		"cpubars",			10,			22},
	{"",		"disk /home",			10,			15},
	{"",		"disk",				10,			15},
	{"",		"astrological",			18000,			18},
	{"",		"weather",			18000,			5},
	{"",		"mailbox",			0,			12},
	{"",		"nettraf",			1,			16},
	{"",		"volume",			0,			10},
	{"",		"battery",			5,			3},
	{"",		"clock",			0,			1},
	//{"",		"sip",				10,			2},
	//{"",		"vpnstat express",		30,			21},
	{"",		"internet",			5,			4},
	//{"",		"cord",				0,			23},
	//{"",		"help-icon",			0,			20},
};

//sets delimeter between status commands. NULL character ('\0') means no delimeter.
static char *delim = " ";

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
 * 23 cord
 * 22 cpubar
 */

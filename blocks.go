package main

var (
	/* config */
	Shell = "sh" /* shell that will be used to run the commands */
	RunIn = "-c" /* shell opt to read from command string */
	Delim = " " /* the delimiter that will be used */
	Receivers = 6 /* They amount of receivers that can receive a blockUpdate | If 0 value is given, it will default to 1 */

	Blocks = []Block {
			     /* command */            /* update interval */    /* update sig */
		Block { Cmd: "cat /tmp/recordingicon 2>/dev/null", UpInt: 0, 	UpSig:  9, },
		//Block { Cmd: 	"music", 	UpInt:		0,	UpSig:		11, },
		Block { Cmd:	"pacpackages",	UpInt:		0,	UpSig:		8, },
		Block { Cmd: 	"news",		UpInt:		0,	UpSig:		6, },
		//Block { Cmd: "georona | cut -d' ' -f1,3", UpInt: 0,	UpSig:		19, },
		//Block { Cmd: 	"crypto",	UpInt:		18000,	UpSig:		17, },
		Block { Cmd: 	"torrent",	UpInt:		20,	UpSig:		7, },
		Block { Cmd: 	"memory",	UpInt:		10,	UpSig:		14, },
		Block { Cmd: 	"cpu",		UpInt:		10,	UpSig:		13, },
		Block { Cmd: 	"cpubars",	UpInt:		10,	UpSig:		22, },
		Block { Cmd: 	"disk /home",	UpInt:		10,	UpSig:		15, },
		Block { Cmd: 	"disk",		UpInt:		10,	UpSig:		15, },
		Block { Cmd: 	"astrological",	UpInt:		18000,	UpSig:		18, },
		Block { Cmd: 	"weather",	UpInt:		18000,	UpSig:		5, },
		Block { Cmd: 	"mailbox",	UpInt:		0,	UpSig:		12, },
		Block { Cmd: 	"nettraf",	UpInt:		1,	UpSig:		16, },
		Block { Cmd: 	"volume",	UpInt:		0,	UpSig:		10, },
		Block { Cmd: 	"battery",	UpInt:		5,	UpSig:		3, },
		Block { Cmd: 	"clock",	UpInt:		0,	UpSig:		1, },
		//Block { Cmd: 	"sip",		UpInt:		10,	UpSig:		2, },
		//Block { Cmd: 	"vpnstat express", UpInt:	0,	UpSig:		21, },
		Block { Cmd: 	"internet",	UpInt:		5,	UpSig:		4, },
		//Block { Cmd: 	"cord",		UpInt:		0,	UpSig:		23, },
		//Block { Cmd: 	"help-icon",	UpInt:		0,	UpSig:		20, },
		}
)

[![Go Report Card](https://goreportcard.com/badge/github.com/lordrusk/gocaudices)](https://goreportcard.com/report/github.com/lordrusk/gocaudices)

# Simple dwmblocks alternative
Gocaudices is a dwmblocks alternative written in GO using [xgb](https://github.com/jezek/xgb).

## About
I wrote this alternative to be simple, fast, and elegant. This project is never meant to exceed 100 SLOC in the base build.

## How-To
First download gocaudices with: `git clone https://github.com/lordrusk/gocaudices`. To make sure you have all dependencies installed you can run `go mod tidy`. You can (re)compile with: `go install`

* Configure

Gocaudices can be configured through adding `blocks` in the `blocks.go` I've left an example of a normal `blocks.go` config file that works with [my dotfiles](https://github.com/lordrusk/artixdwm).

* Shell commands

To run shell commands, add `inSh: true,` to the block in `blocks.go`.

* Update a module

The `block` definition of `{cmd: "volume", upSig: 10},` would be updated like `kill -44 $(pidof gocaudices)`. A dwm volume mute keybind might look like `{ 0, XF86XK_AudioMute, spawn, SHCMD("pamixer -t; kill -44 $(pidof gocaudices)") },`.

## Patches
Patches are hosted in this repo in `patches/*patch*`. To apply patches: `patch -p1 < path/to/patch.diff`. To create a proper patch, refer to [hacking](https://suckless.org/hacking/). If you would like to contribute a patch or feature, create a pull request.

* [complexdelim](https://github.com/LordRusk/gocaudices/tree/master/patches/complexdelim)
* [status2d](https://github.com/LordRusk/gocaudices/tree/master/patches/status2d)
* [statuscmd](https://github.com/LordRusk/gocaudices/tree/master/patches/statuscmd)
* [statuscolors](https://github.com/LordRusk/gocaudices/tree/master/patches/statuscolors)
* [icons](https://github.com/LordRusk/gocaudices/tree/master/patches/icons)

## (Non)-Features
+ Multiple blocks can have the same update signal.

## FQA -- Frequently Questioned Answers
+ Does it have bar click-ability?

	• Yes.

## AWESOME BARS
dwm bars that I think are awesome! check them out and give them a star!

• [sysmon](https://github.com/blmayer/sysmon/tree/main) I would use this if I hadn't made zara

• [spoon](https://git.2f30.org/spoon/) I don't know much C but this is great

• [rsblocks](https://github.com/MustafaSalih1993/rsblocks) I don't know much Rust, but featureful and well starred, makes me wanna get my status emoji game up to par

• [mblocks](https://gitlab.com/mhdy/mblocks) another great rusty bar

• [integrated-status-text](https://dwm.suckless.org/patches/integrated-status-text) the way god intended
  
• [gods](https://github.com/schachmat/gods) ICONIC

• [dwmblocks-async](https://github.com/UtkarshVerma/dwmblocks-async) Awesome! I wrote this project because dwmblocks wasn't async...and I've lived without bar clickability since...maybe should have gone with this and learned C!

•[Luke Smith's Dwmblocks](https://github.com/LukeSmithxyz/dwmblocks) how could I forget where it all began?

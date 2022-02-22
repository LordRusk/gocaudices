[![Go Report Card](https://goreportcard.com/badge/github.com/lordrusk/gocaudices)](https://goreportcard.com/report/github.com/lordrusk/gocaudices)

# Simple dwmblocks alternative
Gocaudices is a dwmblocks alternative written in GO using [xgb](https://github.com/jezek/xgb).

## About
I wrote this alternative to be simple, fast, and elegant. This project is never meant to exceed 100 SLOC in the base build.

## How-To
First download gocaudices with: `git clone https://github.com/lordrusk/gocaudices`. To make sure you have all dependencies installed you can run `go mod tidy`. You can (re)compile with: `go install`

- Configure

Gocaudices can be configured through adding `blocks` in the `blocks.go` I've left an example of a normal `blocks.go` config file that works with [my dotfiles](https://github.com/lordrusk/artixdwm).

- Shell commands

To run shell commands, add `inSh: true,` to the block in `blocks.go`.

- Update a module

The `block` definition of `{cmd: "volume", upSig: 10},` would be updated like `kill -44 $(pidof gocaudices)`. A dwm volume mute keybind might look like `{ 0, XF86XK_AudioMute, spawn, SHCMD("pamixer -t; kill -44 $(pidof gocaudices)") },`.

## Patches
Patches are hosted in this repo in `patches/*patch*`. To apply patches: `patch -p1 < path/to/patch.diff`. To create a proper patch, refer to [hacking](https://suckless.org/hacking/). If you would like to contribute a patch or feature, create a pull request.

[complexdelim](https://github.com/LordRusk/gocaudices/tree/master/patches/complexdelim)

[status2d](https://github.com/LordRusk/gocaudices/tree/master/patches/status2d)

## (Non)-Features
+ Multiple blocks can have the same update signal.

## FQA -- Frequently Questioned Answers
+ Does it have bar click-ability?

	• No, but if you'd like to create a patch, create a pull request!

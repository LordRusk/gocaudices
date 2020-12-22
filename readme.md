# An extremely simple dwmblocks alternative
Gocaudices is a dwmblocks alternative written in GO in less than 100 SLOC using [xgb](https://github.com/BurntSushi/xgb).

## About
I wrote this alternative to be simple, fast, and elegant. This project is never meant to exceed 100 SLOC in the base build.

## How-To
Gocaudices is configured through editing it's source code. This makes it extremely slim, fast, and secure. Since `main.go` is only ~100 SLOC, it makes it extremely simple to edit and add features.

- How To

First download gocaudices with: `git clone https://github.com/lordrusk/gocaudices` You can (re)compile with: `go install`

- Configure

Gocaudices can be configured through adding `Blocks` in the `blocks.go` file. Add their command, their update intriguers, and their update signals. I've left an example of a normal `blocks.go` config file that works with [my dotfiles](https://github.com/lordrusk/artixdwm).

- Shell commands

To run shell commands, add `inSh: true,` to the block in `blocks.go`. You can also set the shell in `blocks.go`.

- Update a module

The `Block` definition of `{cmd: "volume", upSig: 10},` would be updated like `kill -44 $(pidof gocaudices)`. A dwm volume mute keybind might look like `{ 0, XF86XK_AudioMute, spawn, SHCMD("pamixer -t; kill -44 $(pidof gocaudices)") },`.

## (Non)-Features
+ Multiple blocks can have the same update signal.

## FQA -- Frequently Questioned Answers
+ Does it have bar click-ability?

	• No, but if you'd like to add that feature or create a patch, create a pull request!

## License
Gocaudices, along with all software created by me, is in the public domain, as all software should be.

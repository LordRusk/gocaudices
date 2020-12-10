# An extremely simple dwmblocks alternative
Gocaudices is a dwmblocks alternative written in GO in less than 100 SLOC using [xgb](https://github.com/BurntSushi/xgb).

## Why?
Dwmblocks in incredibly poorly written, it gets the job done, but not very well. I wrote this alternative to be simple, fast, and elegant. This project is never meant to exceed 100 SLOC in the base build.

## How-To
Gocaudices is configured through editing it's source code. This makes it extremely slim, fast, and secure. Since `main.go` is only ~100 SLOC, it makes it extremely simple to edit and add features.

- Download

First download gocaudices with `git clone https://github.com/lordrusk/gocaudices`. You can (re)compile with `go install`.

- Configure

Gocaudices can be configured through adding `Blocks` in the `blocks.go` file. Add individual scripts or commands, their update intriguers (undefined or 0 value = only update on signal), and their update signals (undefined or 0 value = no update signal). I've left an example of a normal `blocks.go` config file that works with [my dotfiles](https://github.com/lordrusk/artixdwm).

- Shell commands

Commands are not run in any shell, rather with `os/exec`. Because of this, a `block` defined as `Block{Cmd: "georona | cut -d' ' -f1,3"}, UpSig: 19},` will not work. If you want to run a shell command, something like this `Block{Cmd: "dash", Args: []string{"-c", "georona | cut -d' ' -f1,3"}, UpSig: 19},` will work.

- Update a module

Note that multiple blocks can have the same update signal. The `Block` definition of `Block{Cmd: "volume", UpSig: 10},` would be updated like `kill -44 $(pidof gocaudices)`. A dwm volume mute keybind might look like `{ 0, XF86XK_AudioMute, spawn, SHCMD("pamixer -t; kill -44 $(pidof gocaudices)") },`.

## (Non)-Features
+ Gocaudices automatically trims raw bytes from the end of block outputs on a block by block basis. This keeps the bar looking nice.

## FQA -- Frequently Questioned Answers
+ Does it have bar click-ability?

	• No, but if you'd like to add that feature or create a patch, create a pull request!

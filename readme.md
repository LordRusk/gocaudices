# An extremely simple dwmblocks alternative
Gocaudices is a dwmblocks alternative written in GO in less than 100 SLOC using [xgb](https://github.com/BurntSushi/xgb).

## Why?
Dwmblocks in incredibly poorly written, it gets the job done, but not very well. I wrote this alternative to be simple, fast, and elegant.

## How-To
Gocaudices is configured through editing it's source code. This makes it extremly slim, fast, and secure. Since `main.go` is only ~100 SLOC, it makes it extremely simple to edit and add features.

- Download

first download gocaudices with `git clone https://github.com/lordrusk/gocaudices`. You can (re)compile with `go install .`.

- Configure

Gocaudices can be configured through adding `Blocks` in the `blocks.go` file. Add individual scripts or commands, their update intriguers (0 means it will only update on signal), and their update signals. I've left an example of a normal `blocks.go` config file that works with [my dotfiles](https://github.com/lordrusk/artixdwm).

- Update a module

The `Block` definition of `Block { Cmd: "volume", UpInt: 0, UpSig: 10, },` would be updated like `kill -$((34+10)) $(pidof gocaudices)`. A dwm volume mute keybind might look like `{ 0, XF86XK_AudioMute, spawn, SHCMD("pamixer -t; kill -$((34+10)) $(pidof gocaudices)") },`.

## (Non)-Features

+ Gocaudices automatically trims raw bytes from the end of block outputs on a block by block basis. this keeps the bar looking nice.

## FQA -- Frequently Questioned Answers

+ Does it have bar click-ability?

	• Not right now, after I work out a few kinks I'll write a patch that will be compatible with the patch already on suckless.org for bar click-ablity.

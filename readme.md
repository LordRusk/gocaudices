# A dwmblocks alternative written in go
Goblocks is a dwmblocks alternative written in GO.

## Why?
Dwmblocks in incredibly poorly written, it gets the job done, but not very well. I wrote this alternative to be simple, fast, and elegant.

## Example
I've left an example of a normal `blocks.h` config file, and the equivalent `blocks.go` found in the file of the same name.

## How-To
- Configure goblocks

Goblocks can be configured through adding `Blocks` in the `blocks.go` file. Add individual scripts, their update intriguers (0 means it will only update on signal), and their update signals.

- Update a module

The `Block` definition of `Block { Cmd: "volume", UpInt: 0, UpSig: 10, },` would be updated like `kill -((34+10)) $(pidof goblocks)`. A dwm volume mute keybind might look like `{ 0, XF86XK_AudioMute, spawn, SHCMD("pamixer -t; kill -$((34+10)) $(pidof goblocks)") },`.

## FQA -- Frequently Questioned Answers

+ Does it have bar click-ability?

	 Not right now, after I work out a few kinks I'll write a patch that will be compatible with the patch already on suckless.org for bar click-ablity.

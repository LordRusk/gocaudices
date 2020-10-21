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

The `Block` definition of `Block{Cmd: "volume", UpSig: 10},` would be updated like `kill -$((34+10)) $(pidof gocaudices)`. A dwm volume mute keybind might look like `{ 0, XF86XK_AudioMute, spawn, SHCMD("pamixer -t; kill -$((34+10)) $(pidof gocaudices)") },`.

- Deeper explanation.

The `block` struct has 5 fields, `Cmd string`, `Args []string`, `UpInt int`, `UpSig int`, and `Pos int`. `Cmd` defines the command, which internally is split into an array of strings, `Cmd` is redefined as the first item in this array, and `Args` gets what's left over. The reason for doing this is so commands like `disk /home` will be properly ran, while keeping the `blocks.go` from looking ugly. If `Arg` is defined in `block` than it won't do the previous process and handle running it just as defined. This can be used, as shown above, to run shell commands. `UpInt` and `UpSig` are integrule parts of gocaudices. `UpInt` sets the update intigure. If this value isn't 0 when initializing all the blocks, it will start a loop where every `time.Duration(UpInt) * time.Second)` it will update the `block`. `UpSig` sets the update signal + 34. If the value is 0, it won't be externally updatable. I'm not sure why you would do this for a block, but its there if you please. The way gocaudices keeps track of block output is with a slice with length of `len(Blocks)`. `Pos` is used to keep track of where in the array the new output text should go. `Pos` should never be defined in the `Blocks.go` file. Even if it is, it will be redefined when all the blocks get initialized.

## (Non)-Features

+ Gocaudices automatically trims raw bytes from the end of block outputs on a block by block basis. This keeps the bar looking nice.

## FQA -- Frequently Questioned Answers

+ Does it have bar click-ability?

	• Not right now, after I work out a few kinks I'll write a patch that will be compatible with the patch already on suckless.org for bar click-ablity.

# statuscmd
HOOORAYYYY BAR CLICKABILITY!!!!!!!!!

# why this patch is horrific
There are a lot of roadblocks with making this patch. So many it took me over a year of hacking on and off to get a working model. This is that working model. First, we have no access to the `siginfo_t` struct from go and I don't [think we will for a while...](https://github.com/golang/go/issues/9764). Secondly, [only async-signal-safe](https://github.com/golang/go/issues/45499) code can run from a C signal handler...0% of go code is async-signal-safe. This means we can't even do something like calling an exported go function from C using cgo (very possible and almost how this worked).

Alright, so now that you know what you are working with, just take a look at the absolute hack of a patch I came up with for this bullshit. This should be prone to crashes...I think....

If we ever get access to the `siginfo_t` struct from pure go than I can restore the simplicity of gocaudices inside this patch. Until then, this mutilates gocaudices and only exists because I've been trying to write this patch for over a year.

Use at your own risk.

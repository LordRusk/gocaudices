# Complex Delimiter
This patch adds the ability to have a more complex delimiter.

Old way: delimiter: `|` would look like `block|block|block`

New way: delimiter: `[` `] [` `]` would look like `[block] [block] [block]`

## TODO
Figure an efficient way to get rid of the `append` to get `finalBytes`. Append copy's the current slice, and doing that every time we build the bar text isn't a good idea. With my Xrate set high, holding my volume up button, though no lag, does cause about double the cpu usage as the normal delimiter.

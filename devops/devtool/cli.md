## Cli tool

[The Art of Command Line](https://github.com/jlevy/the-art-of-command-line)


wget

curl

tree
- http://mama.indstate.edu/users/ice/tree

jq
- Command-line JSON processor
- https://github.com/jqlang/jq
- https://github.com/antonmedv/fx

ripgrep
- line-oriented search tool
- https://github.com/BurntSushi/ripgrep

fzf
- command-line fuzzy finder
- github: https://github.com/junegunn/fzf
- usage: https://oskernellab.com/2021/02/15/2021/0215-0001-Using_FZF_to_Improve_Productivity/

autojump
- a faster way to navigate your filesystem
- https://github.com/wting/autojump



### date time
- Convert Unix timestamp to a date string
    - `date -d "@$TIMESTAMP"` or `date -r "$TIMESTAMP"`
    - https://stackoverflow.com/questions/3249827/convert-unix-timestamp-to-a-date-string
- generate Unix timestamps
    - date +%s


### base64
- encode
    - echo -n 'hello' | base64
- decode
    - echo -n 'aGVsbG8=' | base64 --decode


### Count code
- https://github.com/XAMPPRocky/tokei

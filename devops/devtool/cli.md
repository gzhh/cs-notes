## Cli tool

[The Art of Command Line](https://github.com/jlevy/the-art-of-command-line)


wget

curl

tree
- http://mama.indstate.edu/users/ice/tree

jq
- Command-line JSON processor
- https://github.com/jqlang/jq

ripgrep
- line-oriented search tool
- https://github.com/BurntSushi/ripgrep


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



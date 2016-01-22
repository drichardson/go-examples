# escape analysis examples
View escape analysis information with the `-m` option to the `go tool compile` command.
Flags to the compile command can also be passed via `go build -gcflags '-m'`.

For example:

    go tool compile -m FILENAME

or

    go build -gcflags -m

You can also stop inlining with `-l` like so:


    go tool compile -m -l

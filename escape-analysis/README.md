# escape analysis examples
View escape analysis information with:

    go build -gcflags -m

You can also stop inlining with `-l` like so:


    go build -gcflags '-m -l'

#! /bin/bash

case "$1" in
    run)
        go build -o game.o ./main.go && ./game.o
        ;;
    build)
        go build -o game.o ./main.go
        ;;
    windows)
        GOOS=windows GOARCH=386 go build -o game.exe main.go
        zip windows.zip assets/* game.exe
        ;;
    todo)
        grep TODO -RIn --exclude={README.md,jam.sh} *
        ;;
    clean)
        rm -f *.{exe,o,zip}
        ;;
    *)
        echo "Usage $0 {run|build|windows|clean|todo}"
        exit 1
esac

#!/bin/bash

if [ -d "build" ]; then
    rm -rf "build"
fi

build_for_platform() {
    GOOS=$1
    GOARCH=$2

    if [ "$GOARCH" == "arm" ]; then
        HUMANARCH="arm"
    elif [ "$GOARCH" == "386" ]; then
        HUMANARCH="x86"
    else
        HUMANARCH="$GOARCH"
    fi

    echo "Building for $GOOS $HUMANARCH..."
    GOOS=$GOOS 
    GOARCH=$GOARCH 
    go build -ldflags="-s -w -X main.isCLI=true" -o "build/$GOOS-$HUMANARCH/controlify_cli" controlify.go
    go build -ldflags="-s -w" -o "build/$GOOS-$HUMANARCH/controlify_tray" controlify.go
    cp config.json "build/$GOOS-$HUMANARCH/config.json"
}

build_for_platform "linux" "386"
build_for_platform "linux" "arm"

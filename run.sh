#!/bin/bash

export PATH=/Users/isaac/bin:/opt/homebrew/bin:/opt/homebrew/sbin:/opt/homebrew/include:$PATH

export CGO_CFLAGS="-I/opt/homebrew/include"
export CGO_LDFLAGS="-L/opt/homebrew/lib"
export PKG_CONFIG_PATH="/opt/homebrew/lib/pkgconfig"

go run main.go

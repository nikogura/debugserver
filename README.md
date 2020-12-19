# debugserver

Very simple HTTP server for debugging clients that don't always do what their doc say they do.

## Installation

    go get github.com/nikogura/debugserver

## Usage

Standard port 8888

    debugserver

Custom port via flags:

    debugserver -p 8989

Custom port via args:

    debugserver 9999
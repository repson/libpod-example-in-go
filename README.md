# podman-example-in-go

In this repository you can see some tests made with the podman library in Go.

Install dependences:

```
$ go get .
```

Execute:

```
$ go run main.go
```

A snippet of a sample run looks like:

```
Connected to Podman
Pulling Alpine Nginx image...
Trying to pull quay.io/libpod/alpine_nginx:latest...
Getting image source signatures
Copying blob sha256:d2c7362ca710ad35a846a34571a7c3450ea3cce04efcbcb4d3af276eda154ade
Copying blob sha256:df9b9388f04ad6279a7410b85cedfdcb2208c0a003da7ab5613af71079148139
Copying blob sha256:71895e83ea49901b7b752bbf3ca19a54148a5f4ab5fdff3dca9bcd59d44c59e3
Copying config sha256:ecea49d99daa5bd62ebaef1338f6bc4c948bf2651b139160404f9c1c48fcd85c
Writing manifest to image destination
WARNING: image platform (linux/amd64) does not match the expected platform (linux/arm64)
Images:
[quay.io/libpod/alpine_nginx:latest]
Container created with ID: a9da8b59a19719721af80db27e66b73154c07002557b9ee3c6d0017a0a17d21f
Starting container...
Container started
Containers: [{false [nginx -g daemon off;] 2024-06-30 19:59:57.230113862 +0200 CEST   false -62135596800 0 map[80:[tcp]] a9da8b59a19719721af80db27e66b73154c07002557b9ee3c6d0017a0a17d21f quay.io/libpod/alpine_nginx:latest ecea49d99daa5bd62ebaef1338f6bc4c948bf2651b139160404f9c1c48fcd85c false map[] [] [foobar] {      } [podman] 11296   [] 0 <nil> 1719770397 running starting}]
Container Inspect Name: quay.io/libpod/alpine_nginx:latest
Container Inspect Status: running
Stopping container...
Container Inspect Status: exited
```

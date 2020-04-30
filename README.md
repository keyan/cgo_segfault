# cgo_segfault

Example code demonstrating setup of a global constructor to install a signal handler to catch signals occuring within [cgo](https://golang.org/cmd/cgo/) called C/C++ functions.

## Background
Recently I ran into an issue where I was calling into a C++ function, accessing invalid memory, and causing a segmentation fault which results in the operating system issuing a `SIGSEGV` signal to the parent Go process, which then crashes. Unfortunately, the Go signal handler doesn't know how to unwind the C stack and generate a panic so instead it just outputs something like:

```
go fatal error: unexpected signal during runtime execution
```

This code installs a signal handler in C before the Go process starts, intentionally casuses a segmentation fault, handles the resulting OS signal, and exits. It doesn't properly dump a backtrace, for reasons which are explained in more detail here: https://keyanp.com/cgo-segfault.html

# moprogress

The progress bar from the docker/moby project, repackaged  as a standalone library.

Have you ever wanted a progress bar for your golang-based CLI tool only to find that 
the library either does not work on all platforms or does not have multibar support?  
Me too!  I've always liked the `docker` cli progress bar so I went ahead and broke
it out into it's own library.

Note that since I am not the original author of this code I won't accept new features;
this code is considered "complete".

# Installation

```
go get github.com/pcj/moprogress@latest
```

# Usage

Create a `moprogess.Output` instance over an `io.Writer`:

```go
out := moprogress.NewOut(os.Stdout)

progress := moprogress.NewProgressOutput(out)
```

Then, send progress updates to the output:

```go
progress.WriteProgress(progress.Progress{
    ID:      "some.tar.gz",
    Action:  "downloading",
    Total:   int64(bytesTotal),
    Current: int64(bytesSent),
    Units:   "bytes",
})
```

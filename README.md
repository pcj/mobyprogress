# mobyprogress

[![ci](https://github.com/pcj/mobyprogress/actions/workflows/ci.yaml/badge.svg?branch=master&event=push)](https://github.com/pcj/mobyprogress/actions/workflows/ci.yaml)
[![godoc](https://godoc.org/github.com/pcj/mobyprogress?status.svg)](https://godoc.org/github.com/pcj/mobyprogress)

The progress bar from the docker/moby project, repackaged  as a standalone library.

Have you ever wanted a progress bar for your golang-based CLI tool only to find that 
the library either does not work on all platforms or does not have multibar support?  
Me too!  I've always liked the `docker` cli progress bar so I went ahead and broke
it out into it's own library.

Note that since I am not the original author of this code I won't accept new features;
this code is considered "complete".

# Installation

```
go get github.com/pcj/mobyprogress@latest
```

# Usage

Create a `mobyprogess.Output` instance over an `io.Writer`:

```go
out := mobyprogress.NewOut(os.Stdout)

progress := mobyprogress.NewProgressOutput(out)
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

The `progress.ID` field is used as the key to identify the progress bar to update. 
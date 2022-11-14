package mobyprogress

import (
	"fmt"
	"io"
)

const streamNewline = "\r\n"

func NewProgressOutput(out io.Writer) Output {
	return &progressOutput{sf: &rawProgressFormatter{}, out: out, newLines: true}
}

type rawProgressFormatter struct{}

func (sf *rawProgressFormatter) formatStatus(id, format string, a ...interface{}) []byte {
	return []byte(fmt.Sprintf(format, a...) + streamNewline)
}

func (sf *rawProgressFormatter) formatProgress(id, action string, progress *JSONProgress, aux interface{}) []byte {
	if progress == nil {
		progress = &JSONProgress{}
	}
	endl := "\r"
	if progress.String() == "" {
		endl += "\n"
	}
	return []byte(action + " " + progress.String() + endl)
}

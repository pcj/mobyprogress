package moprogress

import (
	"encoding/json"
	"io"
)

// NewJSONProgressOutput returns a progress.Output that formats output
// using JSON objects
func NewJSONProgressOutput(out io.Writer, newLines bool) Output {
	return &progressOutput{sf: &jsonProgressFormatter{}, out: out, newLines: newLines}
}

type jsonProgressFormatter struct{}

func appendNewline(source []byte) []byte {
	return append(source, []byte(streamNewline)...)
}

func (sf *jsonProgressFormatter) formatStatus(id, format string, a ...interface{}) []byte {
	return FormatStatus(id, format, a...)
}

// formatProgress formats the progress information for a specified action.
func (sf *jsonProgressFormatter) formatProgress(id, action string, progress *JSONProgress, aux interface{}) []byte {
	if progress == nil {
		progress = &JSONProgress{}
	}
	var auxJSON *json.RawMessage
	if aux != nil {
		auxJSONBytes, err := json.Marshal(aux)
		if err != nil {
			return nil
		}
		auxJSON = new(json.RawMessage)
		*auxJSON = auxJSONBytes
	}
	b, err := json.Marshal(&JSONMessage{
		Status:          action,
		ProgressMessage: progress.String(),
		Progress:        progress,
		ID:              id,
		Aux:             auxJSON,
	})
	if err != nil {
		return nil
	}
	return appendNewline(b)
}

package moprogress

import (
	"strings"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestRawProgressFormatterFormatStatus(t *testing.T) {
	sf := rawProgressFormatter{}
	res := sf.formatStatus("ID", "%s%d", "a", 1)
	assert.Check(t, is.Equal("a1\r\n", string(res)))
}

func TestRawProgressFormatterFormatProgress(t *testing.T) {
	sf := rawProgressFormatter{}
	jsonProgress := &JSONProgress{
		Current: 15,
		Total:   30,
		Start:   1,
	}
	res := sf.formatProgress("id", "action", jsonProgress, nil)
	out := string(res)
	assert.Check(t, strings.HasPrefix(out, "action [===="))
	assert.Check(t, is.Contains(out, "15B/30B"))
	assert.Check(t, strings.HasSuffix(out, "\r"))
}

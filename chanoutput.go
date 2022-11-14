package mobyprogress

type chanOutput chan<- Progress

func (out chanOutput) WriteProgress(p Progress) error {
	// FIXME: workaround for panic in #37735
	defer func() {
		recover()
	}()
	out <- p
	return nil
}

// ChanOutput returns an Output that writes progress updates to the
// supplied channel.
func ChanOutput(progressChan chan<- Progress) Output {
	return chanOutput(progressChan)
}

package moprogress

type Output interface {
	WriteProgress(Progress) error
}

package mobyprogress

type Output interface {
	WriteProgress(Progress) error
}

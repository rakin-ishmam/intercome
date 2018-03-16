package invite

type task struct {
	fileLoc string

	err error
}

func (t *task) Read(fileLoc string) {

}

func (t *task) List(f Filter) []Customer {

	return []Customer{}
}

func (t task) Err() error {
	return t.err
}

// NewInviter returns instance of Inviter
func NewInviter() Inviter {
	return &task{}
}

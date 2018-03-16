package invite

import (
	"bufio"
	"encoding/json"
	"os"
)

type task struct {
	fileLoc string

	customers []Customer

	err error
}

func (t task) hasErr() bool {
	if t.err != nil {
		return true
	}

	return false
}

func (t *task) clear() {
	t.err = nil
	t.customers = []Customer{}
}

func (t *task) append(str string) {
	if t.hasErr() {
		return
	}

	cust := Customer{}
	err := json.Unmarshal([]byte(str), &cust)
	if err != nil {
		t.err = err
		return
	}

	t.customers = append(t.customers, cust)
}

func (t *task) Read(fileLoc string) {
	t.clear()

	f, err := os.Open(fileLoc)
	if err != nil {
		t.err = err
		return
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		t.append(s.Text())
	}

}

func (t *task) List(f Filter) []Customer {

	return t.customers
}

func (t task) Err() error {
	return t.err
}

// NewInviter returns instance of Inviter
func NewInviter() Inviter {
	return &task{}
}

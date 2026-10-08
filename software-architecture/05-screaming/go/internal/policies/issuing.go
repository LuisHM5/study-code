package policies

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

// IssueInput son los datos necesarios para emitir una póliza.
type IssueInput struct {
	Number    string
	StartDate time.Time
	EndDate   time.Time
}

// Issue emite una póliza. En el punto de uso se lee: desk.Issue(...).
func (d *Desk) Issue(in IssueInput) (Policy, error) {
	if !in.EndDate.After(in.StartDate) {
		return Policy{}, ErrInvalidTerm
	}
	if d.store.numberTaken(in.Number) {
		return Policy{}, ErrDuplicated
	}

	p := Policy{ID: newID(), Number: in.Number, StartDate: in.StartDate, EndDate: in.EndDate}
	d.store.save(p)
	d.notifier.PolicyIssued(p.Number)
	return p, nil
}

func newID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

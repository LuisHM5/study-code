package policies

import (
	"encoding/json"
	"time"
)

const dateLayout = "2006-01-02"

// Date es una fecha sin hora que se serializa como YYYY-MM-DD.
type Date struct {
	t time.Time
}

func ParseDate(s string) (Date, error) {
	t, err := time.Parse(dateLayout, s)
	return Date{t: t}, err
}

func (d Date) String() string               { return d.t.Format(dateLayout) }
func (d Date) After(other Date) bool        { return d.t.After(other.t) }
func (d Date) AddMonths(months int) Date    { return Date{t: d.t.AddDate(0, months, 0)} }
func (d Date) MarshalJSON() ([]byte, error) { return json.Marshal(d.String()) }

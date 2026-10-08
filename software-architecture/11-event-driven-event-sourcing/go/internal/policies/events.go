package policies

// Los eventos son hechos en pasado e inmutables. Son la fuente de verdad.
type Event interface {
	EventType() string
}

type Issued struct {
	PolicyID  string `json:"policyId"`
	Number    string `json:"number"`
	StartDate Date   `json:"startDate"`
	EndDate   Date   `json:"endDate"`
}

func (Issued) EventType() string { return "PolicyIssued" }

type Renewed struct {
	PolicyID   string `json:"policyId"`
	Months     int    `json:"months"`
	NewEndDate Date   `json:"newEndDate"`
}

func (Renewed) EventType() string { return "PolicyRenewed" }

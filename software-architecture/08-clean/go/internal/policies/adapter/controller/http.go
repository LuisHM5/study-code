// Package controller: círculo 3. Traduce HTTP al Input del caso de uso.
package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"arquitecturas/08-clean/internal/policies/adapter/presenter"
	"arquitecturas/08-clean/internal/policies/usecase"
)

const dateLayout = "2006-01-02"

type PolicyController struct {
	issuePolicy usecase.IssuePolicy
	getPolicy   usecase.GetPolicy
}

func NewPolicyController(issue usecase.IssuePolicy, get usecase.GetPolicy) *PolicyController {
	return &PolicyController{issuePolicy: issue, getPolicy: get}
}

// issuePolicyRequest es el modelo HTTP, propio del adaptador.
type issuePolicyRequest struct {
	Number    string `json:"number"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

func (c *PolicyController) Issue(w http.ResponseWriter, r *http.Request) {
	var req issuePolicyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		presenter.BadRequest(w, "json inválido")
		return
	}
	start, errStart := time.Parse(dateLayout, req.StartDate)
	end, errEnd := time.Parse(dateLayout, req.EndDate)
	if req.Number == "" || errStart != nil || errEnd != nil {
		presenter.BadRequest(w, "number, startDate y endDate (YYYY-MM-DD) son obligatorios")
		return
	}

	out, err := c.issuePolicy.Execute(r.Context(), usecase.IssuePolicyInput{
		Number: req.Number, StartDate: start, EndDate: end,
	})
	if err != nil {
		presenter.Error(w, err)
		return
	}
	presenter.Created(w, out)
}

func (c *PolicyController) Get(w http.ResponseWriter, r *http.Request) {
	out, err := c.getPolicy.Execute(r.Context(), r.PathValue("id"))
	if err != nil {
		presenter.Error(w, err)
		return
	}
	presenter.Policy(w, out)
}

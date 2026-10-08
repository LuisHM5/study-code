// Package presenter: círculo 3. Convierte la salida de los casos de uso en respuestas HTTP.
package presenter

import (
	"encoding/json"
	"errors"
	"net/http"

	"arquitecturas/08-clean/internal/policies/entity"
	"arquitecturas/08-clean/internal/policies/usecase"
)

const dateLayout = "2006-01-02"

func Created(w http.ResponseWriter, out usecase.IssuePolicyOutput) {
	write(w, http.StatusCreated, map[string]string{"id": out.ID})
}

func Policy(w http.ResponseWriter, out usecase.GetPolicyOutput) {
	write(w, http.StatusOK, map[string]string{
		"id":        out.ID,
		"number":    out.Number,
		"startDate": out.StartDate.Format(dateLayout),
		"endDate":   out.EndDate.Format(dateLayout),
	})
}

func BadRequest(w http.ResponseWriter, message string) {
	write(w, http.StatusBadRequest, map[string]string{"message": message})
}

func Error(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	message := "error interno"
	switch {
	case errors.Is(err, entity.ErrInvalidTerm):
		status, message = http.StatusBadRequest, err.Error()
	case errors.Is(err, entity.ErrDuplicated):
		status, message = http.StatusConflict, err.Error()
	case errors.Is(err, entity.ErrNotFound):
		status, message = http.StatusNotFound, err.Error()
	}
	write(w, status, map[string]string{"message": message})
}

func write(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

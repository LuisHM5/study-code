// Package contracts es la frontera entre servicios: el JSON que viaja por la red.
// Es lo único que comparten (equivale al .proto en gRPC). Cambiarlo exige versionar.
package contracts

type PolicyIssuedNotification struct {
	PolicyID string `json:"policyId"`
	Number   string `json:"number"`
}

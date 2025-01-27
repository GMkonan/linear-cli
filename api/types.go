package api

type GraphqlResponse[T any] struct {
	Data struct {
		Operation T `json:"operation"` // This will be unmarshaled into whatever type T is
	} `json:"data"`
}

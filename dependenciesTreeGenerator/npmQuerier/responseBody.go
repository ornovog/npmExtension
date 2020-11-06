package npmQuerier

type responseBody struct {
	Name        string            `json:"name"`
	Package     packageData       `json:"versions"`
}
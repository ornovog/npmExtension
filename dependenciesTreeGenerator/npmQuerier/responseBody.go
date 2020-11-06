package npmQuerier

type responseBody struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	DistTags    map[string]string `json:"dist-tags"`
	Package     packageData       `json:"versions"`
}
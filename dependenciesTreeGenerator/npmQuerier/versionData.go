package npmQuerier

type versionData struct {
	Name string `json:"name"`
	Version string `json:"version"`
	Dependencies map[string]string `json:"dependencies"`
}
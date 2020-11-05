package dependenciesTreeGenerator

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	Url "net/url"
	"path"
)
var url = "https://registry.npmjs.org/"

type ResponseBody struct {
	Name        string             `json:"name"`
	Description string             `json:"description"`
	DistTags    map[string]string  `json:"dist-tags"`
	Versions    map[string]Version `json:"versions"`
}

type Version struct {
	Name string `json:"name"`
	Version string `json:"version"`
	Dependencies map[string]string `json:"dependencies"`
}

func GetNpmDependencies(package_ Package)map[string]string{
	urlStr := extractNpmUrl(package_)

	Response, _ := http.Get(urlStr)

	RespBody, _ := parseResponse(Response)

	version := extractPackageVersion(package_)

	dependencies := RespBody.Versions[version].Dependencies

	return dependencies
}

func parseResponse(Response *http.Response) (ResponseBody, error){
	body, err := ioutil.ReadAll(Response.Body)
	if err != nil{
		return ResponseBody{}, err
	}

	var respBody ResponseBody
	json.Unmarshal(body, &respBody)
	if err!=nil{
		return ResponseBody{}, err
	}

	return respBody, nil
}

func extractNpmUrl(package_ Package) string{
	npmUrl, _ := Url.Parse(url)
	npmUrl.Path = path.Join(npmUrl.Path, package_.Name)
	npmUrlStr := npmUrl.String()
	return npmUrlStr
}

func extractPackageVersion(package_ Package) string {
	version := package_.Version
	if version[0] == '^' {
		version = version[1:]
	}
	return version
}
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
	Name        string            `json:"name"`
	Description string            `json:"description"`
	DistTags    map[string]string `json:"dist-tags"`
	Package     PackageData       `json:"versions"`
}

type VersionData struct {
	Name string `json:"name"`
	Version string `json:"version"`
	Dependencies map[string]string `json:"dependencies"`
}

type PackageData map[string]VersionData

type npmQuerier struct {
	packagesDataCache map[string]PackageData
}

func NewNpmQuerier() npmQuerier {
	q := npmQuerier{}
	q.packagesDataCache = make(map[string]PackageData, 0)
	return q
}

func (q npmQuerier) GetNpmDependencies(package_ Package)map[string]string{
	packageName := package_.Name
	version := extractPackageVersion(package_)

	packageData, ok := q.packagesDataCache[packageName]
	if !ok{
		packageData = queryNpm(packageName)
		q.packagesDataCache[packageName] = packageData
	}

	versionData := packageData[version]
	dependencies := versionData.Dependencies

	return dependencies
}

func queryNpm(packageName string) PackageData {
	urlStr := extractNpmUrl(packageName)
	Response, _ := http.Get(urlStr)
	RespBody, _ := parseResponse(Response)
	packageData := RespBody.Package
	return packageData
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

func extractNpmUrl(packageName string) string{
	npmUrl, _ := Url.Parse(url)
	npmUrl.Path = path.Join(npmUrl.Path, packageName)
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
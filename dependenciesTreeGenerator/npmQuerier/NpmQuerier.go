package npmQuerier

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	Url "net/url"
	inter "npmExtension/dependenciesTreeGenerator"
	"path"
)
var url = "https://registry.npmjs.org/"

type NpmQuerier struct {
	packagesDataCache map[string]packageData
}

func NewNpmQuerier() NpmQuerier {
	q := NpmQuerier{}
	q.packagesDataCache = make(map[string]packageData, 0)
	return q
}

func (q NpmQuerier) GetNpmDependencies(package_ inter.Package)map[string]string{
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

func queryNpm(packageName string) packageData {
	urlStr := extractNpmUrl(packageName)
	Response, _ := http.Get(urlStr)
	RespBody := parseResponse(Response)
	packageData := RespBody.Package
	return packageData
}

func extractPackageVersion(package_ inter.Package) string {
	version := package_.Version
	if version[0] == '^' {
		version = version[1:]
	}
	return version
}

func extractNpmUrl(packageName string) string{
	npmUrl, _ := Url.Parse(url)
	npmUrl.Path = path.Join(npmUrl.Path, packageName)
	npmUrlStr := npmUrl.String()
	return npmUrlStr
}

func parseResponse(Response *http.Response)responseBody{
	body, _ := ioutil.ReadAll(Response.Body)

	var respBody responseBody
	json.Unmarshal(body, &respBody)

	return respBody
}



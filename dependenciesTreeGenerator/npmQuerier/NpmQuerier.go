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
		packageData, _ = queryNpm(packageName)
		q.packagesDataCache[packageName] = packageData
	}

	versionData := packageData[version]
	dependencies := versionData.Dependencies

	return dependencies
}

func queryNpm(packageName string) (packageData, error) {
	urlStr := extractNpmUrl(packageName)

	response, err := http.Get(urlStr)
	if err!=nil{
		return packageData{},err
	}

	respBody, _ := parseResponse(response)
	if err!=nil{
		return packageData{}, err
	}

	packageData := respBody.Package
	return packageData, nil
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

func parseResponse(Response *http.Response)(responseBody, error){
	body, err := ioutil.ReadAll(Response.Body)
	if err!=nil{
		return responseBody{}, err
	}

	var respBody responseBody
	err = json.Unmarshal(body, &respBody)
	if err!=nil{
		return responseBody{}, err
	}

	return respBody, nil
}



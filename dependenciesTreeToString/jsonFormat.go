package dependenciesTreeToString

import (
	"encoding/json"
	treeGen "npmExtension/dependenciesTreeGenerator"
)

type jsonFormat struct {
}

type jsonFormatDependencies struct {
	Name string
	Dependencies map[string]string
}

func NewJsonFormat() jsonFormat{
	return jsonFormat{}
}

func (_ jsonFormat) DependenciesTreeToString(packageNode treeGen.PackageNode) string{
	j := jsonFormatDependencies{Name: packageNode.Package.Name, Dependencies: map[string]string{}}

	j.fillDependenciesFlat(packageNode.Dependencies)

	dependenciesJson, _ := json.Marshal(j)

	return string(dependenciesJson)
}

func(j jsonFormatDependencies) fillDependenciesFlat(dependenciesInCurrentLevel []treeGen.PackageNode) {
	for _, dependency := range dependenciesInCurrentLevel{
		name := dependency.Package.Name
		version := dependency.Package.Version

		j.Dependencies[name] = version
	}
}




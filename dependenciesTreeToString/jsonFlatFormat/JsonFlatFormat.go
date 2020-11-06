package jsonFlatFormat

import (
	"encoding/json"
	inter "npmExtension/dependenciesTreeGenerator"
)

type jsonFlatFormat struct {

}

func NewJsonFlatFormat() jsonFlatFormat {
	return jsonFlatFormat{}
}

func (_ jsonFlatFormat) DependenciesTreeToString(packageNode inter.PackageNode) string {
	packageName := packageNode.Package.Name
	dependencyToVersions := getDependencyNameToVersionsSlices(packageNode)

	 fFD :=  flatFormatData{PackageName: packageName, DependencyToVersions: dependencyToVersions}

	dependenciesJson, _ := json.Marshal(fFD)

	return string(dependenciesJson)
}

func getDependencyNameToVersionsSlices(packageRoot inter.PackageNode) stringToStringSlice {
	firstLevelDependencies := packageRoot.Dependencies
	dependencyNameToVersionsSet := stringToStringSet{}
	fillDependencyNameToVersionsSet(firstLevelDependencies, dependencyNameToVersionsSet)

	versionsSlices := versionsSetsToVersionsSlices(dependencyNameToVersionsSet)

	return versionsSlices
}

func fillDependencyNameToVersionsSet(dependenciesInCurrentLevel []inter.PackageNode,
	                          dependencyNameToVersionsSet stringToStringSet) {
	for _, dependency := range dependenciesInCurrentLevel{
		name := dependency.Package.Name
		version := dependency.Package.Version

		updateDependencyNameToVersionsSet(dependencyNameToVersionsSet, name, version)

		fillDependencyNameToVersionsSet(dependency.Dependencies, dependencyNameToVersionsSet)
	}
}

func updateDependencyNameToVersionsSet(dependencyNameToVersionsSet stringToStringSet, name string, version string) {
	_, ok := dependencyNameToVersionsSet[name]
	if !ok {
		dependencyNameToVersionsSet[name] = stringSet{}
	}
	dependencyNameToVersionsSet[name][version] = true
}

func versionsSetsToVersionsSlices(dependencyNameToVersionsSet stringToStringSet) stringToStringSlice {
	versionsSlices := stringToStringSlice{}

	for dependencyName, versionSet := range dependencyNameToVersionsSet {
		versionSlice := versionSetToVersionSlice(versionSet)
		versionsSlices[dependencyName] = versionSlice
	}

	return versionsSlices
}

func versionSetToVersionSlice(versionSet stringSet) []string{
	versionSlice := make([]string, 0, len(versionSet))

	for version := range versionSet {
		versionSlice = append(versionSlice, version)
	}

	return versionSlice
}

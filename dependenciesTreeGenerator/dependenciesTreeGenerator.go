package dependenciesTreeGenerator

type dependenciesTreeGenerator struct {
}

func NewDependenciesTreeGenerator() dependenciesTreeGenerator{
	return dependenciesTreeGenerator{}
}

func (_ dependenciesTreeGenerator) GetPackageDependenciesTree(package_ Package) PackageNode{
	packageRoot := PackageNode{Package: package_}
	fillDependenciesTreeBfs(&packageRoot)
	return packageRoot
}

func fillDependenciesTreeBfs(packageNode *PackageNode){
	npmPackageDependencies := GetNpmDependencies(packageNode.Package)
	if len(npmPackageDependencies) == 0{
		return
	}

	packageNode.Dependencies = npmDependenciesToPackageNodes(npmPackageDependencies)

	for i, _ := range packageNode.Dependencies {
		fillDependenciesTreeBfs(&packageNode.Dependencies[i])
	}
}

func npmDependenciesToPackageNodes(npmDependencies map[string]string)[]PackageNode{
	dependencies := make([]PackageNode, len(npmDependencies))

	index := 0
	for name, version := range npmDependencies {
		dependencyPackage := Package{Name: name, Version: version}
		dependencies[index] = PackageNode{Package:dependencyPackage}
		index++
	}

	return dependencies
}



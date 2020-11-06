package dependenciesTreeGenerator

type dependenciesTreeGenerator struct {
	querier npmQuerier.npmQuerier
}

func NewDependenciesTreeGenerator() dependenciesTreeGenerator{
	querier := NewNpmQuerier()
	return dependenciesTreeGenerator{querier: querier}
}

func (d dependenciesTreeGenerator) GetPackageDependenciesTree(package_ Package) PackageNode{
	packageRoot := PackageNode{Package: package_}
	d.fillDependenciesTreeBfs(&packageRoot)
	UpdateDependenciesBySemanticVersion(&packageRoot)
	return packageRoot
}

func (d *dependenciesTreeGenerator)fillDependenciesTreeBfs(packageNode *PackageNode){
	npmPackageDependencies := d.querier.GetNpmDependencies(packageNode.Package)
	if len(npmPackageDependencies) == 0{
		return
	}

	packageNode.Dependencies = npmDependenciesToPackageNodes(npmPackageDependencies)
	for i, _ := range packageNode.Dependencies {
		d.fillDependenciesTreeBfs(&packageNode.Dependencies[i])
	}
}

func npmDependenciesToPackageNodes(npmDependencies map[string]string)[]PackageNode{
	dependencies := make([]PackageNode, 0, len(npmDependencies))

	for name, version := range npmDependencies {
		dependencyPackage := Package{Name: name, Version: version}
		packageNode := PackageNode{Package:dependencyPackage}

		dependencies = append(dependencies, packageNode)
	}

	return dependencies
}



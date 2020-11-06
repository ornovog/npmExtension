package treeGeneratorImplementaion

import (
	inter "npmExtension/dependenciesTreeGenerator"
	npmquerier"npmExtension/dependenciesTreeGenerator/npmQuerier"
	"npmExtension/dependenciesTreeGenerator/semanticVersioning"
)

type dependenciesTreeGenerator struct {
	querier npmquerier.NpmQuerier
}

func NewDependenciesTreeGenerator() dependenciesTreeGenerator {
	querier := npmquerier.NewNpmQuerier()
	return dependenciesTreeGenerator{querier: querier}
}

func (d dependenciesTreeGenerator) GetPackageDependenciesTree(package_ inter.Package) inter.PackageNode {
	packageRoot := inter.PackageNode{Package: package_}
	d.fillDependenciesTreeBfs(&packageRoot)
	semanticVersioning.UpdateDependenciesBySemanticVersion(&packageRoot)
	return packageRoot
}

func (d *dependenciesTreeGenerator)fillDependenciesTreeBfs(packageNode *inter.PackageNode){
	npmPackageDependencies := d.querier.GetNpmDependencies(packageNode.Package)
	if len(npmPackageDependencies) == 0{
		return
	}

	packageNode.Dependencies = npmDependenciesToPackageNodes(npmPackageDependencies)
	for i, _ := range packageNode.Dependencies {
		d.fillDependenciesTreeBfs(&packageNode.Dependencies[i])
	}
}

func npmDependenciesToPackageNodes(npmDependencies map[string]string)[]inter.PackageNode {
	dependencies := make([]inter.PackageNode, 0, len(npmDependencies))

	for name, version := range npmDependencies {
		dependencyPackage := inter.Package{Name: name, Version: version}
		packageNode := inter.PackageNode{Package: dependencyPackage}

		dependencies = append(dependencies, packageNode)
	}

	return dependencies
}



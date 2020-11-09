package treeGeneratorImplementaion

import (
	inter "npmExtension/dependenciesTreeGenerator"
	npmquerier "npmExtension/dependenciesTreeGenerator/npmQuerier"
	"npmExtension/dependenciesTreeGenerator/semanticVersioning"
)

type dependenciesTreeGenerator struct {
	querier npmquerier.NpmQuerier
}

func NewDependenciesTreeGenerator() dependenciesTreeGenerator {
	querier := npmquerier.NewNpmQuerier()
	return dependenciesTreeGenerator{querier: querier}
}

func (d dependenciesTreeGenerator) GetPackageDependenciesTree(package_ inter.Package) (inter.PackageNode,error) {
	packageRoot := inter.PackageNode{Package: package_}
	err := d.fillDependenciesTreeBfs(&packageRoot)
	if err!= nil{
		return inter.PackageNode{},err
	}

	semanticVersioning.UpdateDependenciesBySemanticVersion(&packageRoot)
	return packageRoot,nil
}

func (d *dependenciesTreeGenerator)fillDependenciesTreeBfs(packageNode *inter.PackageNode)error{
	npmPackageDependencies:= d.querier.GetNpmDependencies(packageNode.Package)

	if len(npmPackageDependencies) == 0{
		return nil
	}

	packageNode.Dependencies = npmDependenciesToPackageNodes(npmPackageDependencies)
	for i, _ := range packageNode.Dependencies {
		d.fillDependenciesTreeBfs(&packageNode.Dependencies[i])
	}

	return nil
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



package dependenciesTreeGenerator

type IDependenciesTreeGenerator interface {
	GetPackageDependenciesTree(package_ Package)PackageNode
}

type Package struct{
	Name string
	Version string
}

type PackageNode struct{
	Package Package
	Dependencies []PackageNode
}
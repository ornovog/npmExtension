package dependenciesTreeGenerator

type IDependenciesTreeGenerator interface {
	GetPackageDependenciesTree(p Package)PackageNode
}

type Package struct{
	Name string
	Version string
}

type PackageNode struct{
	Package Package
	Dependencies []PackageNode
}
package dependenciesTreeGenerator

type IDependenciesTreeGenerator interface {
	GetPackageDependenciesTree(package_ Package)PackageNode
}

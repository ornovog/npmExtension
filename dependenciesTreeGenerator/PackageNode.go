package dependenciesTreeGenerator

type PackageNode struct{
	Package Package
	Dependencies []PackageNode
}
package dependenciesTreeToString

import (
	treeGenerator "npmExtension/dependenciesTreeGenerator"
)

type IDependenciesTreeToString interface {
	DependenciesTreeToString(packageNode treeGenerator.PackageNode) string
}
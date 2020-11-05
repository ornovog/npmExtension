package dependenciesTreeToString

import (
	"github.com/xlab/treeprint"
	treeGen "npmExtension/dependenciesTreeGenerator"
)

type treeFormat struct {
}

func NewTreeFormat() treeFormat{
	return treeFormat{}
}

func (_ treeFormat) DependenciesTreeToString(packageNode treeGen.PackageNode) string{
	printedTree := treeprint.New()
	fillPrintedTreeDfs(packageNode, &printedTree)
	return printedTree.String()
}

func fillPrintedTreeDfs(packageNode treeGen.PackageNode, printedTreeBranch *treeprint.Tree){
	name := packageNode.Package.Name
	version := packageNode.Package.Version
	branch := (*printedTreeBranch).AddBranch(name + "-" + version)

	for _, dependencyNode:= range packageNode.Dependencies{
		fillPrintedTreeDfs(dependencyNode, &branch)
	}
}

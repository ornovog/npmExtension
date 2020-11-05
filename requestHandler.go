package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	treeGen "npmExtension/dependenciesTreeGenerator"
	treeStr "npmExtension/dependenciesTreeToString"
)

type logicHandler struct{
	dependenciesTreeGenerator treeGen.IDependenciesTreeGenerator
	printedTreeGenerator treeStr.IDependenciesTreeToString
}

func (handler logicHandler) HandleRequest(response http.ResponseWriter, _ *http.Request, params httprouter.Params){
	package_ := extractPackageFromRoutParams(params)

	dependenciesTree := handler.dependenciesTreeGenerator.GetPackageDependenciesTree(package_)

	dependenciesTreeStr := handler.printedTreeGenerator.DependenciesTreeToString(dependenciesTree)

	response.Write([]byte(dependenciesTreeStr))
}

func (handler logicHandler) HandleTreeFormatRequest(response http.ResponseWriter, _ *http.Request, params httprouter.Params){
	handler.printedTreeGenerator = treeStr.NewTreeFormat()
	handler.handleRequest(response,nil,params)
	handler.printedTreeGenerator = treeStr.NewJsonFormat()
}

func extractPackageFromRoutParams(params httprouter.Params) treeGen.Package{
	name := params.ByName("name")
	version := params.ByName("version")

	PackageRequest := treeGen.Package{Name: name, Version: version}
	return PackageRequest
}

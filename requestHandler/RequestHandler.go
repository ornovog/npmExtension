package requestHandler

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	treeGen "npmExtension/dependenciesTreeGenerator"
	"npmExtension/dependenciesTreeGenerator/treeGeneratorImplementaion"
	treeStr "npmExtension/dependenciesTreeToString"
	jsonForm "npmExtension/dependenciesTreeToString/jsonFlatFormat"
	treeForm "npmExtension/dependenciesTreeToString/treeFormat"
)

type requestHandler struct{
	dependenciesTreeGenerator treeGen.IDependenciesTreeGenerator
}

func NewRequestHandler() requestHandler{
	dependenciesTreeGenerator := treeGeneratorImplementaion.NewDependenciesTreeGenerator()
	return requestHandler{dependenciesTreeGenerator: dependenciesTreeGenerator}
}

func (handler requestHandler) HandleJsonFormatRequest(response http.ResponseWriter, _ *http.Request, params httprouter.Params){
	treeFormat := jsonForm.NewJsonFlatFormat()
	handler.handleRequest(response, params, treeFormat)
}

func (handler requestHandler) HandleTreeFormatRequest(response http.ResponseWriter, _ *http.Request, params httprouter.Params){
	treeFormat := treeForm.NewTreeFormat()
	handler.handleRequest(response, params, treeFormat)
}

func (handler requestHandler) handleRequest(response http.ResponseWriter, params httprouter.Params,
											printedTreeGenerator treeStr.IDependenciesTreeToString){
	package_ := extractPackageFromRoutParams(params)

	dependenciesTree := handler.dependenciesTreeGenerator.GetPackageDependenciesTree(package_)

	dependenciesTreeStr := printedTreeGenerator.DependenciesTreeToString(dependenciesTree)

	response.Write([]byte(dependenciesTreeStr))
}

func extractPackageFromRoutParams(params httprouter.Params) treeGen.Package{
	name := params.ByName("name")
	version := params.ByName("version")

	PackageRequest := treeGen.Package{Name: name, Version: version}
	return PackageRequest
}


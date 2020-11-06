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
	treeFormat treeStr.IDependenciesTreeToString
	jsonFormat treeStr.IDependenciesTreeToString
}

func NewRequestHandler() requestHandler{
	dependenciesTreeGenerator := treeGeneratorImplementaion.NewDependenciesTreeGenerator()
	treeFormat := treeForm.NewTreeFormat()
	jsonFormat := jsonForm.NewJsonFlatFormat()

	handler:= requestHandler{dependenciesTreeGenerator: dependenciesTreeGenerator,
	                      treeFormat: treeFormat,
	                      jsonFormat: jsonFormat}
	return handler
}

func (handler requestHandler) HandleJsonFormatRequest(response http.ResponseWriter, _ *http.Request,
													  params httprouter.Params){
	handler.handleRequest(response, params, handler.jsonFormat)
}

func (handler requestHandler) HandleTreeFormatRequest(response http.ResponseWriter, _ *http.Request, params httprouter.Params){
	handler.handleRequest(response, params, handler.treeFormat)
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


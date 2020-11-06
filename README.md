For running the code, enter to the project folder (in terminal) and write "go run main.go",
it will upload a server that will listen on localhost:8080.
The server reveals  two endpoints :
1) /package/:name/:version - for flat json results
2) /package/:name/:version/tree-format - for tree structure results

Semantic-Versioning:
In a case where there is a dependency on two (or more) different versions of the same package -
If the difference is major, returns the both versions.
If the difference is not major, returns the more up-to-date version.

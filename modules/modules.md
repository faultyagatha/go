# Modules


`go mod init github.com/faultyagatha/go/modules/moduleone`

- the `go mod init` command creates a go.mod file to track the code's dependencies
- as you add dependencies, the go.mod file will list the versions your code depends on. 
- this keeps builds reproducible and gives you direct control over which module versions to use.


- if the repository is not yet published, the module path can be relative to the project
`go mod edit -replace github.com/faultyagatha/go/modules/moduleone=../moduleone`
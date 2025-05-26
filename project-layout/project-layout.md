# Project Organisation: Best Practices

## Option 1

- based on [Eli Bendersky's structure](https://eli.thegreenplace.net/2019/simple-go-project-layout-with-modules/)

### Naming 

> Idiomatic for Go projects to be named by their GitHub path --> the `project path` is the `module name`.

Example: github.com/your-handle/your-project 
`github.com/your-handle/your-project`

```go
// go.mod file
module github.com/eliben/modlib
```
### Project layout

Example of one project with different modules:
```
├── LICENSE
├── README.md
├── config.go
├── go.mod
├── go.sum
├── clientlib
│   ├── lib.go
│   └── lib_test.go
├── cmd
│   ├── modlib-client
│   │   └── main.go
│   └── modlib-server
│       └── main.go
├── internal
│   └── auth
│       ├── auth.go
│       └── auth_test.go
└── serverlib
    └── lib.go
```

`go.mod` 
- is the module definition file
- contains the module name, a Go version and dependencies
- keep it checked into source control

`go.sum` 
- contains all the dependency checksums
- is managed by the go tools
- keep it checked into source control

`config.go` 
- contains a single trivial function:

```go
package modlib

func Config() string {
  return "modlib config"
}
```
- is what you get when you just `import github.com/eliben/modlib`

```go
package main

import "fmt"
import "github.com/eliben/modlib"

func main() {
  fmt.Println(modlib.Config())
}
```

> If your module provides a single package, or you want to export code from the top-level package of the module, place all the code for this at the top-level directory of the module, and name the package as the last part of the module's path (unless you're using vanity imports, in which case it's more flexible).

`clientlib`:
- `clientlib/lib.go` 

```go
package clientlib

func Hello() string {
  return "clientlib hello"
}
```
This code should be imported with `github.com/eliben/modlib/clientlib` like this:

```go
package main

import "fmt"
import "github.com/eliben/modlib"
import "github.com/eliben/modlib/clientlib"

func main() {
  fmt.Println(modlib.Config())
  fmt.Println(clientlib.Hello())
}
```

`serverlib` follows the same logic

`cmd`:
- conventional location of all the command-line programs made available by the project
- such commands can be installed by the user using the go tool as follows:

```bash
$ go install github.com/eliben/modlib/cmd/cmd-name@latest

# Go downloads, builds and installs cmd-name into the default location.
# You can also pick a specific version after the @ sign, instead of "latest".
# The bin/ directory in the default location is often in $PATH, so we can
# just invoke cmd-name now

$ cmd-name ...
```

Example of the code in the cmd/modlib-server:

```go
package main

// Note absolute path even if importing from sub-dirs!! 
import (
  "fmt"

  "github.com/eliben/modlib"
  "github.com/eliben/modlib/internal/auth"
  "github.com/eliben/modlib/serverlib"
)

func main() {
  fmt.Println("Running server")
  fmt.Println("Config:", modlib.Config())
  fmt.Println("Auth:", auth.GetAuth())
  fmt.Println(serverlib.Hello())
}
```

> In Go, `absolute imports` are the way to go.

`internal`:

- conventional location for all private packages that are used internally by a project but which we don't want to export to users

> Everything exported by your project in v1 becomes a public API, and has to abide by semantic versioning compatibility guarantees. --> It's imperative to `export only the minimal API surface` that's essential for users of your project. All the other code which your package needs for its implementation should live in internal.

> Rule of thumb: in doubt, `better internal`. 

It's easy to take an internal API and export it to users - just a quick renaming/refactoring commit. `It's very painful to take an external API and un-export it` (user code may depend on it); at stable module versions (v1 and beyond), this requires a major version bump to break compatibility.

Example: if the repository contains the source code of the website of the project, place that in `internal/website`.

- put here any internal tools or scripts needed to work on the project

> The idea is that the `root directory of a project should be minimal and clear to users`. In a way, it's `self-documentation`. A user looking at my project's GitHub page should get an immediate sense of where the things they need are located. Since users don't typically really need the stuff used to develop the project, hiding it in internal makes sense.


- Go tooling recognizes internal as a special path

Example with the current folder structure: standalone go projects

```
go/
│
├── basics/
│   ├── go.mod
│   └── arrays/
│       └── main.go
│   └── closures/
│       └── main.go
│
├── concurrency/
│   ├── go.mod
│   └── buffered/
│       └── main.go
│   └── channels/
│       └── main.go
│
├── modules/
│   ├── go.mod
│   └── moduleone/
│       └── moduleone.go
│   └── main/
│       └── main.go
│
├── go.work
```
**Go Mod vs Go Work**

`go mod` 

— required per module
- defines a Go module (unit of code + its dependencies)
- every Go project that you go run, go build, or go test must have a go.mod
- the core of Go's dependency system since Go 1.11
- `go mod init github.com/you/project-name`

`go work` 

- optional helper for multiple modules
- introduced in Go 1.18+
- lets you combine multiple Go modules into a single workspace, so local modules can import each other without requiring a version/tag/push.
- good for monorepos, learning repos, or polyglot projects with multiple Go tools.
- `go work init ./mod1 ./mod2`
- Use it if:
  - You have multiple go.mod files in the same repo.
  - One Go module depends on another.
  - You want to develop and test them together easily.

Rule of Thumb!!
> `Always use go mod`. Use go work when you have multiple modules and want to work with them together locally.

Example:
```
my-projects/
│
├── mod-a/          ← has go.mod
├── mod-b/          ← has go.mod (imports mod-a)
├── go.work         ← optional, links both

```

```bash
cd my-projects
cd mod-a
go mod init github.com/yourname/mod-a
cd ../mod-b
go mod init github.com/yourname/mod-b
go get github.com/yourname/mod-a@v0.0.0
cd ..

go work init ./mod-a ./mod-b
# or if the go mod already exist, just add the project go work use ./path/to/module
go work use ./mod-a
go work use ./mod-b

go run ./mod-b
```


### Package Nesting

- can go as deep as you need
- the package name visible to users is determined by the relative path from the module root. 

For example, if we have a subdirectory called clientlib/tokens with some code in the tokens package, the user will import that with import "github.com/eliben/modlib/clientlib/tokens.

- for some modules a single top-level package is sufficient. 

In the case of modlib this would mean no subdirectories with user-importable packages, but all code being in the top directory in a single or multiple Go files all in package modlib.

> Sidenote: In many projects, there is a `pkg/` dir. In the majority of cases it's an antipattern. 

The code you place in pkg/ should almost certainly be in internal/ instead. If your project is a large top-level application, it shouldn't have importable packages; instead, importable packages should be split out to separate repositories which are small, self-contained and reusable.

[standard layout golang](https://github.com/golang-standards/project-layout)
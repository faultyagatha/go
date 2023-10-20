# Tips. Tricks, Interview Questions

## How to swap values

simple: a,b, = b,a

```go
func swap(a, b string) (string, string) {
  return b, a
}
```

## Structs and pointers
```go
type Vertex struct {
	X int
	Y int
}

//returns a copy of a struct
func funcOne() Vertex {
  return Vertex{X: 1}
}

//returns a pointer to the struct
func funcTwo() *Vertex {
  return &Vertex{}
}

//overrides a value in a struct passed to a func
func funcThree(v *Vertex) {
  v.X = 1
}
```
## Concat strings
```go
import (
    "strings"
    "fmt"
)

func main() {
  var str strings.Builder
  for i := 0; i < 10; i++ {
    str.WriteString("hello")
  }
}
```

## Check if map contains a key

```go
if val, ok := dict["someval"]; ok {
  //do something
}
```

```go
type Node struct {
  Next  *Node
  Value interface{}
}

var first *Node

visited := make(map[*Node]bool)
for n := first; n != nil; n = n.Next {
  if visited[n] {
    fmt.Println("cycle detected")
    break
  }

  visited[n] = true
  fmt.Println(n.Value)
}
```

5. Copy map:
```go
mapOne := map[string]bool{"A": true, "B": true}
mapTwo := make(map[string]bool)
for i, v:= range mapOne {
  mapTwo[i] = v
}
```

## Reverse a slice of ints

```go
func reverse(s []int) {
  for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
    s[i], s[j] = s[j], s[i]
  }
}
```

## Print const (iotas)

```go
type State int

//integers under the hood
const (
	OffHook State = iota
	Connecting
	Connected
	OnHold
	OnHook
)

// ---------------
// String allows to handle
// const ints as strings
//  ---------------
func (s State) String() string {
	switch s {
	case OffHook:
		return "OffHook"
	case Connecting:
		return "Connecting"
	case Connected:
		return "Connected"
	case OnHold:
		return "OnHold"
	case OnHook:
		return "OnHook"
	}
	return "Unknown"
}

func main() {
  var state State //do something with it
  ...
  fmt.Println("The state is currently", state)
}
```
## FIFO Queue: Push / Pop

```go
var queue = []int
//push
queue = append(queue, 1)
queue = append(queue, 9)
queue = append(queue, 19)
//pop
var first int
first, queue = queue[0], queue[1:]
```

## LIFO Stack: Push / Pop

```go
var stack = []int
//push
stack = append(stack, 1)
stack = append(stack, 9)
stack = append(stack, 19)
//pop
var last int
last, stack = stack[len(stack) - 1], stack[:len(stack) - 1]
```

## Convert string 

- to a slice of bytes:

```go
// convert from a string to a slice of bytes 
// because of UTF-8 encoding, some characters in the string may 
// end up in 1, 2, 3 or 4 bytes
mystring := "hello this is string"
byteslice := []byte(mystring)

- to a slice of runes:

// convert from a string to a slice of runes
runeslice  := []rune(mystring)

// from a slice of bytes or runes to a string.
b := []byte{'h','e','l','l','o'} // Composite literal.
s := string(b)
i := []rune{257,1024,65}
r := string(i)
```

## When to use string vs []bytes

- se `string` by default when you're working with text. 
- use `[]byte` instead if one of the following conditions applies:
  - the mutability of a []byte will significantly reduce the number of allocations needed
  - you are dealing with an API that uses []byte, and avoiding a conversion to string will simplify your code.

## DB Querying 

```go
func doQuery() {
  dbTables := []string{"job_m", "job_s", "job_i"}

  for _, table := range dbTables {
    err := db.QueryRow("SELECT id FROM ? WHERE target = ?", table, target).Scan(&jobID)
    if errors.Is(err, sql.ErrNoRows) {
      continue
    }

    if err != nil {
       return fmt.Errorf("whatever")
    }
    // if found, do whatever
  }
}
```

// TODO: use this into doit app
// TodoService manages the TODO data, and the TodoController handles web requests and responses. 
```go
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    todoService := NewTodoService()

    todoController := NewTodoController(todoService)
    todoController.RegisterRoutes(r)

    r.Run(":8080")
}

// todoservice.go
package main

type TodoService struct {
    todos   []Todo
    lastID  int
}

func NewTodoService() *TodoService {
    return &TodoService{}
}

func (s *TodoService) CreateTodo(todo Todo) (Todo, error) {
    s.lastID++
    todo.ID = s.lastID
    s.todos = append(s.todos, todo)
    return todo, nil
}

func (s *TodoService) ListTodos() []Todo {
    return s.todos
}

func (s *TodoService) GetTodoByID(id int) (Todo, error) {
    for _, todo := range s.todos {
        if todo.ID == id {
            return todo, nil
        }
    }
    return Todo{}, ErrTodoNotFound
}

func (s *TodoService) UpdateTodoByID(id int, updatedTodo Todo) (Todo, error) {
    for i, todo := range s.todos {
        if todo.ID == id {
            updatedTodo.ID = id
            s.todos[i] = updatedTodo
            return updatedTodo, nil
        }
    }
    return Todo{}, ErrTodoNotFound
}

func (s *TodoService) DeleteTodoByID(id int) error {
    for i, todo := range s.todos {
        if todo.ID == id {
            s.todos = append(s.todos[:i], s.todos[i+1:]...)
            return nil
        }
    }
    return ErrTodoNotFound
}

// todocontroller.go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type TodoController struct {
    service *TodoService
}

func NewTodoController(service *TodoService) *TodoController {
    return &TodoController{
        service: service,
    }
}

func (c *TodoController) RegisterRoutes(router *gin.Engine) {
    router.POST("/todos", c.CreateTodo)
    router.GET("/todos", c.ListTodos)
    router.GET("/todos/:id", c.GetTodo)
    router.PUT("/todos/:id", c.UpdateTodo)
    router.DELETE("/todos/:id", c.DeleteTodo)
}

func (c *TodoController) CreateTodo(ctx *gin.Context) {
    var todo Todo
    if err := ctx.ShouldBindJSON(&todo); err != nil {
        ctx.JSON(http.StatusBadRequest, "Invalid input")
        return
    }

    createdTodo, err := c.service.CreateTodo(todo)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, err.Error())
        return
    }

    ctx.JSON(http.StatusCreated, createdTodo)
}

func (c *TodoController) ListTodos(ctx *gin.Context) {
    todos := c.service.ListTodos()
    ctx.JSON(http.StatusOK, todos)
}

func (c *TodoController) GetTodo(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, "Invalid ID")
        return
    }

    todo, err := c.service.GetTodoByID(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, err.Error())
        return
    }

    ctx.JSON(http.StatusOK, todo)
}

func (c *TodoController) UpdateTodo(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, "Invalid ID")
        return
    }

    var updatedTodo Todo
    if err := ctx.ShouldBindJSON(&updatedTodo); err != nil {
        ctx.JSON(http.StatusBadRequest, "Invalid input")
        return
    }

    updatedTodo.ID = id
    todo, err := c.service.UpdateTodoByID(id, updatedTodo)
    if err != nil {
        ctx.JSON(http.StatusNotFound, err.Error())
        return
    }

    ctx.JSON(http.StatusOK, todo)
}

func (c *TodoController) DeleteTodo(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, "Invalid ID")
        return
    }

    err = c.service.DeleteTodoByID(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, err.Error())
        return
    }

    ctx.Status(http.StatusNoContent)
}

// models.go
package main

type Todo struct {
    ID          int
    Description string
    Expiration  string
    Completed   bool
}

var ErrTodoNotFound = errors.New("Todo not found")

// todo_service_test.go
package main

import (
    "reflect"
    "testing"
)

func TestCreateTodo(t *testing.T) {
    service := NewTodoService()
    todo := Todo{
        Description: "Test TODO",
        Expiration:  "2023-12-31",
        Completed:   false,
    }

    createdTodo, err := service.CreateTodo(todo)
    if err != nil {
        t.Errorf("Error creating todo: %v", err)
    }

    if createdTodo.ID != 1 {
        t.Errorf("Expected todo ID 1, but got %d", createdTodo.ID)
    }

    if !reflect.DeepEqual(createdTodo, todo) {
        t.Errorf("Created todo does not match the input: %+v", createdTodo)
    }
}

func TestListTodos(t *testing.T) {
    service := NewTodoService()

    // Add some test todos
    service.CreateTodo(Todo{Description: "Test TODO 1", Expiration: "2023-12-31", Completed: false})
    service.CreateTodo(Todo{Description: "Test TODO 2", Expiration: "2024-01-15", Completed: true})

    todos := service.ListTodos()

    if len(todos) != 2 {
        t.Errorf("Expected 2 todos, but got %d", len(todos))
    }
}

func TestGetTodoByID(t *testing.T) {
    service := NewTodoService()

    // Add a test todo
    createdTodo, _ := service.CreateTodo(Todo{Description: "Test TODO", Expiration: "2023-12-31", Completed: false})

    // Get the todo by ID
    todo, err := service.GetTodoByID(createdTodo.ID)

    if err != nil {
        t.Errorf("Error getting todo by ID: %v", err)
    }

    if !reflect.DeepEqual(todo, createdTodo) {
        t.Errorf("Fetched todo does not match the created todo: %+v", todo)
    }
}

func TestUpdateTodoByID(t *testing.T) {
    service := NewTodoService()

    // Add a test todo
    createdTodo, _ := service.CreateTodo(Todo{Description: "Test TODO", Expiration: "2023-12-31", Completed: false})

    // Update the test todo
    updatedTodo := Todo{Description: "Updated TODO", Expiration: "2024-12-31", Completed: true}
    todo, err := service.UpdateTodoByID(createdTodo.ID, updatedTodo)

    if err != nil {
        t.Errorf("Error updating todo by ID: %v", err)
    }

    if !reflect.DeepEqual(todo, updatedTodo) {
        t.Errorf("Updated todo does not match the input: %+v", todo)
    }
}

func TestDeleteTodoByID(t *testing.T) {
    service := NewTodoService()

    // Add a test todo
    createdTodo, _ := service.CreateTodo(Todo{Description: "Test TODO", Expiration: "2023-12-31", Completed: false})

    // Delete the test todo
    err := service.DeleteTodoByID(createdTodo.ID)

    if err != nil {
        t.Errorf("Error deleting todo by ID: %v", err)
    }

    todos := service.ListTodos()

    if len(todos) != 0 {
        t.Errorf("Expected 0 todos after deletion, but got %d", len(todos))
    }
}

// todocontroller_test.go
package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestTodoController_CreateTodo(t *testing.T) {
	service := NewTodoService()
	controller := NewTodoController(service)

	r := setupRouter(controller)

	payload := []byte(`{"description": "Test TODO", "expiration": "2023-12-31", "completed": false}`)
	req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, but got %d", w.Code)
	}

	var response Todo
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Error unmarshaling JSON: %v", err)
	}

	if response.Description != "Test TODO" || response.Expiration != "2023-12-31" || !response.Completed {
		t.Errorf("Response does not match expectations: %+v", response)
	}
}

func TestTodoController_ListTodos(t *testing.T) {
	service := NewTodoService()
	controller := NewTodoController(service)
	r := setupRouter(controller)

	// Add some test todos
	addTestTodos(service)

	req, _ := http.NewRequest("GET", "/todos", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, but got %d", w.Code)
	}

	var response []Todo
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Error unmarshaling JSON: %v", err)
	}

	if len(response) != 2 {
		t.Errorf("Expected 2 todos, but got %d", len(response))
	}
}

func TestTodoController_GetTodo(t *testing.T) {
	service := NewTodoService()
	controller := NewTodoController(service)
	r := setupRouter(controller)

	// Add a test todo
	addTestTodos(service)

	req, _ := http.NewRequest("GET", "/todos/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, but got %d", w.Code)
	}

	var response Todo
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Error unmarshaling JSON: %v", err)
	}

	if response.ID != 1 {
		t.Errorf("Expected TODO ID 1, but got %d", response.ID)
	}
}

func TestTodoController_UpdateTodo(t *testing.T) {
	service := NewTodoService()
	controller := NewTodoController(service)
	r := setupRouter(controller)

	// Add a test todo
	addTestTodos(service)

	payload := []byte(`{"description": "Updated TODO", "expiration": "2024-12-31", "completed": true}`)
	req, _ := http.NewRequest("PUT", "/todos/1", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, but got %d", w.Code)
	}

	var response Todo
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Error unmarshaling JSON: %v", err)
	}

	if response.Description != "Updated TODO" || response.Expiration != "2024-12-31" || !response.Completed {
		t.Errorf("Response does not match expectations: %+v", response)
	}
}

func TestTodoController_DeleteTodo(t *testing.T) {
	service := NewTodoService()
	controller := NewTodoController(service)
	r := setupRouter(controller)

	// Add a test todo
	addTestTodos(service)

	req, _ := http.NewRequest("DELETE", "/todos/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("Expected status 204, but got %d", w.Code)
	}
}

func setupRouter(controller *TodoController) *gin.Engine {
	r := gin.Default()
	controller.RegisterRoutes(r)
	return r
}

func addTestTodos(service *TodoService) {
	service.CreateTodo(Todo{Description: "Test TODO 1", Expiration: "2023-12-31", Completed: false})
	service.CreateTodo(Todo{Description: "Test TODO 2", Expiration: "2024-01-15", Completed: true})
}


func addTestTodos(service *TodoService) {
    testTodos := []Todo{
        {
            Description: "Test TODO 1",xw
            Expiration:  "2023-12-31",
            Completed:   false,
        },
        {
            Description: "Test TODO 2",
            Expiration:  "2024-01-15",
            Completed:   true,
        },
    }

    for _, todo := range testTodos {
        service.CreateTodo(todo)
    }
}

func TestSomething(t *testing.T) {
    service := NewTodoService()
    addTestTodos(service)

    // Your test logic here
}
```
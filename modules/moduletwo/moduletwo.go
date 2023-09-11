package moduletwo

import (
	"fmt"

	"github.com/faultyagatha/go/modules/moduleone"
)

func main() {
    // Get a greeting message and print it.
    message := moduleone.DoSomething("Gladys")
    fmt.Println(message)
}
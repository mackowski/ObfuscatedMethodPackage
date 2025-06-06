
package reflection

import (
    "fmt"
    "reflect"
)

// InspectVariable uses reflection to inspect a variable.
func InspectVariable(x interface{}) {
    v := reflect.ValueOf(x)
    fmt.Println("Value:", v)
}

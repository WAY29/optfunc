# OPTFUNC
***Optional Parameters for go functions***

## Examples

```go
package main

import (
	"fmt"
	. "github.com/WAY29/optfunc"
)

/*
Use optfun.With for options

address string "127.0.0.1"

port string "80"
*/
func NewClient(opts ...Options) {
	o := NewOptions(OptionsParams{"address": "127.0.0.1", "port": "80"}).Apply(opts...)
	fmt.Printf("Address is %s:%s\n", o.Get("address"), o.Get("port"))
}

func main() {
	NewClient()
	NewClient(With("address", "192.168.1.1"))
	// outputs
	// Address is 127.0.0.1:80
	// Address is 192.168.1.1:80
}

```
# Go Modules

New modules can be create with the following command:
> go mod init github.com/emidev98/test

Modules can be added to the project by using the following command adding the route to that module:
> go get github.com/donvito/hellomod

A version can be defined at the end of the module in case you want to use a different version than the default:
> go get github.com/donvito/hellomod/v2

There is also a posibility of having both modules coexisting in the same program by defining an alias to one of them:

```go
package main

import (
	"github.com/donvito/hellomod"
	hellomod2 "github.com/donvito/hellomod/v2" // alias for the second version of the program
)

func main() {
	hellomod.SayHello()
	hellomod2.SayHello("emidev98")
}
```

When a module is not used in the project you can automatically delete that module by executing the following command:
> go mod tidy

To view the list of the modules that are stored locally:
> go mod download -json


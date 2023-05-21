[![integration](https://github.com/toel-app/biteship/actions/workflows/e2e.yml/badge.svg)](https://github.com/toel-app/biteship/actions/workflows/e2e.yml)
[![unit-tests](https://github.com/toel-app/biteship/actions/workflows/go.yml/badge.svg)](https://github.com/toel-app/biteship/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/toel-app/biteship)](https://goreportcard.com/report/github.com/toel-app/biteship)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/stretchr/testify)](https://pkg.go.dev/github.com/toel-app/biteship)




### Installation
```
go get -u github.com/toel-app/biteship-go
```

### Usage
```go
package main

import (
	"fmt"
	"github.com/toel-app/biteship"
	"log"
)

func main() {
	b := biteship.New(
		biteship.WithSecret("somesecret"),
		biteship.WithUrl("https://api.biteship.com"),
	)

	response, err := b.GetCouriers()
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(response.Couriers)
}
```

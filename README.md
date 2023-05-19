[![integration](https://github.com/toel-app/biteship/actions/workflows/e2e.yml/badge.svg)](https://github.com/toel-app/biteship/actions/workflows/e2e.yml)
[![unit-tests](https://github.com/toel-app/biteship/actions/workflows/go.yml/badge.svg)](https://github.com/toel-app/biteship/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/toel-app/biteship)](https://goreportcard.com/report/github.com/toel-app/biteship)
## Unstable

[Official Docs](https://biteship.com/en/docs/intro)

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
	secretKey := "abcdefg"
	biteshipApp := biteship.New(secretKey)

	resp, err := biteshipApp.GetCourier()
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
```

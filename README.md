# nulltype

[![License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://github.com/r-fujiyama/nulltype/blob/master/LICENSE)
[![CI](https://github.com/r-fujiyama/nulltype/workflows/CI/badge.svg)](https://github.com/r-fujiyama/nulltype/actions?query=workflow%3ACI)
[![Codecov](https://img.shields.io/codecov/c/github/r-fujiyama/nulltype/.svg?style=flat)](https://codecov.io/gh/r-fujiyama/nulltype/)
[![Go Report Card](https://goreportcard.com/badge/github.com/r-fujiyama/nulltype?style=flat)](https://goreportcard.com/report/github.com/r-fujiyama/nulltype)

## Guide

## Installation

```sh
$ go get -u github.com/r-fujiyama/nulltype
```

### Usage

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/r-fujiyama/nulltype"
)

func main() {

	person := struct {
		Name null.String `json:"name"`
	}{
		Name: null.String{String: "John Smith", Valid: true},
	}
	value, _ := json.Marshal(person)
	fmt.Println(string(value)) // {"name":"John Smith"}

	person = struct {
		Name null.String `json:"name"`
	}{
		Name: null.String{String: "", Valid: false},
	}
	value, _ = json.Marshal(person)
	fmt.Println(string(value)) // {"name":null}
}
```

## License

[MIT](https://github.com/r-fujiyama/nulltype/blob/master/LICENSE)

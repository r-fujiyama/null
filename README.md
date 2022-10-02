# null

[![License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://github.com/r-fujiyama/null/blob/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/r-fujiyama/null.svg)](https://pkg.go.dev/github.com/r-fujiyama/null)
[![CI](https://github.com/r-fujiyama/null/workflows/CI/badge.svg)](https://github.com/r-fujiyama/null/actions?query=workflow%3ACI)
[![Codecov](https://img.shields.io/codecov/c/github/r-fujiyama/null/.svg?style=flat)](https://codecov.io/gh/r-fujiyama/null/)
[![Go Report Card](https://goreportcard.com/badge/github.com/r-fujiyama/null)](https://goreportcard.com/report/github.com/r-fujiyama/null)

## Installation

```sh
go get github.com/r-fujiyama/null
```

## Example

```go
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/r-fujiyama/null"
)

func main() {
	// JSON Encode
	s1 := null.String{String: "foo", Valid: true}
	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(s1)
	fmt.Print(buf.String()) // "foo"

	buf.Reset()
	s1 = null.String{String: "", Valid: true}
	_ = json.NewEncoder(&buf).Encode(s1)
	fmt.Print(buf.String()) // ""

	buf.Reset()
	s1 = null.String{String: "", Valid: false}
	_ = json.NewEncoder(&buf).Encode(s1)
	fmt.Print(buf.String()) // null

	// JSON Decode
	s2 := null.String{}
	_ = json.NewDecoder(strings.NewReader(`"foo"`)).Decode(&s2)
	fmt.Printf("%#v\n", s2) // null.String{String:"foo", Valid:true}

	_ = json.NewDecoder(strings.NewReader(`""`)).Decode(&s2)
	fmt.Printf("%#v\n", s2) // null.String{String:"", Valid:true}

	_ = json.NewDecoder(strings.NewReader("null")).Decode(&s2)
	fmt.Printf("%#v\n", s2) // null.String{String:"", Valid:false}
}
```

## License

[MIT](https://github.com/r-fujiyama/null/blob/master/LICENSE)

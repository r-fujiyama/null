# nulltype

[![License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://github.com/r-fujiyama/nulltype/blob/master/LICENSE)
[![CI](https://github.com/r-fujiyama/nulltype/workflows/CI/badge.svg)](https://github.com/r-fujiyama/nulltype/actions?query=workflow%3ACI)
[![Codecov](https://img.shields.io/codecov/c/github/r-fujiyama/nulltype/.svg?style=flat)](https://codecov.io/gh/r-fujiyama/nulltype/)
[![Go Report Card](https://goreportcard.com/badge/github.com/r-fujiyama/nulltype)](https://goreportcard.com/report/github.com/r-fujiyama/nulltype)

## Guide

## Installation

```sh
$ go get -u github.com/r-fujiyama/nulltype
```

### Example

```go
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/r-fujiyama/nulltype"
)

func main() {
	// JSON Encode
	s1 := nulltype.String{String: "foo", Valid: true}
	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(s1)
	fmt.Print(buf.String()) // "foo"

	buf.Reset()
	s1 = nulltype.String{String: "", Valid: true}
	_ = json.NewEncoder(&buf).Encode(s1)
	fmt.Print(buf.String()) // ""

	buf.Reset()
	s1 = nulltype.String{String: "", Valid: false}
	_ = json.NewEncoder(&buf).Encode(s1)
	fmt.Print(buf.String()) // null

	// JSON Decode
	s2 := nulltype.String{}
	_ = json.NewDecoder(strings.NewReader(`"foo"`)).Decode(&s2)
	fmt.Printf("%#v\n", s2) // nulltype.String{String:"foo", Valid:true}

	_ = json.NewDecoder(strings.NewReader(`""`)).Decode(&s2)
	fmt.Printf("%#v\n", s2) // nulltype.String{String:"", Valid:true}

	_ = json.NewDecoder(strings.NewReader("null")).Decode(&s2)
	fmt.Printf("%#v\n", s2) // nulltype.String{String:"", Valid:false}
}
```

## License

[MIT](https://github.com/r-fujiyama/nulltype/blob/master/LICENSE)

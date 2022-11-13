updater - check update for cli tools you created
=======
[![](https://godoc.org/github.com/Code-Hex/updater?status.svg)](http://godoc.org/github.com/Code-Hex/updater)
[![Go Report Card](https://goreportcard.com/badge/github.com/Code-Hex/updater)](https://goreportcard.com/report/github.com/Code-Hex/updater)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
## Installation

    $ go get github.com/Code-Hex/updater

## Synopsis

    updater.Check...(Your organization, Your project name, now version...)

## Example
CheckWithTag using release tag
```Go
import "github.com/Code-Hex/updater"

func main() {
	result, err := updater.CheckWithTag("Code-Hex", "pget", "0.0.1")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(result)
}
```

CheckWithTitle using release title
```Go
result, err := updater.CheckWithTitle("Code-Hex", "pget", "0.0.1")
```

Check using key of json from github api
```Go
result, err := updater.Check("Code-Hex", "pget", "0.0.1", "keyname")
```
## Author

[codehex](https://twitter.com/CodeHex)

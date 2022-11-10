package api

import (
	"fmt"
	"strings"
)

func GetDocVersion(ver string) string {
	for k, v := range versions {
		if k == ver {
			return v
		}
	}
	return "unknown version"
}

func GetFullDoc() string {
	var b strings.Builder
	for _, v := range versions {
		fmt.Fprintf(&b, "%s\n", v)
	}

	return b.String()
}

var versions = map[string]string{
	"1.19": go119,
	"1.18": go118,
}

const (
	go119 = `
Go 1.19 is a major release of Go. Read the Go 1.19 Release Notes (https://go.dev/doc/go1.19) for more information.

Minor revisions:
	- go1.19.1 (released 2022-09-06) includes security fixes to the net/http and net/url packages, 
	as well as bug fixes to the compiler, the go command, the pprof command, 
	the linker, the runtime, and the crypto/tls and crypto/x509 packages. 
	See the Go 1.19.1 milestone (https://github.com/golang/go/issues?q=milestone%3AGo1.19.1+label%3ACherryPickApproved) on our issue tracker for details.

	- go1.19.2 (released 2022-10-04) includes security fixes to the archive/tar, 
	net/http/httputil, and regexp packages, as well as bug fixes to the compiler, the linker, 
	the runtime, and the go/types package. See the Go 1.19.2 milestone (https://github.com/golang/go/issues?q=milestone%3AGo1.19.2+label%3ACherryPickApproved) on our issue tracker for details.
	
	- go1.19.3 (released 2022-11-01) includes security fixes to the os/exec and syscall packages, 
	as well as bug fixes to the compiler and the runtime. 
	See the Go 1.19.3 milestone (https://github.com/golang/go/issues?q=milestone%3AGo1.19.3+label%3ACherryPickApproved) on our issue tracker for details.
`

	go118 = `
Go 1.18 is a major release of Go. Read the Go 1.18 Release Notes for more information.

Minor revisions:
	go1.18.1 (released 2022-04-12) includes security fixes to the crypto/elliptic,
	crypto/x509, and encoding/pem packages, as well as bug fixes to the compiler, linker, 
	runtime, the go command, vet, and the bytes, crypto/x509, and go/types packages. 
	See the Go 1.18.1 milestone on our issue tracker for details.
	
	go1.18.2 (released 2022-05-10) includes security fixes to the syscall package, 
	as well as bug fixes to the compiler, runtime, the go command, 
	and the crypto/x509, go/types, net/http/httptest, reflect, and sync/atomic packages. 
	See the Go 1.18.2 milestone on our issue tracker for details.
	
	go1.18.3 (released 2022-06-01) includes security fixes to the crypto/rand, 
	crypto/tls, os/exec, and path/filepath packages, as well as bug fixes to the compiler, 
	and the crypto/tls and text/template/parse packages. 
	See the Go 1.18.3 milestone on our issue tracker for details.
	
	go1.18.4 (released 2022-07-12) includes security fixes to the compress/gzip, 
	encoding/gob, encoding/xml, go/parser, io/fs, net/http, and path/filepath packages, 
	as well as bug fixes to the compiler, the go command, the linker, the runtime, 
	and the runtime/metrics package. 
	See the Go 1.18.4 milestone on our issue tracker for details.
	
	go1.18.5 (released 2022-08-01) includes security fixes to the encoding/gob and math/big packages, 
	as well as bug fixes to the compiler, the go command, the runtime,and the testing package. 
	See the Go 1.18.5 milestone on our issue tracker for details.
	
	go1.18.6 (released 2022-09-06) includes security fixes to the net/http package, 
	as well as bug fixes to the compiler, the go command, the pprof command, the runtime, 
	and the crypto/tls, encoding/xml, and net packages. 
	See the Go 1.18.6 milestone on our issue tracker for details.
	
	go1.18.7 (released 2022-10-04) includes security fixes to the archive/tar, 
	net/http/httputil, and regexp packages, as well as bug fixes to the compiler, the linker, 
	and the go/types package. See the Go 1.18.7 milestone on our issue tracker for details.
	
	go1.18.8 (released 2022-11-01) includes security fixes to the os/exec and syscall packages, 
	as well as bug fixes to the runtime. 
	See the Go 1.18.8 milestone on our issue tracker for details.
`
)

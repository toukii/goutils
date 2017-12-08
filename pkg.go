package goutils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Pkg struct {
	Dir         string   `json:"Dir"`
	ImportPath  string   `json:"ImportPath"`
	Name        string   `json:"Name"`
	Target      string   `json:"Target"`
	Stale       string   `json:"Stale"`
	StaleReason string   `json:"StaleReason"`
	Root        string   `json:"Root"`
	GoFiles     []string `json:"GoFiles"`
	Imports     []string `json:"Imports"`
	Deps        []string `json:"Deps"`
	Incomplete  bool     `json:"Incomplete"`
	DepsErrors  []struct {
		ImportStack []string `json:"ImportStack"`
		Pos         string   `json:"Pos"`
		Err         string   `json:"Err"`
	} `json:"DepsErrors"`
	TestGoFiles []string `json:"TestGoFiles"`
	TestImports []string `json:"TestImports"`

	NoStdDepOKPkgs    []string
	NoStdDepErrPkgs   []NoStdDepErrPkg
	nostdDepErrPkgMap map[string]bool
}

type NoStdDepErrPkg struct {
	PkgName string
	Tips    string
}

func NewPkg(bs []byte) (*Pkg, error) {
	var pkg Pkg
	err := json.Unmarshal(bs, &pkg)
	if CheckErr(err) {
		return nil, err
	}
	pkg.analyse()
	return &pkg, err
}

func (p *Pkg) analyse() {
	p.analyseDepsErrors()
	p.NoStdDepOKPkgs = make([]string, 0, len(p.Deps)-len(p.DepsErrors))
	for _, dep := range p.Deps {
		if p.nostdDepErrPkgMap[dep] {
			continue
		}
		p.NoStdDepOKPkgs = append(p.NoStdDepOKPkgs, dep)
	}
}

func (p *Pkg) analyseDepsErrors() []*NoStdDepErrPkg {
	desize := len(p.DepsErrors)
	if desize <= 0 {
		return nil
	}
	p.nostdDepErrPkgMap = make(map[string]bool)
	ret := make([]*NoStdDepErrPkg, 0, desize)
	for _, depsErr := range p.DepsErrors {
		tips := depsErr.Err
		for _, imp := range depsErr.ImportStack {
			if imp == p.ImportPath {
				continue
			}
			p.nostdDepErrPkgMap[imp] = true
			ret = append(ret, &NoStdDepErrPkg{
				PkgName: imp,
				Tips:    tips,
			})
		}
	}
	return ret
}

type PkgNames []interface{}

func (p *PkgNames) DisplayString() string {
	ret := ""
	for i, name := range p {
		ret += fmt.Sprintf("\t%d. %s", i, name)
	}
	return ret
}

func TrimGopath(pkg string) string {
	abspath, err := filepath.Abs(os.Getenv("GOPATH"))
	if CheckErr(err) {
		return pkg
	}
	return strings.TrimPrefix(pkg, abspath+"/src/")
}

/*{
	"Dir": "/Users/toukii/PATH/GOPATH/src/gitlab.1dmy.com/ezbuy/goflow/src/github.com/toukii/bkg",
	"ImportPath": "github.com/toukii/bkg",
	"Name": "main",
	"Target": "/Users/toukii/PATH/GOPATH/src/gitlab.1dmy.com/ezbuy/goflow/bin/bkg",
	"Stale": true,
	"StaleReason": "build ID mismatch",
	"Root": "/Users/toukii/PATH/GOPATH/src/gitlab.1dmy.com/ezbuy/goflow",
	"GoFiles": [
		"main.go"
	],
	"Imports": [
		"time"
	],
	"Deps": [
		"bufio",
		"unicode",
		"unicode/utf16",
		"unicode/utf8",
		"unsafe",
		"vendor/golang_org/x/crypto/chacha20poly1305",
		"vendor/golang_org/x/crypto/chacha20poly1305/internal/chacha20",
		"vendor/golang_org/x/crypto/curve25519",
		"vendor/golang_org/x/crypto/poly1305",
		"vendor/golang_org/x/net/http2/hpack",
		"vendor/golang_org/x/text/unicode/norm"
	],
	"Incomplete": true,
	"DepsErrors": [
		{
			"ImportStack": [
				"github.com/toukii/bkg",
				"github.com/everfore/exc"
			],
			"Pos": "main.go:11:2",
			"Err": "cannot find package \"github.com/everfore/exc\" in any of:\n\t/Users/toukii/PATH/Go/go/src/github.com/everfore/exc (from $GOROOT)\n\t/Users/toukii/PATH/GOPATH/src/gitlab.1dmy.com/ezbuy/goflow/src/github.com/everfore/exc (from $GOPATH)"
		}
	],
	"TestGoFiles": [
		"jsnm_test.go",
		"zhihu_test.go"
	],
	"TestImports": [
		"encoding/json",
		"fmt",
		"github.com/toukii/goutils",
		"github.com/toukii/membership/pkg3/go-simplejson",
		"io/ioutil",
		"net/http",
		"reflect",
		"testing"
	]
}
*/

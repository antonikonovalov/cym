package test_test

import (
	"testing"
	"os"
	"io/ioutil"
	"github.com/antonikonovalov/cym/page"
	"github.com/antonikonovalov/cym/parser"
)

var (
	gopath string = os.Getenv("GOPATH")
	fixturesPath string = gopath+"/src/github.com/antonikonovalov/cym/fixtures"
)

func getMainPage() *page.Page {
	content,err := ioutil.ReadFile(fixturesPath+`/index.html`)
	if err != nil {
		panic(err)
	}
	return &page.Page{page.PageTypeCatalogs,string(content)}
}

func TestParserMainPage(t *testing.T) {
	p := getMainPage()
	pageParser,err := parser.New(*p)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	err = pageParser.Do()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

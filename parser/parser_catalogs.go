package parser
import (
	"github.com/antonikonovalov/cym/page"
	"gopkg.in/xmlpath.v2"
)


type pageCommonParser struct {
	node *xmlpath.Node
}


func newPageCatalogsParser(p page.Page) (Parser,error) {
	node,err := xmlpath.ParseHTML(p.Reader())
	if err != nil {
		return nil,err
	}

	return Parser(&pageCommonParser{node}),nil
}



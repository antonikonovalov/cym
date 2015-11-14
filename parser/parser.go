package parser
import "github.com/antonikonovalov/cym/page"


func New(p page.Page) (Parser,error) {
	err := p.Validate()
	if err != nil {
		return nil,err
	}

	var pageParser Parser
	switch p.Type {
	case page.PageTypeCatalogs:
		pageParser,err = newPageCatalogsParser(p)
	}
	return pageParser,err
}

type Parser interface {
	Do() error
}





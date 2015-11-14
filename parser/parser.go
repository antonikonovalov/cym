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
		pageParser = newPageCatalogsParser(p)
	}
	return pageParser,nil
}

type Parser interface {
	Do() error
}





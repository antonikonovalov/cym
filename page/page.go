package page
import (
	"errors"
	"io"
	"strings"
)

type PageType int

const (
	_ PageType = iota
	PageTypeCatalogs
	PageTypeCatalog
	PageTypeProducers
	PageTypeItems
	PageTypeProvider
)

func (t PageType) Supported () error {
	if int(t) > 5 || int(t) < 1 {
		return errors.New("not supported type of page")
	}
	return nil
}

type Page struct {
	Type    PageType
	Content string
}

func (p Page) Reader() io.Reader {
	return strings.NewReader(p.Content)
}

func (p Page) Validate() error {
	if len(p.Content) == 0 {
		return errors.New("int valid Content")
	}
	err := p.Type.Supported()
	if err != nil {
		return err
	}

	return nil
}


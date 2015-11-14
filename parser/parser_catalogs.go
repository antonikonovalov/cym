package parser
import (
	"github.com/antonikonovalov/cym/page"
	"log"
	"github.com/PuerkitoBio/goquery"
"errors"
	"net/url"
)


type pageCommonParser struct {
	root *goquery.Document
}

type pageCatalogsParser struct {
	*pageCommonParser
	pathToCatalogs,pathToCatalogMainName string
}

func (p *pageCatalogsParser) initPaths() {
	p.pathToCatalogs = `/html/body/div[1]/div[3]/div[2]/div[2]/div/div`
	p.pathToCatalogMainName = `/a[href]`
}


//func (p *pageCatalogsParser) getCatalogsElements() *xmlpath.Iter {
//	catalogsPath := xmlpath.MustCompile(p.pathToCatalogs)
//	if catalogsPath.Exists(p.root) {
//		return catalogsPath.Iter(p.root)
//	}
//	return nil
//}

func (p *pageCatalogsParser) Do() error {
	catalogs := p.root.Find(`div.catalog-simple__item > a`)
	//skip first - it sales info
	if catalogs == nil {
		return errors.New("Can't find catalogs")
	}
	log.Println("catalogs,",catalogs)
	for _,node := range catalogs.Nodes {
		for i := range node.Attr {
			if node.Attr[i].Key == `href` {
				u,err := url.Parse(node.Attr[i].Val)
				if err != nil {
					log.Println(err)
				}
				hid := u.Query().Get(`hid`)
				if hid != "" {
					log.Println(node.Attr[i].Val)
				}

			}
		}
	}
	return nil
}

func newPageCatalogsParser(p page.Page) (Parser,error) {
//	root,err := xmlpath.ParseHTML(p.Reader())
//	if err != nil {
//		return nil,err
//	}
	root,err := goquery.NewDocumentFromReader(p.Reader())
	if err != nil {
		return nil,err
	}
	common := &pageCommonParser{root}
	pr := &pageCatalogsParser{common,"",""}
	pr.initPaths()
	return Parser(pr),nil
}



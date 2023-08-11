package global

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"math"
	"net/url"
)

type PageUrl struct {
	Page int
	Url  string
}

type Pager struct {
	Offset     int
	Size       int
	Current    int
	Total      int64
	ReqPath    string
	ReqQuery   url.Values
	PrevUrl    string
	NextUrl    string
	FirstUrl   string
	EndUrl     string
	CurrentUrl string
	PagesUrl   []PageUrl
	TotalPage  int
	ctx        iris.Context
}

const DefaultPageSize = 20

func NewPager(ctx iris.Context) *Pager {
	p := &Pager{}
	p.Current = ctx.URLParamIntDefault("Page", 1)
	p.Size = ctx.URLParamIntDefault("Size", DefaultPageSize)
	if p.Current <= 0 {
		p.Current = 1
	}
	p.ctx = ctx
	req := ctx.Request()
	p.ReqPath = req.URL.Path
	p.ReqQuery = req.URL.Query()
	p.CurrentUrl = req.URL.String()
	p.PrevUrl = ""
	p.NextUrl = ""
	if p.Size == 0 {
		p.Size = DefaultPageSize
	}
	p.Offset = (p.Current - 1) * p.Size
	return p
}

func (p *Pager) SetTotal(total int64) {
	p.Total = total
	if p.Total == 0 {
		return
	}
	p.TotalPage = int(math.Ceil(float64(p.Total) / float64(p.Size)))
	pageUrlNum := 10
	p.PagesUrl = make([]PageUrl, 0, pageUrlNum)
	req := p.ctx.Request()
	var pagesUrl PageUrl
	var n = p.Current % pageUrlNum
	for k := 1; k <= pageUrlNum; k++ {
		if n == 0 {
			pagesUrl = PageUrl{Page: p.Current - pageUrlNum + k}
		} else if k < n {
			pagesUrl = PageUrl{Page: p.Current - n + k}
		} else {
			pagesUrl = PageUrl{Page: p.Current + (k - n)}
		}
		if pagesUrl.Page > p.TotalPage {
			break
		}
		query := req.URL.Query()
		query.Set("Page", fmt.Sprintf("%d", pagesUrl.Page))
		pagesUrl.Url = p.ReqPath + "?" + query.Encode()
		p.PagesUrl = append(p.PagesUrl, pagesUrl)
	}
	query := req.URL.Query()
	if p.Current > 1 {
		query.Set("Page", fmt.Sprintf("%d", p.Current-1))
		p.PrevUrl = p.ReqPath + "?" + query.Encode()
	}
	if p.Current < p.TotalPage {
		query.Set("Page", fmt.Sprintf("%d", p.Current+1))
		p.NextUrl = p.ReqPath + "?" + query.Encode()
	}
	query.Set("Page", fmt.Sprintf("%d", 1))
	p.FirstUrl = p.ReqPath + "?" + query.Encode()

	query.Set("Page", fmt.Sprintf("%d", p.TotalPage))
	p.EndUrl = p.ReqPath + "?" + query.Encode()
}

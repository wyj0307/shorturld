package handler

import (
	"strings"

	errs "crazyant.com/groot/pkg/errors"
	"golang.org/x/net/context"

	"shorturld/shorten"

	"crazyant.com/groot/pbd/go/shorturld"
)

type Shorturl struct{}

func (hand *Shorturl) ConvertUrlShort(ctx context.Context, req *shorturld.UrlReq, rsp *shorturld.UrlResp) error {
	url := strings.TrimSpace(req.LongUrl)
	provider := strings.TrimSpace(req.Provider)
	if len(url) == 0 {
		return errs.BadThroughRequest("url为空")
	}

	if len(provider) == 0 {
		provider = "gggg"
	}
	c := shorten.NewClient(provider)
	if c == nil {
		return errs.InternalServerError("申请 shorten client 失败")
	}

	u, err := c.Shorten(url)
	if err == nil {
		rsp.DesUrl = string(u)
	}

	return err
}

func (hand *Shorturl) ConvertUrlOriginal(ctx context.Context, req *shorturld.ShorturlReq, rsp *shorturld.UrlResp) error {
	url := strings.TrimSpace(req.Shorturl)
	if len(url) == 0 {
		return errs.BadThroughRequest("短网址url为空")
	}

	u, err := shorten.Expanden(url)
	if err == nil {
		rsp.DesUrl = string(u)
	}
	return err
}

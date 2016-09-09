package shorten

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"strings"

	"github.com/subosito/shorturl"
)

type Shorten struct {
	Provider string
}

func NewClient(provider string) *Shorten {
	return &Shorten{Provider: provider}
}

func (c *Shorten) Shorten(url string) ([]byte, error) {
	if !strings.HasPrefix(url, "http") {
		return nil, errors.New("原始url网址缺少http头")
	} else if err := urlRegular(url); err != nil {
		return nil, errors.New(fmt.Sprintf("原始url不合理,err:%v", err))
	}

	u, err := shorturl.Shorten(url, c.Provider)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("用[%s]方式进行短网址转换时出错,err:%v", c.Provider, err))
	}
	return u, nil
}

func Expanden(sUrl string) ([]byte, error) {
	if !strings.HasPrefix(sUrl, "http") {
		return nil, errors.New("原始url网址缺少http头")
	} else if err := urlRegular(sUrl); err != nil {
		return nil, errors.New(fmt.Sprintf("原始url不合理,err:%v", err))
	}

	return shorturl.Expand(sUrl)
}

func urlRegular(oriUrl string) error {
	//判断url是否合理
	host, err := url.ParseRequestURI(oriUrl)
	if err != nil {
		return err
	}

	//判断是否能解析到对应的host记录
	_, err = net.LookupIP(host.Host)
	if err != nil {
		return err
	}
	return nil
}

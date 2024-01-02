package generators

import (
	"fmt"
	"io"
	"net"
	"net/url"
	"strings"

	"github.com/twoFiveOneTen/gena"
)

// Generator is interface of template generator
type Generator interface {
	Run(cfg *gena.Config, writer io.Writer)
}

// favicon return favicon of url
func favicon(rawurl string) string {
	base := "https://navi.zkk.me/icon/%s.png"
	rawurl = strings.TrimSpace(rawurl)
	u, err := url.Parse(rawurl)
	if err != nil {
		return fmt.Sprintf(base, "o.oo")
	}
	host := u.Host
	if strings.Contains(host, ":") {
		host, _, _ = net.SplitHostPort(host)
	}
	return fmt.Sprintf(base, host)
}

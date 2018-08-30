package actions

import (
	"fmt"
	"net"
	"strings"

	"github.com/gobuffalo/buffalo"
)

var UserAgents = [4]string{"WindowsPowerShell", "curl", "fetch", "Wget"}

// RootHandler is a default handler to serve up
// a home page.
func RootHandler(c buffalo.Context) error {

	for _, v := range UserAgents {
		if strings.Contains(c.Request().UserAgent(), v) {
			ip, _, err := net.SplitHostPort(c.Request().RemoteAddr)
			if err != nil {
				fmt.Print("err")
			}
			return c.Render(200, r.String(ip))
		}
	}
	return c.Render(200, r.JSON(map[string]string{"message": fmt.Sprintf("Welcome to Buffalo! %s", c.Request().UserAgent())}))
}

func AllHeadersHandler(c buffalo.Context) error {
	// Create return string
	var request []string
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", c.Request().Host))
	// Loop through headers
	for name, headers := range c.Request().Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}
	return c.Render(200, r.String(strings.Join(request, "\n")))
}

func HeaderHandler(c buffalo.Context) error {
	var request string
	for name, headers := range c.Request().Header {
		name = strings.ToLower(name)
		reqHeader := strings.ToLower(c.Param("header"))
		if name == reqHeader {
			for _, h := range headers {
				request = fmt.Sprintf("%s: %s", name, h)
			}
		}
	}
	if len(request) == 0 {
		return c.Render(200, r.String(""))
	}
	return c.Render(200, r.String(request))
}

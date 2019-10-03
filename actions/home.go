package actions

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/gobuffalo/buffalo"
)

var userAgents = [4]string{"WindowsPowerShell", "curl", "fetch", "Wget"}

// RootHandler is a default handler to serve up
// a home page.
func RootHandler(c buffalo.Context) error {

	// Eventually use some sort of middlware to determine CLI or Browser
	for _, v := range userAgents {
		if strings.Contains(c.Request().UserAgent(), v) {
			ip, _, err := net.SplitHostPort(c.Request().RemoteAddr)
			if err != nil {
				fmt.Print("err")
			}
			return c.Render(200, r.String(ip))
		}
	}
	c.Set("title", TITLE)
	c.Set("hostname", HOSTNAME)
	c.Set("headers", GetAllHeaders(c.Request()))
	return c.Render(200, r.HTML("index.plush.html"))
}

// AllHeadersHandler will return all the headers to the client
func AllHeadersHandler(c buffalo.Context) error {
	// Create return string
	request := GetAllHeaders(c.Request())
	var headerList []string
	for name, header := range request {
		headerList = append(headerList, fmt.Sprintf("%v: %v", name, header))
	}
	return c.Render(200, r.String(strings.Join(headerList, "\n")))
}

// HeaderHandler returns requested header
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

// GetAllHeaders gathers all the headers
func GetAllHeaders(r *http.Request) map[string]string {
	var request = make(map[string]string)
	// Add the host
	request["Host"] = r.Host
	// Loop through headers
	for name, headers := range r.Header {
		fmt.Print(name)
		name = strings.ToLower(name)
		for _, h := range headers {
			request[name] = h
		}
	}
	return request
}

// Copyright (c) 2015 Jeevanandam M (jeeva@myjeeva.com), All rights reserved.
// resty source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package resty

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"golang.org/x/net/publicsuffix"
)

// The default Client
var DefaultClient *Client

func init() {
	DefaultClient = New()
}

// New creates a new go-resty client
// 		For Example:
// 		client1 := resty.New()
//
//		client2 := resty.New()
func New() *Client {
	cookieJar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})

	c := &Client{
		HostURL:    "",
		QueryParam: url.Values{},
		FormData:   url.Values{},
		Header:     http.Header{},
		UserInfo:   nil,
		Token:      "",
		Cookies:    make([]*http.Cookie, 0),
		Debug:      false,
		Log:        getLogger(nil),
		httpClient: &http.Client{CheckRedirect: NoRedirectPolicy, Jar: cookieJar},
		transport:  &http.Transport{},
	}

	// default before request middlewares
	c.beforeRequest = []func(*Client, *Request) error{
		parseRequestURL,
		parseRequestHeader,
		parseRequestBody,
		createHTTPRequest,
		addCredentials,
		requestLogger,
	}

	// default after response middlewares
	c.afterResponse = []func(*Client, *Response) error{
		responseLogger,
		parseResponseBody,
	}

	return c
}

// R creates a new resty request object, it is used form a HTTP request
// such as GET, POST, PUT, DELETE, HEAD, PATCH and OPTIONS.
func R() *Request {
	return DefaultClient.R()
}
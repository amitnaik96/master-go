package main

import (
	"fmt"
	"net"
	"net/url"
)

func parseUrl() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, _ := url.Parse(s)

	fmt.Println(u.Scheme)

	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	fmt.Println(u.User.Password())

	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	fmt.Println(u.RawQuery)
}

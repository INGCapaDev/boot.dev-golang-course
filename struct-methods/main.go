package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (r authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", r.username, r.password)
}

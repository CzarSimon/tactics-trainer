package role

import "github.com/CzarSimon/tactics-trainer/gopkg/auth/scope"

const Anonymous = "ANONYMOUS"

var anonymousRole = Role{
	Name:   Anonymous,
	scopes: map[scope.Scope]bool{},
}

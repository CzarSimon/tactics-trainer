package role

import (
	"github.com/CzarSimon/tactics-trainer/gopkg/auth/scope"
)

// Role represents a users role and permissions in the system.
type Role struct {
	Name   string
	scopes map[scope.Scope]bool
}

// HasScope checks if a role has a specific scope
func (r Role) HasScope(scope scope.Scope) bool {
	_, ok := r.scopes[scope]
	return ok
}

func Get(name string) (Role, bool) {
	switch name {
	case User:
		return userRole, true
	case Anonymous:
		return anonymousRole, true
	default:
		return Role{}, false
	}
}

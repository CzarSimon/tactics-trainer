package role

import "github.com/CzarSimon/tactics-trainer/gopkg/auth/scope"

const User = "USER"

var userRole = Role{
	Name: User,
	scopes: map[scope.Scope]bool{
		scope.ListProblemSets:  true,
		scope.ReadProblemSet:   true,
		scope.CreateProblemSet: true,
		scope.UpdateProblemSet: true,
		scope.DeleteProblemSet: true,
	},
}

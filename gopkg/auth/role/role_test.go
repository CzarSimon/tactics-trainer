package role_test

import (
	"testing"

	"github.com/CzarSimon/tactics-trainer/gopkg/auth/role"
	"github.com/CzarSimon/tactics-trainer/gopkg/auth/scope"
	"github.com/stretchr/testify/assert"
)

func TestGetRole(t *testing.T) {
	assert := assert.New(t)

	_, ok := role.Get(role.User)
	assert.True(ok)

	_, ok = role.Get(role.Anonymous)
	assert.True(ok)

	_, ok = role.Get("missing")
	assert.False(ok)
}

func TestUserRole(t *testing.T) {
	expectedScopes := []scope.Scope{
		scope.CreateProblemSet,
		scope.DeleteProblemSet,
		scope.ListProblemSets,
		scope.ReadProblemSet,
		scope.UpdateProblemSet,
		scope.ListProblemSetCycles,
		scope.CreateProblemSetCycle,
		scope.ReadCycle,
		scope.UpdateCycle,
	}

	notExpectedScopes := []scope.Scope{}

	testRole(t, role.User, expectedScopes, notExpectedScopes)
}

func TestAnonymousRole(t *testing.T) {
	expectedScopes := []scope.Scope{}

	notExpectedScopes := []scope.Scope{
		scope.CreateProblemSet,
		scope.DeleteProblemSet,
		scope.ListProblemSets,
		scope.ReadProblemSet,
		scope.UpdateProblemSet,
		scope.ListProblemSetCycles,
		scope.CreateProblemSetCycle,
		scope.ReadCycle,
		scope.UpdateCycle,
	}

	testRole(t, role.Anonymous, expectedScopes, notExpectedScopes)
}

func testRole(t *testing.T, roleName string, expectedScopes []scope.Scope, notExpectedScopes []scope.Scope) {
	assert := assert.New(t)

	r, ok := role.Get(roleName)
	assert.True(ok)

	for _, es := range expectedScopes {
		assert.True(r.HasScope(es), "%s did not have scope: %s", r.Name, es)
	}

	for _, nes := range notExpectedScopes {
		assert.False(r.HasScope(nes), "%s did have scope: %s but should not have", r.Name, nes)
	}
}

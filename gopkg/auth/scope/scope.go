package scope

// Scope defines a permission attached to a role an required to perform an operation.
type Scope string

// Problem set scopes
const (
	ListProblemSets  Scope = "problem-sets:list"
	ReadProblemSet   Scope = "problem-sets:read:one"
	CreateProblemSet Scope = "problem-sets:create:one"
	UpdateProblemSet Scope = "problem-sets:update:one"
	DeleteProblemSet Scope = "problem-sets:delete:one"
)

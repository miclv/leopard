package object

// NewEnclosedEnvironment creates a new environment with an outer environment for variable scoping.
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// NewEnvironment creates a new, empty environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// Environment represents a scope for storing variables, with optional outer scope support.
type Environment struct {
	store map[string]Object
	outer *Environment
}

// Get retrieves the value of a variable by its name, checking outer environments if needed.
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set assigns a value to a variable in the current environment.
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

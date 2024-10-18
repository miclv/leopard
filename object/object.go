/*
Package object defines the data structures and interfaces for the objects in the Leopard programming language.

This package includes representations for various object types, such as integers, booleans, strings, functions, arrays,
and hashes. Each object type implements the Object interface, which provides methods to retrieve the object's type
and string representation.
*/
package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"leopard/ast"
	"strings"
)

// ObjectType represents the type of an object in the language.
type ObjectType string

// Supported object types
const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
	HASH_OBJ         = "HASH"
)

// Object is an interface for all objects in the language.
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer represents an integer value.
type Integer struct {
	Value int64
}

// Type and Inspect methods for Integer.
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

// Boolean represents a boolean value
type Boolean struct {
	Value bool
}

// Type and Inspect methods for Boolean.
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

// Null represents a null value.
type Null struct{}

// Type and Inspect methods for Null.
func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

// ReturnValue wraps a value returned from a function.
type ReturnValue struct {
	Value Object
}

// Type and Inspect methods for ReturnValue.
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

// Error represents an error message.
type Error struct {
	Message string
}

// Type and Inspect methods for Error.
func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }

// Function represents a user-defined function
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Type and Inspect methods for Function.
func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

// String represents a string value.
type String struct {
	Value string
}

// Type and Inspect methods for String
func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }

// Builtin represents a built-in function
type Builtin struct {
	Fn BuiltinFunction
}

// Type and Inspect methods for Builtin.
func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }

// BuiltinFunction defines a function signature for built-in functions.
type BuiltinFunction func(args ...Object) Object

// Array represents a collection of objects
type Array struct {
	Elements []Object
}

// Type and Inspect methods for Array.
func (ao *Array) Type() ObjectType { return ARRAY_OBJ }
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// HashKey represents a key-value pair in a Hash.
type HashKey struct {
	Type  ObjectType
	Value uint64
}

// HashPair represents a key-value pair in a Hash.
type HashPair struct {
	Key   Object
	Value Object
}

// Hash represents a key-value pair in a Hash.
type Hash struct {
	Pairs map[HashKey]HashPair
}

// Type and Inspect methods for Hash.
func (h *Hash) Type() ObjectType { return HASH_OBJ }
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

// Hashable is an interface for objects that can be used as hash keys.
type Hashable interface {
	HashKey() HashKey
}

// HashKey generates a hash key for Boolean.
func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}
	return HashKey{Type: b.Type(), Value: value}
}

// HashKey generates a hash key for Integer.
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

// HashKey generates a hash key for String.
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

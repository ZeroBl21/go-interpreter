package evaluator

import (
	"github.com/ZeroBl21/go-interpreter/object"
)

var builtins = map[string]*object.Builtin{
	"print": object.GetBuiltinByName("print"),
	"len":   object.GetBuiltinByName("len"),
	"first": object.GetBuiltinByName("first"),
	"last": object.GetBuiltinByName("last"),
	"rest": object.GetBuiltinByName("rest"),
	"push": object.GetBuiltinByName("push"),
}

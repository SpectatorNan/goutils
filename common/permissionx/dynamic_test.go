package permissionx

import (
	"strings"
	"testing"
)

type FunctionType struct {
	TypeID        int
	TypeName      string
	ConditionRule string
}

type Function struct {
	FunctionID    int
	FunctionType  int
	FunctionValue string
}

type User struct {
	ID       int
	Location string
	Grade    string
	Action   string
}

type Context struct {
	User     User
	Function Function
}

var functionTypes = []FunctionType{
	{TypeID: 1, TypeName: "normal", ConditionRule: "action in ['create', 'read', 'update', 'delete']"},
	{TypeID: 2, TypeName: "location", ConditionRule: "user.location == function_value"},
	{TypeID: 3, TypeName: "grade", ConditionRule: "user.grade == function_value"},
}

var functions = []Function{
	{FunctionID: 1, FunctionType: 1, FunctionValue: "create"},
	{FunctionID: 2, FunctionType: 2, FunctionValue: "main_campus"},
	{FunctionID: 3, FunctionType: 3, FunctionValue: "grade_1"},
}

func getFunctionTypeByID(id int) *FunctionType {
	for _, ft := range functionTypes {
		if ft.TypeID == id {
			return &ft
		}
	}
	return nil
}

func checkPermission(ctx Context) bool {
	function := ctx.Function
	user := ctx.User

	functionType := getFunctionTypeByID(function.FunctionType)
	if functionType == nil {
		return false
	}

	rule := functionType.ConditionRule

	// Simplified rule evaluation, replace with a proper expression evaluator for complex rules
	switch functionType.TypeName {
	case "normal":
		return strings.Contains(rule, user.Action)
	case "location":
		return user.Location == function.FunctionValue
	case "grade":
		return user.Grade == function.FunctionValue
	default:
		return false
	}
}

func TestCheckPermission(t *testing.T) {
	ctx := Context{
		User:     User{ID: 1, Location: "main_campus", Grade: "grade_1", Action: "create"},
		Function: Function{FunctionID: 1, FunctionType: 1, FunctionValue: "create"},
	}
	if !checkPermission(ctx) {
		t.Errorf("expected permission to be granted")
	}

	ctx = Context{
		User:     User{ID: 1, Location: "main_campus", Grade: "grade_1", Action: "create"},
		Function: Function{FunctionID: 1, FunctionType: 1, FunctionValue: "read"},
	}
	if checkPermission(ctx) {
		t.Errorf("expected permission to be denied")
	}

	ctx = Context{
		User:     User{ID: 1, Location: "main_campus", Grade: "grade_1", Action: "create"},
		Function: Function{FunctionID: 2, FunctionType: 2, FunctionValue: "main_campus"},
	}
	if !checkPermission(ctx) {
		t.Errorf("expected permission to be granted")
	}

	ctx = Context{
		User:     User{ID: 1, Location: "main_campus", Grade: "grade_1", Action: "create"},
		Function: Function{FunctionID: 2, FunctionType: 2, FunctionValue: "branch_campus"},
	}
	if checkPermission(ctx) {
		t.Errorf("expected permission to be denied")
	}

	ctx = Context{
		User:     User{ID: 1, Location: "main_campus", Grade: "grade_1", Action: "create"},
		Function: Function{FunctionID: 3, FunctionType: 3, FunctionValue: "grade_1"},
	}
	if !checkPermission(ctx) {
		t.Errorf("expected permission to be granted")
	}

	ctx = Context{
		User:     User{ID: 1, Location: "main_campus", Grade: "grade_1", Action: "create"},
		Function: Function{FunctionID: 3, FunctionType: 3, FunctionValue: "grade_2"},
	}
	if checkPermission(ctx) {
		t.Errorf("expected permission to be denied")
	}
}

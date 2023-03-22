package utilities_test

import (
	"testing"

	a "github.com/tsanton/goflake-client/goflake/models/assets"
	i "github.com/tsanton/goflake-client/goflake/models/assets/interface"
	u "github.com/tsanton/goflake-client/goflake/utilities"
)

func Test_stack_is_empty(t *testing.T) {
	/* Arrange */
	s := u.Stack[i.ISnowflakeAsset]{}

	/* Act and Assert */
	if !s.IsEmpty() {
		t.FailNow()
	}
}

func Test_stack_put(t *testing.T) {
	/* Arrange */
	s := u.Stack[i.ISnowflakeAsset]{}

	/* Act */
	r := a.RoleRelationship{ChildRoleName: "CHILD", ParentRoleName: "PARENT"}
	s.Put(&r)

	/* Act and Assert */
	if s.IsEmpty() {
		t.FailNow()
	}
}

func Test_stack_order(t *testing.T) {
	/* Arrange */
	s := u.Stack[i.ISnowflakeAsset]{}
	r1 := a.RoleRelationship{ChildRoleName: "A", ParentRoleName: "B"}
	r2 := a.RoleRelationship{ChildRoleName: "C", ParentRoleName: "D"}
	r3 := a.RoleRelationship{ChildRoleName: "E", ParentRoleName: "F"}
	s.Put(&r1)
	s.Put(&r2)
	s.Put(&r3)

	/* Act */
	i1 := s.Get()
	i2 := s.Get()
	i3 := s.Get()

	g1, ok1 := i1.(*a.RoleRelationship)
	g2, ok2 := i2.(*a.RoleRelationship)
	g3, ok3 := i3.(*a.RoleRelationship)

	/* Assert */
	if !ok1 || g1.ChildRoleName != r3.ChildRoleName {
		t.FailNow()
	}

	if !ok2 || g2.ChildRoleName != r2.ChildRoleName {
		t.FailNow()
	}

	if !ok3 || g3.ChildRoleName != r1.ChildRoleName {
		t.FailNow()
	}

	if !s.IsEmpty() {
		t.FailNow()
	}
}

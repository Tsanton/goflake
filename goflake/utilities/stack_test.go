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
	r := a.RoleInheritance{ChildPrincipal: &a.Role{Name: "CHILD"}, ParentPrincipal: &a.Role{Name: "PARENT"}}
	s.Put(&r)

	/* Act and Assert */
	if s.IsEmpty() {
		t.FailNow()
	}
}

func Test_stack_order(t *testing.T) {
	/* Arrange */
	s := u.Stack[i.ISnowflakeAsset]{}
	r1 := a.RoleInheritance{ChildPrincipal: &a.Role{Name: "A"}, ParentPrincipal: &a.Role{Name: "B"}}
	r2 := a.RoleInheritance{ChildPrincipal: &a.Role{Name: "C"}, ParentPrincipal: &a.Role{Name: "D"}}
	r3 := a.RoleInheritance{ChildPrincipal: &a.Role{Name: "E"}, ParentPrincipal: &a.Role{Name: "F"}}
	s.Put(&r1)
	s.Put(&r2)
	s.Put(&r3)

	/* Act */
	i1 := s.Get()
	i2 := s.Get()
	i3 := s.Get()

	g1, ok1 := i1.(*a.RoleInheritance)
	g2, ok2 := i2.(*a.RoleInheritance)
	g3, ok3 := i3.(*a.RoleInheritance)

	/* Assert */
	if !ok1 || g1.ChildPrincipal != r3.ChildPrincipal {
		t.FailNow()
	}

	if !ok2 || g2.ChildPrincipal != r2.ChildPrincipal {
		t.FailNow()
	}

	if !ok3 || g3.ChildPrincipal != r1.ChildPrincipal {
		t.FailNow()
	}

	if !s.IsEmpty() {
		t.FailNow()
	}
}

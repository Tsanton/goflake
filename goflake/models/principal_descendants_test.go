package models_test

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	g "github.com/tsanton/goflake-client/goflake"
	i "github.com/tsanton/goflake-client/goflake/integration"
	d "github.com/tsanton/goflake-client/goflake/models/describables"
	e "github.com/tsanton/goflake-client/goflake/models/entities"
)

func Test_security_admin_descendants(t *testing.T) {
	cli := i.Goflake()
	defer cli.Close()
	/* Act */
	hier, err := g.DescribeOne[e.PrincipalDescendants](cli, &d.PrincipalDescendants{Principal: &d.Role{Name: "SECURITYADMIN"}})

	/* Assert */
	assert.Nil(t, err)

	usr, ok := lo.Find(hier.Descendants, func(i e.Grant) bool { return i.GrantedIdentifier == "USERADMIN" })
	assert.True(t, ok)
	assert.Equal(t, 0, usr.DistanceFromSource)
}

func Test_account_admin_descendants(t *testing.T) {
	cli := i.Goflake()
	defer cli.Close()
	/* Act */
	hier, err := g.DescribeOne[e.PrincipalDescendants](cli, &d.PrincipalDescendants{Principal: &d.Role{Name: "ACCOUNTADMIN"}})

	/* Assert */
	assert.Nil(t, err)

	sec, ok := lo.Find(hier.Descendants, func(i e.Grant) bool { return i.GrantedIdentifier == "SECURITYADMIN" })
	assert.True(t, ok)
	assert.Equal(t, 0, sec.DistanceFromSource)

	sys, ok := lo.Find(hier.Descendants, func(i e.Grant) bool { return i.GrantedIdentifier == "SYSADMIN" })
	assert.True(t, ok)
	assert.Equal(t, 0, sys.DistanceFromSource)
}

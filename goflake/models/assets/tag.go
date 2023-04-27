package assets

import (
	"fmt"
	"strings"

	i "github.com/tsanton/goflake-client/goflake/models/assets/interface"
	e "github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ i.ISnowflakeAsset = &Tag{}
)

type Tag struct {
	DatabaseName string
	SchemaName   string
	TagName      string
	TagValues    []string
	Comment      string
	Owner        i.ISnowflakePrincipal
}

// GetCreateStatement implements ISnowflakeAsset
func (r *Tag) GetCreateStatement() (string, int) {
	var sb strings.Builder
	var principal e.SnowflakePrincipal
	switch x := r.Owner.GetPrincipalType(); x {
	case e.SnowflakePrincipalRole, e.SnowflakePrincipalDatabaseRole:
		principal = x
	default:
		panic("Ownership for this principal type is not implemented")
	}

	sb.WriteString(fmt.Sprintf("CREATE TAG %[1]s.%[2]s.%[3]s", r.DatabaseName, r.SchemaName, r.TagName))

	if len(r.TagValues) > 0 {
		var wrappedVals []string
		for _, val := range r.TagValues {
			wrappedVals = append(wrappedVals, fmt.Sprintf("'%[1]s'", val))
		}
		sb.WriteString(fmt.Sprintf(" ALLOWED_VALUES %[1]s", strings.Join(wrappedVals, ", ")))
	}

	sb.WriteString(fmt.Sprintf(" COMMENT = '%[1]s';\n", r.Comment))

	sb.WriteString(fmt.Sprintf("GRANT OWNERSHIP ON TAG %[1]s.%[2]s.%[3]s TO %[4]s %[5]s;", r.DatabaseName, r.SchemaName, r.TagName, principal.GrantType(), r.Owner.GetIdentifier()))
	return sb.String(), 2
}

// GetDeleteStatement implements ISnowflakeAsset
func (r *Tag) GetDeleteStatement() (string, int) {
	return fmt.Sprintf("DROP TAG %[1]s.%[2]s.%[3]s;", r.DatabaseName, r.SchemaName, r.TagName), 1
}

package entities

type ClassificationTag struct {
	TagDatabaseName string `json:"tag_database"`
	TagSchemaName   string `json:"tag_schema"`
	TagName         string `json:"tag_name"`
	//Either TABLE or COLUMN: indicates if the tag is applied directly to the column or if it is inherited from the table
	DomainLevel string `json:"domain"`
	TagValue    string `json:"tag_value"`
}

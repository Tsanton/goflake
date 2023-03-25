package inter

type ISnowflakeAsset interface {
	GetCreateStatement() (string, int)
	GetDeleteStatement() (string, int)
}

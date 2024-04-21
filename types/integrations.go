package types

type DatabaseType string

const (
	DatabaseTypeSqlite     DatabaseType = "sqlite"
	DatabaseTypeMySQL      DatabaseType = "mysql"
	DatabaseTypePostgreSQL DatabaseType = "postgresql"
)

type QueueType string

const (
	QueueTypeKafka QueueType = "kafka"
)

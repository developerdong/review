package conf

import "os"

type env string

const prefix env = "REVIEW_"

const (
	// See database/sql.Open
	DriverName     = prefix + "DRIVER_NAME"
	DataSourceName = prefix + "DATA_SOURCE_NAME"
)

func GetEnv(key env) string {
	return os.Getenv(string(key))
}

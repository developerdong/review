package conf

import "os"

type env string

const prefix env = "REVIEW_"

const (
	DriverName     = prefix + "DRIVER_NAME"
	DataSourceName = prefix + "DATA_SOURCE_NAME"
) // See database/sql.Open

func GetEnv(key env) string {
	return os.Getenv(string(key))
}

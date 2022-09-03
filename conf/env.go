package conf

import "os"

type env string

const prefix env = "REVIEW_"

const (
	// DriverName is from https://pkg.go.dev/database/sql#Open.
	DriverName     = prefix + "DRIVER_NAME"
	DataSourceName = prefix + "DATA_SOURCE_NAME"
	// URLFilePath is the path of file whose content is the previously selected url.
	URLFilePath = prefix + "URL_FILE_PATH"
)

func GetEnv(key env) string {
	return os.Getenv(string(key))
}

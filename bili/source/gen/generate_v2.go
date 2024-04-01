package main

import (
	"github.com/wordpress-plus/kit-common/gormx/source/gen"
	"os"
	"path"
)

var (
	moduleRoot, _ = os.Getwd()
)

var MySQLDSN string = "wJ0ZxlPJVsTm/iaFWxAIa1M5jF+eYR9JvAFlnCdxzCbWJ+TQY093XQ4OKiqqmjBKzHxFIvOoiZ7/eMzuHBaU7zZ2UzzZyHj0/jbDoLIfqK6qu20k4ibz8BXR7bzUAFykoxSTNW00g1kWUPj5yiBWql/LrtkeKMmCusoreXsRNwII+DvPIEVI9JKIB2ynYkyT"

func main2() {
	gen.G(MySQLDSN, "./source/gen/dal", path.Join(moduleRoot, "./source/gen/relation.yaml"))
}

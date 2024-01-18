package conf_test

import (
	"os"
	"testing"

	"github.com/GeekQk/vblog/conf"
)

func TestLoadFromFile(t *testing.T) {
	err := conf.LoadFromFile("etc/application.toml")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C())
}

func TestLoadFromEnv(t *testing.T) {
	os.Setenv("DATASOURCE_HOST", "127.0.0.1")
	os.Setenv("DATASOURCE_PORT", "3306")
	os.Setenv("DATASOURCE_USERNAME", "root")
	os.Setenv("DATASOURCE_PASSWORD", "test")
	os.Setenv("DB_NAME", "vblog")
	err := conf.LoadFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C())
}

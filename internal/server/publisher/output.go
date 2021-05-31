package publisher

import (
	"context"
	"time"

	"github.com/elastic/go-ucfg"
)

type esConfig struct {
	Urls     []string `config:"hosts" validate:"required"`
	UserName string   `config:"username"`
	PassWord string   `config:"password"`
}

type kafkaConfig struct {
	Urls     []string      `config:"hosts" validate:"required"`
	UserName string        `config:"username"`
	PassWord string        `config:"password"`
	Timeout  time.Duration `config:"timeout"             validate:"min=1"`
}

type HttpConfig struct {
	ApiHost string `config:"host" validate:"required"`
	ApiUri  string `config:"uri" validate:"required"`
}

type influxdbConfig struct {
	Urls     []string `config:"hosts" validate:"required"`
	UserName string   `config:"username"`
	PassWord string   `config:"password"`
}

type fileConfig struct {
	Path string `config:"path" validate:"required"`
}

type TransferConfig struct {
	Concurrency   uint8         `config:"concurrency" validate:"min=1,max=32"`
	Timeout       time.Duration `config:"timeout" validate:"min=1s"`
	Server        string        `config:"server" validate:"required"`
	Port          int           `config:"port" validate:"required"`
	CertFile      string        `config:"cert"`
	Tls           bool          `config:"tls"`
	TlsSkipVerify bool          `config:"tls-skip-verify"`
}

type OutputConfig struct {
	EsConfig       esConfig       `config:"elasticsearch"`
	InfluxdbConfig influxdbConfig `config:"influxdb"`
	FileConfig     fileConfig     `config:"file"`
}

type Outputer interface {
	Name() string
	OutputJson(ctx context.Context, id, index, tpy string, js interface{}, key []byte) error
	Close()
}

type OutputCreater func(config map[string]*ucfg.Config) (Outputer, error)

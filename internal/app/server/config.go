package server

import (
	"encoding/json"
	"io"
	"os"

	"github.com/MishaNiki/go-login-signup/internal/app/mail"
	"github.com/MishaNiki/go-login-signup/internal/app/storage"
	"github.com/MishaNiki/go-login-signup/internal/app/templates"
)

// Config ...
type Config struct {
	BindPort  string `json:"bind_port"`
	Storage   *storage.Config
	Mail      *mail.Config
	Templates *templates.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindPort:  ":8085",
		Storage:   storage.NewConfig(),
		Mail:      mail.NewConfig(),
		Templates: templates.NewConfig(),
	}
}

// DecodeJSONConf ...
func (config *Config) DecodeJSONConf(configPath string) error {
	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	data := make([]byte, 64)

	var lenBuf int
	for {
		len, e := file.Read(data)
		lenBuf += len
		if e == io.EOF {
			break
		}
	}
	err = json.Unmarshal(data[:lenBuf], config)
	if err != nil {
		return err
	}

	return nil
}

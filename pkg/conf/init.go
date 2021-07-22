package conf

import (
	"os"

	"github.com/dapr/kit/config"
	"gopkg.in/yaml.v3"
)

func Init(filePath string) (conf *Config, err error) {
	var data []byte

	if data, err = os.ReadFile(filePath); err != nil {
		return
	}

	cfg := make(map[string]interface{})

	if err = yaml.Unmarshal(data, cfg); err != nil {
		return
	}

	conf = new(Config)
	if err = config.Decode(cfg, conf); err != nil {
		return
	}

	return
}

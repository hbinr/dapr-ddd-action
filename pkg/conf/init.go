package conf

import (
	"log"
	"os"

	"github.com/dapr/kit/config"
	"gopkg.in/yaml.v3"
)

func Init(filePath string) (conf *Config) {
	var (
		data []byte
		err  error
	)

	if data, err = os.ReadFile(filePath); err != nil {
		log.Fatalln("conf: os.ReadFile failed,err:", err)
	}

	cfg := make(map[string]interface{})

	if err = yaml.Unmarshal(data, cfg); err != nil {
		log.Fatalln("conf: yaml.Unmarshal failed, err:", err)
	}

	conf = new(Config)
	if err = config.Decode(cfg, conf); err != nil {
		log.Fatalln("conf: config.Decode failed, err: ", err)
	}

	return
}

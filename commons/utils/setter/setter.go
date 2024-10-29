package setter

import (
	"fmt"
	"go-notifier/commons/utils/logger"
	"io/ioutil"
	"os"

	"github.com/spf13/viper"
)

type Utils struct {
}

func NewUtils() *Utils {
	return &Utils{}
}

func (u *Utils) SetApplicationVersion(key, filename string) (string, error) {
	var (
		version []byte
		err     error
	)
	version, err = ioutil.ReadFile(filename)
	os.Setenv(key, string(version))
	return string(version), err
}

func (u *Utils) SetDefaultProperties(propertiesMap map[string]interface{}) {
	log := logger.New(logger.Warn)
	for key, value := range propertiesMap {
		if !viper.IsSet(key) {
			log.Warnf(fmt.Sprintf("no value set for property :  %s", key))
			viper.SetDefault(key, value)
		}
	}
}

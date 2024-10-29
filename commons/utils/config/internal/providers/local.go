package ConfigCloud

import (
	"encoding/json"
	"os"
)

func GetFile(path string) (ConfigResponse, error) {
	var cloudConfig ConfigResponse
	file, err := os.ReadFile(path)
	if err != nil {
		return ConfigResponse{}, err
	}
	err = json.Unmarshal(file, &cloudConfig)
	if err != nil {
		return ConfigResponse{}, err
	}
	return cloudConfig, nil
}

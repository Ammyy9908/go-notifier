package config

import (
	"go-notifier/commons/constants"
	"time"
)

var PropertiesMap = map[string]interface{}{
	constants.MAX_OPTMISTIC_LOCKING_RETRY_COUNT: 3,
	constants.CACHE_TTL:                         5 * 24 * time.Hour,
	constants.REST_EXECUTE_TIME_OUT_IN_SEC:      20,
	constants.MESSAGE_PUBLISHER_RETRY_COUNT:     5,
	constants.MAX_STARTUP_ATTEMPT:               3,
}

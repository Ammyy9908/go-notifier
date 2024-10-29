package safegoroutine

import "go-notifier/commons/utils/logger"

func SafeGoRoutine(identifier string, fn func(), maxStartupAttemptCount int) {
	log := logger.New(logger.Info)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("%s crashed due to %v. Attempting recovery !!!", identifier, err)
				if maxStartupAttemptCount > 1 {
					log.Errorf("Attempting intialization of %s once again", identifier)
					SafeGoRoutine(identifier, fn, maxStartupAttemptCount-1)
				} else {
					log.Errorf("maxStartupAttemptCount has exhausted. %s go routine will terminate now due to %v", identifier, err)
				}
			}
		}()
		fn()
	}()
}

package main

import (
	"flag"
	"fmt"
	commonConstants "go-notifier/commons/constants"
	"go-notifier/commons/utils/config"
	"go-notifier/commons/utils/gracefulshutdown"
	"go-notifier/commons/utils/logger"
	safegoroutine "go-notifier/commons/utils/safe_go_routine"
	"go-notifier/commons/utils/setter"
	configs "go-notifier/notification-service/config"
	"go-notifier/notification-service/internal/adapter/db"
	"go-notifier/notification-service/internal/pkg"
	"go-notifier/notification-service/internal/port/constants"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

func init() {
	env := os.Getenv("ENV")
	env = "development"
	logger.InitializeLogger("notification-service", env)
	log := logger.GetLogger()

	defer log.Sync()
	//Read the Version file
	_, mainFilePath, _, _ := runtime.Caller(0)
	projectRootDir := filepath.Dir(mainFilePath)

	ver, err := setter.NewUtils().SetApplicationVersion(constants.ENV_IS_VERSION, projectRootDir+"/"+constants.VERSION_FILE_NAME)
	if err != nil {
		log.Error(err)
	} else {
		viper.Set("version", ver)
	}

	log.Info("VERSION: ", ver)

	configs := config.New()
	if env == constants.DEV_ENV { // Remove extra {
		configs = configs.FromFile(projectRootDir + constants.DEV_CONFIG_FILE_NAME)
		if configs.HasErrors() {
			log.Error(configs.Errors)
			log.Error("error while fetching configurations from local config file Error: ", configs.Errors)
			os.Exit(1)
		}
	}
	configs.SetViper()
	log.Info(viper.AllSettings())

}

func main() {
	log := logger.GetLogger()

	//Initialize Database
	db.Init()
	exitChannel := make(chan struct{})

	gracefulShutDownManager := gracefulshutdown.NewManager(log, exitChannel)
	setter.NewUtils().SetDefaultProperties(configs.PropertiesMap)

	StartService(gracefulShutDownManager, exitChannel)

	fmt.Println("Go Notifier!")
}

func StartService(gracefulShutDownManager *gracefulshutdown.Manager, exitChannel chan struct{}) {
	log := logger.GetLogger()
	flag.Usage = func() {
		fmt.Println("Usage: server -s {service_name} -e {environment}")
		os.Exit(1)
	}
	flag.Parse()
	appRouter := pkg.GetRouter()
	port := viper.GetString("server.port")
	httpServer := &http.Server{
		Addr:    port,
		Handler: appRouter,
	}

	safegoroutine.SafeGoRoutine("httpServer", func() {
		err := httpServer.ListenAndServe()
		if err != nil {
			log.Error("listen: ", port)
		}
	}, viper.GetInt(commonConstants.MAX_STARTUP_ATTEMPT))

	gracefulShutDownManager.Shutdown(httpServer)
	<-exitChannel
	log.Info("Server Shutdown")

}

package gracefulshutdown

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"go.uber.org/zap"
)

type Manager struct {
	log                 *zap.SugaredLogger
	ShutdownWaitGroup   *sync.WaitGroup
	ShutdownChannel     chan int
	exitChannelInstance chan struct{}
}

func NewManager(log *zap.SugaredLogger, exitChannel chan struct{}) *Manager {
	return &Manager{
		log:                 log,
		ShutdownWaitGroup:   new(sync.WaitGroup),
		ShutdownChannel:     make(chan int),
		exitChannelInstance: exitChannel,
	}
}

func (m *Manager) Shutdown(srv *http.Server) {
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGTERM, syscall.SIGINT)

	<-quitChannel
	m.log.Info("Quit signal received....")

	contextTimeoutIn := time.Duration(10)
	// Wait for interrupt signal to gracefully shut down the server
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeoutIn*time.Second)
	m.log.Info("Quit signal received, sending shutdown and waiting on HTTP calls...")
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		m.log.Fatal("Error Occurred",
			"Error message", err,
		)
	}
	m.log.Info("HTTP Server, shutdown gracefully.")

	m.log.Info("Quit signal received, sending shutdown and waiting on goroutines...")
	close(m.ShutdownChannel)

	m.ShutdownWaitGroup.Wait()
	m.log.Info("All go routines shutdown gracefully.")

	m.log.Info("main goroutine shutdown triggering...")
	close(m.exitChannelInstance)
}

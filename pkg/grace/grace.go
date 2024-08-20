package grace

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Serve(port string, h http.Handler) (err error) {
	server := &http.Server{
		ReadTimeout:  1000 * time.Second,
		WriteTimeout: 1000 * time.Second,
		Handler:      h,
	}

	listen, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
		<-signals

		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server shutdown : %v", err)
		}

		close(idleConnsClosed)
	}()

	log.Printf("HTTP server is running on port: %v\n", port)
	if err := server.Serve(listen); !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	<-idleConnsClosed
	log.Println("HTTP server shutdown gracefully")
	return nil
}

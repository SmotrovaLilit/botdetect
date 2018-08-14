package examples

import (
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"time"
	"log"
	"context"
	"SmotrovaLilit/botdetect"
)

func main() {

	errChan := make(chan error, 2)

	var (
		httpAddr = flag.String("http.addr", ":8000", "HTTP listen address")
	)

	var server *http.Server
	{
		http.HandleFunc("/", handlerCheckBot)
		server = &http.Server{
			Addr: *httpAddr,
		}

		go func() {
			log.Println("transport", "http", "address", *httpAddr, "msg", "listening")
			errChan <- server.ListenAndServe()
		}()
	}

	go func() {
		signals := make(chan os.Signal)
		signal.Notify(signals, syscall.SIGINT)
		errChan <- fmt.Errorf("%s", <-signals)
	}()

	log.Println("terminated", <-errChan)

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()

	server.Shutdown(ctx)
}


func handlerCheckBot(w http.ResponseWriter, r *http.Request) {
	s := botdetect.NewBotDetect(r, nil)
	if s.IsBot() {
		http.Error(w, "it is bot", http.StatusForbidden)
	}
}

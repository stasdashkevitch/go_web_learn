package main

import (
	"context"
	"log"
	"micro/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hp := handlers.NewProducts(l)
	// handle - регает handler - структура

	// method
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//
	// 	log.Println("Basic")
	// 	fmt.Fprintf(w, "Basic")
	// })

	// handle func - сахар
	sm := http.NewServeMux()
	sm.Handle("/", hp)

	// в реальных проектах мы должны делать настраиваемый сервер
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// go func используется для того чтобы мы не блокировали основной поток
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// основной поток будет ожидать сигнала os
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)
	// нужно для плавного завершения: новые запросы не будут обрабатываться, однако старые будут обрабатываться до конца
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel() // освобождение ресурсов в конце программы

	s.Shutdown(tc)
}

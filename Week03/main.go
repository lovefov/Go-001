package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	eg, ctx := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithCancel(ctx)
	// 整两个 goroutine
	eg.Go(func() error {
		log.Println("server 1 on localhost:9090")
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello world"))
		})
		return server(ctx, mux, ":9090")
	})

	eg.Go(func() error {
		log.Println("server 2 on localhost:8080")
		mux := http.NewServeMux()
		mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		})
		return server(ctx, mux, ":8080")
	})

	quit := make(chan os.Signal, 0)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("graceful shut down ...")
	cancel()

	fmt.Printf("errgroup exiting: %+v\n", eg.Wait())
}

// 包一个 server 函数丢给 goroutine 执行.
func server(ctx context.Context, handler http.Handler, addr string) error {
	s := http.Server{
		Handler: handler,
		Addr:    addr,
	}
	// 收到 ctx 发来的退出信号 调用 http server 的shutdown() 优雅退出.
	go func() {
		<-ctx.Done()
		log.Printf("server will exiting, addr: %s", addr)
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}

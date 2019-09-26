package cmd

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"

	transportHttp "github.com/lmfuture-ma/lmaker/transport/http"

	"github.com/lmfuture-ma/lmaker/example/todolist/server"
	"github.com/lmfuture-ma/lmaker/example/todolist/services"
)

func NewTodolistCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "todolist",
		RunE: func(cmd *cobra.Command, args []string) error {
			// 初始化router
			start, shutdown, err := newRouter()
			if err != nil {
				return err
			}
			start()
			c := make(chan os.Signal, 1)
			signal.Reset(syscall.SIGINT, syscall.SIGTERM)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case <-c:
				shutdown()
				os.Exit(0)
			}
			return nil
		},
	}
	return cmd
}

// Start 启动函数
type Start func()

// Shutdown 停止函数
type Shutdown func()

//NewRouter：初始化路由
func newRouter() (Start, Shutdown, error) {
	e := transportHttp.NewMuxServer()
	endpoint := server.NewServiceImpl()
	services.MakeHandler(endpoint, e)

	// 构建 http server,用来平滑关闭
	svr := http.Server{Addr: "127.0.0.1:8088", Handler: e, ReadTimeout: 15 * time.Second, WriteTimeout: 15 * time.Second}
	// 启动函数
	start := func() {
		go func() {
			if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			}
		}()
	}
	// 关闭函数
	shutdown := func() {
		ctx, cancel := context.WithTimeout(context.Background(), 30)
		defer cancel()
		if err := svr.Shutdown(ctx); err != nil {
		}
	}
	return start, shutdown, nil
}

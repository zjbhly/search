package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	baseerror "search/common"
	"search/internal/config"
	"search/internal/handler"
	"search/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/search-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	//自定义错误
	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, interface{}) {
		switch e := err.(type) {
		case *baseerror.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})
	// 初始化kafka管理
	//km := kafka.GetKafkaManager()
	//controlCtx, cancelFn := context.WithCancel(context.Background())
	//defer cancelFn()
	//km.Init(controlCtx, ctx.Config.KafKa.Hosts)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	fmt.Println("server start!")
	server.Start()
}

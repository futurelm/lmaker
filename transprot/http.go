package transprot

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type HttpServer struct {
	dec DecodeHttpReq
	//todo
	//  ctx      context.Context
	//	before   []RequestFunc
	//	after    []ServerResponseFunc
	handler HttpHandler
}

func NewHttpServer(dec DecodeHttpReq, handler HttpHandler) *HttpServer {
	return &HttpServer{
		dec, handler,
	}
}

func (server *HttpServer) Server(c *gin.Context) {
	structReq, err := server.dec(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := server.handler(c, structReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func NewGinServer() *gin.Engine {
	e := gin.New()
	e.Use(gin.Recovery())
	if gin.Mode() != gin.ReleaseMode {
		e.Use(gin.Logger())
	}
	e.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://*", "https://*"},
		AllowMethods:     []string{"PUT", "GET", "POST", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origins", "X-Requested-With", "Content-Type", "Accept", "access_token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))
	return e
}

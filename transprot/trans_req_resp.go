package transprot

import (
	"github.com/gin-gonic/gin"
)

// get req struct from url,query,body
type DecodeHttpReq func(c *gin.Context) (structReq interface{}, err error)

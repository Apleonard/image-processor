package middleware

import (
	"errors"
	"fmt"
	"os"
	"runtime/debug"

	apphtpp "image-processor/pkg/http"

	"github.com/gin-gonic/gin"
)

func GinRecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			rcv := panicRecover(recover())
			if rcv != nil {
				fmt.Fprintf(os.Stderr, "Panic: %+v\n", rcv)
				debug.PrintStack()
				apphtpp.ResponseInternalServerError(c, rcv)
			}
		}()
		c.Next()
	}
}

func recovery(r interface{}) error {
	var err error
	if r != nil {
		switch t := r.(type) {
		case string:
			err = errors.New(t)
		case error:
			err = t
		default:
			err = errors.New("unknown error")
		}
	}
	return err
}

func panicRecover(rc interface{}) error {
	return recovery(rc)
}

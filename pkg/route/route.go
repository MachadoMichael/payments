package route

import (
	"github.com/MachadoMichael/payments/pkg/handler"
	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()

	v1 := router.Group("/payments")
	{
		v1.POST("/", handler.MakePayment)
	}

	router.Run()

}

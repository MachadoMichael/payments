package handler

import (
	"net/http"

	"github.com/MachadoMichael/payments/schema"
	"github.com/gin-gonic/gin"
)

func MakePayment(ctx *gin.Context) {
	request := schema.Payment{}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}

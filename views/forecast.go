package views

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Person struct {
}

func GetPendingForecasts(context *gin.Context) {
	context.Writer.WriteHeader(http.StatusAccepted)
	context.Writer.WriteString("Shu!")
}

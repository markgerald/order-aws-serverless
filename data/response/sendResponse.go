package response

import (
	"github.com/gin-gonic/gin"
)

func SendErrorResponse(ctx *gin.Context, code int, status string) {
	webresponse := Response{
		Code:   code,
		Status: status,
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, webresponse)
}

func SendSuccessResponse(ctx *gin.Context, code int, status string, data interface{}) {
	webresponse := Response{
		Code:   code,
		Status: status,
		Data:   data,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, webresponse)
}

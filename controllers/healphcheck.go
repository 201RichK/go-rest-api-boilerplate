package controllers

import (
	"net/http"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// Healphcheck implements
// @Summary api healph check
// @Produce  json
// @Success 200 {object} Response{}
// @Failure 500 {object} Response
// @Router /healphcheck [get]
// @tags Healphcheck
func Healphcheck(ctx *gin.Context) {
	HTTPRes(ctx, http.StatusOK, http.StatusText(http.StatusOK), ginI18n.MustGetMessage(
		ctx,
		&i18n.LocalizeConfig{
			MessageID: "hello-world",
			TemplateData: map[string]string{
				"name": ctx.Query("name"),
			},
		}))
}

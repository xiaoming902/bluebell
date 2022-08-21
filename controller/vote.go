package controller

import (
	"bluebell/models"

	"github.com/gin-gonic/gin"
)

func PostVoteController(c *gin.Context) {
	err := c.ShouldBindQuery("name")
	if err != nil {
		return
	}
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		return
	}

}

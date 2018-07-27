package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (rH RouterHandler) commentsGet(c *gin.Context) {
	log := rH.log(c.Request.URL.Path)

	comments, err := rH.ucHandler.CommentsGet(c.Param("slug"))
	if err != nil {
		log(err)
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	c.JSON(http.StatusOK, gin.H{"comments": comments})
}

type commentPostReq struct {
	Comment struct {
		Body string `json:"body,required"`
	} `json:"comment,required"`
}

func (rH RouterHandler) commentPost(c *gin.Context) {
	log := rH.log(c.Request.URL.Path)

	req := &commentPostReq{}
	if err := c.BindJSON(req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	comment, err := rH.ucHandler.CommentsPost(rH.getUserName(c), c.Param("slug"), req.Comment.Body)
	if err != nil {
		log(err)
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"comment": comment})
}

func (rH RouterHandler) commentDelete(c *gin.Context) {
	log := rH.log(c.Request.URL.Path)

	if err := rH.ucHandler.CommentsDelete(rH.getUserName(c), c.Param("slug"), c.Param("id")); err != nil {
		log(err)
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	c.Status(http.StatusOK)
}

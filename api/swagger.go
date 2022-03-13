package api

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func Swagger(c *gin.Context) {
	b, err := ioutil.ReadFile("config/swagger.yml")

	if err != nil {
		log.Fatal(err)
	}

	c.String(http.StatusOK, string(b))
}

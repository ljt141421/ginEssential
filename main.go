package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r = CollectRouter(r)
	panic(r.Run()) // listen and serve on 0.0.0.0:8080

}

package main

import (
	"mechanic/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.Register(r)
	r.Run("localhost:50809")
}


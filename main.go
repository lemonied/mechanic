package main

import (
	"fmt"
	"mechanic/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	defer func()  {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	r := gin.Default()
	routers.Register(r)
	r.Run("localhost:50809")
}


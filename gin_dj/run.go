package gin_dj

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/Zncl2222/gin-dj/core"

	"github.com/gin-gonic/gin"
)

func Run() {
	// Initialize the Gin router
	router := gin.Default()
	if len(os.Args) >= 2 {
		valid := regexp.MustCompile(`^(\d{1,5}):(\d{1,5})$`).MatchString(os.Args[1])
		if valid {
			runServer(router, os.Args[1])
		} else {
			fmt.Println("Error Port Number \nUsage: go run main.go [port] or ./main.go [port]")
			os.Exit(1)
		}
	}
}

func runServer(router *gin.Engine, port string) {
	// Initialize router and middleware with the configuration
	RouterInit(router)
	core.MiddlewareInit(router)

	// Start the web server
	err := router.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}

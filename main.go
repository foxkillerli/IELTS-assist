package main

import (
	"flag"
	"fmt"
	"github.com/foxkillerli/IELTS-assist/config"
	"github.com/foxkillerli/IELTS-assist/route"
	gin "github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	var token string
	flag.StringVar(&token, "token", "", "openai token")
	flag.Parse()
	fmt.Printf("token: %s\n", token)
	config.OPENAI_TOKEN = token
	mux := http.NewServeMux()
	r := route.SetupRouter()
	r.Any("/admin/*resources", gin.WrapH(mux))
	r.Run("0.0.0.0:8080")
	log.Printf("[Debug] initializing backend server on host: %s, port: %d", "0.0.0.0", 8080)
}

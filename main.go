package main

import (
	"flag"
	"fmt"
	"github.com/foxkillerli/IELTS-assist/config"
	"github.com/foxkillerli/IELTS-assist/route"
	"github.com/foxkillerli/IELTS-assist/utils"
	gin "github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	var token string
	var postgresHost string
	var postgresPort int
	var postgresUser string
	var postgresPassword string
	var postgresDBName string
	flag.StringVar(&token, "token", "", "openai token")
	flag.StringVar(&postgresHost, "postgres-host", "", "Postgres Database Host")
	flag.IntVar(&postgresPort, "postgres-port", 5432, "Postgres Database Port")
	flag.StringVar(&postgresUser, "postgres-user", "", "Postgres Database User")
	flag.StringVar(&postgresPassword, "postgres-password", "", "Postgres Database Password")
	flag.StringVar(&postgresDBName, "postgres-db", "", "Postgres Database Name")
	flag.Parse()
	config.OPENAI_TOKEN = token
	config.PostgresHost = postgresHost
	config.PostgresPort = postgresPort
	config.PostgresUser = postgresUser
	config.PostgresPassword = postgresPassword
	config.PostgresDBName = postgresDBName
	fmt.Printf("token: %s\n", token)
	fmt.Printf("postgres host: %s\n", config.PostgresHost)
	fmt.Printf("postgres port: %v\n", config.PostgresPort)
	fmt.Printf("postgres user: %s\n", config.PostgresUser)
	fmt.Printf("postgres password: %s\n", config.PostgresPassword)
	fmt.Printf("postgres db name: %s\n", config.PostgresDBName)

	utils.Migrate()
	mux := http.NewServeMux()
	r := route.SetupRouter()
	r.Any("/admin/*resources", gin.WrapH(mux))
	r.Run("0.0.0.0:8080")
	log.Printf("[Debug] initializing backend server on host: %s, port: %d", "0.0.0.0", 8080)
}

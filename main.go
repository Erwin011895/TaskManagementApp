package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Erwin011895/TaskManagementApp/config"
	"github.com/Erwin011895/TaskManagementApp/handler"
	"github.com/Erwin011895/TaskManagementApp/routes"
	"github.com/gin-gonic/gin"
	sqlx "github.com/jmoiron/sqlx" //make alias name the package to sqlx
	_ "github.com/lib/pq"          //save it into underscore variable
)

func main() {
	conf := config.SetupConfig()

	dataSource := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		conf.Postgres.User, conf.Postgres.Password, conf.Postgres.Database)

	db, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		log.Fatalln(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	h := handler.InitHandler(db, conf)

	router := gin.Default()
	routes.Init(router, h)

	s := &http.Server{
		Addr:           ":" + conf.App.Port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	s.ListenAndServe()
}

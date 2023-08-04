package app

import (
	"log"
	"os"
	"os/signal"

	"github.com/Akmyrzza/blog-api/internal/config"
	"github.com/Akmyrzza/blog-api/internal/handler"
	"github.com/Akmyrzza/blog-api/internal/repository/sqlite"
	"github.com/Akmyrzza/blog-api/internal/service"
	"github.com/Akmyrzza/blog-api/pkg/httpserver"
	"github.com/Akmyrzza/blog-api/pkg/jwttoken"
	_ "github.com/mattn/go-sqlite3"
)

func Run(cfg *config.Config) error {
	//db
	db, err := sqlite.New(cfg.DB.DBName) //sql.Open("sqlite3", "./sqlite3-database")
	if err != nil {
		log.Printf("connection to DB err: %s", err.Error())
	}
	defer db.Sqldb.Close()
	log.Println("db connection success")

	token := jwttoken.New(cfg.Token.SecretKey)
	//service
	srvs := service.New(db, token, cfg)
	//handler
	hndlr := handler.New(srvs)
	//httpserver
	server := httpserver.New(
		hndlr.InitRouter(),
		httpserver.WithPort(cfg.HTTP.Port),
		httpserver.WithReadTimeout(cfg.HTTP.ReadTimeout),
		httpserver.WithWriteTimeout(cfg.HTTP.WriteTimeout),
		httpserver.WithShutdownTimeout(cfg.HTTP.ShutdownTimeout),
	)
	//run

	log.Println("server started")
	server.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case s := <-interrupt:
		log.Printf("signal received: %s", s.String())
	case err = <-server.Notify():
		log.Printf("server notify: %s", err.Error())
	}

	err = server.Shutdown()
	if err != nil {
		log.Printf("server shutdown err: %s", err)
	}

	return nil
}

package main

import (
	"log"
	"os"

	"github.com/nelsonmarro/vapulabs-web/config"
	"github.com/nelsonmarro/vapulabs-web/internal/web/server"
)

func main() {
	cfg := config.LoadConfigOnce()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	s := server.NewServer(cfg, infoLog, errorLog)

	if err := s.Serve(); err != nil {
		errorLog.Println(err)
		panic(err)
	}
}
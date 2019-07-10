package main

import (
	"log"
	"net/http"
	"path"
	"time"

	"videoSecurity/logwriter"

	"github.com/syndtr/goleveldb/leveldb"
)

type program struct {
	Logger     *logwriter.Logger
	httpServer *http.Server
	Db         *leveldb.DB
}

func main() {
	dir, err := getCurrentPath()
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("executing dir:%s", dir)

	p := &program{}

	p.initLog(dir)
	p.initDb(dir)
	p.initRouters()
	p.startHTTP()
}

//InitLog log initialization
func (p *program) initLog(dir string) (err error) {
	logger := logwriter.Logger{
		Path:  getPathToLog(dir, "log.txt"),
		Level: "Debug",
	}
	logger.Init()

	p.Logger = &logger

	return
}

//InitLog database initialization
func (p *program) initDb(dir string) (err error) {
	pathToDb := path.Join(dir, "/levelb")

	db, err := leveldb.OpenFile(pathToDb, nil)
	//defer db.Close()
	if err != nil {
		return err
	}

	p.Db = db
	return
}

//InitLog routers initialization
func (p *program) initRouters() (err error) {
	serviceContainer := ServiceContainer(p.Logger, p.Db)
	return initRouters(serviceContainer, p.Logger, p.Db)
}

//StartHTTP start HTTP server
func (p *program) startHTTP() (err error) {
	p.httpServer = &http.Server{
		Addr:         ":8002",
		Handler:      Router,
		WriteTimeout: time.Second * 150000,
		ReadTimeout:  time.Second * 150000,
		IdleTimeout:  time.Second * 600000,
	}

	p.Logger.Debug("http server start -->", p.httpServer.Addr)

	if err := p.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		p.Logger.Error("start http server error", err)
	}
	p.Logger.Debug("http server end")

	return
}

//getPathToLog get current path to log file
func getPathToLog(dir string, logPath string) string {
	return path.Join(dir, logPath)
}

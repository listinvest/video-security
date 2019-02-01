package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"./auth"
	"./tasks/autosearch"
	"./tasks/manualsearch"
	taskDispatcher "./tasks/task-dispatcher"
	"github.com/gorilla/mux"
)

//autoSearchHandler
var autoSearchHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	bizTask := autosearch.DeviceTask{
		Task: taskDispatcher.BizTask{
			ID:   "1",
			Name: "test",
		},
	}

	dispatcher := taskDispatcher.GetInstance()
	dispatcher.RunTask(&bizTask)

	devices := bizTask.Result.Devices
	j, _ := json.Marshal(devices)
	res := string(j)
	fmt.Fprintf(w, "Find: %s", res)
})

//manualSearchHandler
var manualSearchHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	bizTask := manualsearch.DeviceTask{
		Ips:   "192.168.11.255-192.168.12.2, 192.168.11.8",
		Ports: "8080-8083, 80",
		Task: taskDispatcher.BizTask{
			ID:   "1",
			Name: "test",
		},
	}

	dispatcher := taskDispatcher.GetInstance()
	dispatcher.RunTask(&bizTask)

	devices := bizTask.Result.Devices
	j, _ := json.Marshal(devices)
	res := string(j)
	fmt.Fprintf(w, "Find: %s", res)
})

//helloHandler home handler
var helloHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
})

func main() {

	r := mux.NewRouter()

	v1 := r.PathPrefix("/v1/").Subrouter()
	v1.Handle("/get-token", auth.GetTokenHandler).Methods("GET")
	v1.Handle("/home", auth.MiddlewareHandler(helloHandler)).Methods("GET")
	v1.Handle("/autosearch", autoSearchHandler).Methods("GET")
	v1.Handle("/manualsearch", manualSearchHandler).Methods("GET")

	srv := &http.Server{
		Addr: ":8001",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

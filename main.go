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
	"./videostreamer"
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
		Ips:   "192.168.11.100-192.168.11.185, 192.168.11.186",
		Ports: "80",
		Task: taskDispatcher.BizTask{
			ID:   "1",
			Name: "test",
		},
	}

	dispatcher := taskDispatcher.GetInstance()
	dispatcher.RunAsyncTask(&bizTask)

	fmt.Println("RunAsyncTask")

	start := time.Now()

	for {
		time.Sleep(1 * time.Second)

		t := dispatcher.GetTask(bizTask.Task.ID)
		if t == nil {
			break
		}

		if time.Since(start) > (3 * time.Second) {
			break
		}
	}

	fmt.Println("RunAsyncTask End")

	dispatcher.AbortTask(bizTask.Task.ID)

	devices := bizTask.Result.Devices
	j, _ := json.Marshal(devices)
	res := string(j)
	fmt.Fprintf(w, "Find: %s", res)
})

//helloHandler home handler
var helloHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
})

//videooHandler video handler
var videoHandler = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, "templates/video.html")
})

var videoStreamHandler = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	videostreamer.Run(rw, r, "rtsp://184.72.239.149/vod/mp4:BigBuckBunny_115k.mov", true)
})

func main() {

	r := mux.NewRouter()

	v1 := r.PathPrefix("/v1/").Subrouter()
	v1.Handle("/get-token", auth.GetTokenHandler).Methods("GET")
	v1.Handle("/home", auth.MiddlewareHandler(helloHandler)).Methods("GET")
	v1.Handle("/autosearch", autoSearchHandler).Methods("GET")
	v1.Handle("/manualsearch", manualSearchHandler).Methods("GET")
	v1.Handle("/video", videoHandler).Methods("GET")
	v1.Handle("/videostream", videoStreamHandler).Methods("GET")

	srv := &http.Server{
		Addr: ":8002",
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

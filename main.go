package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"./auth"
	"./tasks/autosearch"
	"./tasks/manualsearch"
	taskDispatcher "./tasks/task-dispatcher"
	vs "./videostreamer"

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

//VideoStreamHandler lk
func VideoStreamHandler(a *appContext, rw http.ResponseWriter, r *http.Request) (int, error) {
	rw.Header().Set("Content-Type", "video/mp4")
	rw.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	err := a.videoStreamer.Run(rw, r, "rtsp://admin:1q2w3e4r5t6y@192.168.11.131:554/cam/realmonitor?channel=1&subtype=1", false)
	if err != nil {
		log.Printf("error: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		_, _ = rw.Write([]byte("<h1>500 Internal server error</h1>"))
	}

	return 200, err
}

type appContext struct {
	videoStreamer *vs.VideoStreamer
}

type appHandler struct {
	*appContext
	h func(*appContext, http.ResponseWriter, *http.Request) (int, error)
}

func (ah appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Updated to pass ah.appContext as a parameter to our handler type.
	status, err := ah.h(ah.appContext, w, r)
	if err != nil {
		log.Printf("HTTP %d: %q", status, err)
		switch status {
		case http.StatusNotFound:
			http.NotFound(w, r)
			// And if we wanted a friendlier error page:
			// err := ah.renderTemplate(w, "http_404.tmpl", nil)
		case http.StatusInternalServerError:
			http.Error(w, http.StatusText(status), status)
		default:
			http.Error(w, http.StatusText(status), status)
		}
	}
}

func main() {
	ac := &appContext{
		videoStreamer: &vs.VideoStreamer{
			Mutex:      &sync.RWMutex{},
			Dispatcher: &vs.StreamerDispatcher{},
		},
	}

	r := mux.NewRouter()

	v1 := r.PathPrefix("/v1/").Subrouter()
	v1.Handle("/get-token", auth.GetTokenHandler).Methods("GET")
	v1.Handle("/home", auth.MiddlewareHandler(helloHandler)).Methods("GET")
	v1.Handle("/autosearch", autoSearchHandler).Methods("GET")
	v1.Handle("/manualsearch", manualSearchHandler).Methods("GET")
	v1.Handle("/video", videoHandler).Methods("GET")
	v1.Handle("/videostream", appHandler{ac, VideoStreamHandler}).Methods("GET")

	srv := &http.Server{
		Addr: ":8002",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 150000, // timeout write to io
		ReadTimeout:  time.Second * 150000,
		IdleTimeout:  time.Second * 600000,
		Handler:      r,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

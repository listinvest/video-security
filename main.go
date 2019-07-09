package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"
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
			Name: "auto search devices",
		},
	}

	dispatcher := taskDispatcher.GetInstance()
	dispatcher.RunTask(&bizTask)

	devices := bizTask.Result.Devices

	w.Header().Set("Content-Type", "application/json")
	json, _ := json.Marshal(devices)
	w.Write(json)
})

//manualSearchHandler
var manualSearchHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	queryValues := r.URL.Query()

	bizTask := manualsearch.DeviceTask{
		Ips:   queryValues.Get("ips"),
		Ports: queryValues.Get("ports"),
		Task: taskDispatcher.BizTask{
			ID:   "1",
			Name: "manual search devices",
		},
	}

	dispatcher := taskDispatcher.GetInstance()
	dispatcher.RunAsyncTask(&bizTask)
	dispatcher.Wait(&bizTask)

	devices := bizTask.Result.Devices

	w.Header().Set("Content-Type", "application/json")
	json, _ := json.Marshal(devices)
	w.Write(json)
})

//indexHandler index handler
var indexHandler = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	tpl := template.New("index").Delims("[[", "]]")
	tpl, err := tpl.ParseFiles("www/index.html")

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tpl.ExecuteTemplate(rw, "index.html", nil)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
})

//videoStreamHandler lk
func videoStreamHandler(a *appContext, rw http.ResponseWriter, r *http.Request) (int, error) {
	rw.Header().Set("Content-Type", "video/mp4")
	rw.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	url := getUrl(r)
	err := a.videoStreamer.Run(rw, r, url, false)
	if err != nil {
		log.Printf("error: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		_, _ = rw.Write([]byte("<h1>500 Internal server error</h1>"))
	}

	return 200, err
}

//getUrl url from request
func getUrl(r *http.Request) string {
	defaultURL := "rtsp://admin:1q2w3e4r5t6y@192.168.11.131:554/cam/realmonitor?channel=1&subtype=1"

	urls, ok := r.URL.Query()["url"]

	if !ok {
		log.Println("Url Param 'key' is missing")
		return defaultURL
	}

	if len(urls[0]) > 0 {
		url := strings.Replace(urls[0], "+", "?", -1)
		return strings.Replace(url, "$", "&", -1)
	}

	return defaultURL
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

	v1 := r.PathPrefix("/v1").Subrouter()

	v1.Handle("/js/{{s+}.js}", http.StripPrefix("/v1/js/", http.FileServer(http.Dir("www/js"))))
	v1.Handle("/css/{{s+}.css}", http.StripPrefix("/v1/css/", http.FileServer(http.Dir("www/css"))))

	v1.Handle("/", indexHandler)
	v1.Handle("/get-token", auth.GetTokenHandler)
	v1.Handle("/autosearch", autoSearchHandler)
	v1.Handle("/manualsearch", manualSearchHandler)
	v1.Handle("/videostream", appHandler{ac, videoStreamHandler})

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

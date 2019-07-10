package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"

	"videoSecurity/models"

	"github.com/gin-gonic/gin"
)

func middlewareResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/v1/videostream" {
			return
		}

		var wb *responseBuffer
		if w, ok := c.Writer.(gin.ResponseWriter); ok {
			wb = newResponseBuffer(w)
			c.Writer = wb
			c.Next()
		} else {
			c.Next()
			return
		}

		if wb.Response.Status() == 404 {
			c.AbortWithError(404, fmt.Errorf("Page not found"))
		}

		if wb.Header().Get("Content-Type") == "" || strings.Index(wb.Header().Get("Content-Type"), "/json") >= 0 {
			status := wb.status
			data := wb.Body.Bytes()
			wb.Body.Reset()

			errorMessage := ""
			if c.Errors != nil {
				for _, e := range c.Errors {
					errorMessage = e.Error()
					break
				}
			}

			resp := &models.WrapperResponse{
				IsError:      len(errorMessage) > 0 || status != 200,
				ErrorMessage: errorMessage,
				Data:         string(data),
			}

			body, _ := json.Marshal(resp)

			wb.Body.Write(body)
		}

		wb.Flush()
	}
}

type responseBuffer struct {
	Response gin.ResponseWriter // the actual ResponseWriter to flush to
	status   int                // the HTTP response code from WriteHeader
	Body     *bytes.Buffer      // the response content body
	Flushed  bool
}

func (w *responseBuffer) Header() http.Header {
	return w.Response.Header() // use the actual response header
}

func (w *responseBuffer) Write(buf []byte) (int, error) {
	w.Body.Write(buf)
	return len(buf), nil
}

func (w *responseBuffer) WriteString(s string) (n int, err error) {
	n, err = w.Write([]byte(s))
	return
}

func (w *responseBuffer) Written() bool {
	return w.Body.Len() != noWritten
}

func (w *responseBuffer) WriteHeader(status int) {
	w.status = status
}

func (w *responseBuffer) WriteHeaderNow() {

}

func (w *responseBuffer) Status() int {
	return w.status
}

func (w *responseBuffer) Size() int {
	return w.Body.Len()
}

func (w *responseBuffer) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.Response.(http.Hijacker).Hijack()
}

func (w *responseBuffer) CloseNotify() <-chan bool {
	return w.Response.(http.CloseNotifier).CloseNotify()
}

// Fake Flush
// TBD
func (w *responseBuffer) Flush() {
	w.realFlush()
}

func (w *responseBuffer) realFlush() {
	if w.Flushed {
		return
	}
	w.Response.WriteHeader(w.status)
	if w.Body.Len() > 0 {
		bytes := w.Body.Bytes()
		_, err := w.Response.Write(bytes)
		if err != nil {
			panic(err)
		}
		w.Body.Reset()
		bytes = nil
	}
	w.Flushed = true
}

func (w *responseBuffer) Pusher() http.Pusher {
	return w.Pusher()
}

const (
	noWritten     = -1
	defaultStatus = 200
)

func newResponseBuffer(w gin.ResponseWriter) *responseBuffer {
	return &responseBuffer{
		Response: w, status: defaultStatus, Body: &bytes.Buffer{},
	}
}

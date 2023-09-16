package controller

import (
	"bytes"
	"net/http"
)

func bufWriteTo(buf *bytes.Buffer, w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	if buf == nil {
		w.Write([]byte("Response Error"))
	}
	if _, err := buf.WriteTo(w); err != nil {
		w.Write([]byte(err.Error()))
	}
}

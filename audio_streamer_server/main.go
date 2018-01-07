package main

import (
	"encoding/binary"
	"github.com/gordonklaus/portaudio"
	"net/http"
)

func main() {

	portaudio.Initialize()
	defer portaudio.Terminate()
	buffer := make([]float32, 44100)
	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(buffer), func(in []float32) {
		for i := range buffer {
			buffer[i] = in[i]
		}
	})
	chk(err)

	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		chk(stream.Start())
		defer stream.Close()
	})
	http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		chk(stream.Stop())
	})
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			panic("expected http.ResponseWriter to be an http.Flusher")
		}

		w.Header().Set("Connection", "Keep-Alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Transfer-Encoding", "chunked")
		w.Header().Set("Content-Type", "audio/wave")
		for true {
			if len(buffer) == len(buffer) {
				binary.Write(w, binary.BigEndian, &buffer)
				flusher.Flush() // Trigger "chunked" encoding and send a chunk...
				return
			}
		}
	})

	http.ListenAndServe(":8080", nil)
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}

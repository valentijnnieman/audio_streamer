package main

import (
	"github.com/gordonklaus/portaudio"
	//"time"
  "net/http"
  "encoding/binary"
  "fmt"
)

func main() {

	portaudio.Initialize()
	defer portaudio.Terminate()
  buffer := make([]float32, 44100)
  stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(buffer), func(in []float32) {
    //out = in
    //buffer = in
    for i := range buffer{
      buffer[i] = in[i]
    }
  })
  chk(err)
  chk(stream.Start())
  //chk(stream.Stop())
  defer stream.Close()

  http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Transfer-Encoding", "chunked")
    w.Header().Set("Content-Type", "audio/wave")
    for true {
      if(len(buffer) == len(buffer)) {
        fmt.Println("Sending binary!")
        fmt.Println(buffer)
        binary.Write(w, binary.BigEndian, &buffer)
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

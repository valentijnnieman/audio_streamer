package main

import (
	"github.com/gordonklaus/portaudio"
  "net/http"
  "fmt"
  "io/ioutil"
  "time"
  "encoding/binary"
  "bytes"
)

func main() {
	portaudio.Initialize()
	defer portaudio.Terminate()
  buffer := make([]float32, 44100)

  stream, err := portaudio.OpenDefaultStream(0, 1, 44100, len(buffer), func(out []float32) {
    resp, err := http.Get("http://localhost:8080/test")
    chk(err)
    body, _ := ioutil.ReadAll(resp.Body)
    responseReader := bytes.NewReader(body)
    binary.Read(responseReader, binary.BigEndian, &buffer)
    fmt.Println("buffer: ", buffer)
    for i := range out{
      out[i] = buffer[i]
    }
  })
  chk(err)
  chk(stream.Start())
  time.Sleep(time.Second * 40)
  chk(stream.Stop())
  defer stream.Close()

  if err != nil {
    fmt.Println(err)
  }

}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}

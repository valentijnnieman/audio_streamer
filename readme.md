# This is a audio streaming server and client example written in Golang

It uses Portaudio under the hood to get audio I/O (and should therefore work on most platforms) and sends/receives via HTTP. This is an example I wrote for [this article](https://medium.com/@valentijnnieman_79984/how-to-build-an-audio-streaming-server-in-go-part-1-1676eed93021).

# How to run

Define module. It can be your github repo name. Something like, github.com/prisar/audio_streamer_server


    go mod init github.com/prisar/audio_streamer_server

then, do tidy

    go mod tidy

After that run

    go run main.go

Similarly, you run the client.

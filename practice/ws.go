package main

import (
  "log"
  "os"
  "os/signal"
  "net/url"
  "time"
  "github.com/gorilla/websocket"
)

//var addr =  flag.String("addr", "", "http service address")

const wsurl = "wss://stream.binance.com:9443"

func main() {
  interrupt := make(chan os.Signal, 1)
  signal.Notify(interrupt, os.Interrupt)

  u := url.URL{Scheme: "wss", Host: "stream.binance.com:9443", Path: "/ws/!ticker@arr"}
  log.Printf("connecting to %s", u.String())

  c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
  if err != nil {
    log.Fatal("dial:", err)
  }
  defer c.Close() //close the connection at the end of the function

  done := make(chan struct{})

  go func() {
    defer close(done) //something with that channel up there
    for {
      _, message, err := c.ReadMessage()
      if err != nil {
        log.Println("read:", err)
        return
      }
      log.Printf("recv: %s", message)
    }
  }()

  ticker := time.NewTicker(time.Second) //no clue
  defer ticker.Stop()

  for {
    select {
    case <- done:
      return
    case t:= <-ticker.C:
      err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
      if err != nil {
        log.Println("write:", err)
      }
    case <-interrupt:
      log.Println("interrupt")

      // os sent interrupt so..
      // cleanly close the connection by sending a close message and then
      // waiting (with timeout) for the server to close the connection.
      err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
      if err != nil {
        log.Println("write close:", err)
        return
      }
      select {
      case <-done:
      case <-time.After(time.Second):
      }
      return
    }
  }
}

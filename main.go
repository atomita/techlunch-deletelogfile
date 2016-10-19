package main

import (
  "log"
  "./app"
)


func main() {
  app.Init(false)
  
  log.Printf("trace: this is a trace log.")
  log.Printf("debug: this is a debug log.")
  log.Printf("info: this is an info log.")
  log.Printf("warn: this is a warning log.")
  log.Printf("error: this is an error log.")
  log.Printf("alert: this is an alert log.")

  log.Printf("this is a default level log.")
}

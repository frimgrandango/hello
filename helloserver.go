package main

import (
  "net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
  message := "<html><h1>Hello Internet</h1></html> "

  w.Write( []byte(message))
}

func sayHola(w http.ResponseWriter, r *http.Request) {
  message := "<html><h1>Hola Muchachos</h1></html> "

  w.Write( []byte(message))
}

func main() {
  http.HandleFunc("/hello", sayHello)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic("El web has encoutred an error - No Bueno")
  }
  http.HandleFunc("/hola", sayHola)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic("El web has encoutred an error - No Bueno")
  }
}

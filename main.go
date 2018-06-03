package main

import (
  "net/http"
  "encoding/json"
  "strconv"
  "./jankenpon"
)

func choose (w http.ResponseWriter, r *http.Request) {
  player, _ := r.URL.Query()["player"]
  shape, _ := r.URL.Query()["shape"]

  playerInt, _ := strconv.Atoi(player[0])
  shapeInt, _ := strconv.Atoi(shape[0])

  jankenpon.Choose(playerInt, shapeInt)
  w.WriteHeader(200)
}

func reset (w http.ResponseWriter, r *http.Request) {
  jankenpon.Reset()

  w.WriteHeader(200)
}

func status (w http.ResponseWriter, r *http.Request) {
  status := jankenpon.Status()
  js, err := json.Marshal(status)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}

func main () {
  http.HandleFunc("/choose", choose)
  http.HandleFunc("/reset", reset)
  http.HandleFunc("/status", status)

  if err := http.ListenAndServe(":3000", nil); err != nil {
    panic(err)
  }
}

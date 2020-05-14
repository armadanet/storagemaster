package storagemaster

import (
  "github.com/gorilla/mux"
  "net/http"
  "log"
  "encoding/json"
)

type FileLocation struct {
  FileName  string
  Location  string
}

type Handler struct {
  Files map[string]FileLocation
}

func get(handler *Handler) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    filename := vars["filename"]

    log.Println("Showing:", filename)
    log.Println("Location:", handler.Files[filename])
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(handler.Files[filename]); err != nil {
      panic(err)
    }
  }
}

func put(handler *Handler) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    filename := vars["filename"]
    filelocation := vars["filelocation"]
    handler.Files[filename] = FileLocation{
      FileName: filename,
      Location: filelocation,
    }
    log.Println("Added:", filename)
    log.Println("To:", filelocation)
    w.WriteHeader(http.StatusOK)
  }
}

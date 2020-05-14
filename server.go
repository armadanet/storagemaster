package storagemaster

import (
  "github.com/gorilla/mux"
  "github.com/phayes/freeport"
  "net/http"
  "log"
  "strconv"
)

// Server for the Nebula Spinner
type Server interface {
  // Given a port of 0, assigns a free port to the server.
  Run(port int)
}

type server struct {
  router    *mux.Router
  handler   *Handler
}

// Produces a new Server interface of struct server
func New() Server {
  handler := &Handler{
    Files: make(map[string]FileLocation),
  }
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/put/{filename}/{filelocation}", put(handler)).Name("Put")
  router.HandleFunc("/get/{filename}", get(handler)).Name("Get")
  return &server{
    router: router,
    handler: handler,
  }
}

// Runs the spinner server.
// If given a port value of 0, then finds a free port.
func (s *server) Run(port int) {
  var err error
  if port == 0 {
    port, err = freeport.GetFreePort()
    if err != nil {log.Println(err); return}
  }
  log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), s.router))
}

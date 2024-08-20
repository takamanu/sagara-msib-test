package infrastructures

import (
	"github.com/rs/cors"
	"net/http"
	"sagara-msib-test/pkg/grace"
)

type BajuHandler interface {
	HandleClient(w http.ResponseWriter, r *http.Request)
}

type Server struct {
	server       *http.Server
	BajuDelivery BajuHandler
}

func (s *Server) Serve(port string) (err error) {
	handler := cors.AllowAll().Handler(s.InitiateRouter())
	return grace.Serve(port, handler)
}

package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jnewmano/public-notices/internal/checker"
	"github.com/jnewmano/public-notices/internal/processor"
)

type Server struct {
	checker   *checker.Checker
	processor *processor.Processor
}

func New(addr string, checker *checker.Checker, p *processor.Processor) error {

	s := Server{
		checker:   checker,
		processor: p,
	}

	http.HandleFunc("/checkTarget", s.CheckTargetHandler)
	http.HandleFunc("/upcomingActions", s.UpcomingActionHandler)

	http.Handle("/", http.FileServer(http.Dir("./www")))

	fmt.Println("Starting HTTP server on", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		return err
	}

	return nil

}

func (s *Server) CheckTargetHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := s.checkTarget(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func (s *Server) checkTarget(ctx context.Context) error {

	fmt.Println("checking target")
	_, err := s.checker.Do(ctx, "", "")
	if err != nil {
		return err
	}

	return nil
}

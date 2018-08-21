package service

func NewService() Service {
	return Service{make(chan bool), make(chan bool)}
}

// Service is a struct containing channels to stop or reload the blog
type Service struct {
	Stop   chan bool
	Reload chan bool
}

// DoStop adds "true" to the Stop channel
func (s *Service) DoStop() {
	s.Stop <- true
}

// DoReload adds "true" to the Reload channel
func (s *Service) DoReload() {
	s.Reload <- true
}

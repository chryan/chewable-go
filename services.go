package cbl

type Services struct {
	services map[string]interface{}
}

func NewServices() *Services {
	return &Services{
		services: make(map[string]interface{}),
	}
}

func (s *Services) Add(name string, service interface{}) bool {
	if _, ok := s.services[name]; !ok {
		s.services[name] = service
		return true
	}
	return false
}

func (s *Services) Del(name string) bool {
	if _, ok := s.services[name]; ok {
		delete(s.services, name)
		return true
	}
	return false
}

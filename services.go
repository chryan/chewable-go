package cbl

type Services struct {
	services map[string]interface{}
}

func NewServices() *Services {
	return &Services{
		services: make(map[string]interface{}),
	}
}

// Adds a new service to our service map.
func (s *Services) Add(name string, service interface{}) bool {
	if _, ok := s.services[name]; !ok {
		s.services[name] = service
		return true
	}
	return false
}

// Finds and returns a service.
func (s *Services) Get(name string) interface{} {
	if service, ok := s.services[name]; ok {
		return service
	}
	return nil
}

// Deletes a service and returns a pointer to the service if successful.
func (s *Services) Del(name string) interface{} {
	if service, ok := s.services[name]; ok {
		delete(s.services, name)
		return service
	}
	return nil
}

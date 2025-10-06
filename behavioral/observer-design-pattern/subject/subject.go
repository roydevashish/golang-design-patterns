package subject

import "github.com/roydevashish/golang-design-patterns/observer/observer"

type Subject struct {
	obs []observer.Observer
}

func NewSubject() *Subject {
	return &Subject{}
}

func (s *Subject) AddObserver(obs observer.Observer) {
	s.obs = append(s.obs, obs)
}

func (s *Subject) RemoveObserver(obs observer.Observer) {
	// remove the particular observer
}

func (s *Subject) NotifyObserver() {
	for _, v := range s.obs {
		v.Update()
	}
}

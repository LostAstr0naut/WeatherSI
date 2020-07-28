package location

import "errors"

// Service represents a location interface.
type Service interface {
	GetCoordinates(name string) (*Location, error)
}

type service struct {
	locations []Location
}

// New returns a new instance of location service.
func New() Service {
	instance := &service{}
	instance.ensureLocations()
	return instance
}

// Location represents a struct with data of location.
type Location struct {
	Name string
	X    int32
	Y    int32
}

func (s *service) GetCoordinates(name string) (*Location, error) {
	for _, location := range s.locations {
		if location.Name == name {
			return &location, nil
		}
	}

	return nil, errors.New("location not found")
}

func (s *service) ensureLocations() {
	s.locations = append(s.locations, createLocation("BO", 224, 290))
	s.locations = append(s.locations, createLocation("GO", 237, 374))
	s.locations = append(s.locations, createLocation("KP", 249, 464))
	s.locations = append(s.locations, createLocation("ID", 296, 364))
	s.locations = append(s.locations, createLocation("JE", 300, 268))
	s.locations = append(s.locations, createLocation("PO", 325, 417))
	s.locations = append(s.locations, createLocation("KR", 346, 311))
	s.locations = append(s.locations, createLocation("LJ", 373, 350))
	s.locations = append(s.locations, createLocation("KO", 424, 445))
	s.locations = append(s.locations, createLocation("TB", 454, 332))
	s.locations = append(s.locations, createLocation("SG", 459, 252))
	s.locations = append(s.locations, createLocation("NM", 474, 410))
	s.locations = append(s.locations, createLocation("CE", 484, 310))
	s.locations = append(s.locations, createLocation("KK", 521, 374))
	s.locations = append(s.locations, createLocation("MB", 545, 247))
	s.locations = append(s.locations, createLocation("PT", 581, 270))
	s.locations = append(s.locations, createLocation("MS", 625, 216))
}

func createLocation(name string, x, y int32) Location {
	return Location{
		Name: name,
		X:    x,
		Y:    y,
	}
}

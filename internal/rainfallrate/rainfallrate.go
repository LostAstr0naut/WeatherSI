package rainfallrate

import (
	"errors"
	"image/color"
)

// Service represents service interface for rainfall rate levels.
type Service interface {
	GetLevelByRGBA(r, g, b uint16) (Level, error)
}

type service struct {
	Levels []Level
}

// Level represents a rainfall rate level.
type Level struct {
	Color color.RGBA64
	Value float64
}

// New return a new rainfallrate service.
func New() Service {
	intance := &service{}
	intance.ensureLevels()
	return intance
}

// GetLevelByRGBA returns a warning level in rgba searching by rgba.
func (s *service) GetLevelByRGBA(r, g, b uint16) (Level, error) {
	for _, level := range s.Levels {
		if level.Color.R == r && level.Color.G == g && level.Color.B == b {
			return level, nil
		}
	}
	return Level{}, errors.New("level not found")
}

func (s *service) ensureLevels() {
	s.Levels = append(s.Levels, createLevel(2056, 23130, 65278, 0.25))
	s.Levels = append(s.Levels, createLevel(0, 35980, 65278, 0.5))
	s.Levels = append(s.Levels, createLevel(0, 44718, 65021, 0.75))
	s.Levels = append(s.Levels, createLevel(0, 51400, 65278, 1))
	s.Levels = append(s.Levels, createLevel(1028, 55512, 33667, 1.5))
	s.Levels = append(s.Levels, createLevel(16962, 60395, 16962, 2))
	s.Levels = append(s.Levels, createLevel(27756, 63993, 0, 3.5))
	s.Levels = append(s.Levels, createLevel(47288, 64250, 0, 5))
	s.Levels = append(s.Levels, createLevel(63993, 64250, 0, 10))
	s.Levels = append(s.Levels, createLevel(65278, 50886, 0, 15))
	s.Levels = append(s.Levels, createLevel(65278, 33924, 0, 32.5))
	s.Levels = append(s.Levels, createLevel(65535, 15934, 257, 50))
	s.Levels = append(s.Levels, createLevel(54227, 0, 0, 75))
	s.Levels = append(s.Levels, createLevel(46517, 771, 771, 100))
	s.Levels = append(s.Levels, createLevel(52171, 0, 52428, 150))
}

func createLevel(r, g, b uint16, value float64) Level {
	return Level{
		Color: color.RGBA64{
			R: r,
			G: g,
			B: b,
			A: 0, // Disregard alpha.
		},
		Value: value,
	}
}

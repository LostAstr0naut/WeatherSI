package rainfall

import "image/color"

// Mapper represents a mapper interface for rainfall rate levels.
type Mapper interface {
	GetLevelByRGBA(r, g, b uint16) float64
}

type mapper struct {
	Levels []Level
}

// Level represents a rainfall rate level.
type Level struct {
	Color color.RGBA64
	Value float64
}

// New return a new rainfallrate mapper.
func New() Mapper {
	helper := &mapper{}
	helper.ensureLevels()
	return helper
}

// GetLevelByRGBA returns a warning level in rgba searching by rgba.
func (w *mapper) GetLevelByRGBA(r, g, b uint16) float64 {
	for _, item := range w.Levels {
		if item.Color.R == r && item.Color.G == g && item.Color.B == b {
			return item.Value
		}
	}
	return 0
}

func (w *mapper) ensureLevels() {
	w.Levels = append(w.Levels, createLevel(2056, 23130, 65278, 0.25))
	w.Levels = append(w.Levels, createLevel(0, 35980, 65278, 0.5))
	w.Levels = append(w.Levels, createLevel(0, 44718, 65021, 0.75))
	w.Levels = append(w.Levels, createLevel(0, 51400, 65278, 1))
	w.Levels = append(w.Levels, createLevel(1028, 55512, 33667, 1.5))
	w.Levels = append(w.Levels, createLevel(16962, 60395, 16962, 2))
	w.Levels = append(w.Levels, createLevel(27756, 63993, 0, 3.5))
	w.Levels = append(w.Levels, createLevel(47288, 64250, 0, 5))
	w.Levels = append(w.Levels, createLevel(63993, 64250, 0, 10))
	w.Levels = append(w.Levels, createLevel(65278, 50886, 0, 15))
	w.Levels = append(w.Levels, createLevel(65278, 33924, 0, 32.5))
	w.Levels = append(w.Levels, createLevel(65535, 15934, 257, 50))
	w.Levels = append(w.Levels, createLevel(54227, 0, 0, 75))
	w.Levels = append(w.Levels, createLevel(46517, 771, 771, 100))
	w.Levels = append(w.Levels, createLevel(52171, 0, 52428, 150))
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

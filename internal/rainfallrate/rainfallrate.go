package rainfallrate

import (
	"errors"
	"image/color"
)

// Level represents a rainfall rate level.
type Level struct {
	Color       color.RGBA64
	Value       float64
	Description string
}

// LevelByRGBA returns a warning level in rgba.
func LevelByRGBA(r, g, b uint16) (Level, error) {
	for _, level := range colorLevels() {
		if level.Color.R == r && level.Color.G == g && level.Color.B == b {
			return level, nil
		}
	}
	return Level{}, errors.New("level not found")
}

func colorLevels() []Level {
	return []Level{
		{
			Color: color.RGBA64{
				R: 2056,
				G: 23130,
				B: 65278,
				A: 0, // Disregard alpha on all.
			},
			Value:       0.25,
			Description: "Light rainfall.",
		},
		{
			Color: color.RGBA64{
				R: 0,
				G: 35980,
				B: 65278,
				A: 0,
			},
			Value:       0.5,
			Description: "Light rainfall.",
		},
		{
			Color: color.RGBA64{
				R: 0,
				G: 44718,
				B: 65021,
				A: 0,
			},
			Value:       0.75,
			Description: "Light rainfall.",
		},
		{
			Color: color.RGBA64{
				R: 0,
				G: 51400,
				B: 65278,
				A: 0,
			},
			Value:       1,
			Description: "Light rainfall.",
		},
		{
			Color: color.RGBA64{
				R: 1028,
				G: 55512,
				B: 33667,
				A: 0,
			},
			Value:       1.5,
			Description: "Moderate rainfall.",
		},
		{
			Color: color.RGBA64{
				R: 16962,
				G: 60395,
				B: 16962,
				A: 0,
			},
			Value:       2,
			Description: "Moderate rainfall.",
		},
		{
			Color: color.RGBA64{
				R: 27756,
				G: 63993,
				B: 0,
				A: 0,
			},
			Value:       3.5,
			Description: "Moderate rainfall.",
		},
		{
			Color: color.RGBA64{
				R: 47288,
				G: 64250,
				B: 0,
				A: 0,
			},
			Value:       5,
			Description: "Moderate rainfall.",
		},
		{
			Color: color.RGBA64{
				R: 63993,
				G: 64250,
				B: 0,
				A: 0,
			},
			Value:       10,
			Description: "Heavy rainfall.",
		},
		{
			Color: color.RGBA64{
				R: 65278,
				G: 50886,
				B: 0,
				A: 0,
			},
			Value:       15,
			Description: "Heavy rainfall.",
		},
		{
			Color: color.RGBA64{
				R: 65278,
				G: 33924,
				B: 0,
				A: 0,
			},
			Value:       32.5,
			Description: "Heavy rainfall.",
		},
		{
			Color: color.RGBA64{
				R: 65535,
				G: 15934,
				B: 257,
				A: 0,
			},
			Value:       50,
			Description: "Very heavy rainfall.",
		},
		{
			Color: color.RGBA64{
				R: 54227,
				G: 0,
				B: 0,
				A: 0,
			},
			Value:       75,
			Description: "Very heavy rainfall.",
		},
		{
			Color: color.RGBA64{
				R: 46517,
				G: 771,
				B: 771,
				A: 0,
			},
			Value:       100,
			Description: "Very heavy rainfall.",
		},
		{
			Color: color.RGBA64{
				R: 52171,
				G: 0,
				B: 52428,
				A: 0,
			},
			Value:       150,
			Description: "Probable hail.",
		},
	}
}

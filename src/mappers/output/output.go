package output

// Output represents a model for output mapping.
type Output struct {
	AreaLevel     float64
	LocationLevel float64
}

// Mapper represents an output mapper.
type Mapper interface {
	ToOutputModel(areaLevel, locationLevel float64) Output
}

type mapper struct{}

// New returns a new instance of output mapper.
func New() Mapper {
	return mapper{}
}

func (m mapper) ToOutputModel(areaLevel, locationLevel float64) Output {
	return Output{
		AreaLevel:     areaLevel,
		LocationLevel: locationLevel,
	}
}

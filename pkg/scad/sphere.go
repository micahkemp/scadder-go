package scad

import "log"

type Sphere struct {
	Name             string
	Diameter, Radius float64
}

func (s Sphere) SCADWriter() SCADWriter {
	var f fields

	if s.Radius != 0 && s.Diameter != 0 {
		log.Fatalf("Sphere (%s) with radius (%f) and diameter (%f) disallowed", s.Name, s.Radius, s.Diameter)
	} else if s.Radius != 0 {
		f = fields{"r": shortFloat(s.Radius)}
	} else if s.Diameter != 0 {
		f = fields{"d": shortFloat(s.Diameter)}
	} else {
		log.Fatalf("Sphere (%s) with zero radius and diameter disallowed", s.Name)
	}

	return SCADWriter{
		Name: Name{
			Specified: s.Name,
			Default:   "sphere_module",
		},
		templateString: shapeTemplate,
		Function:       "sphere",
		Fields:         f,
	}
}

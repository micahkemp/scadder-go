package scad

import (
	"fmt"
	"log"
	"strconv"
)

type Cylinder struct {
	Name      string
	Height    float64
	Radius    float64
	Radius1   float64
	Radius2   float64
	Diameter  float64
	Diameter1 float64
	Diameter2 float64
	Center    bool
}

func (c Cylinder) setsRadius() bool {
	if c.Radius != 0 || c.Radius1 != 0 || c.Radius2 != 0 {
		return true
	}

	return false
}

func (c Cylinder) validRadius() bool {
	if (c.Radius != 0 && (c.Radius1 == 0 && c.Radius2 == 0)) ||
		(c.Radius == 0 && (c.Radius1 != 0 && c.Radius2 != 0)) {
		return true
	}

	return false
}

func (c Cylinder) radiusFields() fields {
	if c.Radius != 0 {
		return fields{
			"r": shortFloat(c.Radius),
		}
	}

	return fields{
		"r1": shortFloat(c.Radius1),
		"r2": shortFloat(c.Radius2),
	}
}

func (c Cylinder) radiusMessage() string {
	return fmt.Sprintf("Radius %f, Radius1 %f, Radius2 %f",
		c.Radius,
		c.Radius1,
		c.Radius2)
}

func (c Cylinder) setsDiameter() bool {
	if c.Diameter != 0 || c.Diameter1 != 0 || c.Diameter2 != 0 {
		return true
	}

	return false
}

func (c Cylinder) validDiameter() bool {
	if (c.Diameter != 0 && (c.Diameter1 == 0 && c.Diameter2 == 0)) ||
		(c.Diameter == 0 && (c.Diameter1 != 0 && c.Diameter2 != 0)) {
		return true
	}

	return false
}

func (c Cylinder) diameterMessage() string {
	return fmt.Sprintf("Diameter %f, Diameter1 %f, Diameter2 %f",
		c.Diameter,
		c.Diameter,
		c.Diameter)
}

func (c Cylinder) diameterFields() fields {
	if c.Diameter != 0 {
		return fields{
			"d": shortFloat(c.Diameter),
		}
	}

	return fields{
		"d1": shortFloat(c.Diameter1),
		"d2": shortFloat(c.Diameter2),
	}
}

func (c Cylinder) fields() fields {
	f := fields{}

	if c.setsRadius() && !c.setsDiameter() {
		if !c.validRadius() {
			log.Fatalf("Invalid Radius: %s", c.radiusMessage())
		}
		for name, value := range c.radiusFields() {
			f[name] = value
		}
	} else if c.setsDiameter() && !c.setsRadius() {
		if !c.validDiameter() {
			log.Fatalf("Invalid Diameter: %s", c.diameterMessage())
		}
		for name, value := range c.diameterFields() {
			f[name] = value
		}
	} else {
		log.Fatalf("Invalid Radius/Diameter:\n%s\n%s", c.radiusMessage(), c.diameterMessage())
	}

	f["h"] = shortFloat(c.Height)
	f["center"] = strconv.FormatBool(c.Center)

	return f
}

func (c Cylinder) SCADWriter() SCADWriter {
	return SCADWriter{
		Name: Name{
			Specified: c.Name,
			Default:   "cylinder_component",
		},
		templateString: shapeTemplate,
		Function:       "cylinder",
		Fields:         c.fields(),
	}
}

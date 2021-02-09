package tz

import (
	"encoding/json"
)

// FeatureCollection ...
type FeatureCollection struct {
	featureCollection
}

type featureCollection struct {
	Features []*Feature
}

//Feature ...
type Feature struct {
	feature
}

type feature struct {
	Geometry   Geometry
	Properties struct {
		Tzid string
	}
}

//Geometry ...
type Geometry struct {
	geometry
}

type geometry struct {
	Coordinates   [][]Point
	BoundingBoxes [][]Point
}

var jPolyType struct {
	Type       string
	Geometries []*Geometry
}

var jPolygon struct {
	Coordinates [][][]float64
}

var jMultiPolygon struct {
	Coordinates [][][][]float64
}

//UnmarshalJSON ...
func (g *Geometry) UnmarshalJSON(data []byte) (err error) {
	if err := json.Unmarshal(data, &jPolyType); err != nil {
		return err
	}

	if jPolyType.Type == "Polygon" {
		if err := json.Unmarshal(data, &jPolygon); err != nil {
			return err
		}
		pol := make([]Point, len(jPolygon.Coordinates[0]))
		for i, v := range jPolygon.Coordinates[0] {
			pol[i].Lon = v[0]
			pol[i].Lat = v[1]
		}
		b := getBoundingBox(pol)
		g.BoundingBoxes = append(g.BoundingBoxes, b)
		g.Coordinates = append(g.Coordinates, pol)
		return nil
	}

	if jPolyType.Type == "MultiPolygon" {
		if err := json.Unmarshal(data, &jMultiPolygon); err != nil {
			return err
		}
		g.BoundingBoxes = make([][]Point, len(jMultiPolygon.Coordinates))
		g.Coordinates = make([][]Point, len(jMultiPolygon.Coordinates))
		for j, poly := range jMultiPolygon.Coordinates {
			pol := make([]Point, len(poly[0]))
			for i, v := range poly[0] {
				pol[i].Lon = v[0]
				pol[i].Lat = v[1]
			}
			b := getBoundingBox(pol)
			g.BoundingBoxes[j] = b
			g.Coordinates[j] = pol
		}
		return nil
	}
	return nil
}

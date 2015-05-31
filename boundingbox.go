// Copyright 2015 Simon HEGE. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package geographic

//BoundingBox represents a bounding box in WGS84
type BoundingBox struct {
	LatitudeMinDeg  float64
	LongitudeMinDeg float64
	LatitudeMaxDeg  float64
	LongitudeMaxDeg float64
}

//NewBoundingBox creates a new BoundingBox containing all the given points
func NewBoundingBox(points ...Point) BoundingBox {
	var b BoundingBox
	if len(points) == 0 {
		return b
	}

	b.LatitudeMinDeg = points[0].LatitudeDeg
	b.LongitudeMinDeg = points[0].LongitudeDeg
	b.LatitudeMaxDeg = points[0].LatitudeDeg
	b.LongitudeMaxDeg = points[0].LongitudeDeg

	for i := 1; i < len(points); i++ {
		if points[i].LatitudeDeg < b.LatitudeMinDeg {
			b.LatitudeMinDeg = points[i].LatitudeDeg
		}
		if points[i].LatitudeDeg > b.LatitudeMaxDeg {
			b.LatitudeMaxDeg = points[i].LatitudeDeg
		}
		if points[i].LongitudeDeg < b.LongitudeMinDeg {
			b.LongitudeMinDeg = points[i].LongitudeDeg
		}
		if points[i].LongitudeDeg > b.LongitudeMinDeg {
			b.LongitudeMaxDeg = points[i].LongitudeDeg
		}
	}

	return b
}

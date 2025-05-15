// In a new file called extended_types.go in the appropriate package
package flightroutetrajectory

import (
	"encoding/xml"
	"strconv"

	"github.com/edancain/FIXMSchema.git/generated/fixm/base"
)

// ExtendedRouteTrajectoryElementType extends the generated type
type ExtendedRouteTrajectoryElementType struct {
	*RouteTrajectoryElementType     // Embed the original type
	SeqNum                      int // Add sequence number
}

// ExtendedDistanceType extends the base DistanceType
type ExtendedDistanceType struct {
	*base.DistanceType        // Embed the original type
	UOM                string // Add UOM field
}

// Create custom unmarshalers for the extended types
func (e *ExtendedRouteTrajectoryElementType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Get seqNum attribute
	for _, attr := range start.Attr {
		if attr.Name.Local == "seqNum" {
			if val, err := strconv.Atoi(attr.Value); err == nil {
				e.SeqNum = val
			}
		}
	}

	// Initialize the embedded type if it's nil
	if e.RouteTrajectoryElementType == nil {
		e.RouteTrajectoryElementType = &RouteTrajectoryElementType{}
	}

	// Continue with normal unmarshaling
	return d.DecodeElement(e.RouteTrajectoryElementType, &start)
}

// Similar for distance type
func (e *ExtendedDistanceType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Get UOM attribute
	for _, attr := range start.Attr {
		if attr.Name.Local == "uom" {
			e.UOM = attr.Value
		}
	}

	// Initialize the embedded type if it's nil
	if e.DistanceType == nil {
		e.DistanceType = &base.DistanceType{}
	}

	// Continue with normal unmarshaling
	return d.DecodeElement(e.DistanceType, &start)
}

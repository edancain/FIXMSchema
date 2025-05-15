package base

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// MarshalXML implements the xml.Marshaler interface for LatLongPosType
func (pos LatLongPosType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// Convert the slice of float64 to a space-separated string
	parts := make([]string, len(pos))
	for i, v := range pos {
		parts[i] = fmt.Sprintf("%f", v)
	}
	return e.EncodeElement(strings.Join(parts, " "), start)
}

// UnmarshalXML implements the xml.Unmarshaler interface for LatLongPosType
func (pos *LatLongPosType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	// Split the string by whitespace
	parts := strings.Fields(s)
	result := make([]float64, len(parts))

	for i, part := range parts {
		// Parse each part as a float64
		val, err := strconv.ParseFloat(part, 64)
		if err != nil {
			return err
		}
		result[i] = val
	}

	*pos = result
	return nil
}

// UnmarshalXML implements the xml.Unmarshaler interface for DurationType
func (dt *DurationType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	// Parse ISO 8601 duration
	duration, err := parseDuration(s)
	if err != nil {
		return err
	}

	*dt = DurationType(duration)
	return nil
}

// parseDuration parses ISO 8601 duration strings (e.g., P0Y0M0DT0H2M35S)
func parseDuration(durationStr string) (time.Duration, error) {
	// Regular expression to extract components from the duration string
	re := regexp.MustCompile(`P(?:(\d+)Y)?(?:(\d+)M)?(?:(\d+)D)?T(?:(\d+)H)?(?:(\d+)M)?(?:(\d+)S)?`)
	matches := re.FindStringSubmatch(durationStr)

	if len(matches) == 0 {
		return 0, fmt.Errorf("invalid duration format: %s", durationStr)
	}

	// Initialize duration components
	var years, months, days, hours, minutes, seconds int64

	// Extract values from regex matches (skip the first match which is the entire string)
	if len(matches) > 1 && matches[1] != "" {
		years, _ = strconv.ParseInt(matches[1], 10, 64)
	}
	if len(matches) > 2 && matches[2] != "" {
		months, _ = strconv.ParseInt(matches[2], 10, 64)
	}
	if len(matches) > 3 && matches[3] != "" {
		days, _ = strconv.ParseInt(matches[3], 10, 64)
	}
	if len(matches) > 4 && matches[4] != "" {
		hours, _ = strconv.ParseInt(matches[4], 10, 64)
	}
	if len(matches) > 5 && matches[5] != "" {
		minutes, _ = strconv.ParseInt(matches[5], 10, 64)
	}
	if len(matches) > 6 && matches[6] != "" {
		seconds, _ = strconv.ParseInt(matches[6], 10, 64)
	}

	// Convert to time.Duration (approximating years and months)
	durationNano := time.Duration(years)*365*24*time.Hour +
		time.Duration(months)*30*24*time.Hour +
		time.Duration(days)*24*time.Hour +
		time.Duration(hours)*time.Hour +
		time.Duration(minutes)*time.Minute +
		time.Duration(seconds)*time.Second

	return durationNano, nil
}

// LogAttributes logs the attributes of an XML start element for debugging
func LogAttributes(start xml.StartElement) {
	for _, attr := range start.Attr {
		fmt.Printf("Attribute: %s=%s\n", attr.Name.Local, attr.Value)
	}
}

// GetSeqNumFromXML is a helper function to extract sequence numbers from XML elements
func GetSeqNumFromXML(start xml.StartElement) (int, bool) {
	for _, attr := range start.Attr {
		if attr.Name.Local == "seqNum" {
			seqNum, err := strconv.Atoi(attr.Value)
			if err == nil {
				return seqNum, true
			}
			break
		}
	}
	return 0, false
}

// GetUomFromXML is a helper function to extract UOM attributes from XML elements
func GetUomFromXML(start xml.StartElement) (string, bool) {
	for _, attr := range start.Attr {
		if attr.Name.Local == "uom" {
			return attr.Value, true
		}
	}
	return "", false
}

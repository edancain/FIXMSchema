package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/edancain/FIXMSchema/generated/fixm/base"
	"github.com/edancain/FIXMSchema/generated/fixm/flight/flightdata"
	"github.com/edancain/FIXMSchema/generated/fixm/flight/flightroutetrajectory"
)

// Document structure to parse different FIXM document types
type FIXMDocument struct {
	XMLName              xml.Name                                      `xml:"Document"`
	FlightElement        *flightdata.FlightType                        `xml:"flight"`
	RouteTrajectoryGroup *flightdata.RouteTrajectoryGroupContainerType `xml:"routeTrajectoryGroup"`
}

func main() {
	// Parse command line arguments
	verbose := flag.Bool("v", false, "Enable verbose output")
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Usage: fixmparser [-v] <file.xml> OR fixmparser [-v] <directory>")
		listSampleFiles()
		os.Exit(1)
	}

	// Check if a file or directory was provided
	path := args[0]
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Fatalf("Error accessing path %s: %v", path, err)
	}

	if fileInfo.IsDir() {
		// Process all XML files in directory
		processDirectory(path, *verbose)
	} else {
		// Process single file
		processFile(path, *verbose)
	}
}

func listSampleFiles() {
	fmt.Println("\nAvailable sample files:")

	// Walk sample_messages directory if it exists
	samplesDir := "sample_messages"
	if _, err := os.Stat(samplesDir); err == nil {
		filepath.Walk(samplesDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".xml") {
				fmt.Printf("  %s\n", path)
			}
			return nil
		})
	} else {
		fmt.Println("  No sample files found in ./sample_messages directory")
	}
}

func processDirectory(dirPath string, verbose bool) {
	fmt.Printf("Processing directory: %s\n", dirPath)

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(strings.ToLower(entry.Name()), ".xml") {
			filePath := filepath.Join(dirPath, entry.Name())
			fmt.Printf("\n=== Processing: %s ===\n", entry.Name())
			processFile(filePath, verbose)
		}
	}
}

func processFile(filePath string, verbose bool) {
	// Read the file
	xmlData, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Print file size info if verbose
	if verbose {
		fmt.Printf("File size: %d bytes\n", len(xmlData))
	}

	// First try to determine if it's a complete document or just a fragment
	xmlStr := string(xmlData)
	isFragment := !strings.Contains(xmlStr, "<?xml")

	if isFragment {
		if verbose {
			fmt.Println("Detected XML fragment, adding XML wrapper...")
		}
		xmlStr = addXMLWrapper(xmlStr)
		xmlData = []byte(xmlStr)
	}

	// Try to parse as a standard document
	var doc FIXMDocument
	err = xml.Unmarshal(xmlData, &doc)
	if err != nil {
		log.Printf("Error parsing document: %v\n", err)

		// For a fragment, try more specific parsing
		if isFragment {
			if verbose {
				fmt.Println("Attempting specialized fragment parsing...")
			}
			parseFragment(xmlData, verbose)
		}
		return
	}

	// Process based on what we found
	if doc.FlightElement != nil {
		fmt.Println("Found Flight element in document")
		printFlightInfo(doc.FlightElement)
		return
	}

	if doc.RouteTrajectoryGroup != nil {
		fmt.Println("Found RouteTrajectoryGroup element in document")
		printRouteTrajectoryGroupInfo(doc.RouteTrajectoryGroup, nil)
		return
	}

	fmt.Println("No recognized FIXM elements found at the root level")
}

func parseFragment(xmlData []byte, verbose bool) {
	xmlStr := string(xmlData)

	// Try to parse the route trajectory group
	if strings.Contains(xmlStr, "<fx:routeTrajectoryGroup") ||
		strings.Contains(xmlStr, "<routeTrajectoryGroup") {
		wrappedXML := `<?xml version="1.0" encoding="UTF-8"?>
<root xmlns:fx="http://www.fixm.aero/flight/4.3" 
      xmlns:fb="http://www.fixm.aero/base/4.3">` + xmlStr + `</root>`

		// Modified wrapper struct to use extended types
		var wrapper struct {
			XMLName              xml.Name `xml:"root"`
			RouteTrajectoryGroup struct {
				// Use extended element type
				Element []flightroutetrajectory.ExtendedRouteTrajectoryElementType `xml:"element"`
				// Other fields remain the same
				ClimbProfile     *flightroutetrajectory.PerformanceProfileType     `xml:"climbProfile"`
				ClimbSchedule    *flightroutetrajectory.SpeedScheduleType          `xml:"climbSchedule"`
				DescentProfile   *flightroutetrajectory.PerformanceProfileType     `xml:"descentProfile"`
				DescentSchedule  *flightroutetrajectory.SpeedScheduleType          `xml:"descentSchedule"`
				Extension        []base.RouteTrajectoryGroupExtensionType          `xml:"extension"`
				RouteInformation *flightroutetrajectory.FlightRouteInformationType `xml:"routeInformation"`
				TakeoffMass      *base.MassType                                    `xml:"takeoffMass"`
			} `xml:"routeTrajectoryGroup"`
		}

		if err := xml.Unmarshal([]byte(wrappedXML), &wrapper); err == nil {
			fmt.Println("Successfully parsed RouteTrajectoryGroup fragment")

			// Convert back to standard type for compatibility with existing code
			rtg := flightdata.RouteTrajectoryGroupContainerType{
				Desired: &flightroutetrajectory.RouteTrajectoryGroupType{
					ClimbProfile:     wrapper.RouteTrajectoryGroup.ClimbProfile,
					ClimbSchedule:    wrapper.RouteTrajectoryGroup.ClimbSchedule,
					DescentProfile:   wrapper.RouteTrajectoryGroup.DescentProfile,
					DescentSchedule:  wrapper.RouteTrajectoryGroup.DescentSchedule,
					Extension:        wrapper.RouteTrajectoryGroup.Extension,
					RouteInformation: wrapper.RouteTrajectoryGroup.RouteInformation,
					TakeoffMass:      wrapper.RouteTrajectoryGroup.TakeoffMass,
				},
			}

			// Copy elements and their sequence numbers
			elements := make([]flightroutetrajectory.RouteTrajectoryElementType, len(wrapper.RouteTrajectoryGroup.Element))
			extendedElements := make([]flightroutetrajectory.ExtendedRouteTrajectoryElementType, len(wrapper.RouteTrajectoryGroup.Element))
			for i, extElem := range wrapper.RouteTrajectoryGroup.Element {
				elements[i] = *extElem.RouteTrajectoryElementType
				extendedElements[i] = extElem
				if verbose && extElem.SeqNum != 0 {
					fmt.Printf("Element %d has sequence number: %d\n", i, extElem.SeqNum)
				}
			}
			rtg.Desired.Element = elements

			printRouteTrajectoryGroupInfo(&rtg, extendedElements)
			return
		} else if verbose {
			fmt.Printf("Error parsing RouteTrajectoryGroup: %v\n", err)
		}
	}

	fmt.Println("Unable to parse the XML fragment")
}

func addXMLWrapper(xmlFragment string) string {
	return `<?xml version="1.0" encoding="UTF-8"?>
<fixm:Document xmlns:fixm="http://www.fixm.aero/fixm/4.3"
               xmlns:fx="http://www.fixm.aero/flight/4.3"
               xmlns:fb="http://www.fixm.aero/base/4.3"
               xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
` + xmlFragment + `
</fixm:Document>`
}

func printFlightInfo(flight *flightdata.FlightType) {
	fmt.Println("\nFlight Information:")

	// Print flight ID if available
	if flight.FlightIdentification != nil {
		if flight.FlightIdentification.AircraftIdentification != nil {
			fmt.Printf("Aircraft ID: %s\n", *flight.FlightIdentification.AircraftIdentification)
		}

		// Print GUFI if available
		if flight.FlightIdentification.Gufi != nil {
			fmt.Printf("GUFI: %v\n", flight.FlightIdentification.Gufi)
		}
	}

	// Print departure/arrival if available
	if flight.Departure != nil {
		fmt.Println("\nDeparture Information:")

		if flight.Departure.ActualTimeOfDeparture != nil && flight.Departure.ActualTimeOfDeparture.Time != nil {
			fmt.Printf("  Actual Departure Time: %v\n", *flight.Departure.ActualTimeOfDeparture.Time)
		}

		// Only check DepartureAerodrome as per the XSD schema
		if flight.Departure.DepartureAerodrome != nil {
			fmt.Println("  Departure Aerodrome:")
			if flight.Departure.DepartureAerodrome.LocationIndicator != nil {
				fmt.Printf("    Location Indicator: %s\n", *flight.Departure.DepartureAerodrome.LocationIndicator)
			}
			if flight.Departure.DepartureAerodrome.Name != nil {
				fmt.Printf("    Name: %s\n", *flight.Departure.DepartureAerodrome.Name)
			}
			if flight.Departure.DepartureAerodrome.ReferencePoint != nil &&
				flight.Departure.DepartureAerodrome.ReferencePoint.Pos != nil {
				// Access Pos directly without dereferencing it
				fmt.Printf("    Position: %v\n", flight.Departure.DepartureAerodrome.ReferencePoint.Pos)
			}
		}
	}

	if flight.Arrival != nil {
		fmt.Println("\nArrival Information:")
		if flight.Arrival.ActualTimeOfArrival != nil {
			fmt.Printf("  Actual Arrival Time: %v\n", *flight.Arrival.ActualTimeOfArrival)
		}

		if flight.Arrival.ArrivalAerodrome != nil {
			fmt.Println("  Arrival Aerodrome:")
			if flight.Arrival.ArrivalAerodrome.LocationIndicator != nil {
				fmt.Printf("    Location Indicator: %s\n", *flight.Arrival.ArrivalAerodrome.LocationIndicator)
			}
			if flight.Arrival.ArrivalAerodrome.Name != nil {
				fmt.Printf("    Name: %s\n", *flight.Arrival.ArrivalAerodrome.Name)
			}
		}
	}

	// Print route trajectory if available
	if flight.RouteTrajectoryGroup != nil {
		printRouteTrajectoryGroupInfo(flight.RouteTrajectoryGroup, nil)
	}
}

func printRouteTrajectoryGroupInfo(rtg *flightdata.RouteTrajectoryGroupContainerType, extendedElements []flightroutetrajectory.ExtendedRouteTrajectoryElementType) {
	fmt.Println("\nRoute Trajectory Information:")
	fmt.Println("DEBUG: In printRouteTrajectoryGroupInfo")

	// Print desired route trajectory if available
	if rtg.Desired != nil {
		fmt.Println("Desired Route:")
		printRouteTrajectoryGroup(rtg.Desired, extendedElements)
	}
	// ... other code
	fmt.Println("DEBUG: Exiting printRouteTrajectoryGroupInfo")
}

func printRouteTrajectoryGroup(rtg *flightroutetrajectory.RouteTrajectoryGroupType, extendedElements []flightroutetrajectory.ExtendedRouteTrajectoryElementType) {
	if rtg == nil {
		fmt.Println("  No route trajectory information available")
		return
	}

	fmt.Println("  Route Trajectory Details:")

	// Create a map to easily find extended elements by their index
	extendedMap := make(map[int]flightroutetrajectory.ExtendedRouteTrajectoryElementType)
	for i, extElem := range extendedElements {
		if i < len(rtg.Element) {
			extendedMap[i] = extElem
		}
	}

	// Print elements with their enhanced data
	for i, elem := range rtg.Element {
		// Get sequence number if available
		seqNumInfo := ""
		if extElem, ok := extendedMap[i]; ok && extElem.SeqNum != 0 {
			seqNumInfo = fmt.Sprintf(" (seqNum: %d)", extElem.SeqNum)
		}

		fmt.Printf("\n  Element %d%s:\n", i, seqNumInfo)

		// Print along route distance with UOM if available
		if elem.AlongRouteDistance != nil {
			fmt.Printf("    AlongRouteDistance: %v", *elem.AlongRouteDistance)
			// If we had extended distance info, we could add UOM here
			fmt.Println()
		}

		// Print element start point info
		if elem.ElementStartPoint != nil {
			fmt.Println("    Start Point:")

			// Check fields in AerodromeReferencePoint directly instead of comparing struct to nil
			if elem.ElementStartPoint.AerodromeReferencePoint.LocationIndicator != nil {
				fmt.Printf("      Aerodrome: %s\n", *elem.ElementStartPoint.AerodromeReferencePoint.LocationIndicator)

				// Print coordinates if available
				if elem.ElementStartPoint.AerodromeReferencePoint.ReferencePoint != nil &&
					len(elem.ElementStartPoint.AerodromeReferencePoint.ReferencePoint.Pos) >= 2 {
					pos := elem.ElementStartPoint.AerodromeReferencePoint.ReferencePoint.Pos
					fmt.Printf("      Position: Lat %.6f, Lon %.6f\n", pos[0], pos[1])
				}
			}

			// Check fields in Navaid directly
			if elem.ElementStartPoint.Navaid.Designator != "" {
				fmt.Printf("      Navaid: %s\n", elem.ElementStartPoint.Navaid.Designator)

				// Print coordinates if available
				if elem.ElementStartPoint.Navaid.Position != nil &&
					len(elem.ElementStartPoint.Navaid.Position.Pos) >= 2 {
					pos := elem.ElementStartPoint.Navaid.Position.Pos
					fmt.Printf("      Position: Lat %.6f, Lon %.6f\n", pos[0], pos[1])
				}
			}

			// Check fields in DesignatedPoint directly
			if elem.ElementStartPoint.DesignatedPoint.Designator != "" {
				fmt.Printf("      Designated Point: %s\n", elem.ElementStartPoint.DesignatedPoint.Designator)

				// Print coordinates if available
				if elem.ElementStartPoint.DesignatedPoint.Position != nil &&
					len(elem.ElementStartPoint.DesignatedPoint.Position.Pos) >= 2 {
					pos := elem.ElementStartPoint.DesignatedPoint.Position.Pos
					fmt.Printf("      Position: Lat %.6f, Lon %.6f\n", pos[0], pos[1])
				}
			}
		}

		// Print 4D point info
		if elem.Point4D != nil {
			fmt.Println("    4D Point Information:")

			// Print level/altitude with UOM if available
			if elem.Point4D.Level != nil {
				if elem.Point4D.Level.IsAltitudeSet() {
					fmt.Printf("      Altitude: %v", *elem.Point4D.Level.Altitude)
					if elem.Point4D.Level.Altitude.UOM != "" {
						fmt.Printf(" %s", elem.Point4D.Level.Altitude.UOM)
					}
					fmt.Println()
				} else if elem.Point4D.Level.IsFlightLevelSet() {
					fmt.Printf("      Flight Level: FL%v\n", *elem.Point4D.Level.FlightLevel)
				}
			}

			// Print position
			if elem.Point4D.Position != nil &&
				len(elem.Point4D.Position.Pos) >= 2 {
				pos := elem.Point4D.Position.Pos
				fmt.Printf("      Position: Lat %.6f, Lon %.6f\n", pos[0], pos[1])
			}

			// Print time
			if elem.Point4D.Time != nil {
				if elem.Point4D.Time.IsAbsoluteTimeSet() {
					fmt.Printf("      Absolute Time: %s\n", elem.Point4D.Time.AbsoluteTime.Format())
				}
				if elem.Point4D.Time.IsRelativeTimeSet() {
					// Convert from nanoseconds to a more readable format
					duration := time.Duration(elem.Point4D.Time.RelativeTimeFromInitialPredictionPoint)
					fmt.Printf("      Relative Time: %v\n", formatDuration(duration))
				}
			}

			// Print predicted airspeed with UOM
			if elem.Point4D.PredictedAirspeed != nil {
				fmt.Printf("      Predicted Airspeed: %v", *elem.Point4D.PredictedAirspeed)
				if elem.Point4D.PredictedAirspeed.UOM != "" {
					fmt.Printf(" %s", elem.Point4D.PredictedAirspeed.UOM)
				}
				fmt.Println()
			}

			// Print point properties
			if len(elem.Point4D.PointProperty) > 0 {
				fmt.Println("      Point Properties:")
				for _, prop := range elem.Point4D.PointProperty {
					if prop.PropertyType != nil {
						fmt.Printf("        - %s\n", *prop.PropertyType)
						if prop.Description != nil {
							fmt.Printf("          Description: %s\n", *prop.Description)
						}
					}
				}
			}
		}

		// Print route designator
		if elem.RouteDesignatorToNextElement != nil {
			rte := elem.RouteDesignatorToNextElement
			fmt.Println("    Route to Next Element:")

			if rte.IsOtherRouteDesignatorSet() {
				fmt.Printf("      Type: %s\n", *rte.OtherRouteDesignator)
			} else if rte.IsRouteDesignatorSet() {
				fmt.Printf("      Airway: %s\n", rte.RouteDesignator.Value)
			} else if rte.IsStandardInstrumentDepartureSet() {
				fmt.Printf("      SID: %s\n", rte.StandardInstrumentDeparture.Designator)
			} else if rte.IsStandardInstrumentArrivalSet() {
				fmt.Printf("      STAR: %s\n", rte.StandardInstrumentArrival.Designator)
			}
		}

		// Print constraints
		if len(elem.Constraint) > 0 {
			fmt.Printf("    Constraints: %d\n", len(elem.Constraint))
			for j, constraint := range elem.Constraint {
				fmt.Printf("      Constraint %d:\n", j+1)

				// Print level constraint
				if constraint.Level != nil {
					fmt.Println("        Level Constraint:")
					if constraint.Level.Level != nil {
						if constraint.Level.Level.IsAltitudeSet() {
							fmt.Printf("          Altitude: %v\n", *constraint.Level.Level.Altitude)
						} else if constraint.Level.Level.IsFlightLevelSet() {
							fmt.Printf("          Flight Level: FL%v\n", *constraint.Level.Level.FlightLevel)
						}
					}
				}

				// Print speed constraint
				if constraint.Speed != nil {
					fmt.Println("        Speed Constraint:")
					if constraint.Speed.Speed != nil {
						fmt.Printf("          Speed: %v\n", *constraint.Speed.Speed)
					}
				}

				// Print time constraint
				if constraint.Time != nil && constraint.Time.TimeSpecification != nil {
					fmt.Println("        Time Constraint:")

					// Use the IsZero method for TimeSpecification
					if !constraint.Time.TimeSpecification.IsZero() {
						fmt.Printf("          Time: %s\n", constraint.Time.TimeSpecification.Format())
					}
				}
			}
		}
	}

	// Print route information
	if rtg.RouteInformation != nil {
		ri := rtg.RouteInformation
		fmt.Println("\n  Route Information:")

		if ri.RouteText != nil {
			fmt.Printf("    Route Text: %s\n", *ri.RouteText)
		}

		if ri.CruisingLevel != nil {
			fmt.Println("    Cruising Level:")

			if ri.CruisingLevel.IsAltitudeSet() {
				fmt.Printf("      Altitude: %v\n", *ri.CruisingLevel.Altitude)
			} else if ri.CruisingLevel.IsFlightLevelSet() {
				fmt.Printf("      Flight Level: FL%v\n", *ri.CruisingLevel.FlightLevel)
			} else if ri.CruisingLevel.IsVFRSet() {
				fmt.Printf("      VFR: %s\n", ri.CruisingLevel.VisualFlightRules)
			}
		}

		if ri.CruisingSpeed != nil {
			fmt.Printf("    Cruising Speed: %v", *ri.CruisingSpeed)
			if ri.CruisingSpeed.UOM != "" {
				fmt.Printf(" %s", ri.CruisingSpeed.UOM)
			}
			fmt.Println()
		}

		if ri.TotalEstimatedElapsedTime != nil {
			duration := time.Duration(*ri.TotalEstimatedElapsedTime)
			fmt.Printf("    Total EET: %s\n", formatDuration(duration))
		}
	}

	// Print climb/descent profiles
	if rtg.ClimbProfile != nil {
		fmt.Println("\n  Climb Profile:")
		// Print climb profile details...
	}

	if rtg.DescentProfile != nil {
		fmt.Println("\n  Descent Profile:")
		// Print descent profile details...
	}

	// Print takeoff mass
	if rtg.TakeoffMass != nil {
		fmt.Printf("\n  Takeoff Mass: %v", rtg.TakeoffMass.Value)
		if rtg.TakeoffMass.UOM != "" {
			fmt.Printf(" %s", rtg.TakeoffMass.UOM)
		}
		fmt.Println()
	}
}

// Helper function to format duration in a readable way
func formatDuration(d time.Duration) string {
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60

	if hours > 0 {
		return fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
	} else if minutes > 0 {
		return fmt.Sprintf("%dm %ds", minutes, seconds)
	} else {
		return fmt.Sprintf("%ds", seconds)
	}
}

// prettyPrint recursively prints a struct with proper indentation
func prettyPrint(v interface{}, indent string) {
	// Get the value of the interface
	val := reflect.ValueOf(v)

	// Handle pointers - dereference if not nil
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			fmt.Printf("%s<nil>\n", indent)
			return
		}
		val = val.Elem()
	}

	// Handle nil interfaces
	if !val.IsValid() {
		fmt.Printf("%s<nil>\n", indent)
		return
	}

	// Handle different kinds of values
	switch val.Kind() {
	case reflect.Struct:
		fmt.Printf("%s{\n", indent)
		nextIndent := indent + "  "
		typ := val.Type()

		// Print each field
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fieldName := typ.Field(i).Name

			// Skip unexported fields
			if !field.CanInterface() {
				continue
			}

			// Print field name
			fmt.Printf("%s%s: ", nextIndent, fieldName)

			// Print field value recursively
			fieldValue := field.Interface()

			// Handle special cases
			if field.Kind() == reflect.Slice {
				if field.Len() == 0 {
					fmt.Println("[]")
				} else {
					fmt.Printf("[\n")
					sliceIndent := nextIndent + "  "

					// Limit the number of elements to display for large slices
					maxItems := 5
					numItems := field.Len()

					for j := 0; j < numItems; j++ {
						if j >= maxItems && numItems > maxItems+2 {
							fmt.Printf("%s... (%d more items)\n", sliceIndent, numItems-maxItems)
							break
						}

						item := field.Index(j).Interface()
						fmt.Printf("%s", sliceIndent)
						prettyPrint(item, sliceIndent)
					}

					fmt.Printf("%s]\n", nextIndent)
				}
			} else if fieldName == "Pos" && field.Kind() == reflect.Slice {
				// Special handling for coordinate positions
				if field.Len() >= 2 {
					lat := field.Index(0).Float()
					lon := field.Index(1).Float()
					fmt.Printf("[Lat: %f, Lon: %f]\n", lat, lon)
				} else {
					prettyPrint(fieldValue, nextIndent)
				}
			} else {
				prettyPrint(fieldValue, nextIndent)
			}
		}

		fmt.Printf("%s}\n", indent)

	case reflect.Slice:
		if val.Len() == 0 {
			fmt.Println("[]")
		} else {
			fmt.Printf("[\n")
			sliceIndent := indent + "  "

			// Limit the number of elements to display for large slices
			maxItems := 5
			numItems := val.Len()

			for i := 0; i < numItems; i++ {
				if i >= maxItems && numItems > maxItems+2 {
					fmt.Printf("%s... (%d more items)\n", sliceIndent, numItems-maxItems)
					break
				}

				item := val.Index(i).Interface()
				fmt.Printf("%s", sliceIndent)
				prettyPrint(item, sliceIndent)
			}

			fmt.Printf("%s]\n", indent)
		}

	case reflect.Map:
		if val.Len() == 0 {
			fmt.Println("{}")
		} else {
			fmt.Printf("{\n")
			mapIndent := indent + "  "

			for _, key := range val.MapKeys() {
				fmt.Printf("%s%v: ", mapIndent, key.Interface())
				item := val.MapIndex(key).Interface()
				prettyPrint(item, mapIndent)
			}

			fmt.Printf("%s}\n", indent)
		}

	case reflect.String:
		fmt.Printf("\"%s\"\n", val.String())

	default:
		fmt.Printf("%v\n", val.Interface())
	}
}

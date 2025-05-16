FIXM Schema to Go Code Conversion: Issues and Solutions

Summary
This is an explanation to list and detail the issues encountered in the XML Schema to Go code conversion for the FIXM (Flight Information Exchange Model) and provide what I hope is a comprehensive analysis of the solutions I implemented. 

The primary issues included incorrect handling of XML choice elements, poor representation of simple types with attributes, lack of proper helper methods, and struct vs. nil comparison issues.

1. XML Schema Concepts and Their Go Representations
    1.1 Simple Types with Attributes (complexType with simpleContent)

    Problem
    In XML Schema, a "simple type with attributes" is represented as a complexType with simpleContent. This allows a type to have both a text value and attributes. 

    For example:
    xml<xs:complexType name="AltitudeType">
    <xs:simpleContent>
        <xs:extension base="fb:VerticalDistanceType">
        <xs:attribute name="uom" use="required" type="fb:UomAltitudeType" />
        </xs:extension>
    </xs:simpleContent>
    </xs:complexType>

    This defines a type that has a decimal value (from VerticalDistanceType) and a "uom" (unit of measure) attribute.

    The original generated Go code (and what will still be generated using the /generate/main.go file) did not correctly represent this pattern:

    type AltitudeType struct {
        UOM *UomAltitudeType `xml:"uom"`// Unit of measure
    }

    This representation lost the actual value of the type, keeping only the attribute.

    Solution
    I fixed this by adding a Value field to capture the element's text content:

    type AltitudeType struct {
        Value decimal.Decimal `xml:",chardata"`
        UOM   UomAltitudeType `xml:"uom,attr"`
    }

    The xml:",chardata" tag instructs the XML marshaler to map the element's text content to this field. The xml:"uom,attr" tag correctly identifies UOM as an attribute rather than a child element.

    I also added custom XML marshalers/unmarshalers to handle the conversion between XML and Go struct:

    func (a *AltitudeType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
        var s struct {
            Value decimal.Decimal `xml:",chardata"`
            UOM   UomAltitudeType `xml:"uom,attr"`
        }
        
        if err := d.DecodeElement(&s, &start); err != nil {
            return err
        }
        
        a.Value = s.Value
        a.UOM = s.UOM
        return nil
    }

    This pattern I applied to several types, including FlightLevelType, IndicatedAirspeedType, TrueAirspeedType, and MassType.

    1.2 Choice Elements (xs:choice)
    Problem
    XML Schema uses the xs:choice element to indicate that only one of several possible child elements should be present. 

    For example:
    xml<xs:complexType name="FlightLevelOrAltitudeChoiceType">
    <xs:choice>
        <xs:element name="altitude" type="fb:AltitudeType" minOccurs="1" maxOccurs="1" />
        <xs:element name="flightLevel" type="fb:FlightLevelType" minOccurs="1" maxOccurs="1" />
    </xs:choice>
    </xs:complexType>

    This indicates that either an altitude or a flightLevel element should be present, but not both.

    The generated Go code represented this as a struct with fields for all possible choices:

    type FlightLevelOrAltitudeChoiceType struct {
        Altitude    *AltitudeType    `xml:"altitude,omitempty"`
        FlightLevel *FlightLevelType `xml:"flightLevel,omitempty"`
    }

    While this structure is capable of representing the choice, it lacked the helper methods to determine which choice was actually set in the XML, so the choice was never determined.

    Solution
    I added helper methods to check which choice is set:

    func (c *FlightLevelOrAltitudeChoiceType) IsAltitudeSet() bool {
        return c != nil && c.Altitude != nil
    }

    func (c *FlightLevelOrAltitudeChoiceType) IsFlightLevelSet() bool {
        return c != nil && c.FlightLevel != nil
    }

    I also added more sophisticated helper methods for types like TimeChoiceType:

    func (t *TimeChoiceType) IsZero() bool {
        if t == nil {
            return true
        }
        
        if t.IsValueSet() {
            return t.Value.IsZero()
        }
        
        if t.IsRangeSet() {
            // If both earliest and latest are zero or nil
            if (t.Range.Earliest == nil || t.Range.Earliest.IsZero()) &&
            (t.Range.Latest == nil || t.Range.Latest.IsZero()) {
                return true
            }
            return false
        }
        
        return true
    }

    func (t *TimeChoiceType) Format() string {
        // Custom formatting logic
    }

    I also implemented custom XML unmarshaling for choice types to ensure that when one choice is set, the others are explicitly nulled out:

    func (c *FlightLevelOrAltitudeChoiceType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
        // Handle choice between altitude and flightLevel
        for {
            token, err := d.Token()
            if err != nil {
                return err
            }

            switch se := token.(type) {
            case xml.StartElement:
                switch se.Name.Local {
                case "altitude":
                    var alt AltitudeType
                    if err := d.DecodeElement(&alt, &se); err != nil {
                        return err
                    }
                    c.Altitude = &alt
                    c.FlightLevel = nil  // Explicitly null out other choice
                case "flightLevel":
                    var fl FlightLevelType
                    if err := d.DecodeElement(&fl, &se); err != nil {
                        return err
                    }
                    c.FlightLevel = &fl
                    c.Altitude = nil  // Explicitly null out other choice
                default:
                    // Skip other elements
                    if err := d.Skip(); err != nil {
                        return err
                    }
                }
            case xml.EndElement:
                if se.Name == start.Name {
                    return nil
                }
            }
        }
    }

    1.3 Date/Time Types and Helper Methods
    Problem
    The generated code included types like DateTimeUtcType but lacked helper methods for checking if the time is zero or for formatting the time as a string.

    Solution
    I added helper methods to these types:

    func (dt *DateTimeUtcType) IsZero() bool {
        if dt == nil {
            return true
        }
        
        // Since DateTimeUtcType is essentially a time.Time
        t := time.Time(*dt)
        return t.IsZero()
    }

    func (dt *DateTimeUtcType) Format() string {
        if dt == nil || dt.IsZero() {
            return "undefined"
        }
        
        // Convert to time.Time and format
        t := time.Time(*dt)
        return t.Format("2006-01-02T15:04:05Z")
    }

    These methods made / make it easier to work with date/time values, especially when determining if a value is set or formatting it for display.

    1.4 Struct vs. Nil Comparison Issues
    Problem
    In Go, only pointer types, slices, maps, channels, functions, and interfaces can be nil. Struct types cannot be nil - they always have a valid zero value. However, the code had comparisons like:

    if elem.ElementStartPoint.AerodromeReferencePoint != nil {
        // Code...
    }

    Where AerodromeReferencePoint is a struct type, not a pointer. This causes a compilation error because you can't compare a struct to nil.

    Solution
    I changed the checks to examine fields within the struct instead:
    if elem.ElementStartPoint.AerodromeReferencePoint.LocationIndicator != nil {
        // Code...
    }

    This correctly checks if a specific field (which is a pointer) is non-nil, rather than trying to check if the struct itself is nil.

    1.5 Inconsistent Type Implementations
    Problem
    The generated code had inconsistent implementations for similar types. Some types like RouteDesignatorType had a Value field with proper XML tags, while others like AltitudeType didn't. Some types used pointer fields for optional elements, while others used value types.

    Solution
    I standardized the approach as follows:

    All types representing XML elements with both text content and attributes got a Value field with xml:",chardata" tag
    All attributes were tagged with xml:"name,attr"
    I added helper methods to consistently handle all choice types
    Custom XML marshalers/unmarshalers were added where needed.

    For example, RouteDesignatorType was properly implemented by:

    type RouteDesignatorType struct {
        Value string `xml:",chardata"`
        Href  string `xml:"href,attr,omitempty"`
    }

    But I had to add similar structures to a greeat deal of other types.

2. Main.go Fixes
    The main.go file needed several updates to work with the fixed types:

    2.1 Accessing Value Fields
    I changed code like:

    fmt.Printf("Altitude: %v", *elem.Point4D.Level.Altitude)

    To:
    fmt.Printf("Altitude: %v", elem.Point4D.Level.Altitude.Value)

    This correctly accesses the Value field in types that have text content.

    2.2 Using Helper Methods for Choice Types
    I updated the code to use the helper methods:

    if elem.Point4D.Level.IsAltitudeSet() {
        fmt.Printf("Altitude: %v", elem.Point4D.Level.Altitude.Value)
    } else if elem.Point4D.Level.IsFlightLevelSet() {
        fmt.Printf("Flight Level: FL%v", elem.Point4D.Level.FlightLevel.Value)
    }

    2.3 UOM Field Handling
    I fixed UOM field handling (I puzzled this over for awhile!):

    // Changed from:
    if rtg.TakeoffMass.UOM != nil {
        fmt.Printf(" %s", *rtg.TakeoffMass.UOM)
    }

    // To:
    if rtg.TakeoffMass.UOM != "" {
        fmt.Printf(" %s", rtg.TakeoffMass.UOM)
    }

    This correctly treats UOM as a string, not a pointer!

    2.4 Inconsistent Type Access
    I discovered that some types had different implementations than others:

    // Some types like RouteDesignator had a Value field:
    fmt.Printf("Airway: %s\n", rte.RouteDesignator.Value)

    // While others like OtherRouteDesignator were pointers to strings:
    fmt.Printf("Type: %s\n", *rte.OtherRouteDesignator)

    // And others like SidStarDesignator were just strings:
    fmt.Printf("SID: %s\n", rte.StandardInstrumentDeparture.Designator)

    I updated the code to handle each type correctly.

3. Why These Patterns Matter in Go
    3.1 Representing XML in Go
    XML and Go have different data models. XML is hierarchical with elements, attributes, and text content, while Go has structs, fields, and tags. Mapping between them requires careful approach to thedesign:

    Simple vs. Complex Content:

    Simple content (just text) maps to basic Go types like string or int
    Complex content (child elements) maps to nested structs
    Simple content with attributes requires a struct with a Value field plus attribute fields


    Choices vs. Structs:

    XML choices (pick one of several elements) don't map directly to Go constructs.
    I used structs with pointer fields (so they can be nil) plus helper methods
    Custom unmarshaling ensures only one choice is set at a time.


    Nil vs. Zero Values:

    In Go, only pointer types can be nil; structs always have a zero value.
    To represent optional elements, use pointers
    To check if a struct field is "set", examine specific pointer fields within it.


    3.2 XML Marshaling in Go
    Go's standard library provides XML marshaling, but it has limitations:

    Field Tags:

    xml:"name" maps a field to an element named "name"
    xml:"name,attr" maps a field to an attribute named "name"
    xml:",chardata" maps a field to the element's text content
    xml:",omitempty" omits the field if it's the zero value


    Custom Marshaling:

    Implementing MarshalXML and UnmarshalXML methods gives complete control
    Essential for complex mappings like choice types
    Also useful for types with both text content and attributes



4. Generator Improvements for Future Integration
    The XML Schema to Go code generator has several deficiencies that should be addressed:
    4.1 Issues to Fix in the Generator

    Simple Types with Attributes:

    The generator should create a struct with a Value field and appropriate attribute fields
    It should add proper XML tags for both the value and attributes
    It should generate custom marshalers and unmarshalers when needed


    Choice Elements:

    The generator should add helper methods like IsXXXSet() for each possible choice
    It should add methods like IsZero() and Format() where appropriate
    It should generate custom unmarshalers that explicitly null out other choices


    Date/Time Types:

    All date/time types should have IsZero() and Format() methods
    Custom marshalers should be generated to handle formatting


    Consistency:

    The generator should ensure consistent patterns across similar types
    Field types should be consistent (pointer vs. value)
    XML tags should be consistently applied


    Nil Checks:

    The generator should avoid generating code that compares structs to nil
    Code should be structured to check pointer fields instead



    4.2 Implementation Strategy for Generator Improvements
    To fix the generator, we would need to:

    Analyze the XML Schema to properly identify:

    Complex types with simple content (simple types with attributes)
    Choice elements
    Date/time types
    Required vs. optional elements


    Create Templates for:

    Simple types with attributes
    Choice types with helper methods
    Custom marshalers and unmarshalers


    Apply Consistent Patterns:

    Use pointer types for optional elements
    Use value types for required elements
    Add proper XML tags
    Generate helper methods


    Generate Utility Functions:

    Create utility functions for common operations
    Generate test code to verify marshaling/unmarshaling



    4.3 Key Topics for Future Discussion
    In a future conversation to improve the generator, we should focus on:

    Schema Analysis:

    How to properly identify complex types with simple content
    How to detect choice elements and their properties
    How to determine which types need custom marshaling


    Go Type Mapping:

    Best practices for mapping XML Schema types to Go types
    Handling of type restrictions and facets
    Representation of extension and restriction type derivation


    Code Generation:

    Templates for different XML Schema constructs
    Generation of helper methods
    Generation of custom marshalers/unmarshalers


    Testing:

    Generation of test code to verify marshaling/unmarshaling
    Validation of generated code against the XML Schema


    Documentation:

    Adding proper documentation to generated code
    Explaining XML Schema concepts in Go terms



5. Conclusion
    The issues encountered in the FIXM Schema to Go code conversion highlights the challenges of mapping between XML Schema and Go. The solutions implemented provides a foundation (it needs more work, all fixes afterward have been manual) for better handling of these conversions. 

    By understanding the patterns and techniques discussed in this report, we should be able to create a more robust and maintainable code solution for processing FIXM data.

    What I think the key insights are:

    Proper XML Mapping:

    Use a Value field with xml:",chardata" for simple content
    Use xml:"name,attr" for attributes
    Use custom marshalers for complex cases


    Choice Handling:

    Use helper methods to check which choice is set
    Use custom unmarshalers to ensure only one choice is set
    Add methods like IsZero() and Format() for convenience


    Nil vs. Zero Values:

    Remember that structs cannot be nil in Go
    Check pointer fields to determine if a field is "set"
    Use pointer types for optional elements



    Applying these patterns I think should enable me to create a more robust XML Schema to Go code generator to produce more efficient, maintainable code, and save me from so much manual code corrections. This intention being that once the generator is improved, that it then becomes useful not only for FIXM, but AIXM and iWXXM data schema to go conversion. 
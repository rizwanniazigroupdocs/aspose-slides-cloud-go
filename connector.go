/*
 * --------------------------------------------------------------------------------------------------------------------
 * <copyright company="Aspose">
 *   Copyright (c) 2018 Aspose.Slides for Cloud
 * </copyright>
 * <summary>
 *   Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 * 
 *  The above copyright notice and this permission notice shall be included in all
 *  copies or substantial portions of the Software.
 * 
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *  SOFTWARE.
 * </summary>
 * --------------------------------------------------------------------------------------------------------------------
 */

package asposeslidescloud

import (
	"encoding/json"
)

// Represents Connector resource.
type IConnector interface {

	// Gets or sets the link to this resource.
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	// List of alternate links.
	getAlternateLinks() []ResourceUri
	setAlternateLinks(newValue []ResourceUri)

	// Gets or sets the name.
	getName() string
	setName(newValue string)

	// Gets or sets the width.
	getWidth() float64
	setWidth(newValue float64)

	// Gets or sets the height.
	getHeight() float64
	setHeight(newValue float64)

	// Gets or sets the alternative text.
	getAlternativeText() string
	setAlternativeText(newValue string)

	// The title of alternative text associated with the shape.
	getAlternativeTextTitle() string
	setAlternativeTextTitle(newValue string)

	// Gets or sets a value indicating whether this ShapeBase is hidden.
	getHidden() bool
	setHidden(newValue bool)

	// Gets or sets the X
	getX() float64
	setX(newValue float64)

	// Gets or sets the Y.
	getY() float64
	setY(newValue float64)

	// Gets z-order position of shape
	getZOrderPosition() int32
	setZOrderPosition(newValue int32)

	// Gets or sets the link to shapes.
	getShapes() IResourceUriElement
	setShapes(newValue IResourceUriElement)

	// Gets or sets the fill format.
	getFillFormat() IFillFormat
	setFillFormat(newValue IFillFormat)

	// Gets or sets the effect format.
	getEffectFormat() IEffectFormat
	setEffectFormat(newValue IEffectFormat)

	// Gets or sets the line format.
	getLineFormat() ILineFormat
	setLineFormat(newValue ILineFormat)

	// Shape type.
	getType() string
	setType(newValue string)

	// Shape type.
	getShapeType() string
	setShapeType(newValue string)

	// Geometry shape type.
	getGeometryShapeType() string
	setGeometryShapeType(newValue string)

	// Start shape link.
	getStartShapeConnectedTo() IResourceUri
	setStartShapeConnectedTo(newValue IResourceUri)

	// Start shape index.
	getStartShapeConnectedToIndex() int32
	setStartShapeConnectedToIndex(newValue int32)

	// End shape link.
	getEndShapeConnectedTo() IResourceUri
	setEndShapeConnectedTo(newValue IResourceUri)

	// End shape index.
	getEndShapeConnectedToIndex() int32
	setEndShapeConnectedToIndex(newValue int32)
}

type Connector struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []ResourceUri `json:"AlternateLinks,omitempty"`

	// Gets or sets the name.
	Name string `json:"Name,omitempty"`

	// Gets or sets the width.
	Width float64 `json:"Width,omitempty"`

	// Gets or sets the height.
	Height float64 `json:"Height,omitempty"`

	// Gets or sets the alternative text.
	AlternativeText string `json:"AlternativeText,omitempty"`

	// The title of alternative text associated with the shape.
	AlternativeTextTitle string `json:"AlternativeTextTitle,omitempty"`

	// Gets or sets a value indicating whether this ShapeBase is hidden.
	Hidden bool `json:"Hidden,omitempty"`

	// Gets or sets the X
	X float64 `json:"X,omitempty"`

	// Gets or sets the Y.
	Y float64 `json:"Y,omitempty"`

	// Gets z-order position of shape
	ZOrderPosition int32 `json:"ZOrderPosition"`

	// Gets or sets the link to shapes.
	Shapes IResourceUriElement `json:"Shapes,omitempty"`

	// Gets or sets the fill format.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	// Gets or sets the effect format.
	EffectFormat IEffectFormat `json:"EffectFormat,omitempty"`

	// Gets or sets the line format.
	LineFormat ILineFormat `json:"LineFormat,omitempty"`

	// Shape type.
	Type_ string `json:"Type"`

	// Shape type.
	ShapeType string `json:"ShapeType"`

	// Geometry shape type.
	GeometryShapeType string `json:"GeometryShapeType"`

	// Start shape link.
	StartShapeConnectedTo IResourceUri `json:"StartShapeConnectedTo,omitempty"`

	// Start shape index.
	StartShapeConnectedToIndex int32 `json:"StartShapeConnectedToIndex,omitempty"`

	// End shape link.
	EndShapeConnectedTo IResourceUri `json:"EndShapeConnectedTo,omitempty"`

	// End shape index.
	EndShapeConnectedToIndex int32 `json:"EndShapeConnectedToIndex,omitempty"`
}

func (this Connector) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this Connector) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this Connector) getAlternateLinks() []ResourceUri {
	return this.AlternateLinks
}

func (this Connector) setAlternateLinks(newValue []ResourceUri) {
	this.AlternateLinks = newValue
}
func (this Connector) getName() string {
	return this.Name
}

func (this Connector) setName(newValue string) {
	this.Name = newValue
}
func (this Connector) getWidth() float64 {
	return this.Width
}

func (this Connector) setWidth(newValue float64) {
	this.Width = newValue
}
func (this Connector) getHeight() float64 {
	return this.Height
}

func (this Connector) setHeight(newValue float64) {
	this.Height = newValue
}
func (this Connector) getAlternativeText() string {
	return this.AlternativeText
}

func (this Connector) setAlternativeText(newValue string) {
	this.AlternativeText = newValue
}
func (this Connector) getAlternativeTextTitle() string {
	return this.AlternativeTextTitle
}

func (this Connector) setAlternativeTextTitle(newValue string) {
	this.AlternativeTextTitle = newValue
}
func (this Connector) getHidden() bool {
	return this.Hidden
}

func (this Connector) setHidden(newValue bool) {
	this.Hidden = newValue
}
func (this Connector) getX() float64 {
	return this.X
}

func (this Connector) setX(newValue float64) {
	this.X = newValue
}
func (this Connector) getY() float64 {
	return this.Y
}

func (this Connector) setY(newValue float64) {
	this.Y = newValue
}
func (this Connector) getZOrderPosition() int32 {
	return this.ZOrderPosition
}

func (this Connector) setZOrderPosition(newValue int32) {
	this.ZOrderPosition = newValue
}
func (this Connector) getShapes() IResourceUriElement {
	return this.Shapes
}

func (this Connector) setShapes(newValue IResourceUriElement) {
	this.Shapes = newValue
}
func (this Connector) getFillFormat() IFillFormat {
	return this.FillFormat
}

func (this Connector) setFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this Connector) getEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this Connector) setEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this Connector) getLineFormat() ILineFormat {
	return this.LineFormat
}

func (this Connector) setLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this Connector) getType() string {
	return this.Type_
}

func (this Connector) setType(newValue string) {
	this.Type_ = newValue
}
func (this Connector) getShapeType() string {
	return this.ShapeType
}

func (this Connector) setShapeType(newValue string) {
	this.ShapeType = newValue
}
func (this Connector) getGeometryShapeType() string {
	return this.GeometryShapeType
}

func (this Connector) setGeometryShapeType(newValue string) {
	this.GeometryShapeType = newValue
}
func (this Connector) getStartShapeConnectedTo() IResourceUri {
	return this.StartShapeConnectedTo
}

func (this Connector) setStartShapeConnectedTo(newValue IResourceUri) {
	this.StartShapeConnectedTo = newValue
}
func (this Connector) getStartShapeConnectedToIndex() int32 {
	return this.StartShapeConnectedToIndex
}

func (this Connector) setStartShapeConnectedToIndex(newValue int32) {
	this.StartShapeConnectedToIndex = newValue
}
func (this Connector) getEndShapeConnectedTo() IResourceUri {
	return this.EndShapeConnectedTo
}

func (this Connector) setEndShapeConnectedTo(newValue IResourceUri) {
	this.EndShapeConnectedTo = newValue
}
func (this Connector) getEndShapeConnectedToIndex() int32 {
	return this.EndShapeConnectedToIndex
}

func (this Connector) setEndShapeConnectedToIndex(newValue int32) {
	this.EndShapeConnectedToIndex = newValue
}

func (this *Connector) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valSelfUri, ok := objMap["selfUri"]; ok {
		if valSelfUri != nil {
			var valueForSelfUri ResourceUri
			err = json.Unmarshal(*valSelfUri, &valueForSelfUri)
			if err != nil {
				return err
			}
			this.SelfUri = valueForSelfUri
		}
	}
	if valSelfUriCap, ok := objMap["SelfUri"]; ok {
		if valSelfUriCap != nil {
			var valueForSelfUri ResourceUri
			err = json.Unmarshal(*valSelfUriCap, &valueForSelfUri)
			if err != nil {
				return err
			}
			this.SelfUri = valueForSelfUri
		}
	}
	
	if valAlternateLinks, ok := objMap["alternateLinks"]; ok {
		if valAlternateLinks != nil {
			var valueForAlternateLinks []ResourceUri
			err = json.Unmarshal(*valAlternateLinks, &valueForAlternateLinks)
			if err != nil {
				return err
			}
			this.AlternateLinks = valueForAlternateLinks
		}
	}
	if valAlternateLinksCap, ok := objMap["AlternateLinks"]; ok {
		if valAlternateLinksCap != nil {
			var valueForAlternateLinks []ResourceUri
			err = json.Unmarshal(*valAlternateLinksCap, &valueForAlternateLinks)
			if err != nil {
				return err
			}
			this.AlternateLinks = valueForAlternateLinks
		}
	}
	
	if valName, ok := objMap["name"]; ok {
		if valName != nil {
			var valueForName string
			err = json.Unmarshal(*valName, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}
	if valNameCap, ok := objMap["Name"]; ok {
		if valNameCap != nil {
			var valueForName string
			err = json.Unmarshal(*valNameCap, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}
	
	if valWidth, ok := objMap["width"]; ok {
		if valWidth != nil {
			var valueForWidth float64
			err = json.Unmarshal(*valWidth, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	if valWidthCap, ok := objMap["Width"]; ok {
		if valWidthCap != nil {
			var valueForWidth float64
			err = json.Unmarshal(*valWidthCap, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	
	if valHeight, ok := objMap["height"]; ok {
		if valHeight != nil {
			var valueForHeight float64
			err = json.Unmarshal(*valHeight, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}
	if valHeightCap, ok := objMap["Height"]; ok {
		if valHeightCap != nil {
			var valueForHeight float64
			err = json.Unmarshal(*valHeightCap, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}
	
	if valAlternativeText, ok := objMap["alternativeText"]; ok {
		if valAlternativeText != nil {
			var valueForAlternativeText string
			err = json.Unmarshal(*valAlternativeText, &valueForAlternativeText)
			if err != nil {
				return err
			}
			this.AlternativeText = valueForAlternativeText
		}
	}
	if valAlternativeTextCap, ok := objMap["AlternativeText"]; ok {
		if valAlternativeTextCap != nil {
			var valueForAlternativeText string
			err = json.Unmarshal(*valAlternativeTextCap, &valueForAlternativeText)
			if err != nil {
				return err
			}
			this.AlternativeText = valueForAlternativeText
		}
	}
	
	if valAlternativeTextTitle, ok := objMap["alternativeTextTitle"]; ok {
		if valAlternativeTextTitle != nil {
			var valueForAlternativeTextTitle string
			err = json.Unmarshal(*valAlternativeTextTitle, &valueForAlternativeTextTitle)
			if err != nil {
				return err
			}
			this.AlternativeTextTitle = valueForAlternativeTextTitle
		}
	}
	if valAlternativeTextTitleCap, ok := objMap["AlternativeTextTitle"]; ok {
		if valAlternativeTextTitleCap != nil {
			var valueForAlternativeTextTitle string
			err = json.Unmarshal(*valAlternativeTextTitleCap, &valueForAlternativeTextTitle)
			if err != nil {
				return err
			}
			this.AlternativeTextTitle = valueForAlternativeTextTitle
		}
	}
	
	if valHidden, ok := objMap["hidden"]; ok {
		if valHidden != nil {
			var valueForHidden bool
			err = json.Unmarshal(*valHidden, &valueForHidden)
			if err != nil {
				return err
			}
			this.Hidden = valueForHidden
		}
	}
	if valHiddenCap, ok := objMap["Hidden"]; ok {
		if valHiddenCap != nil {
			var valueForHidden bool
			err = json.Unmarshal(*valHiddenCap, &valueForHidden)
			if err != nil {
				return err
			}
			this.Hidden = valueForHidden
		}
	}
	
	if valX, ok := objMap["x"]; ok {
		if valX != nil {
			var valueForX float64
			err = json.Unmarshal(*valX, &valueForX)
			if err != nil {
				return err
			}
			this.X = valueForX
		}
	}
	if valXCap, ok := objMap["X"]; ok {
		if valXCap != nil {
			var valueForX float64
			err = json.Unmarshal(*valXCap, &valueForX)
			if err != nil {
				return err
			}
			this.X = valueForX
		}
	}
	
	if valY, ok := objMap["y"]; ok {
		if valY != nil {
			var valueForY float64
			err = json.Unmarshal(*valY, &valueForY)
			if err != nil {
				return err
			}
			this.Y = valueForY
		}
	}
	if valYCap, ok := objMap["Y"]; ok {
		if valYCap != nil {
			var valueForY float64
			err = json.Unmarshal(*valYCap, &valueForY)
			if err != nil {
				return err
			}
			this.Y = valueForY
		}
	}
	
	if valZOrderPosition, ok := objMap["zOrderPosition"]; ok {
		if valZOrderPosition != nil {
			var valueForZOrderPosition int32
			err = json.Unmarshal(*valZOrderPosition, &valueForZOrderPosition)
			if err != nil {
				return err
			}
			this.ZOrderPosition = valueForZOrderPosition
		}
	}
	if valZOrderPositionCap, ok := objMap["ZOrderPosition"]; ok {
		if valZOrderPositionCap != nil {
			var valueForZOrderPosition int32
			err = json.Unmarshal(*valZOrderPositionCap, &valueForZOrderPosition)
			if err != nil {
				return err
			}
			this.ZOrderPosition = valueForZOrderPosition
		}
	}
	
	if valShapes, ok := objMap["shapes"]; ok {
		if valShapes != nil {
			var valueForShapes ResourceUriElement
			err = json.Unmarshal(*valShapes, &valueForShapes)
			if err != nil {
				return err
			}
			this.Shapes = valueForShapes
		}
	}
	if valShapesCap, ok := objMap["Shapes"]; ok {
		if valShapesCap != nil {
			var valueForShapes ResourceUriElement
			err = json.Unmarshal(*valShapesCap, &valueForShapes)
			if err != nil {
				return err
			}
			this.Shapes = valueForShapes
		}
	}
	
	if valFillFormat, ok := objMap["fillFormat"]; ok {
		if valFillFormat != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormat, &valueForFillFormat)
			if err != nil {
				return err
			}
			this.FillFormat = valueForFillFormat
		}
	}
	if valFillFormatCap, ok := objMap["FillFormat"]; ok {
		if valFillFormatCap != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormatCap, &valueForFillFormat)
			if err != nil {
				return err
			}
			this.FillFormat = valueForFillFormat
		}
	}
	
	if valEffectFormat, ok := objMap["effectFormat"]; ok {
		if valEffectFormat != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormat, &valueForEffectFormat)
			if err != nil {
				return err
			}
			this.EffectFormat = valueForEffectFormat
		}
	}
	if valEffectFormatCap, ok := objMap["EffectFormat"]; ok {
		if valEffectFormatCap != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormatCap, &valueForEffectFormat)
			if err != nil {
				return err
			}
			this.EffectFormat = valueForEffectFormat
		}
	}
	
	if valLineFormat, ok := objMap["lineFormat"]; ok {
		if valLineFormat != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormat, &valueForLineFormat)
			if err != nil {
				return err
			}
			this.LineFormat = valueForLineFormat
		}
	}
	if valLineFormatCap, ok := objMap["LineFormat"]; ok {
		if valLineFormatCap != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormatCap, &valueForLineFormat)
			if err != nil {
				return err
			}
			this.LineFormat = valueForLineFormat
		}
	}
	this.Type_ = "TYPE__CONNECTOR"
	if valType, ok := objMap["type"]; ok {
		if valType != nil {
			var valueForType string
			err = json.Unmarshal(*valType, &valueForType)
			if err != nil {
				return err
			}
			this.Type_ = valueForType
		}
	}
	if valTypeCap, ok := objMap["Type"]; ok {
		if valTypeCap != nil {
			var valueForType string
			err = json.Unmarshal(*valTypeCap, &valueForType)
			if err != nil {
				return err
			}
			this.Type_ = valueForType
		}
	}
	this.ShapeType = "Custom"
	if valShapeType, ok := objMap["shapeType"]; ok {
		if valShapeType != nil {
			var valueForShapeType string
			err = json.Unmarshal(*valShapeType, &valueForShapeType)
			if err != nil {
				return err
			}
			this.ShapeType = valueForShapeType
		}
	}
	if valShapeTypeCap, ok := objMap["ShapeType"]; ok {
		if valShapeTypeCap != nil {
			var valueForShapeType string
			err = json.Unmarshal(*valShapeTypeCap, &valueForShapeType)
			if err != nil {
				return err
			}
			this.ShapeType = valueForShapeType
		}
	}
	this.GeometryShapeType = "Custom"
	if valGeometryShapeType, ok := objMap["geometryShapeType"]; ok {
		if valGeometryShapeType != nil {
			var valueForGeometryShapeType string
			err = json.Unmarshal(*valGeometryShapeType, &valueForGeometryShapeType)
			if err != nil {
				return err
			}
			this.GeometryShapeType = valueForGeometryShapeType
		}
	}
	if valGeometryShapeTypeCap, ok := objMap["GeometryShapeType"]; ok {
		if valGeometryShapeTypeCap != nil {
			var valueForGeometryShapeType string
			err = json.Unmarshal(*valGeometryShapeTypeCap, &valueForGeometryShapeType)
			if err != nil {
				return err
			}
			this.GeometryShapeType = valueForGeometryShapeType
		}
	}
	
	if valStartShapeConnectedTo, ok := objMap["startShapeConnectedTo"]; ok {
		if valStartShapeConnectedTo != nil {
			var valueForStartShapeConnectedTo ResourceUri
			err = json.Unmarshal(*valStartShapeConnectedTo, &valueForStartShapeConnectedTo)
			if err != nil {
				return err
			}
			this.StartShapeConnectedTo = valueForStartShapeConnectedTo
		}
	}
	if valStartShapeConnectedToCap, ok := objMap["StartShapeConnectedTo"]; ok {
		if valStartShapeConnectedToCap != nil {
			var valueForStartShapeConnectedTo ResourceUri
			err = json.Unmarshal(*valStartShapeConnectedToCap, &valueForStartShapeConnectedTo)
			if err != nil {
				return err
			}
			this.StartShapeConnectedTo = valueForStartShapeConnectedTo
		}
	}
	
	if valStartShapeConnectedToIndex, ok := objMap["startShapeConnectedToIndex"]; ok {
		if valStartShapeConnectedToIndex != nil {
			var valueForStartShapeConnectedToIndex int32
			err = json.Unmarshal(*valStartShapeConnectedToIndex, &valueForStartShapeConnectedToIndex)
			if err != nil {
				return err
			}
			this.StartShapeConnectedToIndex = valueForStartShapeConnectedToIndex
		}
	}
	if valStartShapeConnectedToIndexCap, ok := objMap["StartShapeConnectedToIndex"]; ok {
		if valStartShapeConnectedToIndexCap != nil {
			var valueForStartShapeConnectedToIndex int32
			err = json.Unmarshal(*valStartShapeConnectedToIndexCap, &valueForStartShapeConnectedToIndex)
			if err != nil {
				return err
			}
			this.StartShapeConnectedToIndex = valueForStartShapeConnectedToIndex
		}
	}
	
	if valEndShapeConnectedTo, ok := objMap["endShapeConnectedTo"]; ok {
		if valEndShapeConnectedTo != nil {
			var valueForEndShapeConnectedTo ResourceUri
			err = json.Unmarshal(*valEndShapeConnectedTo, &valueForEndShapeConnectedTo)
			if err != nil {
				return err
			}
			this.EndShapeConnectedTo = valueForEndShapeConnectedTo
		}
	}
	if valEndShapeConnectedToCap, ok := objMap["EndShapeConnectedTo"]; ok {
		if valEndShapeConnectedToCap != nil {
			var valueForEndShapeConnectedTo ResourceUri
			err = json.Unmarshal(*valEndShapeConnectedToCap, &valueForEndShapeConnectedTo)
			if err != nil {
				return err
			}
			this.EndShapeConnectedTo = valueForEndShapeConnectedTo
		}
	}
	
	if valEndShapeConnectedToIndex, ok := objMap["endShapeConnectedToIndex"]; ok {
		if valEndShapeConnectedToIndex != nil {
			var valueForEndShapeConnectedToIndex int32
			err = json.Unmarshal(*valEndShapeConnectedToIndex, &valueForEndShapeConnectedToIndex)
			if err != nil {
				return err
			}
			this.EndShapeConnectedToIndex = valueForEndShapeConnectedToIndex
		}
	}
	if valEndShapeConnectedToIndexCap, ok := objMap["EndShapeConnectedToIndex"]; ok {
		if valEndShapeConnectedToIndexCap != nil {
			var valueForEndShapeConnectedToIndex int32
			err = json.Unmarshal(*valEndShapeConnectedToIndexCap, &valueForEndShapeConnectedToIndex)
			if err != nil {
				return err
			}
			this.EndShapeConnectedToIndex = valueForEndShapeConnectedToIndex
		}
	}

    return nil
}

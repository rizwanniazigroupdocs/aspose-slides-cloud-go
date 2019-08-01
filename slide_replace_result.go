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

// Represents slide replace result DTO.
type ISlideReplaceResult interface {

	// Gets or sets the link to this resource.
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	// List of alternate links.
	getAlternateLinks() []ResourceUri
	setAlternateLinks(newValue []ResourceUri)

	// Gets or sets the width.
	getWidth() float64
	setWidth(newValue float64)

	// Gets or sets the height.
	getHeight() float64
	setHeight(newValue float64)

	// Specifies if shapes of the master slide should be shown on the slide. True by default.
	getShowMasterShapes() bool
	setShowMasterShapes(newValue bool)

	// Gets or sets the  link to the layout slide.
	getLayoutSlide() IResourceUriElement
	setLayoutSlide(newValue IResourceUriElement)

	// Gets or sets the  link to list of top-level shapes.
	getShapes() IResourceUriElement
	setShapes(newValue IResourceUriElement)

	// Gets or sets the link to theme.
	getTheme() IResourceUriElement
	setTheme(newValue IResourceUriElement)

	// Gets or sets the  link to placeholders.
	getPlaceholders() IResourceUriElement
	setPlaceholders(newValue IResourceUriElement)

	// Gets or sets the link to images.
	getImages() IResourceUriElement
	setImages(newValue IResourceUriElement)

	// Gets or sets the link to comments.
	getComments() IResourceUriElement
	setComments(newValue IResourceUriElement)

	// Get or sets the link to slide's background
	getBackground() IResourceUriElement
	setBackground(newValue IResourceUriElement)

	// Get or sets the link to notes slide.
	getNotesSlide() IResourceUriElement
	setNotesSlide(newValue IResourceUriElement)

	// Gets or sets the number of matches 
	getMatches() int32
	setMatches(newValue int32)
}

type SlideReplaceResult struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []ResourceUri `json:"AlternateLinks,omitempty"`

	// Gets or sets the width.
	Width float64 `json:"Width"`

	// Gets or sets the height.
	Height float64 `json:"Height"`

	// Specifies if shapes of the master slide should be shown on the slide. True by default.
	ShowMasterShapes bool `json:"ShowMasterShapes"`

	// Gets or sets the  link to the layout slide.
	LayoutSlide IResourceUriElement `json:"LayoutSlide,omitempty"`

	// Gets or sets the  link to list of top-level shapes.
	Shapes IResourceUriElement `json:"Shapes,omitempty"`

	// Gets or sets the link to theme.
	Theme IResourceUriElement `json:"Theme,omitempty"`

	// Gets or sets the  link to placeholders.
	Placeholders IResourceUriElement `json:"Placeholders,omitempty"`

	// Gets or sets the link to images.
	Images IResourceUriElement `json:"Images,omitempty"`

	// Gets or sets the link to comments.
	Comments IResourceUriElement `json:"Comments,omitempty"`

	// Get or sets the link to slide's background
	Background IResourceUriElement `json:"Background,omitempty"`

	// Get or sets the link to notes slide.
	NotesSlide IResourceUriElement `json:"NotesSlide,omitempty"`

	// Gets or sets the number of matches 
	Matches int32 `json:"Matches"`
}

func (this SlideReplaceResult) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this SlideReplaceResult) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this SlideReplaceResult) getAlternateLinks() []ResourceUri {
	return this.AlternateLinks
}

func (this SlideReplaceResult) setAlternateLinks(newValue []ResourceUri) {
	this.AlternateLinks = newValue
}
func (this SlideReplaceResult) getWidth() float64 {
	return this.Width
}

func (this SlideReplaceResult) setWidth(newValue float64) {
	this.Width = newValue
}
func (this SlideReplaceResult) getHeight() float64 {
	return this.Height
}

func (this SlideReplaceResult) setHeight(newValue float64) {
	this.Height = newValue
}
func (this SlideReplaceResult) getShowMasterShapes() bool {
	return this.ShowMasterShapes
}

func (this SlideReplaceResult) setShowMasterShapes(newValue bool) {
	this.ShowMasterShapes = newValue
}
func (this SlideReplaceResult) getLayoutSlide() IResourceUriElement {
	return this.LayoutSlide
}

func (this SlideReplaceResult) setLayoutSlide(newValue IResourceUriElement) {
	this.LayoutSlide = newValue
}
func (this SlideReplaceResult) getShapes() IResourceUriElement {
	return this.Shapes
}

func (this SlideReplaceResult) setShapes(newValue IResourceUriElement) {
	this.Shapes = newValue
}
func (this SlideReplaceResult) getTheme() IResourceUriElement {
	return this.Theme
}

func (this SlideReplaceResult) setTheme(newValue IResourceUriElement) {
	this.Theme = newValue
}
func (this SlideReplaceResult) getPlaceholders() IResourceUriElement {
	return this.Placeholders
}

func (this SlideReplaceResult) setPlaceholders(newValue IResourceUriElement) {
	this.Placeholders = newValue
}
func (this SlideReplaceResult) getImages() IResourceUriElement {
	return this.Images
}

func (this SlideReplaceResult) setImages(newValue IResourceUriElement) {
	this.Images = newValue
}
func (this SlideReplaceResult) getComments() IResourceUriElement {
	return this.Comments
}

func (this SlideReplaceResult) setComments(newValue IResourceUriElement) {
	this.Comments = newValue
}
func (this SlideReplaceResult) getBackground() IResourceUriElement {
	return this.Background
}

func (this SlideReplaceResult) setBackground(newValue IResourceUriElement) {
	this.Background = newValue
}
func (this SlideReplaceResult) getNotesSlide() IResourceUriElement {
	return this.NotesSlide
}

func (this SlideReplaceResult) setNotesSlide(newValue IResourceUriElement) {
	this.NotesSlide = newValue
}
func (this SlideReplaceResult) getMatches() int32 {
	return this.Matches
}

func (this SlideReplaceResult) setMatches(newValue int32) {
	this.Matches = newValue
}

func (this *SlideReplaceResult) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valSelfUri, ok := objMap["SelfUri"]; ok {
		if valSelfUri != nil {
			var valueForSelfUri ResourceUri
			err = json.Unmarshal(*valSelfUri, &valueForSelfUri)
			if err != nil {
				return err
			}
			this.SelfUri = valueForSelfUri
		}
	}
	
	if valAlternateLinks, ok := objMap["AlternateLinks"]; ok {
		if valAlternateLinks != nil {
			var valueForAlternateLinks []ResourceUri
			err = json.Unmarshal(*valAlternateLinks, &valueForAlternateLinks)
			if err != nil {
				return err
			}
			this.AlternateLinks = valueForAlternateLinks
		}
	}
	
	if valWidth, ok := objMap["Width"]; ok {
		if valWidth != nil {
			var valueForWidth float64
			err = json.Unmarshal(*valWidth, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	
	if valHeight, ok := objMap["Height"]; ok {
		if valHeight != nil {
			var valueForHeight float64
			err = json.Unmarshal(*valHeight, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}
	
	if valShowMasterShapes, ok := objMap["ShowMasterShapes"]; ok {
		if valShowMasterShapes != nil {
			var valueForShowMasterShapes bool
			err = json.Unmarshal(*valShowMasterShapes, &valueForShowMasterShapes)
			if err != nil {
				return err
			}
			this.ShowMasterShapes = valueForShowMasterShapes
		}
	}
	
	if valLayoutSlide, ok := objMap["LayoutSlide"]; ok {
		if valLayoutSlide != nil {
			var valueForLayoutSlide ResourceUriElement
			err = json.Unmarshal(*valLayoutSlide, &valueForLayoutSlide)
			if err != nil {
				return err
			}
			this.LayoutSlide = valueForLayoutSlide
		}
	}
	
	if valShapes, ok := objMap["Shapes"]; ok {
		if valShapes != nil {
			var valueForShapes ResourceUriElement
			err = json.Unmarshal(*valShapes, &valueForShapes)
			if err != nil {
				return err
			}
			this.Shapes = valueForShapes
		}
	}
	
	if valTheme, ok := objMap["Theme"]; ok {
		if valTheme != nil {
			var valueForTheme ResourceUriElement
			err = json.Unmarshal(*valTheme, &valueForTheme)
			if err != nil {
				return err
			}
			this.Theme = valueForTheme
		}
	}
	
	if valPlaceholders, ok := objMap["Placeholders"]; ok {
		if valPlaceholders != nil {
			var valueForPlaceholders ResourceUriElement
			err = json.Unmarshal(*valPlaceholders, &valueForPlaceholders)
			if err != nil {
				return err
			}
			this.Placeholders = valueForPlaceholders
		}
	}
	
	if valImages, ok := objMap["Images"]; ok {
		if valImages != nil {
			var valueForImages ResourceUriElement
			err = json.Unmarshal(*valImages, &valueForImages)
			if err != nil {
				return err
			}
			this.Images = valueForImages
		}
	}
	
	if valComments, ok := objMap["Comments"]; ok {
		if valComments != nil {
			var valueForComments ResourceUriElement
			err = json.Unmarshal(*valComments, &valueForComments)
			if err != nil {
				return err
			}
			this.Comments = valueForComments
		}
	}
	
	if valBackground, ok := objMap["Background"]; ok {
		if valBackground != nil {
			var valueForBackground ResourceUriElement
			err = json.Unmarshal(*valBackground, &valueForBackground)
			if err != nil {
				return err
			}
			this.Background = valueForBackground
		}
	}
	
	if valNotesSlide, ok := objMap["NotesSlide"]; ok {
		if valNotesSlide != nil {
			var valueForNotesSlide ResourceUriElement
			err = json.Unmarshal(*valNotesSlide, &valueForNotesSlide)
			if err != nil {
				return err
			}
			this.NotesSlide = valueForNotesSlide
		}
	}
	
	if valMatches, ok := objMap["Matches"]; ok {
		if valMatches != nil {
			var valueForMatches int32
			err = json.Unmarshal(*valMatches, &valueForMatches)
			if err != nil {
				return err
			}
			this.Matches = valueForMatches
		}
	}

    return nil
}

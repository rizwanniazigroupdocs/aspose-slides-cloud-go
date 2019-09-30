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

// Represents Slide's theme 
type ITheme interface {

	// Gets or sets the link to this resource.
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	// List of alternate links.
	getAlternateLinks() []ResourceUri
	setAlternateLinks(newValue []ResourceUri)

	// Name.
	getName() string
	setName(newValue string)

	// Color scheme.
	getColorScheme() IResourceUriElement
	setColorScheme(newValue IResourceUriElement)

	// Font scheme.
	getFontScheme() IResourceUriElement
	setFontScheme(newValue IResourceUriElement)

	// Format scheme.
	getFormatScheme() IResourceUriElement
	setFormatScheme(newValue IResourceUriElement)
}

type Theme struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []ResourceUri `json:"AlternateLinks,omitempty"`

	// Name.
	Name string `json:"Name,omitempty"`

	// Color scheme.
	ColorScheme IResourceUriElement `json:"ColorScheme,omitempty"`

	// Font scheme.
	FontScheme IResourceUriElement `json:"FontScheme,omitempty"`

	// Format scheme.
	FormatScheme IResourceUriElement `json:"FormatScheme,omitempty"`
}

func (this Theme) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this Theme) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this Theme) getAlternateLinks() []ResourceUri {
	return this.AlternateLinks
}

func (this Theme) setAlternateLinks(newValue []ResourceUri) {
	this.AlternateLinks = newValue
}
func (this Theme) getName() string {
	return this.Name
}

func (this Theme) setName(newValue string) {
	this.Name = newValue
}
func (this Theme) getColorScheme() IResourceUriElement {
	return this.ColorScheme
}

func (this Theme) setColorScheme(newValue IResourceUriElement) {
	this.ColorScheme = newValue
}
func (this Theme) getFontScheme() IResourceUriElement {
	return this.FontScheme
}

func (this Theme) setFontScheme(newValue IResourceUriElement) {
	this.FontScheme = newValue
}
func (this Theme) getFormatScheme() IResourceUriElement {
	return this.FormatScheme
}

func (this Theme) setFormatScheme(newValue IResourceUriElement) {
	this.FormatScheme = newValue
}

func (this *Theme) UnmarshalJSON(b []byte) error {
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
	
	if valColorScheme, ok := objMap["colorScheme"]; ok {
		if valColorScheme != nil {
			var valueForColorScheme ResourceUriElement
			err = json.Unmarshal(*valColorScheme, &valueForColorScheme)
			if err != nil {
				return err
			}
			this.ColorScheme = valueForColorScheme
		}
	}
	if valColorSchemeCap, ok := objMap["ColorScheme"]; ok {
		if valColorSchemeCap != nil {
			var valueForColorScheme ResourceUriElement
			err = json.Unmarshal(*valColorSchemeCap, &valueForColorScheme)
			if err != nil {
				return err
			}
			this.ColorScheme = valueForColorScheme
		}
	}
	
	if valFontScheme, ok := objMap["fontScheme"]; ok {
		if valFontScheme != nil {
			var valueForFontScheme ResourceUriElement
			err = json.Unmarshal(*valFontScheme, &valueForFontScheme)
			if err != nil {
				return err
			}
			this.FontScheme = valueForFontScheme
		}
	}
	if valFontSchemeCap, ok := objMap["FontScheme"]; ok {
		if valFontSchemeCap != nil {
			var valueForFontScheme ResourceUriElement
			err = json.Unmarshal(*valFontSchemeCap, &valueForFontScheme)
			if err != nil {
				return err
			}
			this.FontScheme = valueForFontScheme
		}
	}
	
	if valFormatScheme, ok := objMap["formatScheme"]; ok {
		if valFormatScheme != nil {
			var valueForFormatScheme ResourceUriElement
			err = json.Unmarshal(*valFormatScheme, &valueForFormatScheme)
			if err != nil {
				return err
			}
			this.FormatScheme = valueForFormatScheme
		}
	}
	if valFormatSchemeCap, ok := objMap["FormatScheme"]; ok {
		if valFormatSchemeCap != nil {
			var valueForFormatScheme ResourceUriElement
			err = json.Unmarshal(*valFormatSchemeCap, &valueForFormatScheme)
			if err != nil {
				return err
			}
			this.FormatScheme = valueForFormatScheme
		}
	}

    return nil
}

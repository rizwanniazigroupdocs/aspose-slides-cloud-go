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

// Slide's color scheme DTO
type IColorScheme interface {

	// Gets or sets the link to this resource.
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	// List of alternate links.
	getAlternateLinks() []ResourceUri
	setAlternateLinks(newValue []ResourceUri)

	// First accent color.
	getAccent1() string
	setAccent1(newValue string)

	// Second accent color.
	getAccent2() string
	setAccent2(newValue string)

	// Third accent color.
	getAccent3() string
	setAccent3(newValue string)

	// Fourth accent color.
	getAccent4() string
	setAccent4(newValue string)

	// Fifth accent color.
	getAccent5() string
	setAccent5(newValue string)

	// Sixth accent color.
	getAccent6() string
	setAccent6(newValue string)

	// First dark color.
	getDark1() string
	setDark1(newValue string)

	// Second dark color.
	getDark2() string
	setDark2(newValue string)

	// Visited hyperlink color.
	getFollowedHyperlink() string
	setFollowedHyperlink(newValue string)

	// Hyperlink color/
	getHyperlink() string
	setHyperlink(newValue string)

	// First light color.
	getLight1() string
	setLight1(newValue string)

	// Second light color.
	getLight2() string
	setLight2(newValue string)
}

type ColorScheme struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []ResourceUri `json:"AlternateLinks,omitempty"`

	// First accent color.
	Accent1 string `json:"Accent1,omitempty"`

	// Second accent color.
	Accent2 string `json:"Accent2,omitempty"`

	// Third accent color.
	Accent3 string `json:"Accent3,omitempty"`

	// Fourth accent color.
	Accent4 string `json:"Accent4,omitempty"`

	// Fifth accent color.
	Accent5 string `json:"Accent5,omitempty"`

	// Sixth accent color.
	Accent6 string `json:"Accent6,omitempty"`

	// First dark color.
	Dark1 string `json:"Dark1,omitempty"`

	// Second dark color.
	Dark2 string `json:"Dark2,omitempty"`

	// Visited hyperlink color.
	FollowedHyperlink string `json:"FollowedHyperlink,omitempty"`

	// Hyperlink color/
	Hyperlink string `json:"Hyperlink,omitempty"`

	// First light color.
	Light1 string `json:"Light1,omitempty"`

	// Second light color.
	Light2 string `json:"Light2,omitempty"`
}

func (this ColorScheme) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this ColorScheme) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this ColorScheme) getAlternateLinks() []ResourceUri {
	return this.AlternateLinks
}

func (this ColorScheme) setAlternateLinks(newValue []ResourceUri) {
	this.AlternateLinks = newValue
}
func (this ColorScheme) getAccent1() string {
	return this.Accent1
}

func (this ColorScheme) setAccent1(newValue string) {
	this.Accent1 = newValue
}
func (this ColorScheme) getAccent2() string {
	return this.Accent2
}

func (this ColorScheme) setAccent2(newValue string) {
	this.Accent2 = newValue
}
func (this ColorScheme) getAccent3() string {
	return this.Accent3
}

func (this ColorScheme) setAccent3(newValue string) {
	this.Accent3 = newValue
}
func (this ColorScheme) getAccent4() string {
	return this.Accent4
}

func (this ColorScheme) setAccent4(newValue string) {
	this.Accent4 = newValue
}
func (this ColorScheme) getAccent5() string {
	return this.Accent5
}

func (this ColorScheme) setAccent5(newValue string) {
	this.Accent5 = newValue
}
func (this ColorScheme) getAccent6() string {
	return this.Accent6
}

func (this ColorScheme) setAccent6(newValue string) {
	this.Accent6 = newValue
}
func (this ColorScheme) getDark1() string {
	return this.Dark1
}

func (this ColorScheme) setDark1(newValue string) {
	this.Dark1 = newValue
}
func (this ColorScheme) getDark2() string {
	return this.Dark2
}

func (this ColorScheme) setDark2(newValue string) {
	this.Dark2 = newValue
}
func (this ColorScheme) getFollowedHyperlink() string {
	return this.FollowedHyperlink
}

func (this ColorScheme) setFollowedHyperlink(newValue string) {
	this.FollowedHyperlink = newValue
}
func (this ColorScheme) getHyperlink() string {
	return this.Hyperlink
}

func (this ColorScheme) setHyperlink(newValue string) {
	this.Hyperlink = newValue
}
func (this ColorScheme) getLight1() string {
	return this.Light1
}

func (this ColorScheme) setLight1(newValue string) {
	this.Light1 = newValue
}
func (this ColorScheme) getLight2() string {
	return this.Light2
}

func (this ColorScheme) setLight2(newValue string) {
	this.Light2 = newValue
}

func (this *ColorScheme) UnmarshalJSON(b []byte) error {
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
	
	if valAccent1, ok := objMap["Accent1"]; ok {
		if valAccent1 != nil {
			var valueForAccent1 string
			err = json.Unmarshal(*valAccent1, &valueForAccent1)
			if err != nil {
				return err
			}
			this.Accent1 = valueForAccent1
		}
	}
	
	if valAccent2, ok := objMap["Accent2"]; ok {
		if valAccent2 != nil {
			var valueForAccent2 string
			err = json.Unmarshal(*valAccent2, &valueForAccent2)
			if err != nil {
				return err
			}
			this.Accent2 = valueForAccent2
		}
	}
	
	if valAccent3, ok := objMap["Accent3"]; ok {
		if valAccent3 != nil {
			var valueForAccent3 string
			err = json.Unmarshal(*valAccent3, &valueForAccent3)
			if err != nil {
				return err
			}
			this.Accent3 = valueForAccent3
		}
	}
	
	if valAccent4, ok := objMap["Accent4"]; ok {
		if valAccent4 != nil {
			var valueForAccent4 string
			err = json.Unmarshal(*valAccent4, &valueForAccent4)
			if err != nil {
				return err
			}
			this.Accent4 = valueForAccent4
		}
	}
	
	if valAccent5, ok := objMap["Accent5"]; ok {
		if valAccent5 != nil {
			var valueForAccent5 string
			err = json.Unmarshal(*valAccent5, &valueForAccent5)
			if err != nil {
				return err
			}
			this.Accent5 = valueForAccent5
		}
	}
	
	if valAccent6, ok := objMap["Accent6"]; ok {
		if valAccent6 != nil {
			var valueForAccent6 string
			err = json.Unmarshal(*valAccent6, &valueForAccent6)
			if err != nil {
				return err
			}
			this.Accent6 = valueForAccent6
		}
	}
	
	if valDark1, ok := objMap["Dark1"]; ok {
		if valDark1 != nil {
			var valueForDark1 string
			err = json.Unmarshal(*valDark1, &valueForDark1)
			if err != nil {
				return err
			}
			this.Dark1 = valueForDark1
		}
	}
	
	if valDark2, ok := objMap["Dark2"]; ok {
		if valDark2 != nil {
			var valueForDark2 string
			err = json.Unmarshal(*valDark2, &valueForDark2)
			if err != nil {
				return err
			}
			this.Dark2 = valueForDark2
		}
	}
	
	if valFollowedHyperlink, ok := objMap["FollowedHyperlink"]; ok {
		if valFollowedHyperlink != nil {
			var valueForFollowedHyperlink string
			err = json.Unmarshal(*valFollowedHyperlink, &valueForFollowedHyperlink)
			if err != nil {
				return err
			}
			this.FollowedHyperlink = valueForFollowedHyperlink
		}
	}
	
	if valHyperlink, ok := objMap["Hyperlink"]; ok {
		if valHyperlink != nil {
			var valueForHyperlink string
			err = json.Unmarshal(*valHyperlink, &valueForHyperlink)
			if err != nil {
				return err
			}
			this.Hyperlink = valueForHyperlink
		}
	}
	
	if valLight1, ok := objMap["Light1"]; ok {
		if valLight1 != nil {
			var valueForLight1 string
			err = json.Unmarshal(*valLight1, &valueForLight1)
			if err != nil {
				return err
			}
			this.Light1 = valueForLight1
		}
	}
	
	if valLight2, ok := objMap["Light2"]; ok {
		if valLight2 != nil {
			var valueForLight2 string
			err = json.Unmarshal(*valLight2, &valueForLight2)
			if err != nil {
				return err
			}
			this.Light2 = valueForLight2
		}
	}

    return nil
}

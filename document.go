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

// Represents document DTO.
type IDocument interface {

	// Gets or sets the link to this resource.
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	// List of alternate links.
	getAlternateLinks() []IResourceUri
	setAlternateLinks(newValue []IResourceUri)

	// Link to Document properties.
	getDocumentProperties() IResourceUriElement
	setDocumentProperties(newValue IResourceUriElement)

	// Link to Document properties.
	getViewProperties() IResourceUriElement
	setViewProperties(newValue IResourceUriElement)

	// Link to slides collection.
	getSlides() IResourceUriElement
	setSlides(newValue IResourceUriElement)

	// Link to images collection.
	getImages() IResourceUriElement
	setImages(newValue IResourceUriElement)

	// Link to layout slides collection.
	getLayoutSlides() IResourceUriElement
	setLayoutSlides(newValue IResourceUriElement)

	// Link to master slides collection.
	getMasterSlides() IResourceUriElement
	setMasterSlides(newValue IResourceUriElement)
}

type Document struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// Link to Document properties.
	DocumentProperties IResourceUriElement `json:"DocumentProperties,omitempty"`

	// Link to Document properties.
	ViewProperties IResourceUriElement `json:"ViewProperties,omitempty"`

	// Link to slides collection.
	Slides IResourceUriElement `json:"Slides,omitempty"`

	// Link to images collection.
	Images IResourceUriElement `json:"Images,omitempty"`

	// Link to layout slides collection.
	LayoutSlides IResourceUriElement `json:"LayoutSlides,omitempty"`

	// Link to master slides collection.
	MasterSlides IResourceUriElement `json:"MasterSlides,omitempty"`
}

func (this *Document) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *Document) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *Document) getAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *Document) setAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *Document) getDocumentProperties() IResourceUriElement {
	return this.DocumentProperties
}

func (this *Document) setDocumentProperties(newValue IResourceUriElement) {
	this.DocumentProperties = newValue
}
func (this *Document) getViewProperties() IResourceUriElement {
	return this.ViewProperties
}

func (this *Document) setViewProperties(newValue IResourceUriElement) {
	this.ViewProperties = newValue
}
func (this *Document) getSlides() IResourceUriElement {
	return this.Slides
}

func (this *Document) setSlides(newValue IResourceUriElement) {
	this.Slides = newValue
}
func (this *Document) getImages() IResourceUriElement {
	return this.Images
}

func (this *Document) setImages(newValue IResourceUriElement) {
	this.Images = newValue
}
func (this *Document) getLayoutSlides() IResourceUriElement {
	return this.LayoutSlides
}

func (this *Document) setLayoutSlides(newValue IResourceUriElement) {
	this.LayoutSlides = newValue
}
func (this *Document) getMasterSlides() IResourceUriElement {
	return this.MasterSlides
}

func (this *Document) setMasterSlides(newValue IResourceUriElement) {
	this.MasterSlides = newValue
}

func (this *Document) UnmarshalJSON(b []byte) error {
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
			this.SelfUri = &valueForSelfUri
		}
	}
	if valSelfUriCap, ok := objMap["SelfUri"]; ok {
		if valSelfUriCap != nil {
			var valueForSelfUri ResourceUri
			err = json.Unmarshal(*valSelfUriCap, &valueForSelfUri)
			if err != nil {
				return err
			}
			this.SelfUri = &valueForSelfUri
		}
	}
	
	if valAlternateLinks, ok := objMap["alternateLinks"]; ok {
		if valAlternateLinks != nil {
			var valueForAlternateLinks []ResourceUri
			err = json.Unmarshal(*valAlternateLinks, &valueForAlternateLinks)
			if err != nil {
				return err
			}
			valueForIAlternateLinks := make([]IResourceUri, len(valueForAlternateLinks))
			for i, v := range valueForAlternateLinks {
				valueForIAlternateLinks[i] = IResourceUri(&v)
			}
			this.AlternateLinks = valueForIAlternateLinks
		}
	}
	if valAlternateLinksCap, ok := objMap["AlternateLinks"]; ok {
		if valAlternateLinksCap != nil {
			var valueForAlternateLinks []ResourceUri
			err = json.Unmarshal(*valAlternateLinksCap, &valueForAlternateLinks)
			if err != nil {
				return err
			}
			valueForIAlternateLinks := make([]IResourceUri, len(valueForAlternateLinks))
			for i, v := range valueForAlternateLinks {
				valueForIAlternateLinks[i] = IResourceUri(&v)
			}
			this.AlternateLinks = valueForIAlternateLinks
		}
	}
	
	if valDocumentProperties, ok := objMap["documentProperties"]; ok {
		if valDocumentProperties != nil {
			var valueForDocumentProperties ResourceUriElement
			err = json.Unmarshal(*valDocumentProperties, &valueForDocumentProperties)
			if err != nil {
				return err
			}
			this.DocumentProperties = &valueForDocumentProperties
		}
	}
	if valDocumentPropertiesCap, ok := objMap["DocumentProperties"]; ok {
		if valDocumentPropertiesCap != nil {
			var valueForDocumentProperties ResourceUriElement
			err = json.Unmarshal(*valDocumentPropertiesCap, &valueForDocumentProperties)
			if err != nil {
				return err
			}
			this.DocumentProperties = &valueForDocumentProperties
		}
	}
	
	if valViewProperties, ok := objMap["viewProperties"]; ok {
		if valViewProperties != nil {
			var valueForViewProperties ResourceUriElement
			err = json.Unmarshal(*valViewProperties, &valueForViewProperties)
			if err != nil {
				return err
			}
			this.ViewProperties = &valueForViewProperties
		}
	}
	if valViewPropertiesCap, ok := objMap["ViewProperties"]; ok {
		if valViewPropertiesCap != nil {
			var valueForViewProperties ResourceUriElement
			err = json.Unmarshal(*valViewPropertiesCap, &valueForViewProperties)
			if err != nil {
				return err
			}
			this.ViewProperties = &valueForViewProperties
		}
	}
	
	if valSlides, ok := objMap["slides"]; ok {
		if valSlides != nil {
			var valueForSlides ResourceUriElement
			err = json.Unmarshal(*valSlides, &valueForSlides)
			if err != nil {
				return err
			}
			this.Slides = &valueForSlides
		}
	}
	if valSlidesCap, ok := objMap["Slides"]; ok {
		if valSlidesCap != nil {
			var valueForSlides ResourceUriElement
			err = json.Unmarshal(*valSlidesCap, &valueForSlides)
			if err != nil {
				return err
			}
			this.Slides = &valueForSlides
		}
	}
	
	if valImages, ok := objMap["images"]; ok {
		if valImages != nil {
			var valueForImages ResourceUriElement
			err = json.Unmarshal(*valImages, &valueForImages)
			if err != nil {
				return err
			}
			this.Images = &valueForImages
		}
	}
	if valImagesCap, ok := objMap["Images"]; ok {
		if valImagesCap != nil {
			var valueForImages ResourceUriElement
			err = json.Unmarshal(*valImagesCap, &valueForImages)
			if err != nil {
				return err
			}
			this.Images = &valueForImages
		}
	}
	
	if valLayoutSlides, ok := objMap["layoutSlides"]; ok {
		if valLayoutSlides != nil {
			var valueForLayoutSlides ResourceUriElement
			err = json.Unmarshal(*valLayoutSlides, &valueForLayoutSlides)
			if err != nil {
				return err
			}
			this.LayoutSlides = &valueForLayoutSlides
		}
	}
	if valLayoutSlidesCap, ok := objMap["LayoutSlides"]; ok {
		if valLayoutSlidesCap != nil {
			var valueForLayoutSlides ResourceUriElement
			err = json.Unmarshal(*valLayoutSlidesCap, &valueForLayoutSlides)
			if err != nil {
				return err
			}
			this.LayoutSlides = &valueForLayoutSlides
		}
	}
	
	if valMasterSlides, ok := objMap["masterSlides"]; ok {
		if valMasterSlides != nil {
			var valueForMasterSlides ResourceUriElement
			err = json.Unmarshal(*valMasterSlides, &valueForMasterSlides)
			if err != nil {
				return err
			}
			this.MasterSlides = &valueForMasterSlides
		}
	}
	if valMasterSlidesCap, ok := objMap["MasterSlides"]; ok {
		if valMasterSlidesCap != nil {
			var valueForMasterSlides ResourceUriElement
			err = json.Unmarshal(*valMasterSlidesCap, &valueForMasterSlides)
			if err != nil {
				return err
			}
			this.MasterSlides = &valueForMasterSlides
		}
	}

    return nil
}

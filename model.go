package javatools

import (
	"encoding/xml"
)

type javaUpdateMap struct {
	XMLName  xml.Name  `xml:"java-update-map"`
	Version  string    `xml:"version,attr"`
	Mappings []mapping `xml:"mapping"`
}

type mapping struct {
	// XMLName xml.Name `xml:"mapping"`
	Version string `xml:"version"`
	Url     string `xml:"url"`
}

type javaUpdate struct {
	XMLName      xml.Name      `xml:"java-update"`
	Informations []information `xml:"information"`
}

type information struct {
	Version string `xml:"version"`
}

type JavaVersion struct {
	Major  int
	Minor  int
	Patch  int
	Update int
	Build  int
}

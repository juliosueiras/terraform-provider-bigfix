package bigfix

import (
	"encoding/xml"
	"log"
)

//=====================================================================================

// BES : Contains XML Structure for post request
type BESFixletRequest struct {
	XMLName                       xml.Name `xml:"BES"`
	Xmlns_xsi                     string   `xml:"xmlns:xsi,attr,omitempty"`
	Xsi_noNamespaceSchemaLocation string   `xml:"xsi:noNamespaceSchemaLocation,attr,omitempty"`
	Text                          string   `xml:",chardata"`
	Fixlet                        Fixlet   `xml:"Fixlet"`
}

// Fixlet :
type Fixlet struct {
	XMLName           xml.Name       `xml:"Fixlet"`
	Text              string         `xml:",chardata"`
	Title             string         `xml:"Title"`
	Description       string         `xml:"Description"`
	Relevances        []Relevance    `xml:"Relevance"`
	Category          string         `xml:"Category"`
	DownloadSize      int            `xml:"DownloadSize,omitempty"`
	Source            string         `xml:"Source,omitempty"`
	SourceID          string         `xml:"SourceID"`
	SourceReleaseDate string         `xml:"SourceReleaseDate,omitempty"`
	SourceSeverity    string         `xml:"SourceSeverity"`
	CVENames          string         `xml:"CVENames"`
	SANSID            string         `xml:"SANSID"`
	MIMEField         []MIMEField    `xml:"MIMEField,omitempty"`
	Domain            string         `xml:"Domain,omitempty"`
	Delay             string         `xml:"Delay,omitempty"`
	DefaultAction     FixletAction   `xml:"DefaultAction,omitempty"`
	Actions           []FixletAction `xml:"Action,omitempty"`
}

type Relevance struct {
	Text string `xml:",chardata"`
}

type MIMEField struct {
	Name  string `xml:"Name"`
	Value string `xml:"Value"`
}

type FixletAction struct {
	//XMLName      xml.Name     `xml:"Action"`
	ID           string            `xml:"ID,attr"`
	Description  ActionDescription `xml:"Description"`
	ActionScript ActionScript      `xml:"ActionScript,omitempty	"`
}

type ActionDescription struct {
	XMLName  xml.Name `xml:"Description"`
	PreLink  string   `xml:"PreLink"`
	Link     string   `xml:"Link"`
	PostLink string   `xml:"PostLink"`
}

type ActionScript struct {
	MIMEType string `xml:"MIMEType,attr"`
	Text     string `xml:",chardata"`
}

// ParseMAGXMLMarshal : will create XML Structure of MAG
func ParseFixletXMLMarshal(requestbody BESFixletRequest) []byte {

	fixletBody, err := xml.MarshalIndent(&requestbody, "", "\t")
	if err != nil {
		log.Println(err)
	}

	return []byte(xml.Header + string(fixletBody))
}

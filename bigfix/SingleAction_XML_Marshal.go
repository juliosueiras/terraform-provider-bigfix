package bigfix

import (
	"encoding/xml"
	"log"
)

//=====================================================================================

// BES : Contains XML Structure for post request
type BESSingleActionRequest struct {
	XMLName                       xml.Name     `xml:"BES"`
	Xmlns_xsi                     string       `xml:"xmlns:xsi,attr,omitempty"`
	Xsi_noNamespaceSchemaLocation string       `xml:"xsi:noNamespaceSchemaLocation,attr,omitempty"`
	Text                          string       `xml:",chardata"`
	SingleAction                  SingleAction `xml:"SingleAction"`
}

// SingleAction :
type SingleAction struct {
	XMLName           xml.Name             `xml:"SingleAction"`
	Text              string               `xml:",chardata"`
	Title             string               `xml:"Title"`
	Description       string               `xml:"Description"`
	Relevances        []Relevance          `xml:"Relevance"`
	Category          string               `xml:"Category"`
	DownloadSize      int                  `xml:"DownloadSize,omitempty"`
	Source            string               `xml:"Source,omitempty"`
	SourceID          string               `xml:"SourceID"`
	SourceReleaseDate string               `xml:"SourceReleaseDate,omitempty"`
	SourceSeverity    string               `xml:"SourceSeverity"`
	CVENames          string               `xml:"CVENames"`
	SANSID            string               `xml:"SANSID"`
	MIMEField         []MIMEField          `xml:"MIMEField,omitempty"`
	Domain            string               `xml:"Domain,omitempty"`
	Delay             string               `xml:"Delay,omitempty"`
	DefaultAction     SingleActionAction   `xml:"DefaultAction,omitempty"`
	Actions           []SingleActionAction `xml:"Action,omitempty"`
}

type SingleActionAction struct {
	//XMLName      xml.Name     `xml:"Action"`
	ID           string            `xml:"ID,attr"`
	Description  ActionDescription `xml:"Description"`
	ActionScript ActionScript      `xml:"ActionScript,omitempty	"`
}

// ParseMAGXMLMarshal : will create XML Structure of MAG
func ParseSingleActionXMLMarshal(requestbody BESSingleActionRequest) []byte {

	fixletBody, err := xml.MarshalIndent(&requestbody, "", "\t")
	if err != nil {
		log.Println(err)
	}

	return []byte(xml.Header + string(fixletBody))
}

package bigfix

import (
	"encoding/xml"
	"log"
)

//=====================================================================================

// BES : Contains XML Structure for post request
type BESTaskRequest struct {
	XMLName                       xml.Name `xml:"BES"`
	Xmlns_xsi                     string   `xml:"xmlns:xsi,attr,omitempty"`
	Xsi_noNamespaceSchemaLocation string   `xml:"xsi:noNamespaceSchemaLocation,attr,omitempty"`
	Text                          string   `xml:",chardata"`
	Task                          Task     `xml:"Task"`
}

// Task :
type Task struct {
	XMLName           xml.Name       `xml:"Task"`
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

func ParseTaskXMLMarshal(requestbody BESTaskRequest) []byte {

	taskBody, err := xml.MarshalIndent(&requestbody, "", "\t")
	if err != nil {
		log.Println(err)
	}

	return []byte(xml.Header + string(taskBody))
}

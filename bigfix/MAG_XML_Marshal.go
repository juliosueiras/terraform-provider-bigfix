package bigfix

import (
	"encoding/xml"
	"log"
)

//=====================================================================================

// BES : Contains XML Structure for post request
type BES struct {
	//XMLName                       xml.Name            `xml:"BES"`
	Xmlns_xsi                     string              `xml:"xmlns:xsi,attr,omitempty"`
	Xsi_noNamespaceSchemaLocation string              `xml:"xsi:noNamespaceSchemaLocation,attr,omitempty"`
	Text                          string              `xml:",chardata"`
	MultipleActionGroup           MultipleActionGroup `xml:"MultipleActionGroup,omitempty	"`
}

// MultipleActionGroup :
type MultipleActionGroup struct {
	XMLName             xml.Name              `xml:"MultipleActionGroup"`
	Text                string                `xml:",chardata"`
	Title               string                `xml:"Title"`
	Relevance           string                `xml:"Relevance,omitempty"`
	SourcedMemberAction []SourcedMemberAction `xml:"SourcedMemberAction,omitempty	"`
	Target              Target                `xml:"Target,omitempty"`
}

// SourcedMemberAction :
type SourcedMemberAction struct {
	XMLName      xml.Name     `xml:"SourcedMemberAction"`
	Text         string       `xml:",chardata"`
	SourceFixlet SourceFixlet `xml:"SourceFixlet,omitempty	"`
}

// SourceFixlet :
type SourceFixlet struct {
	XMLName  xml.Name `xml:"SourceFixlet"`
	Text     string   `xml:",chardata"`
	Sitename string   `xml:"Sitename"`
	FixletID string   `xml:"FixletID"`
	Action   string   `xml:"Action"`
}

//Target :
type Target struct {
	XMLName    xml.Name `xml:"Target"`
	Text       string   `xml:",chardata"`
	ComputerID string   `xml:"ComputerID"`
}

// Parameter :
type Parameter struct {
	XMLName xml.Name `xml:"Parameter"`
	Text    string   `xml:",chardata"`
	Name    string   `xml:"Name,attr"`
}

// ParseMAGXMLMarshal : will create XML Structure of MAG
func ParseMAGXMLMarshal(magDetails MAGFile, target string, siteName []string, memberlistStruct []SourcedMemberAction) []byte {

	requestbody := BES{
		Xmlns_xsi:                     "http://www.w3.org/2001/XMLSchema-instance",
		Xsi_noNamespaceSchemaLocation: "BES.xsd",

		MultipleActionGroup: MultipleActionGroup{
			Title:     magDetails.MultipleActionGroup.Title + " [ " + target + " ]",
			Relevance: magDetails.MultipleActionGroup.Relevance,
			// for multiple members
			SourcedMemberAction: memberlistStruct,
			Target: Target{
				ComputerID: target,
			},
		},
	}

	magBody, err := xml.MarshalIndent(&requestbody, "", "\t")
	if err != nil {
		log.Println(err)
	}

	return magBody
}

package bigfix

import "encoding/xml"

//=====================================================================================

// ComputerDetails : Contains details of computer.
//
// API_Utility : GetComputerDetailAPI()
type ComputerDetails struct {
	XMLName                   xml.Name `xml:"BESAPI"`
	Text                      string   `xml:",chardata"`
	Xsi                       string   `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string   `xml:"noNamespaceSchemaLocation,attr"`
	Query                     struct {
		Text   string `xml:",chardata"`
		Result struct {
			Text  string `xml:",chardata"`
			Tuple struct {
				Text   string `xml:",chardata"`
				Answer []struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"Answer"`
			} `xml:"Tuple"`
		} `xml:"Result"`
		Evaluation struct {
			Text      string `xml:",chardata"`
			Time      string `xml:"Time"`
			Plurality string `xml:"Plurality"`
		} `xml:"Evaluation"`
		Error string `xml:"Error"`
	} `xml:"Query"`
}

//=====================================================================================

// CountComputerStruct : To Store count of computer
//
// API_Utility : GetComputerCountAPI()
type CountComputerStruct struct {
	XMLName                   xml.Name `xml:"BESAPI"`
	Text                      string   `xml:",chardata"`
	Xsi                       string   `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string   `xml:"noNamespaceSchemaLocation,attr"`
	Query                     struct {
		Text     string `xml:",chardata"`
		Resource string `xml:"Resource,attr"`
		Result   struct {
			Text   string `xml:",chardata"`
			Answer struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"Answer"`
		} `xml:"Result"`
		Evaluation struct {
			Text      string `xml:",chardata"`
			Time      string `xml:"Time"`
			Plurality string `xml:"Plurality"`
		} `xml:"Evaluation"`
	} `xml:"Query"`
}

//=====================================================================================

// MAGFile : Structure to store xml file data of MAG
type MAGFile struct {
	XMLName                       xml.Name `xml:"BES"`
	Xmlns_xsi                     string   `xml:"xmlns:xsi,attr,omitempty"`
	Xsi_noNamespaceSchemaLocation string   `xml:"xsi:noNamespaceSchemaLocation,attr,omitempty"`
	MultipleActionGroup           struct {
		Text      string `xml:",chardata"`
		Title     string `xml:"Title"`
		Relevance string `xml:"Relevance"`
	} `xml:"MultipleActionGroup"`
}

//=====================================================================================

// FixletMetaData : Struct for fixlet data
//
// API_Utility : GetFixletMetaDataAPI()
type FixletMetaData struct {
	XMLName                   xml.Name `xml:"BESAPI"`
	Text                      string   `xml:",chardata"`
	Xsi                       string   `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string   `xml:"noNamespaceSchemaLocation,attr"`
	Query                     Query    `xml:"Query"`
}

// Query :
type Query struct {
	Text       string     `xml:",chardata"`
	Resource   string     `xml:"Resource,attr"`
	Result     Result     `xml:"Result"`
	Evaluation Evaluation `xml:"Evaluation"`
}

// Result :
type Result struct {
	XMLName xml.Name `xml:"Result"`
	Text    string   `xml:",chardata"`
	Tuple   Tuple    `xml:"Tuple"`
}

//Tuple :
type Tuple struct {
	XMLName xml.Name `xml:"Tuple"`
	Text    string   `xml:",chardata"`
	Answer  []Answer `xml:"Answer"`
}

//Answer :
type Answer struct {
	Text string `xml:",chardata"`
	Type string `xml:"type,attr"`
}

//Evaluation :
type Evaluation struct {
	XMLName   xml.Name `xml:"Evaluation"`
	Text      string   `xml:",chardata"`
	Time      string   `xml:"Time"`
	Plurality string   `xml:"Plurality"`
}

//=====================================================================================

// ActionCreationResponse : To check whether action is created.
//
// API_Utility : CreateMAGAPI()
type ActionCreationResponse struct {
	XMLName                   xml.Name `xml:"BESAPI"`
	Text                      string   `xml:",chardata"`
	Xsi                       string   `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string   `xml:"noNamespaceSchemaLocation,attr"`
	Action                    struct {
		Text         string `xml:",chardata"`
		Resource     string `xml:"Resource,attr"`
		LastModified string `xml:"LastModified,attr"`
		Name         string `xml:"Name"`
		ID           string `xml:"ID"`
	} `xml:"Action"`
}

type FixletCreationResponse struct {
	XMLName                   xml.Name `xml:"BESAPI"`
	Text                      string   `xml:",chardata"`
	Xsi                       string   `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string   `xml:"noNamespaceSchemaLocation,attr"`
	Fixlet                    struct {
		Text         string `xml:",chardata"`
		Resource     string `xml:"Resource,attr"`
		LastModified string `xml:"LastModified,attr"`
		Name         string `xml:"Name"`
		ID           string `xml:"ID"`
	} `xml:"Fixlet"`
}

type TaskCreationResponse struct {
	XMLName                   xml.Name `xml:"BESAPI"`
	Text                      string   `xml:",chardata"`
	Xsi                       string   `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string   `xml:"noNamespaceSchemaLocation,attr"`
	Task                      struct {
		Text         string `xml:",chardata"`
		Resource     string `xml:"Resource,attr"`
		LastModified string `xml:"LastModified,attr"`
		Name         string `xml:"Name"`
		ID           string `xml:"ID"`
	} `xml:"Task"`
}

type FileCreationResponse struct {
	XMLName                   xml.Name `xml:"BESAPI"`
	Text                      string   `xml:",chardata"`
	Xsi                       string   `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string   `xml:"noNamespaceSchemaLocation,attr"`
	FileUpload                struct {
		Text     string `xml:",chardata"`
		Resource string `xml:"Resource,attr"`
		Name     string `xml:"Name"`
		URL      string `xml:"URL"`
		Size     int    `xml:"Size"`
		SHA1     string `xml:"SHA1"`
		SHA256   string `xml:"SHA256"`
	} `xml:"FileUpload"`
}

type FileReferencesResponse struct {
	XMLName                   xml.Name `xml:"BESAPI"`
	Text                      string   `xml:",chardata"`
	Xsi                       string   `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string   `xml:"noNamespaceSchemaLocation,attr"`
	FileUploadReferences      []struct {
		Text      string `xml:",chardata"`
		Resource  string `xml:"Resource,attr"`
		ID        string `xml:"ID"`
		UserID    string `xml:"UserID"`
		IsPrivate int    `xml:"IsPrivate"`
	} `xml:"FileUploadReference"`
}

//=====================================================================================

// Action : Contains details of action
type Action struct {
	XMLName                   xml.Name `xml:"BES"`
	Text                      string   `xml:",chardata"`
	Xsi                       string   `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string   `xml:"noNamespaceSchemaLocation,attr"`
	MultipleActionGroup       struct {
		Text         string `xml:",chardata"`
		Title        string `xml:"Title"`
		Relevance    string `xml:"Relevance"`
		MemberAction []struct {
			Text         string `xml:",chardata"`
			Title        string `xml:"Title"`
			Relevance    string `xml:"Relevance"`
			ActionScript struct {
				Text     string `xml:",chardata"`
				MIMEType string `xml:"MIMEType,attr"`
			} `xml:"ActionScript"`
			SuccessCriteria struct {
				Text   string `xml:",chardata"`
				Option string `xml:"Option,attr"`
			} `xml:"SuccessCriteria"`
			IncludeInGroupRelevance string `xml:"IncludeInGroupRelevance"`
		} `xml:"MemberAction"`
		Settings struct {
			Text                    string `xml:",chardata"`
			PreActionShowUI         string `xml:"PreActionShowUI"`
			HasRunningMessage       string `xml:"HasRunningMessage"`
			HasTimeRange            string `xml:"HasTimeRange"`
			HasStartTime            string `xml:"HasStartTime"`
			HasEndTime              string `xml:"HasEndTime"`
			EndDateTimeLocalOffset  string `xml:"EndDateTimeLocalOffset"`
			HasDayOfWeekConstraint  string `xml:"HasDayOfWeekConstraint"`
			UseUTCTime              string `xml:"UseUTCTime"`
			ActiveUserRequirement   string `xml:"ActiveUserRequirement"`
			ActiveUserType          string `xml:"ActiveUserType"`
			HasWhose                string `xml:"HasWhose"`
			PreActionCacheDownload  string `xml:"PreActionCacheDownload"`
			Reapply                 string `xml:"Reapply"`
			HasReapplyLimit         string `xml:"HasReapplyLimit"`
			ReapplyLimit            string `xml:"ReapplyLimit"`
			HasReapplyInterval      string `xml:"HasReapplyInterval"`
			HasRetry                string `xml:"HasRetry"`
			HasTemporalDistribution string `xml:"HasTemporalDistribution"`
			ContinueOnErrors        string `xml:"ContinueOnErrors"`
			PostActionBehavior      struct {
				Text     string `xml:",chardata"`
				Behavior string `xml:"Behavior,attr"`
			} `xml:"PostActionBehavior"`
			IsOffer string `xml:"IsOffer"`
		} `xml:"Settings"`
		SettingsLocks struct {
			Text            string `xml:",chardata"`
			ActionUITitle   string `xml:"ActionUITitle"`
			PreActionShowUI string `xml:"PreActionShowUI"`
			PreAction       struct {
				Chardata         string `xml:",chardata"`
				Text             string `xml:"Text"`
				AskToSaveWork    string `xml:"AskToSaveWork"`
				ShowActionButton string `xml:"ShowActionButton"`
				ShowCancelButton string `xml:"ShowCancelButton"`
				DeadlineBehavior string `xml:"DeadlineBehavior"`
				ShowConfirmation string `xml:"ShowConfirmation"`
			} `xml:"PreAction"`
			HasRunningMessage string `xml:"HasRunningMessage"`
			RunningMessage    struct {
				Chardata string `xml:",chardata"`
				Text     string `xml:"Text"`
			} `xml:"RunningMessage"`
			TimeRange              string `xml:"TimeRange"`
			StartDateTimeOffset    string `xml:"StartDateTimeOffset"`
			EndDateTimeOffset      string `xml:"EndDateTimeOffset"`
			DayOfWeekConstraint    string `xml:"DayOfWeekConstraint"`
			ActiveUserRequirement  string `xml:"ActiveUserRequirement"`
			ActiveUserType         string `xml:"ActiveUserType"`
			Whose                  string `xml:"Whose"`
			PreActionCacheDownload string `xml:"PreActionCacheDownload"`
			Reapply                string `xml:"Reapply"`
			ReapplyLimit           string `xml:"ReapplyLimit"`
			RetryCount             string `xml:"RetryCount"`
			RetryWait              string `xml:"RetryWait"`
			TemporalDistribution   string `xml:"TemporalDistribution"`
			ContinueOnErrors       string `xml:"ContinueOnErrors"`
			PostActionBehavior     struct {
				Chardata    string `xml:",chardata"`
				Behavior    string `xml:"Behavior"`
				AllowCancel string `xml:"AllowCancel"`
				Deadline    string `xml:"Deadline"`
				Title       string `xml:"Title"`
				Text        string `xml:"Text"`
			} `xml:"PostActionBehavior"`
			IsOffer              string `xml:"IsOffer"`
			AnnounceOffer        string `xml:"AnnounceOffer"`
			OfferCategory        string `xml:"OfferCategory"`
			OfferDescriptionHTML string `xml:"OfferDescriptionHTML"`
		} `xml:"SettingsLocks"`
		Target struct {
			Text         string `xml:",chardata"`
			AllComputers string `xml:"AllComputers"`
		} `xml:"Target"`
	} `xml:"MultipleActionGroup"`
}

//=====================================================================================

// SourcedMemeberContent : Contains members(relevant fixlets) from particular site.
//
// API_Utility : GetRelevantFixletsAPI()
type SourcedMemeberContent struct {
	XMLName                   xml.Name `xml:"BESAPI"`
	Text                      string   `xml:",chardata"`
	Xsi                       string   `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string   `xml:"noNamespaceSchemaLocation,attr"`
	Query                     struct {
		Text     string `xml:",chardata"`
		Resource string `xml:"Resource,attr"`
		Result   struct {
			Text  string `xml:",chardata"`
			Tuple []struct {
				Text   string `xml:",chardata"`
				Answer []struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"Answer"`
			} `xml:"Tuple"`
		} `xml:"Result"`
		Evaluation struct {
			Text      string `xml:",chardata"`
			Time      string `xml:"Time"`
			Plurality string `xml:"Plurality"`
		} `xml:"Evaluation"`
		Error string `xml:"Error"`
	} `xml:"Query"`
}

//=====================================================================================

// CountofRelevantFixlet : Contains count of relevant fixlets for a particular site.
//
// API_Utility : GetCountofRelevantFixletsAPI()
type CountofRelevantFixlet struct {
	XMLName                   xml.Name `xml:"BESAPI"`
	Text                      string   `xml:",chardata"`
	Xsi                       string   `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string   `xml:"noNamespaceSchemaLocation,attr"`
	Query                     struct {
		Text     string `xml:",chardata"`
		Resource string `xml:"Resource,attr"`
		Result   struct {
			Text   string `xml:",chardata"`
			Answer struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"Answer"`
		} `xml:"Result"`
		Evaluation struct {
			Text      string `xml:",chardata"`
			Time      string `xml:"Time"`
			Plurality string `xml:"Plurality"`
		} `xml:"Evaluation"`
	} `xml:"Query"`
}

//=====================================================================================

// AllSites : Contains list of all sites.
//
// API_Utility : GetAllSitesAPI()
type AllSites struct {
	XMLName                   xml.Name `xml:"BESAPI"`
	Text                      string   `xml:",chardata"`
	Xsi                       string   `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string   `xml:"noNamespaceSchemaLocation,attr"`
	ExternalSite              []struct {
		Text        string `xml:",chardata"`
		Resource    string `xml:"Resource,attr"`
		Name        string `xml:"Name"`
		GatherURL   string `xml:"GatherURL"`
		DisplayName string `xml:"DisplayName"`
	} `xml:"ExternalSite"`
	OperatorSite struct {
		Text      string `xml:",chardata"`
		Resource  string `xml:"Resource,attr"`
		Name      string `xml:"Name"`
		GatherURL string `xml:"GatherURL"`
	} `xml:"OperatorSite"`
	ActionSite struct {
		Text      string `xml:",chardata"`
		Resource  string `xml:"Resource,attr"`
		Name      string `xml:"Name"`
		GatherURL string `xml:"GatherURL"`
	} `xml:"ActionSite"`
}

//=====================================================================================

// ComputerName : Contains name of computer.
//
// API_Utility : GetComputerNameAPI()
type ComputerName struct {
	XMLName                   xml.Name `xml:"BESAPI"`
	Text                      string   `xml:",chardata"`
	Xsi                       string   `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string   `xml:"noNamespaceSchemaLocation,attr"`
	Query                     struct {
		Text     string `xml:",chardata"`
		Resource string `xml:"Resource,attr"`
		Result   struct {
			Text   string `xml:",chardata"`
			Answer struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"Answer"`
		} `xml:"Result"`
		Evaluation struct {
			Text      string `xml:",chardata"`
			Time      string `xml:"Time"`
			Plurality string `xml:"Plurality"`
		} `xml:"Evaluation"`
	} `xml:"Query"`
}

// ActionStatus : Contains name of computer.
//
// API_Utility : GetActionStatus()
type ActionStatus struct {
	XMLName                   xml.Name `xml:"BESAPI"`
	Text                      string   `xml:",chardata"`
	Xsi                       string   `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string   `xml:"noNamespaceSchemaLocation,attr"`
	ActionResults             struct {
		Text               string `xml:",chardata"`
		Resource           string `xml:"Resource,attr"`
		ActionID           string `xml:"ActionID"`
		Status             string `xml:"Status"`
		DateIssued         string `xml:"DateIssued"`
		MemberActionResult []struct {
			Text     string `xml:",chardata"`
			Resource string `xml:"Resource,attr"`
			ActionID string `xml:"ActionID"`
		} `xml:"MemberActionResult"`
	} `xml:"ActionResults"`
}

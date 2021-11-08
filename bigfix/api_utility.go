package bigfix

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

// GetComputerCount : Will return Count of computers
func GetComputerCount(config Config) string {
	var count string
	url := GetComputerCountAPI(config.ServerIP, config.Port)
	log.Println("[DEBUG] Getting Count of total computers...")
	response, err := config.BfxConnection(GET, url, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var countComputerstruct CountComputerStruct
	if response != nil {
		defer response.Body.Close()

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}
		log.Println("[DEBUG] response is :")
		log.Println(string(data))
		xml.Unmarshal([]byte(data), &countComputerstruct)

		count = countComputerstruct.Query.Result.Answer.Text
		return count
	}
	log.Println("[DEBUG] Computer list is empty !!!")
	return ""
}

// DeleteFixlet : Will delete action from console
func DeleteFixlet(config Config, id string, siteName string, siteType string) (bool, error) {
	url := GetDeleteFixletAPI(config.ServerIP, config.Port, id, siteName, siteType)
	response, err := config.BfxConnection(DELETE, url, nil)
	if err != nil {
		return false, err
	}
	if response.StatusCode == http.StatusOK {
		log.Println("[DEBUG] Action Deleted Successfully")
		return true, nil
	}
	data, err1 := ioutil.ReadAll(response.Body)
	if err1 != nil {
		log.Printf("[DEBUG] Error reading response body : %v\n", err1)
		return false, err
	}
	return false, fmt.Errorf(string(data))
}

func DeleteTask(config Config, id string, siteName string, siteType string) (bool, error) {
	url := GetDeleteTaskAPI(config.ServerIP, config.Port, id, siteName, siteType)
	response, err := config.BfxConnection(DELETE, url, nil)
	if err != nil {
		return false, err
	}
	if response.StatusCode == http.StatusOK {
		log.Println("[DEBUG] Action Deleted Successfully")
		return true, nil
	}
	data, err1 := ioutil.ReadAll(response.Body)
	if err1 != nil {
		log.Printf("[DEBUG] Error reading response body : %v\n", err1)
		return false, err
	}
	return false, fmt.Errorf(string(data))
}

// DeleteAction : Will delete action from console
func DeleteAction(config Config, id string) (bool, error) {
	url := GetDeleteActionAPI(config.ServerIP, config.Port, id)
	response, err := config.BfxConnection(DELETE, url, nil)
	if err != nil {
		return false, err
	}
	if response.StatusCode == http.StatusOK {
		log.Println("[DEBUG] Action Deleted Successfully")
		return true, nil
	}
	data, err1 := ioutil.ReadAll(response.Body)
	if err1 != nil {
		log.Printf("[DEBUG] Error reading response body : %v\n", err1)
		return false, err
	}
	return false, fmt.Errorf(string(data))
}

func DeleteUploadFile(config Config, id string) (bool, error) {
	url := GetUploadFileDetailReferencesAPI(config.ServerIP, config.Port, id)
	response, err := config.BfxConnection(GET, url, nil)
	if err != nil {
		return false, err
	}

	// Get action
	var fileResponse FileReferencesResponse
	if response != nil {
		defer response.Body.Close()

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Printf("[DEBUG] Error reading response body : %v\n", err)
			return false, err
		}
		xml.Unmarshal([]byte(data), &fileResponse)
	}
	for _, v := range fileResponse.FileUploadReferences {
		url := v.Resource
		response, err := config.BfxConnection(DELETE, url, nil)
		if err != nil {
			return false, err
		}
		if response.StatusCode == http.StatusOK {
			log.Println("[DEBUG] Action Deleted Successfully")
			continue
		}
		data, err1 := ioutil.ReadAll(response.Body)
		if err1 != nil {
			log.Printf("[DEBUG] Error reading response body : %v\n", err1)
			return false, err
		}
		return false, fmt.Errorf(string(data))
	}

	return true, nil
}

// SetSourcedMemberList : will set members for multiple action group
func SetSourcedMemberList(config Config, site []string, computerID string) ([]SourcedMemberAction, error) {

	var resultStruct []SourcedMemberAction
	var totalSize int = 0
	totalSite := len(site)
	length := make([]int, totalSite)

	length, totalSize = GetCountofRelevantFixlets(config, site, computerID)
	log.Printf("[DEBUG] total length of fixlets is %d", totalSize)
	if totalSize == 0 {
		log.Printf("[DEBUG] No relevant fixlets found from any site !!!")
		return nil, fmt.Errorf("No relevant fixlets found from any site")
	}

	// to store final list of members
	resultStruct = make([]SourcedMemberAction, totalSize)
	var k int = 0

	for i := 0; i < len(site); i++ {

		// to store members from site[i]
		var sourcedMemberContentStruct SourcedMemeberContent
		valid, err := checkSiteValidity(config, site[i])
		if err != nil {
			return nil, err
		}
		if valid {
			url := GetRelevantFixletsAPI(config.ServerIP, config.Port, site[i], computerID)
			response, err := config.BfxConnection(GET, url, nil)
			if err != nil {
				return nil, err

			}

			if response != nil {
				defer response.Body.Close()

				data, err := ioutil.ReadAll(response.Body)
				if err != nil {
					return nil, err
				}
				log.Println("[DEBUG] response is :")
				log.Println(string(data))
				xml.Unmarshal([]byte(data), &sourcedMemberContentStruct)

			}

		}
		fixletCount := length[i]
		for j := 0; j < fixletCount; j++ {
			resultStruct[k] = SourcedMemberAction{
				SourceFixlet: SourceFixlet{
					Sitename: sourcedMemberContentStruct.Query.Result.Tuple[j].Answer[0].Text,
					FixletID: sourcedMemberContentStruct.Query.Result.Tuple[j].Answer[1].Text,
					Action:   "Action1",
				},
			}
			// for final list
			k++
		}
	}
	return resultStruct, nil
}

// GetCountofRelevantFixlets : will return count of total relevent fixlets from multiple sites
func GetCountofRelevantFixlets(config Config, site []string, computerID string) ([]int, int) {

	var result CountofRelevantFixlet

	// length : to count relevant fixlets from each site
	length := make([]int, len(site))

	var total int = 0
	for i := 0; i < len(site); i++ {
		log.Printf("[DEBUG] for Site `%s` setting relevant fixlet members out of %d ", site[i], len(site))
		url := GetCountofRelevantFixletsAPI(config.ServerIP, config.Port, site[i], computerID)
		response, err := config.BfxConnection(GET, url, nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if response != nil {
			defer response.Body.Close()

			data, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println(err)
			}
			// log.Println("[DEBUG] response is :")
			// log.Println(string(data))
			xml.Unmarshal([]byte(data), &result)

			tempLength := result.Query.Result.Answer.Text
			length[i], _ = strconv.Atoi(tempLength)
			log.Printf("[DEBUG] length of %d site is %d", i, length[i])
			total = total + length[i]
		}

	}
	log.Printf("[DEBUG] total length of fixlets is %d", total)
	return length, total
}

func checkActionDestroyed(data []byte, title string) (bool, error) {
	var actionStruct Action
	err := xml.Unmarshal(data, &actionStruct)
	if err != nil {
		return false, fmt.Errorf("Error while unmarshaling data %s", err)
	}
	if actionStruct.MultipleActionGroup.Title == title {
		return true, nil
	}
	return false, nil
}

// checkSiteValidity : Function will check if given site name is valid
func checkSiteValidity(config Config, sitename string) (bool, error) {
	url := GetAllSitesAPI(config.ServerIP, config.Port)
	response, err := config.BfxConnection(GET, url, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if response != nil {
		defer response.Body.Close()

		var sitesStruct AllSites
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return false, err
		}
		log.Println("[DEBUG] response is :")
		log.Println(string(data))
		xml.Unmarshal([]byte(data), &sitesStruct)
		total := len(sitesStruct.ExternalSite)
		for i := 0; i < total; i++ {
			if sitesStruct.ExternalSite[i].Name == sitename {
				return true, nil
			} else if sitesStruct.OperatorSite.Name == sitename {
				return true, nil
			} else if sitesStruct.ActionSite.Name == sitename {
				return true, nil
			}
		}
	}
	// Not Found in list
	return false, fmt.Errorf("Given site Name `%s` is not valid", sitename)
}

// getNameOfComputer : Function will return name of computer
func getNameOfComputer(config Config, id string) string {
	var computerNameStruct ComputerName
	url := GetComputerNameAPI(config.ServerIP, config.Port, id)
	response, err := config.BfxConnection(GET, url, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if response != nil {
		defer response.Body.Close()

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return ""
		}
		log.Println("[DEBUG] response is :")
		log.Println(string(data))
		xml.Unmarshal([]byte(data), &computerNameStruct)
		return computerNameStruct.Query.Result.Answer.Text
	}
	return ""
}

// GetActionStatus : will return status of action
func GetActionStatus(config Config, actionID string) string {
	var actionStatus ActionStatus
	url := GetActionStatusAPI(config.ServerIP, config.Port, actionID)
	response, err := config.BfxConnection(GET, url, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if response != nil {
		defer response.Body.Close()

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return ""
		}
		log.Println("[DEBUG] status response is :")
		log.Println(string(data))
		xml.Unmarshal([]byte(data), &actionStatus)
		log.Println("[DEBUG] ActionStatus :", actionStatus.ActionResults.Status)
		return actionStatus.ActionResults.Status
	}
	return ""
}

package sap_api_caller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sap-api-integrations-equipment-master-creates/SAP_API_Caller/requests"
	sap_api_output_formatter "sap-api-integrations-equipment-master-creates/SAP_API_Output_Formatter"
	"strings"
	"sync"
	"github.com/latonaio/golang-logging-library/logger"
	sap_api_post_header_setup "github.com/latonaio/sap-api-post-header-setup"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	postClient      *sap_api_post_header_setup.SAPPostClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, postClient *sap_api_post_header_setup.SAPPostClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		postClient:      postClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncPostEquipmentMaster(
	equipment         *requests.Equipment,
	partner           *requests.Partner,
	accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	for _, fn := range accepter {
		switch fn {
		case "Equipment":
			func() {
				c.Equipment(equipment)
				wg.Done()
			}()
		case "Partner":
			func() {
				c.Partner(partner)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}
	wg.Wait()
}

func (c *SAPAPICaller) Equipment(equipment *requests.Equipment) {
	outputDataEquipment, err := c.callEquipmentSrvAPIRequirementEquipment("Equipment", equipment)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataEquipment)
}
func (c *SAPAPICaller) callEquipmentSrvAPIRequirementEquipment(api string, equipment *requests.Equipment) (*sap_api_output_formatter.Equipment, error) {
	body, err := json.Marshal(equipment)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_EQUIPMENT", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.postClient.POST(url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}
	data, err := sap_api_output_formatter.ConvertToEquipment(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Partner(partner *requests.Partner) {
	outputDataPartner, err := c.callEquipmentSrvAPIRequirementPartner("EquipmentPartner", partner)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataPartner)
}

func (c *SAPAPICaller) callEquipmentSrvAPIRequirementPartner(api string, partner *requests.Partner) (*sap_api_output_formatter.Partner, error) {
	body, err := json.Marshal(partner)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_EQUIPMENT", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.postClient.POST(url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}
	data, err := sap_api_output_formatter.ConvertToPartner(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) addQuerySAPClient(params map[string]string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["sap-client"] = c.sapClientNumber
	return params
}

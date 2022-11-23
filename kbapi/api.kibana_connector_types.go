package kbapi

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

const (
	basePathKibanaConnectorTypes = "/api/actions/connector_types"
)

type KibanaConnectorType struct {
	ID                     string   `json:"id"`
	Name                   string   `json:"name"`
	Enabled                bool     `json:"enabled"`
	EnabledInConfig        bool     `json:"enabled_in_config"`
	EnabledInLicense       bool     `json:"enabled_in_license"`
	MinimumLicenseRequired string   `json:"minimum_license_required"`
	SupportedFeatureIDs    []string `json:"supported_feature_ids"`
}

// String permit to return KibanaConnectorTypes object as JSON string
func (k KibanaConnectorTypes) String() string {
	json, _ := json.Marshal(k)
	return string(json)
}

type KibanaConnectorTypes []KibanaConnectorType

type KibanaConnectorTypesList func() (KibanaConnectorTypes, error)

func newKibanaConnectorTypesListFunc(c *resty.Client) KibanaConnectorTypesList {
	return func() (KibanaConnectorTypes, error) {
		resp, err := c.R().Get(basePathKibanaConnectorTypes)
		if err != nil {
			return nil, err
		}
		log.Debug("Response: ", resp)
		if resp.StatusCode() >= 300 {
			return nil, NewAPIError(resp.StatusCode(), resp.String())
		}
		kibanaConnectorTypes := make(KibanaConnectorTypes, 0, 1)
		err = json.Unmarshal(resp.Body(), &kibanaConnectorTypes)
		if err != nil {
			return nil, err
		}
		log.Debug("KibanaConnectorTypes: ", kibanaConnectorTypes)

		return kibanaConnectorTypes, nil
	}
}

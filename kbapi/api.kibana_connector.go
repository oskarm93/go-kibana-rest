package kbapi

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

const (
	basePathKibanaConnector = "/api/actions/connector"
)

type KibanaConnectorConfig map[string]interface{}
type KibanaConnectorSecrets map[string]interface{}

type KibanaConnector struct {
	ID                string                `json:"id"`
	Name              string                `json:"name"`
	ConnectorTypeID   string                `json:"connector_type_id"`
	IsPreconfigured   bool                  `json:"is_preconfigured"`
	IsDeprecated      bool                  `json:"is_deprecated"`
	IsMissingSecrets  bool                  `json:"is_missing_secrets,omitempty"`
	ReferencedByCount int                   `json:"referenced_by_count,omitempty"`
	Config            KibanaConnectorConfig `json:"config,omitempty"`
}

type KibanaConnectorCreateParams struct {
	Name            string                 `json:"name"`
	ConnectorTypeID string                 `json:"connector_type_id,omitempty"`
	Config          KibanaConnectorConfig  `json:"config,omitempty"`
	Secrets         KibanaConnectorSecrets `json:"secrets,omitempty"`
}

type KibanaConnectors []KibanaConnector

// KibanaConnectorGet permit to get connector
type KibanaConnectorGet func(id string) (*KibanaConnector, error)

// KibanaConnectorList permit to get all connectors
type KibanaConnectorList func() (KibanaConnectors, error)

// KibanaConnectorCreate permit to create connector
type KibanaConnectorCreate func(kibanaConnectorCreateParams *KibanaConnectorCreateParams) (*KibanaConnector, error)

// KibanaConnectorDelete permit to delete connector
type KibanaConnectorDelete func(id string) error

// KibanaConnectorUpdate permit to update connector
type KibanaConnectorUpdate func(id string, kibanaConnectorCreateParams *KibanaConnectorCreateParams) (*KibanaConnector, error)

// String permit to return KibanaConnector object as JSON string
func (k *KibanaConnector) String() string {
	json, _ := json.Marshal(k)
	return string(json)
}

// newKibanaConnectorGetFunc permit to get the kibana connector with it id
func newKibanaConnectorGetFunc(c *resty.Client) KibanaConnectorGet {
	return func(id string) (*KibanaConnector, error) {
		if id == "" {
			return nil, NewAPIError(600, "You must provide kibana connector ID")
		}
		log.Debug("ID: ", id)

		path := fmt.Sprintf("%s/%s", basePathKibanaConnector, id)
		resp, err := c.R().Get(path)
		if err != nil {
			return nil, err
		}
		log.Debug("Response: ", resp)
		if resp.StatusCode() >= 300 {
			if resp.StatusCode() == 404 {
				return nil, nil
			}
			return nil, NewAPIError(resp.StatusCode(), resp.String())
		}
		KibanaConnector := &KibanaConnector{}
		err = json.Unmarshal(resp.Body(), KibanaConnector)
		if err != nil {
			return nil, err
		}
		log.Debug("KibanaConnector: ", KibanaConnector)

		return KibanaConnector, nil
	}
}

// newKibanaConnectorListFunc permit to get all Kibana connector
func newKibanaConnectorListFunc(c *resty.Client) KibanaConnectorList {
	return func() (KibanaConnectors, error) {

		path := fmt.Sprintf("%s%s", basePathKibanaConnector, "s") // plural
		resp, err := c.R().Get(path)
		if err != nil {
			return nil, err
		}
		log.Debug("Response: ", resp)
		if resp.StatusCode() >= 300 {
			return nil, NewAPIError(resp.StatusCode(), resp.Status())
		}
		KibanaConnectors := make(KibanaConnectors, 0, 1)
		err = json.Unmarshal(resp.Body(), &KibanaConnectors)
		if err != nil {
			return nil, err
		}
		log.Debug("KibanaConnectors: ", KibanaConnectors)

		return KibanaConnectors, nil
	}
}

// newKibanaConnectorCreateFunc permit to create new Kibana connector
func newKibanaConnectorCreateFunc(c *resty.Client) KibanaConnectorCreate {
	return func(kibanaConnectorCreateParams *KibanaConnectorCreateParams) (*KibanaConnector, error) {
		if kibanaConnectorCreateParams == nil {
			return nil, NewAPIError(600, "You must provide kibana connector object")
		}
		log.Debug("KibanaConnector: ", kibanaConnectorCreateParams)

		jsonData, err := json.Marshal(kibanaConnectorCreateParams)
		if err != nil {
			return nil, err
		}
		resp, err := c.R().SetBody(jsonData).Post(basePathKibanaConnector)
		if err != nil {
			return nil, err
		}

		log.Debug("Response: ", resp)
		if resp.StatusCode() >= 300 {
			return nil, NewAPIError(resp.StatusCode(), resp.String())
		}
		kibanaConnector := &KibanaConnector{}
		err = json.Unmarshal(resp.Body(), kibanaConnector)
		if err != nil {
			return nil, err
		}
		log.Debug("KibanaConnector: ", kibanaConnector)

		return kibanaConnector, nil
	}
}

// newKibanaConnectorDeleteFunc permit to delete the kubana connector with id
func newKibanaConnectorDeleteFunc(c *resty.Client) KibanaConnectorDelete {
	return func(id string) error {
		if id == "" {
			return NewAPIError(600, "You must provide kibana connector ID")
		}

		log.Debug("ID: ", id)

		path := fmt.Sprintf("%s/%s", basePathKibanaConnector, id)
		resp, err := c.R().Delete(path)
		if err != nil {
			return err
		}
		log.Debug("Response: ", resp)
		if resp.StatusCode() >= 300 {
			return NewAPIError(resp.StatusCode(), resp.String())
		}
		return nil
	}
}

// newKibanaConnectorUpdateFunc permit to update the Kibana connector
func newKibanaConnectorUpdateFunc(c *resty.Client) KibanaConnectorUpdate {
	return func(id string, kibanaConnectorCreateParams *KibanaConnectorCreateParams) (*KibanaConnector, error) {
		if kibanaConnectorCreateParams == nil {
			return nil, NewAPIError(600, "You must provide kibana connector object")
		}
		log.Debug("kibanaConnectorCreateParams: ", kibanaConnectorCreateParams)

		jsonData, err := json.Marshal(kibanaConnectorCreateParams)
		if err != nil {
			return nil, err
		}
		path := fmt.Sprintf("%s/%s", basePathKibanaConnector, id)
		resp, err := c.R().SetBody(jsonData).Put(path)
		if err != nil {
			return nil, err
		}

		log.Debug("Response: ", resp)
		if resp.StatusCode() >= 300 {
			return nil, NewAPIError(resp.StatusCode(), resp.String())
		}
		kibanaConnector := &KibanaConnector{}
		err = json.Unmarshal(resp.Body(), kibanaConnector)
		if err != nil {
			return nil, err
		}
		log.Debug("kibanaConnector: ", kibanaConnector)

		return kibanaConnector, nil
	}
}

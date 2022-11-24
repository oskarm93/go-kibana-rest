package kbapi

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

const (
	basePathKibanaAlertRule = "/api/alerting/rule"
)

type KibanaAlertRuleSchedule struct {
	Interval string `json:"interval"`
}

type KibanaAlertRuleExecutionStatus map[string]interface{}

type KibanaAlertRuleAction struct {
	Id     string          `json:"id"`
	Group  string          `json:"group"`
	Params json.RawMessage `json:"params"`
}

type KibanaAlertRule struct {
	ID              string                         `json:"id"`
	Name            string                         `json:"name"`
	Consumer        string                         `json:"consumer"`
	Tags            []string                       `json:"tags"`
	Throttle        string                         `json:"throttle,omitempty"`
	Enabled         bool                           `json:"enabled"`
	Schedule        KibanaAlertRuleSchedule        `json:"schedule"`
	Params          json.RawMessage                `json:"params"`
	RuleTypeID      string                         `json:"rule_type_id"`
	CreatedBy       string                         `json:"created_by"`
	UpdatedBy       string                         `json:"updated_by"`
	CreatedAt       string                         `json:"created_at"`
	UpdatedAt       string                         `json:"updated_at"`
	ApiKeyOwner     string                         `json:"api_key_owner"`
	NotifyWhen      string                         `json:"notify_when"`
	MuteAlertIDs    []string                       `json:"muted_alert_ids"`
	MuteAll         bool                           `json:"mute_all"`
	ScheduledTaskID string                         `json:"scheduled_task_id"`
	ExecutionStatus KibanaAlertRuleExecutionStatus `json:"execution_status"`
	Actions         []KibanaAlertRuleAction        `json:"actions"`
}

type KibanaAlertRuleCreateParams struct {
	Name       string                  `json:"name"`
	Consumer   string                  `json:"consumer"`
	Tags       []string                `json:"tags,omitempty"`
	Throttle   string                  `json:"throttle,omitempty"`
	Enabled    bool                    `json:"enabled,omitempty"`
	Schedule   KibanaAlertRuleSchedule `json:"schedule"`
	Params     json.RawMessage         `json:"params"`
	RuleTypeID string                  `json:"rule_type_id"`
	NotifyWhen string                  `json:"notify_when"`
	Actions    []KibanaAlertRuleAction `json:"actions,omitempty""`
}

type KibanaAlertRuleUpdateParams struct {
	Name       string                  `json:"name"`
	Tags       []string                `json:"tags,omitempty"`
	Throttle   string                  `json:"throttle,omitempty"`
	Schedule   KibanaAlertRuleSchedule `json:"schedule"`
	Params     json.RawMessage         `json:"params"`
	NotifyWhen string                  `json:"notify_when"`
	Actions    []KibanaAlertRuleAction `json:"actions,omitempty"`
}

// KibanaAlertRuleGet permit to get alert rule
type KibanaAlertRuleGet func(id string) (*KibanaAlertRule, error)

// KibanaAlertRuleCreate permit to create alert rule
type KibanaAlertRuleCreate func(kibanaAlertRuleCreateParams *KibanaAlertRuleCreateParams) (*KibanaAlertRule, error)

// KibanaAlertRuleDelete permit to delete alert rule
type KibanaAlertRuleDelete func(id string) error

// KibanaAlertRuleUpdate permit to update alert rule
type KibanaAlertRuleUpdate func(id string, kibanaAlertRuleCreateParams *KibanaAlertRuleUpdateParams) (*KibanaAlertRule, error)

// KibanaAlertRuleUpdate permit to enable alert rule
type KibanaAlertRuleEnable func(id string) error

// KibanaAlertRuleUpdate permit to disable alert rule
type KibanaAlertRuleDisable func(id string) error

// String permit to return KibanaAlertRule object as JSON string
func (k *KibanaAlertRule) String() string {
	json, _ := json.Marshal(k)
	return string(json)
}

// newKibanaAlertRuleGetFunc permit to get the kibana alert rule with it id
func newKibanaAlertRuleGetFunc(c *resty.Client) KibanaAlertRuleGet {
	return func(id string) (*KibanaAlertRule, error) {
		if id == "" {
			return nil, NewAPIError(600, "You must provide kibana alert rule ID")
		}
		log.Debug("ID: ", id)

		path := fmt.Sprintf("%s/%s", basePathKibanaAlertRule, id)
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
		KibanaAlertRule := &KibanaAlertRule{}
		err = json.Unmarshal(resp.Body(), KibanaAlertRule)
		if err != nil {
			return nil, err
		}
		log.Debug("KibanaAlertRule: ", KibanaAlertRule)

		return KibanaAlertRule, nil
	}
}

// newKibanaAlertRuleCreateFunc permit to create new kibana alert rule
func newKibanaAlertRuleCreateFunc(c *resty.Client) KibanaAlertRuleCreate {
	return func(kibanaAlertRuleCreateParams *KibanaAlertRuleCreateParams) (*KibanaAlertRule, error) {
		if kibanaAlertRuleCreateParams == nil {
			return nil, NewAPIError(600, "You must provide kibana alert rule object")
		}
		log.Debug("KibanaAlertRule: ", kibanaAlertRuleCreateParams)

		jsonData, err := json.Marshal(kibanaAlertRuleCreateParams)
		if err != nil {
			return nil, err
		}
		resp, err := c.R().SetBody(jsonData).Post(basePathKibanaAlertRule)
		if err != nil {
			return nil, err
		}

		log.Debug("Response: ", resp)
		if resp.StatusCode() >= 300 {
			return nil, NewAPIError(resp.StatusCode(), resp.String())
		}
		kibanaAlertRule := &KibanaAlertRule{}
		err = json.Unmarshal(resp.Body(), kibanaAlertRule)
		if err != nil {
			return nil, err
		}
		log.Debug("KibanaAlertRule: ", kibanaAlertRule)

		return kibanaAlertRule, nil
	}
}

// newKibanaAlertRuleDeleteFunc permit to delete the kubana alert rule with id
func newKibanaAlertRuleDeleteFunc(c *resty.Client) KibanaAlertRuleDelete {
	return func(id string) error {
		if id == "" {
			return NewAPIError(600, "You must provide kibana alert rule ID")
		}

		log.Debug("ID: ", id)

		path := fmt.Sprintf("%s/%s", basePathKibanaAlertRule, id)
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

// newKibanaAlertRuleUpdateFunc permit to update the kibana alert rule
func newKibanaAlertRuleUpdateFunc(c *resty.Client) KibanaAlertRuleUpdate {
	return func(id string, kibanaAlertRuleUpdateParams *KibanaAlertRuleUpdateParams) (*KibanaAlertRule, error) {
		if kibanaAlertRuleUpdateParams == nil {
			return nil, NewAPIError(600, "You must provide kibana alert rule object")
		}
		log.Debug("kibanaAlertRuleUpdateParams: ", kibanaAlertRuleUpdateParams)

		jsonData, err := json.Marshal(kibanaAlertRuleUpdateParams)
		if err != nil {
			return nil, err
		}
		path := fmt.Sprintf("%s/%s", basePathKibanaAlertRule, id)
		resp, err := c.R().SetBody(jsonData).Put(path)
		if err != nil {
			return nil, err
		}

		log.Debug("Response: ", resp)
		if resp.StatusCode() >= 300 {
			return nil, NewAPIError(resp.StatusCode(), resp.String())
		}
		kibanaAlertRule := &KibanaAlertRule{}
		err = json.Unmarshal(resp.Body(), kibanaAlertRule)
		if err != nil {
			return nil, err
		}
		log.Debug("kibanaAlertRule: ", kibanaAlertRule)

		return kibanaAlertRule, nil
	}
}

// newKibanaAlertRuleEnableFunc permit to update the kibana alert rule
func newKibanaAlertRuleEnableFunc(c *resty.Client) KibanaAlertRuleEnable {
	return func(id string) error {
		path := fmt.Sprintf("%s/%s/enable", basePathKibanaAlertRule, id)
		resp, err := c.R().Post(path)
		if err != nil {
			return err
		}

		log.Debug("Response: ", resp)
		if resp.StatusCode() != 200 {
			return NewAPIError(resp.StatusCode(), resp.String())
		}

		return nil
	}
}

// newKibanaAlertRuleDisableFunc permit to update the kibana alert rule
func newKibanaAlertRuleDisableFunc(c *resty.Client) KibanaAlertRuleDisable {
	return func(id string) error {
		path := fmt.Sprintf("%s/%s/disable", basePathKibanaAlertRule, id)
		resp, err := c.R().Post(path)
		if err != nil {
			return err
		}

		log.Debug("Response: ", resp)
		if resp.StatusCode() != 200 {
			return NewAPIError(resp.StatusCode(), resp.String())
		}

		return nil
	}
}

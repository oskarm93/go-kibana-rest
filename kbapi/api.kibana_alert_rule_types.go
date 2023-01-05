package kbapi

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

const (
	basePathKibanaAlertRuleTypes = "/api/alerting/rule_types"
)

type KibanaAlertRuleType struct {
	ID                     string        `json:"id"`
	Name                   string        `json:"name"`
	Producer               string        `json:"producer"`
	DefaultActionGroupID   string        `json:"default_action_group_id"`
	MinimumLicenseRequired string        `json:"minimum_license_required"`
	RuleTaskTimeout        string        `json:"rule_task_timeout"`
	EnabledInLicense       bool          `json:"enabled_in_license"`
	IsExportable           bool          `json:"is_exportable"`
	DoesSetRecoveryContext bool          `json:"does_set_recovery_context"`
	ActionVariables        interface{}   `json:"action_variables"`
	ActionGroups           []interface{} `json:"action_groups"`
	AuthorizedConsumers    interface{}   `json:"authorized_consumers"`
	RecoveryActionGroup    interface{}   `json:"recovery_action_group"`
}

// String permit to return KibanaAlertRuleTypes object as JSON string
func (k KibanaAlertRuleTypes) String() string {
	json, _ := json.Marshal(k)
	return string(json)
}

type KibanaAlertRuleTypes []KibanaAlertRuleType

type KibanaAlertRuleTypesList func() (KibanaAlertRuleTypes, error)

func newKibanaAlertRuleTypesListFunc(c *resty.Client) KibanaAlertRuleTypesList {
	return func() (KibanaAlertRuleTypes, error) {
		resp, err := c.R().Get(basePathKibanaAlertRuleTypes)
		if err != nil {
			return nil, err
		}
		log.Debug("Response: ", resp)
		if resp.StatusCode() >= 300 {
			return nil, NewAPIError(resp.StatusCode(), resp.String())
		}
		kibanaAlertRuleTypes := make(KibanaAlertRuleTypes, 0, 1)
		err = json.Unmarshal(resp.Body(), &kibanaAlertRuleTypes)
		if err != nil {
			return nil, err
		}
		log.Debug("KibanaAlertRuleTypes: ", kibanaAlertRuleTypes)

		return kibanaAlertRuleTypes, nil
	}
}

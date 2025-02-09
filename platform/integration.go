package main

import (
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-jira/global"
	"github.com/opengovern/og-describer-jira/global/maps"
	"github.com/opengovern/og-describer-jira/platform/constants"
	"github.com/opengovern/og-util/pkg/integration"
	"github.com/opengovern/og-util/pkg/integration/interfaces"
)

type Integration struct{}

func (i *Integration) GetConfiguration() (interfaces.IntegrationConfiguration, error) {
	return interfaces.IntegrationConfiguration{
		NatsScheduledJobsTopic:   global.JobQueueTopic,
		NatsManualJobsTopic:      global.JobQueueTopicManuals,
		NatsStreamName:           global.StreamName,
		NatsConsumerGroup:        global.ConsumerGroup,
		NatsConsumerGroupManuals: global.ConsumerGroupManuals,

		SteampipePluginName: "jira",

		UISpec:   constants.UISpec,
		Manifest: constants.Manifest,
		SetupMD:  constants.SetupMd,

		DescriberDeploymentName: constants.DescriberDeploymentName,
		DescriberRunCommand:     constants.DescriberRunCommand,
	}, nil
}

func (i *Integration) HealthCheck(jsonData []byte, providerId string, labels map[string]string, annotations map[string]string) (bool, error) {
	var credentials global.IntegrationCredentials
	err := json.Unmarshal(jsonData, &credentials)
	if err != nil {
		return false, err
	}
	var isLocal bool
	if credentials.Password == "" {
		if credentials.APIKey == "" {
			return false, fmt.Errorf("password or api key must be configured")
		} else {
			isLocal = false
		}
	} else {
		if credentials.APIKey != "" {
			return false, fmt.Errorf("only one of password and api key must be configured")
		} else {
			isLocal = true
		}
	}

	isHealthy, err := JiraIntegrationHealthcheck(credentials.BaseURL, credentials.Username, credentials.APIKey, isLocal)
	return isHealthy, err
}

func (i *Integration) DiscoverIntegrations(jsonData []byte) ([]integration.Integration, error) {
	var credentials global.IntegrationCredentials
	err := json.Unmarshal(jsonData, &credentials)
	if err != nil {
		return nil, err
	}
	var integrations []integration.Integration
	var isLocal bool
	var pass string
	if credentials.Password == "" {
		if credentials.APIKey == "" {
			return nil, fmt.Errorf("password or api key must be configured")
		} else {
			pass = credentials.APIKey
			isLocal = false
		}
	} else {
		if credentials.APIKey != "" {
			return nil, fmt.Errorf("only one of password and api key must be configured")
		} else {
			pass = credentials.Password
			isLocal = true
		}
	}

	jiraInstance, err := JiraIntegrationDiscovery(credentials.BaseURL, credentials.Username, pass, isLocal)
	if err != nil {
		return nil, err
	}
	integrations = append(integrations, integration.Integration{
		ProviderID: jiraInstance.BaseURL,
		Name:       jiraInstance.DisplayURL,
	})

	return integrations, nil
}

func (i *Integration) GetResourceTypesByLabels(labels map[string]string) ([]interfaces.ResourceTypeConfiguration, error) {
	var resourceTypesMap []interfaces.ResourceTypeConfiguration
	for _, resourceType := range maps.ResourceTypesList {
		var resource interfaces.ResourceTypeConfiguration
		if v, ok := maps.ResourceTypeConfigs[resourceType]; ok {
			resource.Description = v.Description
			resource.Params = v.Params
			resource.Name = v.Name
			resource.IntegrationType = v.IntegrationType
			resource.Table = maps.ResourceTypesToTables[v.Name]
			resourceTypesMap = append(resourceTypesMap, resource)

		}
	}
	return resourceTypesMap, nil
}
func (i *Integration) GetResourceTypeFromTableName(tableName string) (string, error) {
	if v, ok := maps.TablesToResourceTypes[tableName]; ok {
		return v, nil
	}

	return "", nil
}

func (i *Integration) GetIntegrationType() (integration.Type, error) {
	return constants.IntegrationName, nil
}

func (i *Integration) ListAllTables() (map[string][]interfaces.CloudQLColumn, error) {
	plugin := global.Plugin()
	tables := make(map[string][]interfaces.CloudQLColumn)
	for tableKey, table := range plugin.TableMap {
		columns := make([]interfaces.CloudQLColumn, 0, len(table.Columns))
		for _, column := range table.Columns {
			columns = append(columns, interfaces.CloudQLColumn{Name: column.Name, Type: column.Type.String()})
		}
		tables[tableKey] = columns
	}

	return tables, nil
}

func (i *Integration) Ping() error {
	return nil
}

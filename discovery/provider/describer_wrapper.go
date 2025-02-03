package provider

import (
	"errors"
	"github.com/andygrunwald/go-jira"
	model "github.com/opengovern/og-describer-jira/discovery/pkg/models"
	"github.com/opengovern/og-util/pkg/describe/enums"
	"golang.org/x/net/context"
)

// DescribeListByJira A wrapper to pass Jira authorization to describers functions
func DescribeListByJira(describe func(context.Context, *jira.Client, *model.StreamSender) ([]model.Resource, error)) model.ResourceDescriber {
	return func(ctx context.Context, cfg model.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
		ctx = WithTriggerType(ctx, triggerType)

		var err error
		// Check for the api_key
		if cfg.APIKey == "" {
			return nil, errors.New("api key must be configured")
		}
		if cfg.Username == "" {
			return nil, errors.New("username must be configured")
		}
		if cfg.BaseURL == "" {
			return nil, errors.New("base url must be configured")
		}

		tp := jira.BasicAuthTransport{
			Username: cfg.Username,
			Password: cfg.APIKey,
		}

		client, err := jira.NewClient(tp.Client(), cfg.BaseURL)
		if err != nil {
			return nil, err
		}

		// Get values from describers
		var values []model.Resource
		result, err := describe(ctx, client, stream)
		if err != nil {
			return nil, err
		}
		values = append(values, result...)
		return values, nil
	}
}

// DescribeSingleByJira A wrapper to pass Jira authorization to describers functions
func DescribeSingleByJira(describe func(context.Context, *jira.Client, string) (*model.Resource, error)) model.SingleResourceDescriber {
	return func(ctx context.Context, cfg model.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, resourceID string, stream *model.StreamSender) (*model.Resource, error) {
		ctx = WithTriggerType(ctx, triggerType)

		var err error
		// Check for the api_key
		if cfg.APIKey == "" {
			return nil, errors.New("api key must be configured")
		}
		if cfg.Username == "" {
			return nil, errors.New("username must be configured")
		}
		if cfg.BaseURL == "" {
			return nil, errors.New("base url must be configured")
		}

		tp := jira.BasicAuthTransport{
			Username: cfg.Username,
			Password: cfg.APIKey,
		}

		client, err := jira.NewClient(tp.Client(), cfg.BaseURL)
		if err != nil {
			return nil, err
		}

		// Get value from describers
		value, err := describe(ctx, client, resourceID)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
}

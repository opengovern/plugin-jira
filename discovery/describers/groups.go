package describers

import (
	"context"
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/opengovern/og-describer-jira/discovery/pkg/models"
	"github.com/opengovern/og-describer-jira/discovery/provider"
	"net/url"
	"strconv"
)

func ListGroups(ctx context.Context, client *jira.Client, stream *models.StreamSender, isLocal bool) ([]models.Resource, error) {
	var groups []provider.GroupJSON
	var groupListResp provider.GroupListResponse

	var baseURL string
	if isLocal {
		baseURL = "rest/api/2/group/bulk"
	} else {
		baseURL = "rest/api/3/group/bulk"
	}

	last := 0

	for {
		params := url.Values{}
		params.Set("startAt", strconv.Itoa(last))
		params.Set("maxResults", "1000")
		finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

		req, err := client.NewRequest("GET", finalURL, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}
		req.Header.Set("Accept", "application/json")

		_, err = client.Do(req, &groupListResp)
		if err != nil {
			return nil, fmt.Errorf("request execution failed: %w", err)
		}

		groups = append(groups, groupListResp.Values...)

		last = groupListResp.StartAt + len(groupListResp.Values)
		if groupListResp.IsLast {
			break
		}
	}

	var values []models.Resource
	for _, group := range groups {
		value := models.Resource{
			ID:   group.GroupID,
			Name: group.Name,
			Description: provider.GroupDescription{
				GroupID: group.GroupID,
				Name:    group.Name,
			},
		}
		if stream != nil {
			if err := (*stream)(value); err != nil {
				return nil, err
			}
		} else {
			values = append(values, value)
		}
	}

	return values, nil
}

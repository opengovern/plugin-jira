package describers

import (
	"context"
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/opengovern/og-describer-jira/discovery/pkg/models"
	"github.com/opengovern/og-describer-jira/discovery/provider"
)

func ListUsers(ctx context.Context, client *jira.Client, stream *models.StreamSender, isLocal bool) ([]models.Resource, error) {
	var users []provider.UserJSON

	var baseURL string
	if isLocal {
		baseURL = "rest/api/2/users/search"
	} else {
		baseURL = "rest/api/3/users/search"
	}

	req, err := client.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	_, err = client.Do(req, &users)
	if err != nil {
		return nil, fmt.Errorf("request execution failed: %w", err)
	}

	var values []models.Resource
	for _, user := range users {
		avatarUrls := provider.AvatarUrls{
			Small16x16:  user.AvatarUrls.Small16x16,
			Small24x24:  user.AvatarUrls.Small24x24,
			Medium32x32: user.AvatarUrls.Medium32x32,
			Large48x48:  user.AvatarUrls.Large48x48,
		}
		value := models.Resource{
			ID:   user.AccountID,
			Name: user.Name,
			Description: provider.UserDescription{
				AccountID:   user.AccountID,
				AccountType: user.AccountType,
				Active:      user.Active,
				AvatarURLs:  avatarUrls,
				DisplayName: user.DisplayName,
				Key:         user.Key,
				Name:        user.Name,
				Self:        user.Self,
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

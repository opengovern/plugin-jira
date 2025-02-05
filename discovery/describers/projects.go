package describers

import (
	"context"
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/opengovern/og-describer-jira/discovery/pkg/models"
	"github.com/opengovern/og-describer-jira/discovery/provider"
)

func ListProjects(ctx context.Context, client *jira.Client, stream *models.StreamSender) ([]models.Resource, error) {
	var project provider.ProjectJSON

	var projectKey string
	projectKeyParam := ctx.Value("project_key")
	if projectKeyParam != nil {
		value, ok := projectKeyParam.(string)
		if ok && value != "" {
			projectKey = value
		} else {
			return nil, fmt.Errorf("project key parameter must be configured")
		}
	} else {
		return nil, fmt.Errorf("project key parameter must be configured")
	}

	baseURL := "rest/api/3/project/"
	finalURL := fmt.Sprintf("%s%s", baseURL, projectKey)

	req, err := client.NewRequest("GET", finalURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	_, err = client.Do(req, &project)
	if err != nil {
		return nil, fmt.Errorf("request execution failed: %w", err)
	}

	avatarUrls := provider.AvatarUrls{
		Small16x16:  project.AvatarUrls.Small16x16,
		Small24x24:  project.AvatarUrls.Small24x24,
		Medium32x32: project.AvatarUrls.Medium32x32,
		Large48x48:  project.AvatarUrls.Large48x48,
	}
	insight := provider.Insight{
		LastIssueUpdateTime: project.Insight.LastIssueUpdateTime,
		TotalIssueCount:     project.Insight.TotalIssueCount,
	}
	projectCategory := provider.ProjectCategory{
		Description: project.ProjectCategory.Description,
		ID:          project.ProjectCategory.ID,
		Name:        project.ProjectCategory.Name,
		Self:        project.ProjectCategory.Self,
	}

	var values []models.Resource

	value := models.Resource{
		ID:   project.ID,
		Name: project.Name,
		Description: provider.ProjectDescription{
			AvatarUrls:      avatarUrls,
			ID:              project.Key,
			Insight:         insight,
			Key:             project.Key,
			Name:            project.Name,
			ProjectCategory: projectCategory,
			Self:            project.Self,
			Simplified:      project.Simplified,
			Style:           project.Style,
		},
	}
	if stream != nil {
		if err := (*stream)(value); err != nil {
			return nil, err
		}
	} else {
		values = append(values, value)
	}

	return values, nil
}

//func ListProjects(ctx context.Context, client *jira.Client, stream *models.StreamSender) ([]models.Resource, error) {
//	var projects []provider.ProjectJSON
//	var projectListResp provider.ProjectListResponse
//	baseURL := "rest/api/3/project/search"
//	last := 0
//
//
//	for {
//		params := url.Values{}
//		params.Set("startAt", strconv.Itoa(last))
//		params.Set("maxResults", "1000")
//		finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
//
//		req, err := client.NewRequest("GET", finalURL, nil)
//		if err != nil {
//			return nil, fmt.Errorf("failed to create request: %w", err)
//		}
//		req.Header.Set("Accept", "application/json")
//
//		_, err = client.Do(req, &projectListResp)
//		if err != nil {
//			return nil, fmt.Errorf("request execution failed: %w", err)
//		}
//
//		projects = append(projects, projectListResp.Values...)
//
//		last = projectListResp.StartAt + len(projectListResp.Values)
//		if projectListResp.IsLast {
//			break
//		}
//	}
//
//	var values []models.Resource
//	for _, project := range projects {
//		avatarUrls := provider.AvatarUrls{
//			Small16x16:  project.AvatarUrls.Small16x16,
//			Small24x24:  project.AvatarUrls.Small24x24,
//			Medium32x32: project.AvatarUrls.Medium32x32,
//			Large48x48:  project.AvatarUrls.Large48x48,
//		}
//		insight := provider.Insight{
//			LastIssueUpdateTime: project.Insight.LastIssueUpdateTime,
//			TotalIssueCount:     project.Insight.TotalIssueCount,
//		}
//		projectCategory := provider.ProjectCategory{
//			Description: project.ProjectCategory.Description,
//			ID:          project.ProjectCategory.ID,
//			Name:        project.ProjectCategory.Name,
//			Self:        project.ProjectCategory.Self,
//		}
//		value := models.Resource{
//			ID:   project.ID,
//			Name: project.Name,
//			Description: provider.ProjectDescription{
//				AvatarUrls:      avatarUrls,
//				ID:              project.Key,
//				Insight:         insight,
//				Key:             project.Key,
//				Name:            project.Name,
//				ProjectCategory: projectCategory,
//				Self:            project.Self,
//				Simplified:      project.Simplified,
//				Style:           project.Style,
//			},
//		}
//		if stream != nil {
//			if err := (*stream)(value); err != nil {
//				return nil, err
//			}
//		} else {
//			values = append(values, value)
//		}
//	}
//
//	return values, nil
//}

func GetProject(ctx context.Context, client *jira.Client, resourceID string) (*models.Resource, error) {
	var project provider.ProjectJSON
	baseURL := "rest/api/3/project/"
	finalURL := fmt.Sprintf("%s%s", baseURL, resourceID)

	req, err := client.NewRequest("GET", finalURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	_, err = client.Do(req, &project)
	if err != nil {
		return nil, fmt.Errorf("request execution failed: %w", err)
	}

	avatarUrls := provider.AvatarUrls{
		Small16x16:  project.AvatarUrls.Small16x16,
		Small24x24:  project.AvatarUrls.Small24x24,
		Medium32x32: project.AvatarUrls.Medium32x32,
		Large48x48:  project.AvatarUrls.Large48x48,
	}
	insight := provider.Insight{
		LastIssueUpdateTime: project.Insight.LastIssueUpdateTime,
		TotalIssueCount:     project.Insight.TotalIssueCount,
	}
	projectCategory := provider.ProjectCategory{
		Description: project.ProjectCategory.Description,
		ID:          project.ProjectCategory.ID,
		Name:        project.ProjectCategory.Name,
		Self:        project.ProjectCategory.Self,
	}

	value := models.Resource{
		ID:   project.ID,
		Name: project.Name,
		Description: provider.ProjectDescription{
			AvatarUrls:      avatarUrls,
			ID:              project.Key,
			Insight:         insight,
			Key:             project.Key,
			Name:            project.Name,
			ProjectCategory: projectCategory,
			Self:            project.Self,
			Simplified:      project.Simplified,
			Style:           project.Style,
		},
	}

	return &value, nil
}

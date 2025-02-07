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

func ListIssues(ctx context.Context, client *jira.Client, stream *models.StreamSender) ([]models.Resource, error) {
	var issues []provider.IssueJSON
	var issueListResp provider.IssueListResponse

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

	var status string
	statusParam := ctx.Value("status")
	if statusParam != nil {
		value, ok := statusParam.(string)
		if ok && value != "" {
			status = value
		} else {
			return nil, fmt.Errorf("status parameter must be configured")
		}
	} else {
		return nil, fmt.Errorf("status parameter must be configured")
	}

	var statusCategory string
	statusCategoryParam := ctx.Value("status_category")
	if statusCategoryParam != nil {
		value, ok := statusCategoryParam.(string)
		if ok && value != "" {
			statusCategory = value
		} else {
			return nil, fmt.Errorf("status category parameter must be configured")
		}
	} else {
		return nil, fmt.Errorf("status category parameter must be configured")
	}

	var fields string
	fieldsParam := ctx.Value("fields")
	if fieldsParam != nil {
		value, ok := fieldsParam.(string)
		if ok && value != "" {
			fields = value
		}
	}

	baseURL := "rest/api/3/search/jql"
	last := 0
	jql := fmt.Sprintf(`project = "%s" AND status = "%s" AND statusCategory = "%s"`, projectKey, status, statusCategory)

	for {
		params := url.Values{}
		params.Set("startAt", strconv.Itoa(last))
		params.Set("maxResults", "1000")
		params.Set("jql", jql)
		if fields != "" {
			params.Set("fields", fields)
		}
		finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

		req, err := client.NewRequest("GET", finalURL, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}
		req.Header.Set("Accept", "application/json")

		_, err = client.Do(req, &issueListResp)
		if err != nil {
			return nil, fmt.Errorf("request execution failed: %w", err)
		}

		issues = append(issues, issueListResp.Issues...)

		last = issueListResp.StartAt + len(issueListResp.Issues)
		if last >= issueListResp.Total {
			break
		}
	}

	var values []models.Resource
	for _, issue := range issues {
		watcher := provider.Watcher{
			IsWatching: issue.Fields.Watcher.IsWatching,
			Self:       issue.Fields.Watcher.Self,
			WatchCount: issue.Fields.Watcher.WatchCount,
		}
		var attachments []provider.Attachment
		for _, attachment := range issue.Fields.Attachment {
			avatarUrls := provider.AvatarUrls{
				Small16x16:  attachment.Author.AvatarUrls.Small16x16,
				Small24x24:  attachment.Author.AvatarUrls.Small24x24,
				Medium32x32: attachment.Author.AvatarUrls.Medium32x32,
				Large48x48:  attachment.Author.AvatarUrls.Large48x48,
			}
			user := provider.User{
				AccountID:   attachment.Author.AccountID,
				AccountType: attachment.Author.AccountType,
				Active:      attachment.Author.Active,
				AvatarUrls:  avatarUrls,
				DisplayName: attachment.Author.DisplayName,
				Key:         attachment.Author.Key,
				Name:        attachment.Author.Name,
				Self:        attachment.Author.Self,
			}
			attachments = append(attachments, provider.Attachment{
				Author:    user,
				Content:   attachment.Content,
				Created:   attachment.Created,
				Filename:  attachment.Filename,
				ID:        attachment.ID,
				MimeType:  attachment.MimeType,
				Self:      attachment.Self,
				Size:      attachment.Size,
				Thumbnail: attachment.Thumbnail,
			})
		}
		var subTasks []provider.SubTask
		for _, subTask := range issue.Fields.SubTasks {
			status := provider.Status{
				IconURL: subTask.OutwardIssue.Fields.Status.IconURL,
				Name:    subTask.OutwardIssue.Fields.Status.Name,
			}
			statusField := provider.StatusField{
				Status: status,
			}
			outwardIssue := provider.OutwardIssue{
				Fields: statusField,
				ID:     subTask.OutwardIssue.ID,
				Key:    subTask.OutwardIssue.Key,
				Self:   subTask.OutwardIssue.Self,
			}
			issueType := provider.IssueType{
				ID:      subTask.Type.ID,
				Inward:  subTask.Type.Inward,
				Name:    subTask.Type.Name,
				Outward: subTask.Type.Outward,
			}
			subTasks = append(subTasks, provider.SubTask{
				ID:           subTask.ID,
				OutwardIssue: outwardIssue,
				Type:         issueType,
			})
		}
		var contentDetails []provider.ContentDetail
		for _, contentDetail := range issue.Fields.Description.Content {
			var textContents []provider.TextContent
			for _, textContent := range contentDetail.Content {
				textContents = append(textContents, provider.TextContent{
					Type: textContent.Type,
					Text: textContent.Text,
				})
			}
			contentDetails = append(contentDetails, provider.ContentDetail{
				Type:    contentDetail.Type,
				Content: textContents,
			})
		}
		description := provider.Content{
			Type:    issue.Fields.Description.Type,
			Version: issue.Fields.Description.Version,
			Content: contentDetails,
		}
		avatarUrls := provider.AvatarUrls{
			Small16x16:  issue.Fields.Project.AvatarUrls.Small16x16,
			Small24x24:  issue.Fields.Project.AvatarUrls.Small24x24,
			Medium32x32: issue.Fields.Project.AvatarUrls.Medium32x32,
			Large48x48:  issue.Fields.Project.AvatarUrls.Large48x48,
		}
		insight := provider.Insight{
			LastIssueUpdateTime: issue.Fields.Project.Insight.LastIssueUpdateTime,
			TotalIssueCount:     issue.Fields.Project.Insight.TotalIssueCount,
		}
		projectCategory := provider.ProjectCategory{
			Description: issue.Fields.Project.ProjectCategory.Description,
			ID:          issue.Fields.Project.ProjectCategory.ID,
			Name:        issue.Fields.Project.ProjectCategory.Name,
			Self:        issue.Fields.Project.ProjectCategory.Self,
		}
		project := provider.ProjectDescription{
			AvatarUrls:      avatarUrls,
			ID:              issue.Fields.Project.Key,
			Insight:         insight,
			Key:             issue.Fields.Project.Key,
			Name:            issue.Fields.Project.Name,
			ProjectCategory: projectCategory,
			Self:            issue.Fields.Project.Self,
			Simplified:      issue.Fields.Project.Simplified,
			Style:           issue.Fields.Project.Style,
		}
		var comments []provider.Comment
		for _, comment := range issue.Fields.Comment {
			authorAvatarUrls := provider.AvatarUrls{
				Small16x16:  comment.Author.AvatarUrls.Small16x16,
				Small24x24:  comment.Author.AvatarUrls.Small24x24,
				Medium32x32: comment.Author.AvatarUrls.Medium32x32,
				Large48x48:  comment.Author.AvatarUrls.Large48x48,
			}
			author := provider.User{
				AccountID:   comment.Author.AccountID,
				AccountType: comment.Author.AccountType,
				Active:      comment.Author.Active,
				AvatarUrls:  authorAvatarUrls,
				DisplayName: comment.Author.DisplayName,
				Key:         comment.Author.Key,
				Name:        comment.Author.Name,
				Self:        comment.Author.Self,
			}
			var bodyContentDetails []provider.ContentDetail
			for _, contentDetail := range comment.Body.Content {
				var bodyTextContents []provider.TextContent
				for _, textContent := range contentDetail.Content {
					bodyTextContents = append(bodyTextContents, provider.TextContent{
						Type: textContent.Type,
						Text: textContent.Text,
					})
				}
				bodyContentDetails = append(bodyContentDetails, provider.ContentDetail{
					Type:    contentDetail.Type,
					Content: bodyTextContents,
				})
			}
			body := provider.Content{
				Type:    comment.Body.Type,
				Version: comment.Body.Version,
				Content: bodyContentDetails,
			}
			updateAuthorAvatarUrls := provider.AvatarUrls{
				Small16x16:  comment.Author.AvatarUrls.Small16x16,
				Small24x24:  comment.Author.AvatarUrls.Small24x24,
				Medium32x32: comment.Author.AvatarUrls.Medium32x32,
				Large48x48:  comment.Author.AvatarUrls.Large48x48,
			}
			updateAuthor := provider.User{
				AccountID:   comment.UpdateAuthor.AccountID,
				AccountType: comment.UpdateAuthor.AccountType,
				Active:      comment.UpdateAuthor.Active,
				AvatarUrls:  updateAuthorAvatarUrls,
				DisplayName: comment.UpdateAuthor.DisplayName,
				Key:         comment.UpdateAuthor.Key,
				Name:        comment.UpdateAuthor.Name,
				Self:        comment.UpdateAuthor.Self,
			}
			visibility := provider.Visibility{
				Identifier: comment.Visibility.Identifier,
				Type:       comment.Visibility.Type,
				Value:      comment.Visibility.Value,
			}
			comments = append(comments, provider.Comment{
				Author:       author,
				Body:         body,
				Created:      comment.Created,
				ID:           comment.ID,
				Self:         comment.Self,
				UpdateAuthor: updateAuthor,
				Updated:      comment.Updated,
				Visibility:   visibility,
			})
		}
		var issueLinks []provider.IssueLink
		for _, issueLink := range issue.Fields.IssueLinks {
			outwardStatus := provider.Status{
				IconURL: issueLink.OutwardIssue.Fields.Status.IconURL,
				Name:    issueLink.OutwardIssue.Fields.Status.IconURL,
			}
			outwardStatusField := provider.StatusField{
				Status: outwardStatus,
			}
			outwardIssue := provider.OutwardIssue{
				Fields: outwardStatusField,
				ID:     issueLink.OutwardIssue.ID,
				Key:    issueLink.OutwardIssue.Key,
				Self:   issueLink.OutwardIssue.Self,
			}
			inwardStatus := provider.Status{
				IconURL: issueLink.InwardIssue.Fields.Status.IconURL,
				Name:    issueLink.InwardIssue.Fields.Status.Name,
			}
			inwardStatusField := provider.StatusField{
				Status: inwardStatus,
			}
			inwardIssue := provider.OutwardIssue{
				Fields: inwardStatusField,
				ID:     issueLink.InwardIssue.ID,
				Key:    issueLink.InwardIssue.Key,
				Self:   issueLink.InwardIssue.Self,
			}
			issueType := provider.IssueType{
				ID:      issueLink.Type.ID,
				Inward:  issueLink.Type.Inward,
				Name:    issueLink.Type.Name,
				Outward: issueLink.Type.Outward,
			}
			issueLinks = append(issueLinks, provider.IssueLink{
				ID:           issueLink.ID,
				OutwardIssue: &outwardIssue,
				InwardIssue:  &inwardIssue,
				Type:         issueType,
			})
		}
		var worklogs []provider.WorklogEntry
		for _, worklog := range issue.Fields.Worklog {
			authorAvatarUrls := provider.AvatarUrls{
				Small16x16:  worklog.Author.AvatarUrls.Small16x16,
				Small24x24:  worklog.Author.AvatarUrls.Small24x24,
				Medium32x32: worklog.Author.AvatarUrls.Medium32x32,
				Large48x48:  worklog.Author.AvatarUrls.Large48x48,
			}
			author := provider.User{
				AccountID:   worklog.Author.AccountID,
				AccountType: worklog.Author.AccountType,
				Active:      worklog.Author.Active,
				AvatarUrls:  authorAvatarUrls,
				DisplayName: worklog.Author.DisplayName,
				Key:         worklog.Author.Key,
				Name:        worklog.Author.Name,
				Self:        worklog.Author.Self,
			}
			var commentContentDetails []provider.ContentDetail
			for _, contentDetail := range worklog.Comment.Content {
				var commentTextContents []provider.TextContent
				for _, textContent := range contentDetail.Content {
					commentTextContents = append(commentTextContents, provider.TextContent{
						Type: textContent.Type,
						Text: textContent.Text,
					})
				}
				commentContentDetails = append(commentContentDetails, provider.ContentDetail{
					Type:    contentDetail.Type,
					Content: commentTextContents,
				})
			}
			comment := provider.Content{
				Type:    worklog.Comment.Type,
				Version: worklog.Comment.Version,
				Content: commentContentDetails,
			}
			updateAuthorAvatarUrls := provider.AvatarUrls{
				Small16x16:  worklog.Author.AvatarUrls.Small16x16,
				Small24x24:  worklog.Author.AvatarUrls.Small24x24,
				Medium32x32: worklog.Author.AvatarUrls.Medium32x32,
				Large48x48:  worklog.Author.AvatarUrls.Large48x48,
			}
			updateAuthor := provider.User{
				AccountID:   worklog.UpdateAuthor.AccountID,
				AccountType: worklog.UpdateAuthor.AccountType,
				Active:      worklog.UpdateAuthor.Active,
				AvatarUrls:  updateAuthorAvatarUrls,
				DisplayName: worklog.UpdateAuthor.DisplayName,
				Key:         worklog.UpdateAuthor.Key,
				Name:        worklog.UpdateAuthor.Name,
				Self:        worklog.UpdateAuthor.Self,
			}
			visiblity := provider.Visibility{
				Identifier: worklog.Visibility.Identifier,
				Type:       worklog.Visibility.Type,
				Value:      worklog.Visibility.Value,
			}
			worklogs = append(worklogs, provider.WorklogEntry{
				Author:           author,
				Comment:          comment,
				ID:               worklog.ID,
				IssueID:          worklog.IssueID,
				Self:             worklog.Self,
				Started:          worklog.Started,
				TimeSpent:        worklog.TimeSpent,
				TimeSpentSeconds: worklog.TimeSpentSeconds,
				UpdateAuthor:     updateAuthor,
				Updated:          worklog.Updated,
				Visibility:       visiblity,
			})
		}
		timeTracking := provider.TimeTracking{
			OriginalEstimate:         issue.Fields.TimeTracking.OriginalEstimate,
			OriginalEstimateSeconds:  issue.Fields.TimeTracking.OriginalEstimateSeconds,
			RemainingEstimate:        issue.Fields.TimeTracking.RemainingEstimate,
			RemainingEstimateSeconds: issue.Fields.TimeTracking.RemainingEstimateSeconds,
			TimeSpent:                issue.Fields.TimeTracking.TimeSpent,
			TimeSpentSeconds:         issue.Fields.TimeTracking.TimeSpentSeconds,
		}
		fields := provider.Fields{
			Watcher:      watcher,
			Attachment:   attachments,
			SubTasks:     subTasks,
			Description:  description,
			Project:      project,
			Comment:      comments,
			IssueLinks:   issueLinks,
			Worklog:      worklogs,
			TimeTracking: timeTracking,
			Updated:      issue.Fields.Updated,
		}
		value := models.Resource{
			ID:   issue.ID,
			Name: issue.Key,
			Description: provider.IssueDescription{
				ID:     issue.ID,
				Key:    issue.Key,
				Self:   issue.Self,
				Fields: fields,
				Expand: issue.Expand,
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

func GetIssue(ctx context.Context, client *jira.Client, resourceID string) (*models.Resource, error) {
	var issue provider.IssueJSON
	baseURL := "rest/api/3/issue/"
	finalURL := fmt.Sprintf("%s%s", baseURL, resourceID)

	req, err := client.NewRequest("GET", finalURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	_, err = client.Do(req, &issue)
	if err != nil {
		return nil, fmt.Errorf("request execution failed: %w", err)
	}

	watcher := provider.Watcher{
		IsWatching: issue.Fields.Watcher.IsWatching,
		Self:       issue.Fields.Watcher.Self,
		WatchCount: issue.Fields.Watcher.WatchCount,
	}
	var attachments []provider.Attachment
	for _, attachment := range issue.Fields.Attachment {
		avatarUrls := provider.AvatarUrls{
			Small16x16:  attachment.Author.AvatarUrls.Small16x16,
			Small24x24:  attachment.Author.AvatarUrls.Small24x24,
			Medium32x32: attachment.Author.AvatarUrls.Medium32x32,
			Large48x48:  attachment.Author.AvatarUrls.Large48x48,
		}
		user := provider.User{
			AccountID:   attachment.Author.AccountID,
			AccountType: attachment.Author.AccountType,
			Active:      attachment.Author.Active,
			AvatarUrls:  avatarUrls,
			DisplayName: attachment.Author.DisplayName,
			Key:         attachment.Author.Key,
			Name:        attachment.Author.Name,
			Self:        attachment.Author.Self,
		}
		attachments = append(attachments, provider.Attachment{
			Author:    user,
			Content:   attachment.Content,
			Created:   attachment.Created,
			Filename:  attachment.Filename,
			ID:        attachment.ID,
			MimeType:  attachment.MimeType,
			Self:      attachment.Self,
			Size:      attachment.Size,
			Thumbnail: attachment.Thumbnail,
		})
	}
	var subTasks []provider.SubTask
	for _, subTask := range issue.Fields.SubTasks {
		status := provider.Status{
			IconURL: subTask.OutwardIssue.Fields.Status.IconURL,
			Name:    subTask.OutwardIssue.Fields.Status.Name,
		}
		statusField := provider.StatusField{
			Status: status,
		}
		outwardIssue := provider.OutwardIssue{
			Fields: statusField,
			ID:     subTask.OutwardIssue.ID,
			Key:    subTask.OutwardIssue.Key,
			Self:   subTask.OutwardIssue.Self,
		}
		issueType := provider.IssueType{
			ID:      subTask.Type.ID,
			Inward:  subTask.Type.Inward,
			Name:    subTask.Type.Name,
			Outward: subTask.Type.Outward,
		}
		subTasks = append(subTasks, provider.SubTask{
			ID:           subTask.ID,
			OutwardIssue: outwardIssue,
			Type:         issueType,
		})
	}
	var contentDetails []provider.ContentDetail
	for _, contentDetail := range issue.Fields.Description.Content {
		var textContents []provider.TextContent
		for _, textContent := range contentDetail.Content {
			textContents = append(textContents, provider.TextContent{
				Type: textContent.Type,
				Text: textContent.Text,
			})
		}
		contentDetails = append(contentDetails, provider.ContentDetail{
			Type:    contentDetail.Type,
			Content: textContents,
		})
	}
	description := provider.Content{
		Type:    issue.Fields.Description.Type,
		Version: issue.Fields.Description.Version,
		Content: contentDetails,
	}
	avatarUrls := provider.AvatarUrls{
		Small16x16:  issue.Fields.Project.AvatarUrls.Small16x16,
		Small24x24:  issue.Fields.Project.AvatarUrls.Small24x24,
		Medium32x32: issue.Fields.Project.AvatarUrls.Medium32x32,
		Large48x48:  issue.Fields.Project.AvatarUrls.Large48x48,
	}
	insight := provider.Insight{
		LastIssueUpdateTime: issue.Fields.Project.Insight.LastIssueUpdateTime,
		TotalIssueCount:     issue.Fields.Project.Insight.TotalIssueCount,
	}
	projectCategory := provider.ProjectCategory{
		Description: issue.Fields.Project.ProjectCategory.Description,
		ID:          issue.Fields.Project.ProjectCategory.ID,
		Name:        issue.Fields.Project.ProjectCategory.Name,
		Self:        issue.Fields.Project.ProjectCategory.Self,
	}
	project := provider.ProjectDescription{
		AvatarUrls:      avatarUrls,
		ID:              issue.Fields.Project.Key,
		Insight:         insight,
		Key:             issue.Fields.Project.Key,
		Name:            issue.Fields.Project.Name,
		ProjectCategory: projectCategory,
		Self:            issue.Fields.Project.Self,
		Simplified:      issue.Fields.Project.Simplified,
		Style:           issue.Fields.Project.Style,
	}
	var comments []provider.Comment
	for _, comment := range issue.Fields.Comment {
		authorAvatarUrls := provider.AvatarUrls{
			Small16x16:  comment.Author.AvatarUrls.Small16x16,
			Small24x24:  comment.Author.AvatarUrls.Small24x24,
			Medium32x32: comment.Author.AvatarUrls.Medium32x32,
			Large48x48:  comment.Author.AvatarUrls.Large48x48,
		}
		author := provider.User{
			AccountID:   comment.Author.AccountID,
			AccountType: comment.Author.AccountType,
			Active:      comment.Author.Active,
			AvatarUrls:  authorAvatarUrls,
			DisplayName: comment.Author.DisplayName,
			Key:         comment.Author.Key,
			Name:        comment.Author.Name,
			Self:        comment.Author.Self,
		}
		var bodyContentDetails []provider.ContentDetail
		for _, contentDetail := range comment.Body.Content {
			var bodyTextContents []provider.TextContent
			for _, textContent := range contentDetail.Content {
				bodyTextContents = append(bodyTextContents, provider.TextContent{
					Type: textContent.Type,
					Text: textContent.Text,
				})
			}
			bodyContentDetails = append(bodyContentDetails, provider.ContentDetail{
				Type:    contentDetail.Type,
				Content: bodyTextContents,
			})
		}
		body := provider.Content{
			Type:    comment.Body.Type,
			Version: comment.Body.Version,
			Content: bodyContentDetails,
		}
		updateAuthorAvatarUrls := provider.AvatarUrls{
			Small16x16:  comment.Author.AvatarUrls.Small16x16,
			Small24x24:  comment.Author.AvatarUrls.Small24x24,
			Medium32x32: comment.Author.AvatarUrls.Medium32x32,
			Large48x48:  comment.Author.AvatarUrls.Large48x48,
		}
		updateAuthor := provider.User{
			AccountID:   comment.UpdateAuthor.AccountID,
			AccountType: comment.UpdateAuthor.AccountType,
			Active:      comment.UpdateAuthor.Active,
			AvatarUrls:  updateAuthorAvatarUrls,
			DisplayName: comment.UpdateAuthor.DisplayName,
			Key:         comment.UpdateAuthor.Key,
			Name:        comment.UpdateAuthor.Name,
			Self:        comment.UpdateAuthor.Self,
		}
		visibility := provider.Visibility{
			Identifier: comment.Visibility.Identifier,
			Type:       comment.Visibility.Type,
			Value:      comment.Visibility.Value,
		}
		comments = append(comments, provider.Comment{
			Author:       author,
			Body:         body,
			Created:      comment.Created,
			ID:           comment.ID,
			Self:         comment.Self,
			UpdateAuthor: updateAuthor,
			Updated:      comment.Updated,
			Visibility:   visibility,
		})
	}
	var issueLinks []provider.IssueLink
	for _, issueLink := range issue.Fields.IssueLinks {
		outwardStatus := provider.Status{
			IconURL: issueLink.OutwardIssue.Fields.Status.IconURL,
			Name:    issueLink.OutwardIssue.Fields.Status.IconURL,
		}
		outwardStatusField := provider.StatusField{
			Status: outwardStatus,
		}
		outwardIssue := provider.OutwardIssue{
			Fields: outwardStatusField,
			ID:     issueLink.OutwardIssue.ID,
			Key:    issueLink.OutwardIssue.Key,
			Self:   issueLink.OutwardIssue.Self,
		}
		inwardStatus := provider.Status{
			IconURL: issueLink.InwardIssue.Fields.Status.IconURL,
			Name:    issueLink.InwardIssue.Fields.Status.Name,
		}
		inwardStatusField := provider.StatusField{
			Status: inwardStatus,
		}
		inwardIssue := provider.OutwardIssue{
			Fields: inwardStatusField,
			ID:     issueLink.InwardIssue.ID,
			Key:    issueLink.InwardIssue.Key,
			Self:   issueLink.InwardIssue.Self,
		}
		issueType := provider.IssueType{
			ID:      issueLink.Type.ID,
			Inward:  issueLink.Type.Inward,
			Name:    issueLink.Type.Name,
			Outward: issueLink.Type.Outward,
		}
		issueLinks = append(issueLinks, provider.IssueLink{
			ID:           issueLink.ID,
			OutwardIssue: &outwardIssue,
			InwardIssue:  &inwardIssue,
			Type:         issueType,
		})
	}
	var worklogs []provider.WorklogEntry
	for _, worklog := range issue.Fields.Worklog {
		authorAvatarUrls := provider.AvatarUrls{
			Small16x16:  worklog.Author.AvatarUrls.Small16x16,
			Small24x24:  worklog.Author.AvatarUrls.Small24x24,
			Medium32x32: worklog.Author.AvatarUrls.Medium32x32,
			Large48x48:  worklog.Author.AvatarUrls.Large48x48,
		}
		author := provider.User{
			AccountID:   worklog.Author.AccountID,
			AccountType: worklog.Author.AccountType,
			Active:      worklog.Author.Active,
			AvatarUrls:  authorAvatarUrls,
			DisplayName: worklog.Author.DisplayName,
			Key:         worklog.Author.Key,
			Name:        worklog.Author.Name,
			Self:        worklog.Author.Self,
		}
		var commentContentDetails []provider.ContentDetail
		for _, contentDetail := range worklog.Comment.Content {
			var commentTextContents []provider.TextContent
			for _, textContent := range contentDetail.Content {
				commentTextContents = append(commentTextContents, provider.TextContent{
					Type: textContent.Type,
					Text: textContent.Text,
				})
			}
			commentContentDetails = append(commentContentDetails, provider.ContentDetail{
				Type:    contentDetail.Type,
				Content: commentTextContents,
			})
		}
		comment := provider.Content{
			Type:    worklog.Comment.Type,
			Version: worklog.Comment.Version,
			Content: commentContentDetails,
		}
		updateAuthorAvatarUrls := provider.AvatarUrls{
			Small16x16:  worklog.Author.AvatarUrls.Small16x16,
			Small24x24:  worklog.Author.AvatarUrls.Small24x24,
			Medium32x32: worklog.Author.AvatarUrls.Medium32x32,
			Large48x48:  worklog.Author.AvatarUrls.Large48x48,
		}
		updateAuthor := provider.User{
			AccountID:   worklog.UpdateAuthor.AccountID,
			AccountType: worklog.UpdateAuthor.AccountType,
			Active:      worklog.UpdateAuthor.Active,
			AvatarUrls:  updateAuthorAvatarUrls,
			DisplayName: worklog.UpdateAuthor.DisplayName,
			Key:         worklog.UpdateAuthor.Key,
			Name:        worklog.UpdateAuthor.Name,
			Self:        worklog.UpdateAuthor.Self,
		}
		visiblity := provider.Visibility{
			Identifier: worklog.Visibility.Identifier,
			Type:       worklog.Visibility.Type,
			Value:      worklog.Visibility.Value,
		}
		worklogs = append(worklogs, provider.WorklogEntry{
			Author:           author,
			Comment:          comment,
			ID:               worklog.ID,
			IssueID:          worklog.IssueID,
			Self:             worklog.Self,
			Started:          worklog.Started,
			TimeSpent:        worklog.TimeSpent,
			TimeSpentSeconds: worklog.TimeSpentSeconds,
			UpdateAuthor:     updateAuthor,
			Updated:          worklog.Updated,
			Visibility:       visiblity,
		})
	}
	timeTracking := provider.TimeTracking{
		OriginalEstimate:         issue.Fields.TimeTracking.OriginalEstimate,
		OriginalEstimateSeconds:  issue.Fields.TimeTracking.OriginalEstimateSeconds,
		RemainingEstimate:        issue.Fields.TimeTracking.RemainingEstimate,
		RemainingEstimateSeconds: issue.Fields.TimeTracking.RemainingEstimateSeconds,
		TimeSpent:                issue.Fields.TimeTracking.TimeSpent,
		TimeSpentSeconds:         issue.Fields.TimeTracking.TimeSpentSeconds,
	}
	fields := provider.Fields{
		Watcher:      watcher,
		Attachment:   attachments,
		SubTasks:     subTasks,
		Description:  description,
		Project:      project,
		Comment:      comments,
		IssueLinks:   issueLinks,
		Worklog:      worklogs,
		TimeTracking: timeTracking,
		Updated:      issue.Fields.Updated,
	}
	value := models.Resource{
		ID:   issue.ID,
		Name: issue.Key,
		Description: provider.IssueDescription{
			ID:     issue.ID,
			Key:    issue.Key,
			Self:   issue.Self,
			Fields: fields,
			Expand: issue.Expand,
		},
	}

	return &value, nil
}

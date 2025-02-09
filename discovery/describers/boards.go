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

func ListBoards(ctx context.Context, client *jira.Client, stream *models.StreamSender, isLocal bool) ([]models.Resource, error) {
	var boards []provider.BoardJSON
	var boardListResp provider.BoardListResponse
	baseURL := "rest/agile/1.0/board"
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

		_, err = client.Do(req, &boardListResp)
		if err != nil {
			return nil, fmt.Errorf("request execution failed: %w", err)
		}

		boards = append(boards, boardListResp.Values...)

		last = boardListResp.StartAt + len(boardListResp.Values)
		if boardListResp.IsLast {
			break
		}
	}

	var values []models.Resource
	for _, board := range boards {
		value := models.Resource{
			ID:   strconv.Itoa(board.ID),
			Name: board.Name,
			Description: provider.BoardDescription{
				ID:   board.ID,
				Name: board.Name,
				Self: board.Self,
				Type: board.Type,
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

func GetBoard(ctx context.Context, client *jira.Client, resourceID string, isLocal bool) (*models.Resource, error) {
	var board provider.BoardJSON
	baseURL := "rest/agile/1.0/board/"
	finalURL := fmt.Sprintf("%s%s", baseURL, resourceID)

	req, err := client.NewRequest("GET", finalURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	_, err = client.Do(req, &board)
	if err != nil {
		return nil, fmt.Errorf("request execution failed: %w", err)
	}

	value := models.Resource{
		ID:   strconv.Itoa(board.ID),
		Name: board.Name,
		Description: provider.BoardDescription{
			ID:   board.ID,
			Name: board.Name,
			Self: board.Self,
			Type: board.Type,
		},
	}

	return &value, nil
}

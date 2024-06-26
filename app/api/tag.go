package api

import (
	"context"
	"net/http"
	"strconv"
)

const (
	// getTagsPath is a path for getting tags.
	getTagsPath = "/tags"
)

// GetTags requests categories
func (c *Client) GetTags(ctx context.Context, page int, limit int) (*http.Response, error) {
	resp, err := c.Do(ctx, http.MethodGet, getTagsPath, map[string]string{"page": strconv.Itoa(page), "limit": strconv.Itoa(limit)}, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

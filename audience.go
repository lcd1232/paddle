package paddle

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type audienceResponse struct {
	baseAPIResponse
	UserID int `json:"user_id"`
}

type audienceRequest struct {
	Email            string     `schema:"email"`
	MarketingConsent customBool `schema:"marketing_consent"`
}

func (c *Client) Audience(ctx context.Context, email string, marketingConsent bool) (userID int, err error) {
	if c.vendorID == "" {
		return 0, errors.New("paddle: vendor_id is required")
	}
	body := audienceRequest{
		Email:            email,
		MarketingConsent: customBool(marketingConsent),
	}
	req, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("1.0/audience/%s/add", c.vendorID), body, true, false)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var data audienceResponse
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return 0, errors.WithStack(err)
		}
		if data.IsSuccess() {
			return data.UserID, nil
		}
		return 0, &data
	}
	return 0, errors.Errorf("got status code: %d", resp.StatusCode)
}

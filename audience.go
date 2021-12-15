package paddle

import "context"

type audienceResponse struct {
	baseAPIResponse
	UserID int `json:"user_id"`
}

type audienceRequest struct {
	Email            string     `schema:"email"`
	MarketingConsent customBool `schema:"marketing_consent"`
}

func (c *Client) Audience(ctx context.Context, email string, marketingConsent bool) (userID int, err error) {
	panic("implement")
}

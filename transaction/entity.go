package transaction

import (
	"govn-crowdfunding/campaign"
	"govn-crowdfunding/user"
	"time"
)

type Transaction struct {
	ID         int64             `json:"id"`
	CampaignID int64             `json:"campaign_id"`
	UserID     int64             `json:"user_id"`
	Amount     int64             `json:"amount"`
	Status     string            `json:"status"`
	Code       string            `json:"code"`
	User       user.User         `json:"user"`
	Campaign   campaign.Campaign `json:"campaign"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
}

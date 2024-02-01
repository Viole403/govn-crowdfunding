package transaction

import "govn-crowdfunding/user"

type GetCampaignTransactionsInput struct {
	ID   int       `uri:"id" binding:"required"`
	User user.User `json:"user"`
}

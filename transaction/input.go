package transaction

import "bwastartup/user"

type GetCampaignTransactionsDetailInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

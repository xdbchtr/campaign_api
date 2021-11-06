package campaign

import "bwastartup/user"

type GetCmpaignDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
type CreateCampaignInput struct {
	Name             string `json:"name" binging:"required"`
	ShortDescription string `json:"short_description" binging:"required"`
	Description      string `json:"description" binging:"required"`
	GoalAmount       int    `json:"goal_amount" binging:"required"`
	Perks            string `json:"perks" binging:"required"`
	User             user.User
}

type CreateCampaignImageInput struct {
	CampaignID int  `form:"campaign_id" binding:"required"`
	IsPrimary  bool `form:"is_primary"`
	User       user.User
}

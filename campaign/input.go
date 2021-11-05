package campaign

type GetCmpaignDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

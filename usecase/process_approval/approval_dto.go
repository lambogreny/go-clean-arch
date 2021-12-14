package process_approval

type ApprovalDtoInput struct {
	user string `form:"user" binding:"required"`
}

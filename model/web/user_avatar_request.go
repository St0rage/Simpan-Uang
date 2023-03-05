package web

type UserAvatarRequest struct {
	Avatar string `json:"avatar" binding:"required"`
}

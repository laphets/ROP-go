package user

type GetInfoResponse struct {

}

type AvatarRequest struct {
	URL string `json:"url" binding:"required"`
}
package model

type UpdateParams struct {
	Username    *string
	Email       *string
	AvatarUrl   *string
	IsStreaming *bool
	StreamKey   *string
	Role        *string
}

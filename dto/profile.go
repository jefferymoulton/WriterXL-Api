package dto

import "writerxl-api/models"

type ProfileDTO struct {
	Email       string `json:"email"`
	Nickname    string `json:"nickname"`
	Name        string `json:"name"`
	Picture     string `json:"picture"`
	Description string `json:"description"`
}

func MapDTO(user models.Profile) ProfileDTO {
	var dto ProfileDTO
	dto.Email = user.Email
	dto.Name = user.Name
	dto.Nickname = user.Nickname
	dto.Picture = user.Picture
	dto.Description = user.Description

	return dto
}

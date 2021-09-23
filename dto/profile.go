package dto

import "writerxl-api/models"

type ProfileDTO struct {
	Email       string `json:"email"`
	Nickname    string `json:"nickname"`
	Name        string `json:"name"`
	Picture     string `json:"picture"`
	Description string `json:"description"`
}

func MapProfileDTO(model models.Profile) ProfileDTO {
	var dto ProfileDTO
	dto.Email = model.Email
	dto.Name = model.Name
	dto.Nickname = model.Nickname
	dto.Picture = model.Picture
	dto.Description = model.Description

	return dto
}

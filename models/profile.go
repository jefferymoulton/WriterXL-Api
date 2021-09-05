package models

type Profile struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Email       string `json:"email" gorm:"unique"`
	Nickname    string `json:"nickname"`
	Name        string `json:"name"`
	Picture     string `json:"picture"`
	Description string `json:"description"`
}

type ProfileInput struct {
	Email       string `json:"email" binding:"required"`
	Nickname    string `json:"nickname" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Picture     string `json:"picture"`
	Description string `json:"description"`
}

func CreateProfile(input ProfileInput) (Profile, error) {
	profile := mapInput(input)

	if err := DB.Create(&profile).Error; err != nil {
		return Profile{}, err
	}

	return profile, nil
}

func UpdateProfile(input ProfileInput) (Profile, error) {
	existing, err := GetProfileByEmail(input.Email)

	if err != nil {
		return Profile{}, err
	}

	profile := mapInput(input)
	profile.ID = existing.ID

	if err := DB.Save(&profile).Error; err != nil {
		return Profile{}, err
	}

	return profile, nil
}

func GetProfileByEmail(email string) (Profile, error) {
	var profile Profile

	if err := DB.Where("email = ?", email).First(&profile).Error; err != nil {
		return Profile{}, err
	}

	return profile, nil
}

func mapInput(input ProfileInput) Profile {
	var profile Profile
	profile.Email = input.Email
	profile.Name = input.Name
	profile.Nickname = input.Nickname
	profile.Picture = input.Picture
	profile.Description = input.Description

	return profile
}

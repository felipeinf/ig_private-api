package scrap

import (
	m "../models"
)

//GetUsernameInfo Get instagram user info
func GetUsernameInfo(username string) (*m.InstagramUser, error) {
	inst, err := GetInstagram(false)

	// Get user profile
	user, err := inst.Profiles.ByName(username)
	if err != nil {
		return nil, err
	}

	userInfo := &m.InstagramUser{
		ID:             user.ID,
		Username:       user.Username,
		FullName:       user.FullName,
		Bio:            user.Biography,
		ProfilePicture: user.ProfilePicURL,
		Private:        user.IsPrivate,
		Followers:      user.FollowerCount,
	}

	return userInfo, nil
}

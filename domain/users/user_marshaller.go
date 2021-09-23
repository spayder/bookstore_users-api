package users

import "encoding/json"

type PublicUser struct {
	Id        int64  `json:"id"`
	CreatedAt string `json:"created_at"`
	Status    string `json:"status"`
}

type PrivateUser struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	Status    string `json:"status"`
}

func (user User) Marshall(isPublic bool) interface{} {
	userJson, _ := json.Marshal(user)
	if isPublic {
		var publicUser PublicUser
		json.Unmarshal(userJson, &publicUser)
		return publicUser
	}

	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}

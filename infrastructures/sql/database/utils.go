package database

import (
	"forum-api/domains"
)

func (u *User) ToDomainsUser() domains.User {
	return domains.User{
		ID:        u.ID.String(),
		Username:  u.Username,
		Password:  u.Password,
		FullName:  u.Fullname,
		CreatedAt: u.CreatedAt.Time,
		UpdatedAt: u.UpdatedAt.Time,
	}
}

func (cur *CreateUserRow) ToSignupResponseData() domains.SignupResponseData {
	return domains.SignupResponseData{
		ID:       cur.ID.String(),
		Username: cur.Username,
		FullName: cur.Fullname,
	}
}

func (ctr *CreateThreadRow) ToAddThreadResponseData() domains.AddThreadResponseData {
	return domains.AddThreadResponseData{
		ID:    ctr.ID.String(),
		Title: ctr.Title,
		Owner: ctr.Owner.String(),
	}
}

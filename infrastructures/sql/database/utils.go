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

func (ctr *CreateThreadRow) ToAddTaskResponseData() domains.AddTaskResponseData {
	return domains.AddTaskResponseData{
		ID:    ctr.ID.String(),
		Title: ctr.Title,
		Owner: ctr.Owner.String(),
	}
}

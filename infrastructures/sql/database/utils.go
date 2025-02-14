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

func (t *GetThreadByIdRow) ToDomainsThread() domains.Thread {
	return domains.Thread{
		ID:    t.ID.String(),
		Title: t.Title,
		Body:  t.Body,
		Date:  t.Date.Time,
		Owner: t.Username,
	}
}

func (ctr *CreateCommentRow) ToAddCommentResponseData() domains.AddCommentResponseData {
	return domains.AddCommentResponseData{
		ID:      ctr.ID.String(),
		Content: ctr.Content,
		Owner:   ctr.Owner.String(),
	}
}

func (gcr *GetCommentsByThreadRow) ToGetThreadComments() domains.GetCommentsByThreadResponseData {
	return domains.GetCommentsByThreadResponseData{
		ID:       gcr.ID.String(),
		Username: gcr.Username,
		Date:     gcr.Date.Time,
		Content:  gcr.Content,
	}
}

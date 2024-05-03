package userModel

import sharedModel "github.com/quinntas/go-rest-template/api/shared/model"

type UserModel struct {
	sharedModel.SharedModel
	Email    string `db:"email"`
	Password string `db:"password"`
	RoleId   int    `db:"roleId"`
}

func NewUserModel(
	id int,
	pid string,
	createdAt string,
	updatedAt string,
	email string,
	password string,
	roleId int,
) UserModel {
	return UserModel{
		sharedModel.NewSharedModel(id, pid, createdAt, updatedAt),
		email,
		password,
		roleId,
	}
}

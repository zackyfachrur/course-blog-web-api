package controllers

import (
	"database/sql"
	"fmt"
	"myproject/src/models"
)

type UserController struct {
	DB *sql.DB
}

func (uc *UserController) InsertUser(user models.User) error {
	query := `INSERT INTO baxaric_user_account (user_username, user_fullname, user_job, user_email, user_phone_number, user_password, isdeleted) VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := uc.DB.Exec(query, user.Username, user.FirstName+" "+user.LastName, user.Job, user.Email, user.PhoneNumber, user.Password, false)

	if err != nil {
		return err
	}

	fmt.Println("User Berhasil di tambahkan")
	return nil
}

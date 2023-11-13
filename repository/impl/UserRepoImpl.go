package impl

import (
	"commonauthsvc/models"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Repository struct {
	DbConn *sql.DB
}

func (r *Repository) AddUser(ctx context.Context, userInfo models.UserInfo) (int, error) {
	query := "insert into Users (name,email, phoneNo, addedOn, updatedOn, encPassword, address, role, isActive, profileImage) values (?, ?, ?, ? , ?, ? , ? , ?, ?, ?)"

	response, err := r.DbConn.Exec(query, userInfo.Name, userInfo.Email, userInfo.PhoneNo, time.Now(), time.Now(),
		userInfo.Password, userInfo.Address, userInfo.Role, userInfo.IsActive, userInfo.ProfilePicture)
	if err != nil {
		fmt.Println("could not insert into db -> ", err)
		return 0, err
	}

	lastId, err := response.LastInsertId()
	if err != nil {
		fmt.Println("could not get lastId -> ", err)
		return 0, err
	}
	return int(lastId), nil
}

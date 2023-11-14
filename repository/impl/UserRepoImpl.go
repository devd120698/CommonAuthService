package impl

import (
	"commonauthsvc/constants"
	"commonauthsvc/models"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type Repository struct {
	DbConn *sqlx.DB
}

func (r *Repository) AddUser(ctx context.Context, userInfo models.UserInfoDB) (int, error) {
	query := "insert into Users (name,email, phoneNo, addedOn, updatedOn, encPassword, address, role, isActive, profileImage) values (?, ?, ?, ? , ?, ? , ? , ?, ?, ?)"

	response, err := r.DbConn.Exec(query, userInfo.Name, userInfo.Email, userInfo.PhoneNo, time.Now(), time.Now(),
		userInfo.Password, userInfo.Address, userInfo.Role, userInfo.IsActive, userInfo.ProfileImage)
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

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*models.UserInfoDB, error) {
	query := `select * from Users where email = ?`
	userInfo := make([]models.UserInfoDB, 0)
	err := r.DbConn.SelectContext(ctx, &userInfo, query, email)
	if err != nil {
		return nil, err
	}
	if len(userInfo) == 0 {
		return nil, &models.BaseError{ErrDetails: "No user with this email ID", ErrType: constants.InvalidRequest}
	}
	return &userInfo[0], nil
}

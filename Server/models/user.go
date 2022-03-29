package models

import (
	"NeChat/dao"
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Username string
	Password string
	Salt     string
}

type MyClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func AddUserInfo(userInfo *User) (err error) {
	sqlStr := "INSERT INTO user_infos (username, password, Salt) VALUES (?, ?, ?)"
	_, err = dao.DB.Exec(sqlStr, userInfo.Username, userInfo.Password, userInfo.Salt)
	if err != nil {
		fmt.Println("[ERROR]Inset data failed,", err)
		return
	}
	return nil
}

// FindPassword Find user's password and salt,and return.
func FindPassword(userInfo *User) (password string, salt string, err error) {
	sqlStr := "SELECT password,salt FROM user_infos WHERE username = ?"
	err = dao.DB.QueryRow(sqlStr, userInfo.Username).Scan(&password, &salt)
	if err != nil {
		fmt.Println("[ERROR]Select data failed,", err)
		return
	}
	return password, salt, nil
}

// FindUser Find user by username.Return "False" if not found.
func FindUser(username string) (bool, error) {
	sqlStr := "SELECT * FROM user_infos WHERE username = ?"
	stmt, err := dao.DB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("[ERROR]Prepare database failed,", err)
		return false, err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(stmt)

	rows, err := stmt.Query(username)
	if err != nil {
		fmt.Println("[ERROR]Query failed", err)
		return false, err
	}
	if rows.Next() {
		return false, nil
	}
	return true, nil
}

func Delete(username string) (err error) {
	sqlStr := "DELETE FROM user_infos WHERE username = ?"
	ret, err := dao.DB.Exec(sqlStr, username)
	if err != nil {
		fmt.Println("[ERROR]Inset data failed,", err)
		return
	}
	_, err = ret.RowsAffected()
	if err != nil {
		fmt.Println("[ERROR]Delete data failed,", err)
		return
	}
	return nil
}

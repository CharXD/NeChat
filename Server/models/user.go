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

func AddUserInfo(userInfo *User) {
	sqlStr := "INSERT INTO user_infos (username, password, Salt) VALUES (?, ?, ?)"
	_, err := dao.DB.Exec(sqlStr, userInfo.Username, userInfo.Password, userInfo.Salt)
	if err != nil {
		fmt.Println("[ERROR]Inset data failed,", err)
		return
	}
}

// FindPassword Find user's password and salt,and return.
func FindPassword(userInfo *User) (password string, salt string) {
	sqlStr := "select password,salt from user_infos where username = ?"
	err := dao.DB.QueryRow(sqlStr, userInfo.Username).Scan(&password, &salt)
	if err != nil {
		fmt.Println("[ERROR]Select data failed,", err)
		return
	}
	return password, salt
}

// FindUser Find user by username.Return "False" if not found.
func FindUser(username string) bool {
	sqlStr := "SELECT * FROM user_infos WHERE username = ?"
	stmt, err := dao.DB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("[ERROR]Prepare database failed,", err)
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
	}
	if rows.Next() {
		return false
	}
	return true
}

func Delete(username string) {
	sqlStr := "DELETE FROM user_infos WHERE username = ?"
	ret, err := dao.DB.Exec(sqlStr, username)
	if err != nil {
		fmt.Println()
	}
	_, err = ret.RowsAffected()
	if err != nil {
		fmt.Println()
		return
	}
}

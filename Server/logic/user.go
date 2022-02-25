package logic

import (
	"NeChat/config"
	"NeChat/models"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"math/rand"
	"regexp"
	"time"
)

func Login(userInfo *models.User) gin.H {
	password, salt := models.FindPassword(userInfo)
	cryptPassword := md5Crypto(userInfo.Password, salt)
	if password == cryptPassword {
		return gin.H{"code": 0, "msg": "Login success.", "token": createToken(userInfo.Username)}
	} else {
		return gin.H{"code": 1, "msg": "Username or password is wrong.", "token": ""}
	}
}

func Register(userInfo *models.User) gin.H {
	//账户名为4-16位, 可以包含字符串a-z A-Z 0-9 - _
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9-_]{4,16}$", userInfo.Username); ok {
		userInfo.Salt = generateSalt()
		userInfo.Password = md5Crypto(userInfo.Password, userInfo.Salt)
		if models.FindUser(userInfo.Username) {
			models.AddUserInfo(userInfo)
			return gin.H{"code": 0, "msg": "Register success.", "token": createToken(userInfo.Username)}
		} else {
			return gin.H{"code": 1, "msg": "Username is taken."}
		}
	} else {
		return gin.H{"code": "1", "msg": "Username is invalid."}
	}
}

func DeleteUser(username string) gin.H {
	models.Delete(username)
	return gin.H{"code": 0, "msg": "Account deleted successfully."}
}

func md5Crypto(password string, salt string) string {
	m := md5.New()
	m.Write([]byte(password))
	m.Write([]byte(salt))
	return hex.EncodeToString(m.Sum(nil))
}

func generateSalt() (salt string) {
	str := "123456789abcdef"
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		salt += string(str[rand.Intn(15)])
	}
	return
}

func createToken(username string) string {
	claim := models.MyClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,      //签发日期(1分钟前
			ExpiresAt: time.Now().Unix() + 60*60*2, //过期时间(2小时
			Issuer:    config.JWTIssuer,            //签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedString, err := token.SignedString([]byte(config.JWTSingedKey))
	if err != nil {
		fmt.Println(err)
	}
	return signedString
}

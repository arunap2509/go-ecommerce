package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arunap2509/ecommerce/models"
	"github.com/arunap2509/ecommerce/token"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
) 

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

func VerifyPassword(userPass string, hasedPass string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(hasedPass), []byte(userPass))
	valid := true
	msg := ""

	if err != nil {
		msg = "username or password is incorrect"
		valid = false
	}

	return  valid, msg
}

func SignUp(ctx *gin.Context) {
	var signinUser models.SignInUser
	var existingUsers []models.User
	var newUser models.User

	if err := ctx.BindJSON(&signinUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error while parsing the request"})
		return
	}

	db.Where("user_name = @userName or email = @email", 
	map[string]interface{}{"userName": signinUser.UserName, "email": signinUser.Email}).Find(&existingUsers)

	if len(existingUsers) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username or email already exists"})
		return
	}

	newUser.FirstName = signinUser.FirstName
	newUser.LastName = signinUser.LastName
	newUser.UserName = signinUser.UserName
	newUser.Email = signinUser.Email
	newUser.Phone = signinUser.Phone
	newUser.Password = HashPassword(signinUser.Password)
	newUser.IsAdmin = signinUser.IsAdmin

	validationErr := newUser.Validate()

	if validationErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "the data you provided is invalid"})
	}

	response := db.Create(&newUser)

	if response.Error != nil || response.RowsAffected == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong pls try again"})
	}

	token, refreshToken, _ := token.TokenGenerator(signinUser.Email, signinUser.FirstName, signinUser.LastName, newUser.Id.String(), newUser.IsAdmin)

	newUser.Token = token
	newUser.RefreshToken = refreshToken

	db.Model(&newUser).Updates(map[string]interface{}{"token": token, "refresh_token": refreshToken})

	ctx.JSON(http.StatusCreated, "successfully signed in the user")

	return
}

func Login(ctx *gin.Context) {
	var loginUser models.LoginUser
	var user models.User
	var err error

	if err = ctx.BindJSON(&loginUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error while parsing the request"})
		return
	}

	if err = db.Where("user_name = ?", loginUser.UserName).First(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "no user found, pls sigin first"})
		return
	}

	passValid, msg := VerifyPassword(loginUser.Password, user.Password)

	if !passValid {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": msg})
		fmt.Println(msg)
		return
	}

	token, refreshToken, _ := token.TokenGenerator(user.Email, user.FirstName, user.LastName, user.Id.String(), user.IsAdmin)

	updateTokenForUser(token, refreshToken, &user, ctx)

	ctx.JSON(http.StatusFound, user)
}

func updateTokenForUser(token string, refreshToken string, user *models.User, ctx *gin.Context) {
	if err := db.Model(&user).Updates(map[string]interface{}{"token": token, "refresh_token": refreshToken}).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		return
	}
}

func Logout(ctx *gin.Context) {
	
}
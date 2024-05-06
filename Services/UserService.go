package Services

import (
	"Authentication/Configs"
	"Authentication/Models"
	"Authentication/Models/Requests"
	"Authentication/Models/Response"
	"Authentication/Orm"
	"Authentication/Utils"
	"github.com/labstack/echo/v4"
	_ "golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Register(c echo.Context) error {
	log.Println("Start Execute Register")
	var user models.UserRequest
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	if user.Firstname == "" || user.Lastname == "" || user.Email == "" || user.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"httpStatus": strconv.Itoa(http.StatusBadRequest),
			"time":       time.Now().Format("2006-01-02 15:04:05"),
			"message":    "field is required"})
	}

	hashPass, err := Utils.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"httpStatus": strconv.Itoa(http.StatusInternalServerError),
			"time":       time.Now().Format("2006-01-02 15:04:05"),
			"message":    err.Error()})
	}

	userModel := models.User{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Password:  hashPass,
	}
	userExist, err := Orm.FindUserByEmail(userModel.Email)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"httpStatus": strconv.Itoa(http.StatusInternalServerError),
			"time":       time.Now().Format("2006-01-02 15:04:05"),
			"message":    err.Error()})
	}
	if userExist == true {
		return c.JSON(http.StatusConflict, map[string]interface{}{
			"httpStatus": strconv.Itoa(http.StatusConflict),
			"time":       time.Now().Format("2006-01-02 15:04:05"),
			"message":    "user already exists",
		})
	}
	if db := Orm.SaveUser(&userModel); db.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"httpStatus": strconv.Itoa(http.StatusInternalServerError),
			"time":       time.Now().Format("2006-01-02 15:04:05"),
			"message":    db.Error.Error()})
	}
	log.Println("End Execute Register")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"httpStatus": strconv.Itoa(http.StatusOK),
		"time":       time.Now().Format("2006-01-02 15:04:05"),
		"username":   userModel.Email,
		"message":    "Success Registered",
	})

}

func Login(c echo.Context) error {
	log.Println("Start Execute Login")
	var user Requests.LoginRequest
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"httpStatus": strconv.Itoa(http.StatusBadRequest),
			"time":       time.Now().Format("2006-01-02 15:04:05"),
			"message":    err.Error()})
	}
	//check user exist
	UserExist, err := Orm.FindUserByEmail(user.Email)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"httpStatus": strconv.Itoa(http.StatusInternalServerError),
			"time":       time.Now().Format("2006-01-02 15:04:05"),
			"message":    err.Error()})
	}

	if UserExist != true {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"httpStatus": strconv.Itoa(http.StatusUnauthorized),
			"time":       time.Now().Format("2006-01-02 15:04:05"),
			"message":    "user not exist",
		})
	}

	u, err := Orm.FindUser(user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"httpStatus": strconv.Itoa(http.StatusInternalServerError),
			"time":       time.Now().Format("2006-01-02 15:04:05"),
			"message":    "user not found"})
	}
	// generate Token
	token, err := Configs.Token(u.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"httpStatus": strconv.Itoa(http.StatusInternalServerError),
			"time":       time.Now().Format("2006-01-02 15:04:05"),
			"message":    err.Error()})
	}
	log.Println("token is generate:", token)

	if err := Utils.ComparePassword(u.Password, user.Password); err != nil {
		log.Println("End Execute Login")
		return c.JSON(http.StatusOK, map[string]interface{}{
			"httpStatus": strconv.Itoa(http.StatusOK),
			"time":       time.Now().Format("2006-01-02 15:04:05"),
			"message":    "Login Failed",
			"user":       user.Email,
		})
	} else {
		log.Println("End Execute Login")
		return c.JSON(http.StatusOK, map[string]interface{}{
			"httpStatus": strconv.Itoa(http.StatusOK),
			"time":       time.Now().Format("2006-01-02 15:04:05"),
			"message":    "Login Success",
			"user":       u.Email,
			"token":      token,
		})
	}
}

func UserAll(c echo.Context) error {
	u, err := Orm.FindAllUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"httpStatus": strconv.Itoa(http.StatusInternalServerError),
			"time":       time.Now().Format("2006-01-02 15:04:05"),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"httpStatus": strconv.Itoa(http.StatusOK),
		"time":       time.Now().Format("2006-01-02 15:04:05"),
		"user":       u,
	})

}

func Profile(c echo.Context) error {
	token := c.Request().Header.Get("token")
	id, err := Configs.PareToken(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"httpStatus": strconv.Itoa(http.StatusInternalServerError),
			"time":       time.Now().Format("2006-01-02 15:04:05"),
			"message":    err.Error(),
		})
	}
	user, err := Orm.FindUserByUserId(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"httpStatus": strconv.Itoa(http.StatusInternalServerError),
			"time":       time.Now().Format("2006-01-02 15:04:05"),
			"message":    err.Error(),
		})
	}
	userResponse := Response.User{
		ID:        user.ID,
		FirstName: user.Firstname,
		LastName:  user.Lastname,
		Email:     user.Email,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"httpStatus": strconv.Itoa(http.StatusOK),
		"time":       time.Now().Format("2006-01-02 15:04:05"),
		"user":       userResponse,
	})
}

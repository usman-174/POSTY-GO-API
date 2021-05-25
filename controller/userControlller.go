package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-playground/validator"
	"github.com/usman-174/database"
	"github.com/usman-174/models"
	"golang.org/x/crypto/bcrypt"
)

type LoginData struct {
	Email    string
	Password string
}

const SecretKey = "secret"

var err error
var Users []models.User
var reqBody LoginData

func Register(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDataBase()
	user := models.User{}
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		respondWithJSON(w, map[string]string{
			"Error": "Registration Failed",
			"Msg":   err.Error(),
		})
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 11)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"msg": err.Error(),
		})
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	user.Password = string(password)

	err = db.Create(&user).Error
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(map[string]string{
			"msg": err.Error(),
		})
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	respondWithJSON(w, user)
}

func empty(s string) bool {
	return len(strings.TrimSpace(s)) < 4
}

func Login(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err = json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		fmt.Println(err)
		respondWithJSON(w, map[string]string{
			"Error": "Login Failed",
			"Msg":   err.Error(),
		})
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	emptyemail := empty(reqBody.Email)
	// emptypass := empty(reqBody.Password)
	if emptyemail {
		respondWithJSON(w, map[string]string{
			"Error": "Please enter valid email",
		})
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	// if emptypass {
	// 	respondWithJSON(w, map[string]string{
	// 		"Error": "Please enter a valid password",
	// 	})
	// 	http.Error(w, "Bad Request", http.StatusBadRequest)
	// 	return
	// }

	db := database.ConnectDataBase()
	err = db.First(&user, "email = ?", reqBody.Email).Error
	if err != nil {
		fmt.Println(err)
		respondWithJSON(w, map[string]string{
			"Error": "User not found.",
			"Msg":   err.Error(),
		})
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	errs := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqBody.Password))
	if errs != nil {
		fmt.Println(err)
		respondWithJSON(w, map[string]string{
			"Error": "Invalid Password",
			"Msg":   err.Error(),
		})
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})
	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		fmt.Println("INVALID SECRET")
		fmt.Println(err)
		respondWithJSON(w, map[string]string{
			"Error": "Could not login",
			"Msg":   err.Error(),
		})
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	cookie := http.Cookie{Name: "token", Value: token, HttpOnly: true, Path: "/", Expires: time.Now().Add(time.Hour * 24)}
	http.SetCookie(w, &cookie)
	respondWithJSON(w, map[string]string{
		"msg": "You are logged In",
	})
}
func GetUser(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value(Mykey).(models.User)
	respondWithJSON(w, user)

}

func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LOGGIG YOU OUT")
	cookie := http.Cookie{Name: "token", Value: "", HttpOnly: true, Path: "/", Expires: time.Now().Add(-time.Hour)}

	http.SetCookie(w, &cookie)
	respondWithJSON(w, map[string]string{
		"Msg": "You are logged out now.",
	})
}
func respondWithJSON(w http.ResponseWriter, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

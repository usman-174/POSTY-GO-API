package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/usman-174/database"
	"github.com/usman-174/models"
)

var like *models.Like

func Likepost(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(Mykey).(models.User)
	request := map[string]uint{}
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Println(err.Error())
		respondWithJSON(w, map[string]string{
			"Error": "Invalid arguments",
			"Msg":   err.Error(),
		})
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	db := database.ConnectDataBase()

	// postId, err := strconv.Atoi(request["id"])
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	respondWithJSON(w, map[string]string{
	// 		"Error": "Invalid req.body",
	// 		"Msg":   err.Error(),
	// 	})
	// 	http.Error(w, "Bad Request", http.StatusBadRequest)
	// 	return
	// }
	foundpost := models.Post{}
	err = db.Find(&foundpost, "id = ?", request["id"]).Error
	if err != nil {
		fmt.Println(err.Error())
		respondWithJSON(w, map[string]string{
			"Error": "Post not found",
			"Msg":   err.Error(),
		})
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	db.Find(&like, "post_id = ?", request["id"])

	if like.ID != 0 && like.UserID == user.ID {
		db.Exec("DELETE FROM likes where id=? ", like.ID)
		like = &models.Like{}
		respondWithJSON(w, map[string]string{
			"msg": "Unliked post",
		})
		return
	}
	like.UserID = user.ID
	like.PostID = foundpost.ID

	err = db.Create(&like).Error
	if err != nil {
		fmt.Println(err.Error())
		respondWithJSON(w, map[string]string{
			"Error": "Couldnt like the post",
			"Msg":   err.Error(),
		})
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	err = db.Preload("Likes").Preload("User").Find(&foundpost, "id = ?", request["id"]).Error
	if err != nil {
		fmt.Println(err.Error())
		respondWithJSON(w, map[string]string{
			"Error": "Post not found",
			"Msg":   err.Error(),
		})
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	respondWithJSON(w, foundpost)
}

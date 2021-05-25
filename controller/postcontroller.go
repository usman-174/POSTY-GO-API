package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/usman-174/database"
	"github.com/usman-174/models"
)

func Post(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(Mykey).(models.User)
	post := models.Post{}
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		fmt.Println(err)
		respondWithJSON(w, map[string]string{
			"Error": "Registration Failed",
			"Msg":   err.Error(),
		})
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	db := database.ConnectDataBase()
	post.UserID = user.ID

	err = db.Create(&post).Error
	if err != nil {
		fmt.Println(err)
		respondWithJSON(w, map[string]string{
			"Error": "Couldn't create post",
			"Msg":   err.Error(),
		})
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	err = db.Preload("User").Find(&post, "id = ?", post.ID).Error
	if err != nil {
		fmt.Println(err)
		respondWithJSON(w, map[string]string{
			"Error": "Couldn't create post",
			"Msg":   err.Error(),
		})
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	respondWithJSON(w, post)
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDataBase()
	getPost := []*models.Post{}
	err = db.Preload("User").Preload("Likes").Find(&getPost).Error
	if err != nil {
		fmt.Println(err)
		respondWithJSON(w, map[string]string{
			"Error": "Could not get posts",
			"Msg":   err.Error(),
		})
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if len(getPost) == 0 {
		respondWithJSON(w, map[string]string{
			"Error": "Not posts found",
		})
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	fmt.Println(getPost)
	respondWithJSON(w, &getPost)

}
func GetPost(w http.ResponseWriter, r *http.Request) {
	post := models.Post{}

	err = json.NewDecoder(r.Body).Decode(&post)
	db := database.ConnectDataBase()
	if err != nil {
		fmt.Println(err.Error())
		respondWithJSON(w, map[string]string{
			"Error": "Invalid Arguments",
			"Msg":   err.Error(),
		})
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	err = db.Preload("User").Preload("Likes").Find(&post, "id = ?", post.ID).Error
	respondWithJSON(w, post)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value(Mykey).(models.User)
	request := map[string]int{}
	err = json.NewDecoder(r.Body).Decode(&request)
	fmt.Printf("The id is = %v\n", request["id"])
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
	foundPost := &models.Post{}
	err = db.Find(foundPost, "id = ?", request["id"]).Error
	if err != nil {
		fmt.Println(err.Error())
		respondWithJSON(w, map[string]string{
			"Error": "Invalid arguments",
			"Msg":   err.Error(),
		})
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if foundPost.UserID != user.ID {
		respondWithJSON(w, map[string]string{
			"Error": "Only the Post Author can delete the post",
		})
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	db.Exec("DELETE FROM posts where id=?", request["id"])

	respondWithJSON(w, map[string]string{
		"msg": "Post Deleted Successfully",
	})

}

func UpdatePost(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value(Mykey).(models.User)
	request := map[string]string{}
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Println(err.Error())
		respondWithJSON(w, map[string]string{
			"Error": "Invalid arguments 1",
			"Msg":   err.Error(),
		})
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	db := database.ConnectDataBase()
	foundPost := &models.Post{}
	postId, err := strconv.Atoi(request["id"])
	if err != nil {
		fmt.Println(err.Error())
		respondWithJSON(w, map[string]string{
			"Error": "Cannot convert string id to int id",
			"Msg":   err.Error(),
		})
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	err = db.Preload("User").Find(foundPost, "id = ?", postId).Error
	if err != nil {
		fmt.Println(err.Error())
		respondWithJSON(w, map[string]string{
			"Error": "Invalid arguments 2",
			"Msg":   err.Error(),
		})
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if foundPost.UserID != user.ID {
		respondWithJSON(w, map[string]string{
			"Error": "Only the Post Author can Update the post",
		})
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	foundPost.Body = request["body"]
	foundPost.Title = request["title"]
	err = db.Save(&foundPost).Error
	if err != nil {
		fmt.Println(err.Error())
		respondWithJSON(w, map[string]string{
			"Error": "Invalid arguments 3",
			"Msg":   err.Error(),
		})
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	respondWithJSON(w, foundPost)

}

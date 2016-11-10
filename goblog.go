package main

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

type Post struct {
  ID string `json:"id,omitempty"`
  Title string `json:"title,omitempty"`
  Body string `json:"body,omitempty"`
}

var posts []Post

func GetPostEndPoint(w http.ResponseWriter, req *http.Request) {
  params := mux.Vars(req)
  for _, item := range posts {
    if item.ID == params["id"] {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Post{})
}

func GetPostsEndPoint(w http.ResponseWriter, req *http.Request) {
  json.NewEncoder(w).Encode(posts)
}

func CreatePostEndPoint(w http.ResponseWriter, req *http.Request) {
  params := mux.Vars(req)
  var post Post
  _ = json.NewDecoder(req.Body).Decode(&post)
  post.ID = params["id"]
  posts = append(posts, post)
  json.NewEncoder(w).Encode(posts)
}

func DeletePostEndPoint(w http.ResponseWriter, req *http.Request) {
  params := mux.Vars(req)
  for index, item := range posts {
    if item.ID == params["id"] {
      posts = append(posts[:index], posts[index+1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(posts)
}

func main() {
  router := mux.NewRouter()
  posts = append(posts, Post{ID: "1", Title : "Hai", Body : "Kamu"})
  router.HandleFunc("/posts", GetPostsEndPoint).Methods("GET")
  //router.HandleFunc("/post/{id}", GetPostEndPoint).Methods("GET")
  router.HandleFunc("/post/{id}", CreatePostEndPoint).Methods("POST")
  router.HandleFunc("/post/{id}", DeletePostEndPoint).Methods("DELETE")
  log.Fatal(http.ListenAndServe(":12345", router))
}

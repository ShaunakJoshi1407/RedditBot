package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type FirstJSONLevel struct {
	Data SecondJSONLevel `json:"data"`
}

type SecondJSONLevel struct {
	Children []ThirdJSONLevel `json:"children"`
}

type ThirdJSONLevel struct {
	Data FinalJSONLevel `json:"data"`
}

type FinalJSONLevel struct {
	Ups   int    `json:"ups"`
	Title string `json:"title"`
	Link  string `json:"permalink"`
}

type Post struct {
	Ups   int
	Title string
	Link  string
}

func main() {

	getPosts()
	//fmt.Println(redditPost)

}

func getPosts() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://reddit.com/r/nosleep/top.json?limit=10&t=month", nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("User-Agent", "Reddit-NoSleepContentBot")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	var jsonResponse FirstJSONLevel

	err = json.Unmarshal(body, &jsonResponse)

	if err != nil {
		log.Fatal(err)
	}

	var postsArray []Post

	for i := range jsonResponse.Data.Children {

		//value := jsonResponse.Data.Children[i].Data.Ups

		jsonResponse.Data.Children[i].Data.Link = "https://reddit.com" + jsonResponse.Data.Children[i].Data.Link

		post := Post{Ups: jsonResponse.Data.Children[i].Data.Ups,
			Title: jsonResponse.Data.Children[i].Data.Title,
			Link:  jsonResponse.Data.Children[i].Data.Link,
		}

		postsArray = append(postsArray, post)

	}
	fmt.Println(postsArray)

	//jsonArray, _ := json.Marshal(postsArray)

	//fmt.Println(string(jsonArray))

	rand.Seed(time.Now().Unix())

	n := rand.Int() % (len(postsArray))

	selectedPost := postsArray[n]

	selectedPostJSON, _ := json.MarshalIndent(selectedPost, "", "\t")

	fmt.Println("Selected post -> ", string(selectedPostJSON))

}

func greaterThan(value int) bool {
	//Adding a < 100000 condition will probably fix the fixed post issue
	if value >= 1000 && value < 50000 {
		return true
	}
	return false
}

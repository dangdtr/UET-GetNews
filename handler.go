package main

import (
	// "container/list"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	// "os"
	"regexp"
	"strings"
	"time"

	"github.com/joho/godotenv"
	// godotenv "github.com/joho/godotenv"
)

const URL = "https://uet.vnu.edu.vn/category/tin-tuc/tin-sinh-vien/"
const REGEX = "<div class=\"item-thumbnail\">\\s*<a href=\"(.*?)\"\\s*title=\"(.*?)\">\\s*<img src=\"(.*?)\""
const NewsLength = 8

var isFirst = true

type News struct {
	Title    string `json:"title"`
	Url      string `json:"url"`
	ImageUrl string `json:"image_Url"`
}

func ReadFromUrl(link string) string {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	content, _ := ioutil.ReadAll(res.Body)
	err = res.Body.Close()
	if err != nil {
		return ""
	}
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

// func ReadFromFile() string {
// 	link := "test.html"
// 	content, err := os.ReadFile(link)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return string(content)
// }

func HandleRegex(str string) []News {
	currentNewsList := make([]News, NewsLength)

	regex, _ := regexp.Compile(REGEX)
	contentRegex, _ := regexp.Compile("\"(.*?)\"")

	findAllString := regex.FindAllString(str, NewsLength)

	for i, s := range findAllString {
		currentNewsList[i] = News{
			strings.Trim(contentRegex.FindAllString(s, 4)[2], "\""),
			strings.Trim(contentRegex.FindAllString(s, 4)[1], "\""),
			strings.Trim(contentRegex.FindAllString(s, 4)[3], "\""),
		}
	}
	return currentNewsList
}

func GetNewsList(preNewsList []News, currentNewsList []News) ([]News, bool) {
	newsList := make([]News, 0)

	for cur := 0; cur < NewsLength; cur++ {
		cmp := strings.Compare(currentNewsList[cur].Url, preNewsList[0].Url)
		if cmp != 0 {
			newsList = append(newsList, currentNewsList[cur])
		} else {
			if cur == 0 && cmp == 0 {
				return newsList, false
			}
			break
		}
	}

	for i := 0; i < NewsLength; i++ {
		preNewsList[i] = currentNewsList[i]
	}
	SavePreNewsList(preNewsList)
	return newsList, true
}

func SavePreNewsList(list []News) {

	file, _ := json.MarshalIndent(list, "", " ")

	_ = ioutil.WriteFile("data.json", file, 0644)
}

func LoadPreNewList() []News {
	file, _ := ioutil.ReadFile("data.json")

	data := []News{}

	_ = json.Unmarshal([]byte(file), &data)
	return data
}

func PrintNews(list []News) {
	for _, e := range list {
		fmt.Println(e)
	}
	fmt.Println()
}

func HandleAndSendMailCmd() {

	// SendMail("Chào mừng bạn đến với UET GET NEWS", "Cảm ơn bạn đã sử dụng")

	// preNewsList := list.New() // load from mongodb
	// fmt.Println(isFirst)

	var preNewsList []News
	if isFirst {
		preNewsList = HandleRegex(ReadFromUrl(URL))
		SavePreNewsList(preNewsList)
		isFirst = false
	} else {
		preNewsList = LoadPreNewList()
	}
	// fmt.Println(isFirst)
	// PrintNews(preNewsList)
	// cheat

	// for {
	fmt.Println(time.Now())
	currentNewsList := HandleRegex(ReadFromUrl(URL))

	newsList, isHasNews := GetNewsList(preNewsList, currentNewsList)

	if isHasNews && len(newsList) != 0 {
		for _, e := range newsList {
			SendEachNews(e)
		}
	}

	time.Sleep(60000 * time.Millisecond)
	// }

}

func Run() {
	HandleAndSendMailCmd()
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

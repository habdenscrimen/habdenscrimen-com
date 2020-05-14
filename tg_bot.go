package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// GithubFile represents type of GitHub file returned by GitHub API
type GithubFile struct {
	Name string
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	githubBaseURL := "https://api.github.com/repos/habdenscrimen/habdenscrimen-com/contents/_posts"

	// get last post file from GitHub repo
	file, err := getLastPostFile(githubBaseURL, "2b561ce69dd3e8671224759da627b04d4fadd1dd")
	if err != nil {
		fmt.Println(err)
		return &events.APIGatewayProxyResponse{
			StatusCode: 400,
		}, err
	}

	// form URL for the last post
	postURL := "https://habdenscrimen.netlify.app/" + strings.Replace(strings.ReplaceAll(file.Name, ".md", ""), "-", "/", 3)
	fmt.Println(postURL)

	// parse chatID from env variable
	chatID, err := strconv.ParseInt(os.Getenv("TELEGRAM_CHAT_ID"), 10, 64)
	if err != nil {
		fmt.Println(err)
		return &events.APIGatewayProxyResponse{
			StatusCode: 400,
		}, err
	}

	// send message with URL of the last post to channel
	err = sendMessageToChannel(os.Getenv("TELEGRAM_API_TOKEN"), postURL, chatID)
	if err != nil {
		fmt.Println(err)
		return &events.APIGatewayProxyResponse{
			StatusCode: 400,
		}, err
	}

	// all is well, relax
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}

func getLastPostFile(githubBaseURL, githubToken string) (GithubFile, error) {
	var file GithubFile

	// create new http client
	client := &http.Client{}

	// form new http request
	req, err := http.NewRequest("GET", githubBaseURL, nil)
	if err != nil {
		return file, err
	}
	// add Authorization header to request
	req.Header.Add("Authorization", "token "+githubToken)

	// send request
	res, err := client.Do(req)
	if err != nil {
		return file, nil
	}
	defer res.Body.Close()

	files := make([]GithubFile, 0)

	// get files from response
	err = json.NewDecoder(res.Body).Decode(&files)
	if err != nil {
		return file, nil
	}

	// return last file (last post)
	return files[len(files)-1], nil
}

func sendMessageToChannel(telegramAPIToken, message string, telegramChatID int64) error {
	// create a bot
	bot, err := tgbotapi.NewBotAPI(telegramAPIToken)
	if err != nil {
		return err
	}

	// form the message to channel
	msg := tgbotapi.NewMessage(telegramChatID, message)

	// send the message to channel
	if _, err := bot.Send(msg); err != nil {
		return err
	}

	// all is well, relax
	return nil
}

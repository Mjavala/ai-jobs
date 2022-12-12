package client

import (
	"bufio"
	"fmt"
	"os"

	"github.com/upwork/golang-upwork/api"
	"github.com/upwork/golang-upwork/api/routers/auth"
)

const cfgFile = "/Users/elo/Code/web2/ai-jobs/pkg/upwork/config.json" // update the path to your config file, or provide properties directly in your code

func Run() {
	client := api.Setup(api.ReadConfig(cfgFile))

	// we need an access token/secret pair in case we haven't received it yet
	if !client.HasAccessToken() {
		aurl := client.GetAuthorizationUrl("")

		// read verifier
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Visit the authorization url and provide oauth_verifier for further authorization")
		fmt.Println(aurl)
		verifier, _ := reader.ReadString('\n')

		token := client.GetAccessToken(verifier)
		fmt.Println(token)
	}

	// http.Response and []byte will be return, you can use any
	_, jsonDataFromHttp1 := auth.New(client).GetUserInfo()

	// here you can Unmarshal received json string, or do any other action(s)
	fmt.Println(string(jsonDataFromHttp1))
}

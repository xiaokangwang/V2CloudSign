package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/xiaokangwang/V2BuildAssist"
	"io/ioutil"
	"net/url"
	"os"
)

type ReturnValue struct {
	Error     string `json:"name"`
	ReleaseID int64  `json:"id"`
}

type InputValue struct {
	Version  string `json:"version"`
	Password string `json:"password"`
}

func HandleRequest(ctx context.Context, req events.APIGatewayV2HTTPRequest) (ReturnValue, error) {
	token := os.Getenv("GITHUB_TOKEN")

	owner := os.Getenv("GITHUB_REPO_OWNER")
	name := os.Getenv("GITHUB_REPO_NAME")

	Sowner := os.Getenv("GITHUB_SREPO_OWNER")
	Sname := os.Getenv("GITHUB_SREPO_NAME")

	Skey := os.Getenv("SIGNING_KEY")

	SPROJ := os.Getenv("SIGNING_PROJ")

	input := &InputValue{}

	data := req.Body

	if req.IsBase64Encoded {
		d := base64.NewDecoder(base64.StdEncoding, bytes.NewReader([]byte(data)))
		w, err := ioutil.ReadAll(d)
		if err != nil {
			return ReturnValue{}, err
		}
		data = string(w)
	}

	v, err := url.ParseQuery(data)
	if err != nil {
		return ReturnValue{}, err
	}
	input.Version = v.Get("version")
	input.Password = v.Get("password")

	d, _, err := V2BuildAssist.RequestForSign(token, Sowner, Sname, SPROJ, owner, name, input.Version, input.Password, Skey)

	return ReturnValue{Error: "OK", ReleaseID: d}, err
}

func main() {
	lambda.Start(HandleRequest)
}

package client

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path"

	credentials "cloud.google.com/go/iam/credentials/apiv1"
	"google.golang.org/api/option"
	csp "google.golang.org/genproto/googleapis/iam/credentials/v1"
)

const (
	SERVER = "http://localhost:4000/api"
)

type TestClient struct {
	token string
}

func New() TestClient {
	return TestClient{
		token: getToken(),
	}
}

func (tc TestClient) Get(url string) (*http.Response, error) {
	fullUrl := fmt.Sprintf("%s%s", SERVER, url)
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tc.token))
	return http.DefaultClient.Do(req)
}

func getToken() string {
	credsFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	ctx := context.Background()
	c, err := credentials.NewIamCredentialsClient(ctx, option.WithCredentialsFile(path.Join("..", credsFile)))
	if err != nil {
		panic(err)
	}
	defer c.Close()

	req := &csp.GenerateIdTokenRequest{
		Name:         "auto-tester@web-stack-starter.iam.gserviceaccount.com",
		Audience:     "web-stack-starter",
		IncludeEmail: true,
	}

	resp, err := c.GenerateIdToken(ctx, req)
	if err != nil {
		panic(err)
	}
	return resp.Token
}

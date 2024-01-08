package main

import (
	"context"
	"fmt"

	nethttplibrary "github.com/microsoft/kiota-http-go"
	auth "github.com/octokit/go-sdk/pkg/authentication"
	"github.com/octokit/go-sdk/pkg/github"
	"github.com/octokit/go-sdk/pkg/github/search"
)

func main() {
	tp := auth.NewTokenProvider()
	adapter, err := nethttplibrary.NewNetHttpRequestAdapter(tp)
	if err != nil {
		panic(err)
	}
	u := "blck-snwmn"
	q := fmt.Sprintf("is:open is:pr archived:false user:%s", u)

	client := github.NewApiClient(adapter)
	resp, err := client.Search().Issues().Get(context.Background(), &search.IssuesRequestBuilderGetRequestConfiguration{
		QueryParameters: &search.IssuesRequestBuilderGetQueryParameters{
			Q: &q,
		},
	})
	if err != nil {
		panic(err)
	}
	for _, item := range resp.GetItems() {
		fmt.Println(*item.GetTitle())
	}
}

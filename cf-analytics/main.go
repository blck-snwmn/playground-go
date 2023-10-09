package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	accountID := os.Getenv("ACCOUNT_ID")
	authToken := os.Getenv("AUTH_TOKEN")

	query := fmt.Sprintf(`{"query":"query{
		viewer{
		  accounts(filter:{
			accountTag:\"%s\"
		  }){
			streamMinutesViewedAdaptiveGroups(
			  filter:{
				date_geq:\"2023-08-01\"
				date_leq:\"2023-09-01\"
			  }
			  limit:100
			){
			  sum{
				minutesViewed
			  }
			  count
			  dimensions{
				datetimeFiveMinutes
			  }
			}
		  }
		}
	  }"}`, accountID)

	query = strings.ReplaceAll(query, "\n", "")
	query = strings.ReplaceAll(query, "\t", "")

	fmt.Println(query)

	req, err := http.NewRequest("POST", "https://api.cloudflare.com/client/v4/graphql", strings.NewReader(query))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

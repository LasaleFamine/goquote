// Really simple quote API wrapper for learning GO
package goquote

import (
    "encoding/json"
		"fmt"
    "net/http"
		"errors"
)

type Item struct {
    Title string
    Content string
		URL string `json:"link"`
}

func (i Item) String() string {
    return fmt.Sprintf("%s\n%s%s", i.Title, i.Content, i.URL)
}

// Pass an URL to the function to get the Item structured response.
func GetQuote(url string) ([]Item, error) {
    // url := fmt.Sprintf("", type)

    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return nil, errors.New(resp.Status)
    }
		var items []Item
		err = json.NewDecoder(resp.Body).Decode(&items)
    if err != nil {
        return nil, err
    }
    return items, nil
}

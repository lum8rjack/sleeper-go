package sleeper

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Verification   interface{} `json:"verification"`
	Username       string      `json:"username"`
	UserID         string      `json:"user_id"`
	Token          interface{} `json:"token"`
	SummonerRegion interface{} `json:"summoner_region"`
	SummonerName   interface{} `json:"summoner_name"`
	Solicitable    interface{} `json:"solicitable"`
	RealName       interface{} `json:"real_name"`
	Phone          interface{} `json:"phone"`
	Pending        interface{} `json:"pending"`
	Notifications  interface{} `json:"notifications"`
	Metadata       interface{} `json:"metadata"`
	IsBot          bool        `json:"is_bot"`
	Email          interface{} `json:"email"`
	DisplayName    string      `json:"display_name"`
	Deleted        interface{} `json:"deleted"`
	DataUpdated    interface{} `json:"data_updated"`
	Currencies     interface{} `json:"currencies"`
	Created        interface{} `json:"created"`
	Cookies        interface{} `json:"cookies"`
	Avatar         string      `json:"avatar"`
}

const (
	userEndpoint string = "/user"
)

// Gets the user's information by their username
func (c *Client) GetUserByUsername(username string) (User, error) {
	url := fmt.Sprintf("%s%s/%s", c.sleeperURL, userEndpoint, username)
	return c.getUser(url)
}

// Gets the user's information by their user id
func (c *Client) GetUserByID(id string) (User, error) {
	url := fmt.Sprintf("%s%s/%s", c.sleeperURL, userEndpoint, id)
	return c.getUser(url)
}

func (c *Client) getUser(url string) (User, error) {
	user := User{}

	data, err := c.getRequest(url)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(data, &user)

	return user, err
}

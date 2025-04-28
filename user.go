package sleeper

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Avatar         string      `json:"avatar"`
	Cookies        interface{} `json:"cookies"`
	Created        interface{} `json:"created"`
	Currencies     interface{} `json:"currencies"`
	DataUpdated    interface{} `json:"data_updated"`
	Deleted        interface{} `json:"deleted"`
	DisplayName    string      `json:"display_name"`
	Email          interface{} `json:"email"`
	IsBot          bool        `json:"is_bot"`
	Metadata       interface{} `json:"metadata"`
	Notifications  interface{} `json:"notifications"`
	Pending        interface{} `json:"pending"`
	Phone          interface{} `json:"phone"`
	RealName       interface{} `json:"real_name"`
	Solicitable    interface{} `json:"solicitable"`
	SummonerName   interface{} `json:"summoner_name"`
	SummonerRegion interface{} `json:"summoner_region"`
	Token          interface{} `json:"token"`
	UserID         string      `json:"user_id"`
	Username       string      `json:"username"`
	Verification   interface{} `json:"verification"`
}

// Get the user's information by their username.
// (GET `https://api.sleeper.app/v1/user/<username>`)
func (c *Client) GetUserByUsername(username string) (User, error) {
	url := fmt.Sprintf("%s/v1/user/%s", c.sleeperURL, username)
	return c.getUser(url)
}

// Get the user's information by their user id.
// (GET `https://api.sleeper.app/v1/user/<user_id>`)
func (c *Client) GetUserByID(id string) (User, error) {
	url := fmt.Sprintf("%s/v1/user/%s", c.sleeperURL, id)
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

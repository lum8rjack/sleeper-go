package sleeper

import "fmt"

const (
	avatarBaseURL string = "https://sleepercdn.com/avatars"
)

// Gets the user's avatar picture
func (c *Client) GetAvatar(avatar_id string) ([]byte, error) {
	// https://sleepercdn.com/avatars/<avatar_id>
	url := fmt.Sprintf("%s/%s", avatarBaseURL, avatar_id)
	return c.getRequest(url)
}

// Gets the user's avatar picture thumbnail
func (c *Client) GetAvatarThumbnail(avatar_id string) ([]byte, error) {
	// https://sleepercdn.com/avatars/thumbs/<avatar_id>
	url := fmt.Sprintf("%s/thumbs/%s", avatarBaseURL, avatar_id)
	return c.getRequest(url)
}

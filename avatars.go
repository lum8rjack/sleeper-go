package sleeper

import "fmt"

const (
	avatarBaseURL string = "https://sleepercdn.com/avatars"
)

// Get the user's avatar picture.
// (GET `https://sleepercdn.com/avatars/<avatar_id>`)
func (c *Client) GetAvatar(avatar_id string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", avatarBaseURL, avatar_id)
	return c.getRequest(url)
}

// Get the user's avatar picture thumbnail.
// (GET `https://sleepercdn.com/avatars/thumbs/<avatar_id>`)
func (c *Client) GetAvatarThumbnail(avatar_id string) ([]byte, error) {
	url := fmt.Sprintf("%s/thumbs/%s", avatarBaseURL, avatar_id)
	return c.getRequest(url)
}

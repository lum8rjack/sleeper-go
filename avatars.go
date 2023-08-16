package sleeper

import "fmt"

const (
	avatarsFullSizeEndpoint  string = "https://sleepercdn.com/avatars/"
	avatarsThumbnailEndpoint string = "https://sleepercdn.com/avatars/thumbs/"
)

// Gets the user's avatar picture
func (c *Client) GetAvatar(avatar_id string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", avatarsFullSizeEndpoint, avatar_id)
	return c.getRequest(url)
}

// Gets the user's avatar picture thumbnail
func (c *Client) GetAvatarThumbnail(avatar_id string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", avatarsThumbnailEndpoint, avatar_id)
	return c.getRequest(url)
}

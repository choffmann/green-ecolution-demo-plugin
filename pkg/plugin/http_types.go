package plugin

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name        string      `json:"name"`
	Path        string      `json:"path"`
	Version     string      `json:"version"`
	Description string      `json:"description"`
	Auth        AuthRequest `json:"auth"`
}

type ClientTokenResponse struct {
	AccessToken      string `json:"access_token"`
	IDToken          string `json:"id_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not_before_policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
}

type Token struct {
	AccessToken  string
	RefreshToken string
}

package twitch_auth_code_grant_flow

type OAuthRequest struct {
	ClientID     string `json:"client_id"`
	ForceVerify  bool   `json:"force_verify,omitempty"`
	RedirectURI  string `json:"redirect_uri"`
	ResponseType string `json:"response_type"`
	Scope        string `json:"scope"`
	StateString  string `json:"state,omitempty"`
}

package twitch_auth_code_grant_flow

import (
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Service) InitRoutes(group *gin.RouterGroup) {
	group.GET("authorize", s.handleAuthorize)
	group.GET("callback", s.handleCallback)
}

func (s *Service) handleAuthorize(c *gin.Context) {
	// Redirect the user to the Twitch authorization URL
	c.Redirect(http.StatusSeeOther, s.buildAuthorizeURL())
}

func (s *Service) buildAuthorizeURL() string {
	cfg, err := s.cfg.GetConfig(s.ConfigFilePath())
	if err != nil {
		s.logger.Error("getting config", "error", err)
		panic(err.Error())
	}

	group := cfg.Group(s.Name())

	const baseURL = "https://id.twitch.tv/oauth2/authorize"

	s.stateString = randomString(16)

	values := url.Values{}
	values.Set("response_type", "code")
	values.Set("client_id", group.GetString("ClientID"))
	values.Set("force_verify", group.GetString("ForceVerify"))
	values.Set("redirect_uri", group.GetString("RedirectURI"))
	values.Set("scope", group.GetString("Scope"))
	values.Set("state", s.stateString)

	fullURL := baseURL + "?" + values.Encode()
	return fullURL
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}

func (s *Service) handleCallback(c *gin.Context) {
	// Get the authorization code from the callback URL query parameters
	authCode := c.DefaultQuery("code", "")
	state := c.DefaultQuery("state", "")

	if state != s.stateString {
		s.Logger().Error("bad state string returned")
	}

	// Exchange the authorization code for an access token
	// You should perform a POST request here to Twitch's token endpoint

	// For the sake of the example, we'll just print the authorization code
	s.Logger().Info("got authorization code", "code", authCode)

	c.String(http.StatusOK, "Authorization code received.")
}

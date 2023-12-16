package twitch_soundboard

import (
	"github.com/gravestench/servicesmesh-examples/services/config_file"
)

const (
	keyOnTwitchConnect                = "OnConnect"
	keyOnTwitchWhisperMessage         = "OnWhisperMessage"
	keyOnTwitchPrivateMessage         = "OnPrivateMessage"
	keyOnTwitchClearChatMessage       = "OnClearChatMessage"
	keyOnTwitchClearMessage           = "OnClearMessage"
	keyOnTwitchRoomStateMessage       = "OnRoomStateMessage"
	keyOnTwitchUserNoticeMessage      = "OnUserNoticeMessage"
	keyOnTwitchUserStateMessage       = "OnUserStateMessage"
	keyOnTwitchGlobalUserStateMessage = "OnGlobalUserStateMessage"
	keyOnTwitchNoticeMessage          = "OnNoticeMessage"
	keyOnTwitchUserJoinMessage        = "OnUserJoinMessage"
	keyOnTwitchUserPartMessage        = "OnUserPartMessage"
	keyOnTwitchReconnectMessage       = "OnReconnectMessage"
	keyOnTwitchNamesMessage           = "OnNamesMessage"
	keyOnTwitchPingMessage            = "OnPingMessage"
	keyOnTwitchPongMessage            = "OnPongMessage"
	keyOnTwitchUnsetMessage           = "OnUnsetMessage"
	keyOnTwitchPingSent               = "OnPingSent"
)

func (s *Service) ConfigFilePath() string {
	return "twitch_integrated_soundboard.json"
}

func (s *Service) Config() (*config_file.Config, error) {
	return s.configManager.GetConfig(s.ConfigFilePath())
}

func (s *Service) DefaultConfig() config_file.Config {
	cfg := config_file.Config{}

	cfg.
		Group(keyOnTwitchPrivateMessage).
		SetDefault("chat trigger message", []string{"/path/to/audio/file.mp3"})

	return cfg
}

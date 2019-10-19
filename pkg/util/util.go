package util

import "github.com/boybird/go_api_development_demo/pkg/setting"

// Setup set jwtsecret
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}

package util

import "github.com/boybird/hello/pkg/setting"

// Setup set jwtsecret
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}

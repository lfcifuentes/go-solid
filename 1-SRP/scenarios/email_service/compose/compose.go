package compose

import "github.com/spf13/viper"

func NewComposeConfig() EmailConfig {
	return EmailConfig{
		From:            viper.GetString("email.from"),
		CharsetEncoding: viper.GetString("email.charset"),
		MaxToRecipients: viper.GetInt("email.max_to"),
	}
}

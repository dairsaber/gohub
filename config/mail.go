package config

import "gohub/pkg/config"

func init() {
	config.Add("mail", func() map[string]any {
		return map[string]any{

			// 默认是 Mailhog 的配置
			"smtp": map[string]any{
				"host": config.Env("MAIL_HOST", "localhost"),
				// 本地的mail默认的1025被占用了 这样会造成mail在send时候无法正常返回,造成卡死,这边使用7788 比较保险
				// MailHog  -smtp-bind-addr "0.0.0.0:7788"
				"port":     config.Env("MAIL_PORT", 7788),
				"username": config.Env("MAIL_USERNAME", ""),
				"password": config.Env("MAIL_PASSWORD", ""),
			},

			"from": map[string]any{
				"address": config.Env("MAIL_FROM_ADDRESS", "gohub@example.com"),
				"name":    config.Env("MAIL_FROM_NAME", "Gohub"),
			},
		}
	})
}

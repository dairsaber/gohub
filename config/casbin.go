package config

import "gohub/pkg/config"

func init() {
	config.Add("casbin", func() map[string]any {

		return map[string]any{
			// 模型路径
			"model_path": config.Env("MODEL_PATH", "config/rbac_model.conf"),
		}
	})
}

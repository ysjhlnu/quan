package response

import "quan/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}

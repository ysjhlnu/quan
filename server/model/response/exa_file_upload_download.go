package response

import "quan/model"

type ExaFileResponse struct {
	File model.ExaFileUploadAndDownload `json:"file"`
}

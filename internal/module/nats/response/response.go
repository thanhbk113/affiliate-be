package response

import (
	"affiliate/internal/util/pjson"

	"git.selly.red/Selly-Modules/natsio"
	"github.com/nats-io/nats.go"
)

// NatsResponse ..
type NatsResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    []byte `json:"data"`
}

type CommonResponseData struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

// StaffDataHelpCenter ...
type StaffDataHelpCenter struct {
	User  interface{} `json:"user"`
	Error string      `json:"error"`
}

// ListStaffChatResponse ...
type ListStaffChatResponse struct {
	Staffs []string `json:"staffs"`
	Error  string   `json:"error"`
}

// ResponseSuccess ...
func ResponseSuccess(msg *nats.Msg, data interface{}) {
	response(msg, NatsResponse{
		Success: true,
		Message: "",
		Data:    pjson.ConvertToBytes(data),
	})
}

// ResponseError ...
func ResponseError(msg *nats.Msg, err error) {
	response(msg, NatsResponse{
		Success: false,
		Message: err.Error(),
	})
}

func response(msg *nats.Msg, data interface{}) {
	_ = natsio.GetServer().Reply(msg, pjson.ConvertToBytes(data))
}

package systems

import (
	"encoding/json"
	"fmt"

	"mono.thienhang.com/pkg/context"
	"mono.thienhang.com/pkg/utils/application"
)

func (h *System) pingHandler(ctx *context.Context) {
	ctx.Data(200, "application/json", []byte("pong"))
}

func (s *System) infoHandler(ctx *context.Context) {
	data := application.GetAppStatus()

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}

	ctx.Data(200, "application/json", jsonData)
}

func (h *System) versionandler(ctx *context.Context) {
	ctx.Data(200, "application/json", []byte(application.GetVersion()))
}

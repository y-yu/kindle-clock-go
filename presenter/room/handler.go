package room

import (
	"context"
	"github.com/sethvargo/go-envconfig"
	"github.com/y-yu/kindle-clock-go/domain"
	"github.com/y-yu/kindle-clock-go/domain/model/config"
	"github.com/y-yu/kindle-clock-go/domain/usecase"
	"log/slog"
	"net/http"
)

type RoomInfoHandler struct {
	authConfig      config.AuthenticationConfiguration
	roomInfoUsecase usecase.GetRoomInfoUsecase
	clock           domain.Clock
}

func NewRoomInfoHandler(
	ctx context.Context,
	roomInfoUsecase usecase.GetRoomInfoUsecase,
	clock domain.Clock,
) *RoomInfoHandler {
	var c config.AuthenticationConfiguration
	if err := envconfig.Process(ctx, &c); err != nil {
		slog.Error("failed to process configuration for NewRoomInfoHandler", "err", err)
	}

	return &RoomInfoHandler{
		authConfig:      c,
		roomInfoUsecase: roomInfoUsecase,
		clock:           clock,
	}
}

func (h *RoomInfoHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if h.authConfig.Token != "" {
		if token := r.URL.Query().Get(h.authConfig.QueryKeyName); token != h.authConfig.Token {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
	}

	ctx := r.Context()
	roomInfo, err := h.roomInfoUsecase.Execute(ctx)
	if err != nil {
		slog.Error("RoomInfoUsecase.Execute failed", "err", err)
		w.Write([]byte("error!"))
	}
	svg, err := GeneratePNGImage(roomInfo, h.clock.Now())
	if err != nil {
		slog.Error("GeneratePNGImage failed", "err", err)
		w.Write([]byte("error!"))
	}
	w.Header().Set("Content-Type", "image/png")
	w.Write(svg.Bytes())
}

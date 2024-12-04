package websocket

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/rocketscienceinc/tictactoe-backend/internal/entity"
	"github.com/rocketscienceinc/tictactoe-backend/internal/pkg"
)

const (
	headerUpgrade            = "Upgrade"
	headerWebSocket          = "websocket"
	headerConnection         = "Connection"
	headerSecWebSocketKey    = "Sec-WebSocket-Key"
	headerSecWebSocketAccept = "Sec-WebSocket-Accept"
)

type gameUseCase interface {
	GetOrCreatePlayer(ctx context.Context, playerID string) (*entity.Player, error)

	GetOrCreateGame(ctx context.Context, playerID, gameType string) (*entity.Game, error)
	GetGameByPlayerID(ctx context.Context, playerID string) (*entity.Game, error)
	CreateOrJoinToPublicGame(ctx context.Context, playerID, gameType string) (*entity.Game, error)
	JoinGameByID(ctx context.Context, gameID, playerID string) (*entity.Game, error)

	MakeTurn(ctx context.Context, playerID string, cell int) (*entity.Game, error)
}

type Server struct {
	logger      *slog.Logger
	gameUseCase gameUseCase

	messageHandlers map[string]func(ctx context.Context, message *Message, w *bufio.ReadWriter) error

	connections map[string]*bufio.ReadWriter
}

func New(logger *slog.Logger, gameUseCase gameUseCase) *Server {
	server := &Server{
		logger:      logger,
		gameUseCase: gameUseCase,

		messageHandlers: make(map[string]func(context.Context, *Message, *bufio.ReadWriter) error),
		connections:     make(map[string]*bufio.ReadWriter),
	}

	server.messageHandlers["connect"] = server.handleConnect
	server.messageHandlers["game:new"] = server.handleNewGame
	server.messageHandlers["game:join"] = server.handleJoinGame
	server.messageHandlers["game:turn"] = server.handleGameTurn

	return server
}

func (that *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	that.upgradeToWebSocket(r.Context(), w, r)
}

// upgradeToWebSocket - upgrades the connection to WebSocket.
func (that *Server) upgradeToWebSocket(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	log := that.logger.With("method", "upgradeToWebSocket")

	if r.Header.Get(headerUpgrade) != headerWebSocket {
		log.Error("not upgrade to websocket")
		http.Error(w, "not a websocket upgrade", http.StatusBadRequest)
		return
	}

	wsKey := r.Header.Get(headerSecWebSocketKey)
	acceptKey := pkg.GenerateAcceptKey(wsKey)

	w.Header().Set(headerUpgrade, headerWebSocket)
	w.Header().Set(headerConnection, headerUpgrade)
	w.Header().Set(headerSecWebSocketAccept, acceptKey)
	w.WriteHeader(http.StatusSwitchingProtocols)

	hijacker, ok := w.(http.Hijacker)
	if !ok {
		log.Error("web server does not support hijacking", "error", http.StatusText(http.StatusInternalServerError))
		return
	}

	conn, bufRW, err := hijacker.Hijack()
	if err != nil {
		log.Error("failed to hijack connection", "error", err)
		return
	}

	defer conn.Close()

	log.Info("WebSocket connection established")

	if err = that.handleMessages(ctx, bufRW); err != nil {
		log.Error("error handling messages", "error", err)
	}
}

// handleMessages - processes messages from the client.
func (that *Server) handleMessages(ctx context.Context, bufRW *bufio.ReadWriter) error {
	log := that.logger.With("method", "HandleMessages")

	for {
		reqBody, err := that.readRequest(bufRW)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Info("Client closed the connection")
				return io.EOF
			}

			log.Error("Error reading message", "error", err)
			return err
		}

		var message Message
		if err = json.Unmarshal(reqBody, &message); err != nil {
			log.Error("failed to unmarshal message", "error", err)
			continue
		}

		handler, ok := that.messageHandlers[message.Action]
		if !ok {
			log.Error("action handler not found")

			err = that.sendErrorResponse(bufRW, message.Action, "action handler not found")
			if err != nil {
				log.Error("failed to send message", "error", err)
			}

			continue
		}

		if err = handler(ctx, &message, bufRW); err != nil {
			log.Error("invalid handle message", "error", err)

			continue
		}
	}
}

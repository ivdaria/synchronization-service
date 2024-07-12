package gateway

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"synchronizationService/internal/config"
	"synchronizationService/internal/convert"
	"synchronizationService/internal/entity"
	er "synchronizationService/internal/errors"
	"synchronizationService/pkg/gateway/model"
)

type deleteClientUseCase interface {
	DeleteClient(ctx context.Context, id int64) error
}

type updateAlgorithmStatusUseCase interface {
	UpdateAlgorithmStatus(ctx context.Context, algorithmStatus *entity.AlgorithmStatus) error
}

type createClientUseCase interface {
	CreateClient(ctx context.Context, client *entity.Client) (int64, error)
}

type updateClientUseCase interface {
	UpdateClient(ctx context.Context, client *entity.Client) error
}

// AppServer веб-сервер
type AppServer struct {
	server                       *http.Server
	createClientUseCase          createClientUseCase
	updateClientUseCase          updateClientUseCase
	updateAlgorithmStatusUseCase updateAlgorithmStatusUseCase
	deleteClientUseCase          deleteClientUseCase
}

// NewAppServer конструктор для AppServer
func NewAppServer(
	cfg *config.Config,
	createClientUseCase createClientUseCase,
	updateClientUseCase updateClientUseCase,
	updateAlgorithmStatus updateAlgorithmStatusUseCase,
	deleteClientUseCase deleteClientUseCase,
) *AppServer {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    cfg.HTTP.ListenAddr,
		Handler: mux,
	}
	appServer := &AppServer{
		server:                       server,
		createClientUseCase:          createClientUseCase,
		updateClientUseCase:          updateClientUseCase,
		updateAlgorithmStatusUseCase: updateAlgorithmStatus,
		deleteClientUseCase:          deleteClientUseCase,
	}

	mux.HandleFunc("POST /clients", appServer.AddClient)
	mux.HandleFunc("POST /clients/{id}/edit", appServer.UpdateClient)
	mux.HandleFunc("POST /clients/{id}/algorithmstatus", appServer.UpdateAlgorithmStatus)
	mux.HandleFunc("DELETE /clients/{id}", appServer.DeleteClient)

	return appServer
}

// Run метод веб-сервера для запуска обслуживания запросов
func (s *AppServer) Run() error {
	if err := s.server.ListenAndServe(); err != nil {
		return fmt.Errorf("run server: %w", err)
	}

	return nil
}

// Shutdown завершение работы сервера
func (s *AppServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

// Ниже приведены хэндлеры, вызывающие бизнес логику

// AddClient хэндлер-обработчик для добавления нового клиента и статуса его алгоритмов
func (s *AppServer) AddClient(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	decoder := json.NewDecoder(r.Body)

	mdl := model.AddClientRequestBody{}

	if err := decoder.Decode(&mdl); err != nil {
		slog.ErrorContext(
			ctx,
			"add client request body error",
			slog.String("error", fmt.Errorf("decode body to model: %w", err).Error()),
		)

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	clientEntity := convert.ClientFromAddClientRequestBody(&mdl)

	id, err := s.createClientUseCase.CreateClient(ctx, clientEntity)
	if err != nil {
		slog.ErrorContext(
			ctx,
			"add client",
			slog.String("error", fmt.Errorf("create client: %w", err).Error()),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseMdl := model.AddClientResponseBody{
		ID: id,
	}
	responseMdlBytes, err := json.Marshal(responseMdl)
	if err != nil {
		slog.ErrorContext(
			ctx,
			"add client",
			slog.String("error", fmt.Errorf("marshall response: %w", err).Error()),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(responseMdlBytes); err != nil {
		slog.ErrorContext(
			ctx,
			"add client",
			slog.String("error", fmt.Errorf("write response: %w", err).Error()),
		)
		return
	}
}

// UpdateClient хэндлер-обработчик для обновления клиента (алгоритмы не обновляются)
func (s *AppServer) UpdateClient(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idString := r.PathValue("id")
	decoder := json.NewDecoder(r.Body)
	mdl := model.UpdateClientRequestBody{}
	if err := decoder.Decode(&mdl); err != nil {
		slog.ErrorContext(
			ctx,
			"update client",
			slog.String("error", fmt.Errorf("decode body to model: %w", err).Error()),
		)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		slog.ErrorContext(
			ctx,
			"update client",
			slog.String("error", fmt.Errorf("parse client id: %w", err).Error()),
		)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	clt := convert.ClientFromUpdateClientRequestBody(id, &mdl)
	err = s.updateClientUseCase.UpdateClient(ctx, clt)
	if err != nil {
		slog.ErrorContext(
			ctx,
			"update client",
			slog.String("error", fmt.Errorf("update client: %w", err).Error()),
		)
		if errors.Is(err, er.ErrNoRowsAffected) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UpdateAlgorithmStatus хэндлер-обработчик для обновления статусов алгоритмов
// (сам клиент не обновляется)
func (s *AppServer) UpdateAlgorithmStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idString := r.PathValue("id")
	decoder := json.NewDecoder(r.Body)
	mdl := model.UpdateStatusRequestBody{}
	if err := decoder.Decode(&mdl); err != nil {
		slog.ErrorContext(
			ctx,
			"update status",
			slog.String("error", fmt.Errorf("decode body to model: %w", err).Error()),
		)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	clientID, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		slog.ErrorContext(
			ctx,
			"update status",
			slog.String("error", fmt.Errorf("parse client id: %w", err).Error()),
		)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	algStatus := convert.StatusFromUpdateStatusRequestBody(clientID, &mdl)

	err = s.updateAlgorithmStatusUseCase.UpdateAlgorithmStatus(ctx, algStatus)
	if err != nil {
		slog.ErrorContext(
			ctx,
			"update client",
			slog.String("error", fmt.Errorf("update client: %w", err).Error()),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteClient хэндлер обработчик для удаления клиента и его статусов
func (s *AppServer) DeleteClient(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idString := r.PathValue("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		slog.ErrorContext(
			ctx,
			"delete client",
			slog.String("error", fmt.Errorf("parse client id: %w", err).Error()),
		)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	err = s.deleteClientUseCase.DeleteClient(ctx, id)
	if err != nil {
		slog.ErrorContext(
			ctx,
			"delete client",
			slog.String("error", fmt.Errorf("delete client: %w", err).Error()),
		)
		if errors.Is(err, er.ErrNoRowsAffected) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

package main

import (
	"context"
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"github.com/jackc/pgx/v5"
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
	"os/signal"
	"synchronizationService/internal/config"
	"synchronizationService/internal/gateway"
	algorithmstatus "synchronizationService/internal/repository/algorithm_status"
	"synchronizationService/internal/repository/client"
	"synchronizationService/internal/service/deployer"
	createclient "synchronizationService/internal/usecase/client/create_client"
	deleteclient "synchronizationService/internal/usecase/client/delete_client"
	updateclient "synchronizationService/internal/usecase/client/update_client"
	updateclientalgorithms "synchronizationService/internal/usecase/client/update_client_algorithms"
	deployworker "synchronizationService/internal/worker/deploy_worker"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// gracefull shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChan
		cancel()
	}()

	// чтение конфига
	yamlFile, err := os.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}
	var cfg config.Config

	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		panic(err)
	}

	// подключение к БД
	dbDSN := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.DBConfig.User,
		cfg.DBConfig.Password,
		cfg.DBConfig.Host,
		cfg.DBConfig.Port,
		cfg.DBConfig.Database,
	)

	conn, err := pgx.Connect(ctx, dbDSN)
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())

	// создание основных репозиториев и юзкейсов
	clientsRepo := client.NewRepo(conn)
	algStatusRepo := algorithmstatus.NewRepo(conn)
	createClientUseCase := createclient.NewUseCase(algStatusRepo, clientsRepo, conn)
	updateClientUseCase := updateclient.NewUseCase(clientsRepo)
	updateAlggorithmStatusUseCase := updateclientalgorithms.NewUseCase(algStatusRepo)
	deleteClientUseCase := deleteclient.NewUseCase(algStatusRepo, clientsRepo, conn)
	appServer := gateway.NewAppServer(
		&cfg,
		createClientUseCase,
		updateClientUseCase,
		updateAlggorithmStatusUseCase,
		deleteClientUseCase,
	)

	// запуск веб-сервера
	go func() {
		if err := appServer.Run(); err != nil {
			slog.ErrorContext(ctx, "app server is closed with error", slog.String("err", err.Error()))
			panic(err)
		}
	}()

	// Создание и запуск планировщика
	s, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	deployerService := deployer.NewService()
	worker := deployworker.NewWorker(deployerService, algStatusRepo)

	// создание джобы на запуск воркера по крону(частоту можно задать в конфиге)
	_, err = s.NewJob(
		gocron.CronJob(cfg.Deploy.CronString, true),
		gocron.NewTask(
			func() {
				worker.Work(ctx)
			},
		),
	)
	if err != nil {
		panic(err)
	}

	// первое выполнение ворекера делаем мгновенно
	go worker.Work(ctx)

	// запуск планировщика
	s.Start()

	slog.Info("server is running and ready to serve")

	<-ctx.Done()
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	if err := appServer.Shutdown(ctx); err != nil {
		slog.ErrorContext(ctx, "app server shutdown", slog.String("err", err.Error()))
	}
}

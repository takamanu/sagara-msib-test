package boot

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"sagara-msib-test/internal/deliveries"
	"sagara-msib-test/internal/infrastructures"
	"sagara-msib-test/internal/repositories"
	"sagara-msib-test/internal/services"
	jaegerLog "sagara-msib-test/pkg/log"
	"sagara-msib-test/pkg/tracing"
)

func BajuInventoryHTTP() (err error) {
	var (
		logger *zap.Logger
		ctx    context.Context
	)

	ctx = context.Background()
	db, err := infrastructures.NewDatabase(ctx)
	if err != nil {
		return err
	}

	logger, _ = zap.NewDevelopment(
		zap.AddStacktrace(zapcore.FatalLevel),
		zap.AddCallerSkip(1),
	)
	zapLogger := logger.With(zap.String("service", "inventory-api"))
	zlogger := jaegerLog.NewFactory(zapLogger)

	tracer, closer := tracing.Init("inventory-api", zlogger)
	defer closer.Close()

	bajuRepos := repositories.NewInventoryBajuRepository(db)
	bajuServices := services.NewInventoryBajuService(bajuRepos)
	bajuHandler := deliveries.NewBajuHandler(bajuServices, tracer, zlogger)

	s := infrastructures.Server{
		BajuDelivery: bajuHandler,
	}

	if err := s.Serve(":8080"); !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return err
}

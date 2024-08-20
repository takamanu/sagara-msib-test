package deliveries

import (
	"context"
	"encoding/json"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go.uber.org/zap"
	"net/http"
	"sagara-msib-test/internal/entities"
	"sagara-msib-test/internal/services"
	jaegerLog "sagara-msib-test/pkg/log"
)

type BajuHandler struct {
	bajuServices services.InventoryBajuServices
	tracer       opentracing.Tracer
	logger       jaegerLog.Factory
}

func NewBajuHandler(bajuServices services.InventoryBajuServices, tracer opentracing.Tracer, logger jaegerLog.Factory) (handler *BajuHandler) {
	handler = &BajuHandler{
		bajuServices: bajuServices,
		tracer:       tracer,
		logger:       logger,
	}

	return handler
}

func (bh *BajuHandler) HandleClient(w http.ResponseWriter, r *http.Request) {
	var (
		baju      entities.Baju
		ctx       context.Context
		err       error
		errorCode int
	)

	spanCtx, _ := bh.tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	span := bh.tracer.StartSpan("HandleClient", ext.RPCServerOption(spanCtx))
	defer span.Finish()

	ctx = opentracing.ContextWithSpan(r.Context(), span)
	bh.logger.For(ctx).Info("HTTP request received", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	if err == nil {
		urlQueryLength := len(r.URL.Query())

		switch r.Method {
		case http.MethodPost:
			switch urlQueryLength {
			default:
				err = json.NewDecoder(r.Body).Decode(&baju)
				if err != nil {
					errorCode = http.StatusBadRequest
				}

				bh.logger.For(ctx).Info("Running service", zap.String("service", "Create New Baju"))
				err = bh.bajuServices.CreateBaju(&baju)
				if err != nil {
					errorCode = http.StatusInternalServerError
				}
			}
		}
	}

	if err != nil {
		http.Error(w, err.Error(), errorCode)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(baju)
}

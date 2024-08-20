package boot

import (
	"errors"
	"net/http"
	"sagara-msib-test/internal/deliveries"
	"sagara-msib-test/internal/infrastructures"
	"sagara-msib-test/internal/repositories"
	"sagara-msib-test/internal/services"
)

func BajuInventoryHTTP() (err error) {
	db, err := infrastructures.NewDatabase()
	if err != nil {
		return err
	}

	bajuRepos := repositories.NewInventoryBajuRepository(db)
	bajuServices := services.NewInventoryBajuService(bajuRepos)
	bajuHandler := deliveries.NewBajuHandler(bajuServices)

	s := infrastructures.Server{
		BajuDelivery: bajuHandler,
	}

	if err := s.Serve(":8080"); !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return err
}

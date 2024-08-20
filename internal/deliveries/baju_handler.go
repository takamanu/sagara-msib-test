package deliveries

import (
	"encoding/json"
	"net/http"
	"sagara-msib-test/internal/entities"
	"sagara-msib-test/internal/services"
)

type BajuHandler struct {
	bajuServices services.InventoryBajuServices
}

func NewBajuHandler(bajuServices services.InventoryBajuServices) (handler *BajuHandler) {
	handler = &BajuHandler{
		bajuServices: bajuServices,
	}

	return handler
}

func (bh *BajuHandler) HandleClient(w http.ResponseWriter, r *http.Request) {
	var (
		baju entities.Baju
	)

	err := json.NewDecoder(r.Body).Decode(&baju)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = bh.bajuServices.CreateBaju(&baju)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(baju)
}

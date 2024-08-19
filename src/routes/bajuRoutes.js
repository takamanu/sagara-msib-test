const express = require('express');

function bajuRoutes(bajuController) {
  const router = express.Router();
  
  router.get('/baju', bajuController.getAllBaju.bind(bajuController));
  router.post('/baju', bajuController.addBaju.bind(bajuController));
  router.get('/baju/:id', bajuController.getBajuById.bind(bajuController));
  router.put('/baju/:id', bajuController.updateBaju.bind(bajuController));
  router.delete('/baju/:id', bajuController.deleteBaju.bind(bajuController));
  router.get('/search', bajuController.searchBajuByWarnaUkuran.bind(bajuController));
  router.put('/baju/adjust-stok/:id', bajuController.adjustStok.bind(bajuController));
  
  return router;
}

module.exports = bajuRoutes;
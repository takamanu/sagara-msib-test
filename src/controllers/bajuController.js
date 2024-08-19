const response = require('../utils/response');

class BajuController {
  async getAllBaju(req, res) {
    try {
      const baju = await this.bajuService.getAllBaju();
      response(200, baju, 'Berhasil mendapatkan data semua baju', res);
    } catch (error) {
      response(500, error.message, 'Gagal mendapatkan data semua baju', res);
    }
  }

  async addBaju(req, res) {
    try {
      const baju = await this.bajuService.addBaju(req.body);
      response(201, baju, 'Berhasil menambahkan data baju baru', res);
    } catch (error) {
      response(500, error.message, 'Gagal menambahkan data baju baru', res);
    }
  }

  async getBajuById(req, res) {
    try {
      const { id } = req.params;
      const baju = await this.bajuService.getBajuById(id);
      if (!baju) {
        return response(404, { error: 'Baju tidak ditemukan' }, 'Gagal mendapatkan baju berdasarkan id', res);
      }
      response(200, baju, 'Data baju berdasarkan id', res);
    } catch (error) {
      response(500, error.message, 'Gagal mendapatkan baju berdasarkan id', res);
    }
  }

  async updateBaju(req, res) {
    try {
      const { id } = req.params;
      const baju = await this.bajuService.updateBaju(id, req.body);
      if (!baju) {
        return response(404, { error: 'Baju tidak ditemukan' }, 'Gagal mengupdate data baju', res);
      }
      response(200, baju, 'Berhasil mengupdate data baju', res);
    } catch (error) {
      response(500, error.message, 'Gagal mengupdate data baju', res);
    }
  }

  async deleteBaju(req, res) {
    try {
      const { id } = req.params;
      const sukses = await this.bajuService.deleteBaju(id);
      if (!sukses) {
        return response(404, { error: 'Baju tidak ditemukan' }, res);
      }
      response(204, '', '', res);
    } catch (error) {
      response(500, error.message, 'Gagal menghapus data baju', res);
    }
  }
}

module.exports = BajuController;
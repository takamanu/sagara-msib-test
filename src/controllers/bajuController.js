const response = require('../utils/response');

class BajuController {
  constructor(bajuService) {
    this.bajuService = bajuService;
  }

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

  async searchBajuByWarnaUkuran(req, res) {
    try {
      const { warna, ukuran } = req.body;
      const baju = await this.bajuService.searchBajuByWarnaUkuran(warna, ukuran);
      response(200, baju, 'Data baju berdasarkan warna dan ukuran', res);
    } catch (error) {
      response(500, error.message, 'Gagal mendapatkan data baju berdasarkan warna dan ukuran', res);
    }
  }

  async adjustStok(req, res) {
    try {
      const { id } = req.params;
      const { quantity } = req.body;
      const baju = await this.bajuService.adjustStok(id, quantity);
      if (!baju) {
        return response(404, { error: 'Baju tidak ditemukan' }, 'Gagal adjust stok baju', res);
      }
      response(200, baju, 'Berhasil adjust stok baju', res);
    } catch (error) {
      response(500, error.message, 'Gagal adjust stok baju', res);
    }
  }

  async stokTersedia(req, res) {
    try {
      const baju = await this.bajuService.cekStok('!=');
      response(200, baju, 'Data baju yang stok nya tersedia', res);
    } catch (error) {
      response(500, error.message, 'Gagal menampilkan stok baju yang tersedia', res);
    }
  }

  async stokHabis(req, res) {
    try {
      const baju = await this.bajuService.cekStok('=');
      response(200, baju, 'Data baju yang stok nya habis', res);
    } catch (error) {
      response(500, error.message, 'Gagal menampilkan stok baju yang habis', res);
    }
  }

  async stokTipis(req, res) {
    try {
      const baju = await this.bajuService.cekStok('<', 5);
      response(200, baju, 'Data baju yang stok nya kurang dari 5', res);
    } catch (error) {
      response(500, error.message, 'Gagal menampilkan stok baju yang kurang dari 5', res);
    }
  }
}

module.exports = BajuController;
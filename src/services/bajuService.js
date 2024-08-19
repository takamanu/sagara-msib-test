class BajuService {
  constructor(bajuRepository) {
    this.bajuRepository = bajuRepository;
  }

  async getAllBaju() {
    return this.bajuRepository.findAll();
  }

  async addBaju(dataBaju) {
    return this.bajuRepository.create(dataBaju);
  }

  async getBajuById(id) {
    return this.bajuRepository.findById(id);
  }

  async updateBaju(id, dataBaju) {
    return this.bajuRepository.update(id, dataBaju);
  }

  async deleteBaju(id) {
    return this.bajuRepository.delete(id);
  }

  async searchBajuByWarnaUkuran(warna, ukuran) {
    return this.bajuRepository.findByWarnaDanUkuran(warna, ukuran);
  }

  async adjustStok(id, quantity) {
    return this.bajuRepository.adjustStok(id, quantity);
  }

  async cekStok(operator, param = 0) {
    return this.bajuRepository.cekStok(operator, param);
  }
}

module.exports = BajuService;
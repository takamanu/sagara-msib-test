class BajuService {
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
}

module.exports = BajuService;
require('dotenv').config();
const express = require('express');
const { Pool } = require('pg');
const bajuRoutes = require('./routes/bajuRoutes');
const BajuRepository = require('./repositories/BajuRepository');
const BajuService = require('./services/bajuService');
const BajuController = require('./controllers/bajuController');
const config = require('./config/database');

const app = express();
app.use(express.json());

const pool = new Pool(config);

async function checkConnection() {
  try {
    const client = await pool.connect();
    console.log('Koneksi berhasil!');
    client.release();
  } catch (err) {
    console.error('Koneksi gagal:', err);
  }
}

checkConnection();

const bajuRepository = new BajuRepository(pool);
const bajuService = new BajuService(bajuRepository);
const bajuController = new BajuController(bajuService);

app.use('/api', bajuRoutes(bajuController));

app.listen(process.env.PORT, () => {
  console.log(`server running on http://localhost:${process.env.PORT}`);
})

module.exports = app;
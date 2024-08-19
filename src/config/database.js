require('dotenv').config();

const config = {
  database: process.env.PG_DATABASE_NAME,
  user: process.env.PG_DATABASE_USER,
  password: process.env.PG_DATABASE_PASSWORD,
  port: 5432,
};

module.exports = config;
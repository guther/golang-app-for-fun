CREATE DATABASE IF NOT EXISTS db_hostgator_challenge;
USE db_hostgator_challenge;
CREATE TABLE IF NOT EXISTS tb_breeds_cache (
  query VARCHAR(500) NOT NULL UNIQUE,
  data MEDIUMTEXT NOT NULL,
  time_creation TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
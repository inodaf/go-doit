-- SQLite schema for the Database;

CREATE TABLE tasks IF NOT EXISTS(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title TEXT NOT NULL,
  description TEXT,

  created_at DATETIME NOT NULL,
  updated_at DATETIME,
  done_at DATETIME,
)

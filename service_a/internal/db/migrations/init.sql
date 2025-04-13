-- Active: 1744549994418@@localhost@443@postgres

CREATE TABLE IF NOT EXISTS outbox (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP DEFAULT NOW(),
  sent_at TIMESTAMP,
  value INT NOT NULL
);

CREATE INDEX idx_outbox_sent_at ON outbox(sent_at);
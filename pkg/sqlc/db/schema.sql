CREATE TABLE users (
  id   BIGSERIAL PRIMARY KEY,
  name text,
  phone_number  text unique,
  otp text,
  otp_expiry_time timestamp
);
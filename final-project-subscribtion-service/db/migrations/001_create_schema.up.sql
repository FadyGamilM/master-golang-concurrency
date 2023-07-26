CREATE OR REPLACE FUNCTION update_timestamp_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = now();
   RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TABLE IF NOT EXISTS plans (
    id BIGSERIAL PRIMARY KEY,
    plan_name VARCHAR(200) NOT NULL,
    plan_amount BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT now() NOT NULL, 
    updated_at TIMESTAMP DEFAULT now() NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR NOT NULL,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    is_active BOOLEAN DEFAULT(TRUE),
    is_admin BOOLEAN DEFAULT(FALSE),
        created_at TIMESTAMP DEFAULT now() NOT NULL, 
    updated_at TIMESTAMP DEFAULT now() NOT NULL
);


CREATE TABLE IF NOT EXISTS users_plans (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    plan_id BIGINT NOT NULL REFERENCES plans(id),
    created_at TIMESTAMP DEFAULT now() NOT NULL, 
    updated_at TIMESTAMP DEFAULT now() NOT NULL
);


CREATE TRIGGER update_timestamp_plans
BEFORE UPDATE ON plans
FOR EACH ROW
EXECUTE PROCEDURE update_timestamp_column();

CREATE TRIGGER update_timestamp_users
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE update_timestamp_column();

CREATE TRIGGER update_timestamp_users_plans
BEFORE UPDATE ON users_plans
FOR EACH ROW
EXECUTE PROCEDURE update_timestamp_column();

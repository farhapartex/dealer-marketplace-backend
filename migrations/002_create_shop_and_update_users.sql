-- Create shop table
CREATE TABLE IF NOT EXISTS shop (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    contact_number VARCHAR(20),
    logo TEXT,
    address TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create index on shop name for faster lookups
CREATE INDEX IF NOT EXISTS idx_shop_name ON shop(name);

-- Add new columns to users table
ALTER TABLE users ADD COLUMN IF NOT EXISTS is_onboard_complete BOOLEAN DEFAULT false;
ALTER TABLE users ADD COLUMN IF NOT EXISTS user_type VARCHAR(50);

-- Create index on user_type for faster lookups
CREATE INDEX IF NOT EXISTS idx_users_user_type ON users(user_type);

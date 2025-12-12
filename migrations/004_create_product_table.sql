-- Create product table
CREATE TABLE IF NOT EXISTS product (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    shop_id UUID NOT NULL,

    -- Basic Information
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,
    short_description VARCHAR(500),
    sku VARCHAR(100) UNIQUE,
    barcode VARCHAR(100),

    -- Pricing
    price DECIMAL(10, 2) NOT NULL,
    cost_price DECIMAL(10, 2),
    compare_at_price DECIMAL(10, 2),
    tax_rate DECIMAL(5, 2) DEFAULT 0,

    -- Inventory Management
    stock_quantity INTEGER DEFAULT 0,
    low_stock_threshold INTEGER DEFAULT 10,
    track_inventory BOOLEAN DEFAULT true,
    allow_backorder BOOLEAN DEFAULT false,

    -- Categorization
    category VARCHAR(100),
    brand VARCHAR(100),
    tags TEXT,

    -- Product Specifications
    weight DECIMAL(8, 2),
    weight_unit VARCHAR(10) DEFAULT 'kg',
    length DECIMAL(8, 2),
    width DECIMAL(8, 2),
    height DECIMAL(8, 2),
    dimension_unit VARCHAR(10) DEFAULT 'cm',

    -- Media
    primary_image TEXT,
    images TEXT,
    video_url TEXT,

    -- Status & Visibility
    status VARCHAR(20) DEFAULT 'draft',
    is_active BOOLEAN DEFAULT true,
    is_featured BOOLEAN DEFAULT false,
    visibility VARCHAR(20) DEFAULT 'public',

    -- SEO
    meta_title VARCHAR(255),
    meta_description TEXT,
    meta_keywords TEXT,

    -- Additional Information
    manufacturer VARCHAR(100),
    warranty_period INTEGER,
    warranty_unit VARCHAR(20) DEFAULT 'months',
    country_of_origin VARCHAR(100),

    -- Timestamps
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    published_at TIMESTAMP,

    -- Foreign Key Constraint
    CONSTRAINT fk_product_shop FOREIGN KEY (shop_id) REFERENCES shop(id) ON DELETE CASCADE
);

-- Create indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_product_shop_id ON product(shop_id);
CREATE INDEX IF NOT EXISTS idx_product_slug ON product(slug);
CREATE INDEX IF NOT EXISTS idx_product_sku ON product(sku);
CREATE INDEX IF NOT EXISTS idx_product_category ON product(category);
CREATE INDEX IF NOT EXISTS idx_product_brand ON product(brand);
CREATE INDEX IF NOT EXISTS idx_product_status ON product(status);
CREATE INDEX IF NOT EXISTS idx_product_is_active ON product(is_active);
CREATE INDEX IF NOT EXISTS idx_product_is_featured ON product(is_featured);
CREATE INDEX IF NOT EXISTS idx_product_name ON product(name);

-- Create composite index for shop products listing
CREATE INDEX IF NOT EXISTS idx_product_shop_status ON product(shop_id, status, is_active);

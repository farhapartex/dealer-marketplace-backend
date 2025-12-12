package dtos

import (
	"time"

	"github.com/google/uuid"
)

// CreateProductRequest represents the payload for creating a product
type CreateProductRequest struct {
	// Basic Information
	Name             string  `json:"name" binding:"required"`
	Description      *string `json:"description"`
	ShortDescription *string `json:"short_description"`
	SKU              *string `json:"sku"`
	Barcode          *string `json:"barcode"`

	// Pricing
	Price          float64  `json:"price" binding:"required,gt=0"`
	CostPrice      *float64 `json:"cost_price" binding:"omitempty,gt=0"`
	CompareAtPrice *float64 `json:"compare_at_price" binding:"omitempty,gt=0"`
	TaxRate        *float64 `json:"tax_rate" binding:"omitempty,gte=0,lte=100"`

	// Inventory Management
	StockQuantity     *int  `json:"stock_quantity" binding:"omitempty,gte=0"`
	LowStockThreshold *int  `json:"low_stock_threshold" binding:"omitempty,gte=0"`
	TrackInventory    *bool `json:"track_inventory"`
	AllowBackorder    *bool `json:"allow_backorder"`

	// Categorization
	Category *string `json:"category"`
	Brand    *string `json:"brand"`
	Tags     *string `json:"tags"`

	// Product Specifications
	Weight        *float64 `json:"weight" binding:"omitempty,gt=0"`
	WeightUnit    *string  `json:"weight_unit" binding:"omitempty,oneof=kg lbs g oz"`
	Length        *float64 `json:"length" binding:"omitempty,gt=0"`
	Width         *float64 `json:"width" binding:"omitempty,gt=0"`
	Height        *float64 `json:"height" binding:"omitempty,gt=0"`
	DimensionUnit *string  `json:"dimension_unit" binding:"omitempty,oneof=cm inch mm m"`

	// Media
	PrimaryImage *string `json:"primary_image" binding:"omitempty,url"`
	Images       *string `json:"images"`
	VideoURL     *string `json:"video_url" binding:"omitempty,url"`

	// Status & Visibility
	Status     *string `json:"status" binding:"omitempty,oneof=draft published archived"`
	IsActive   *bool   `json:"is_active"`
	IsFeatured *bool   `json:"is_featured"`
	Visibility *string `json:"visibility" binding:"omitempty,oneof=public private hidden"`

	// SEO
	MetaTitle       *string `json:"meta_title"`
	MetaDescription *string `json:"meta_description"`
	MetaKeywords    *string `json:"meta_keywords"`

	// Additional Information
	Manufacturer    *string `json:"manufacturer"`
	WarrantyPeriod  *int    `json:"warranty_period" binding:"omitempty,gte=0"`
	WarrantyUnit    *string `json:"warranty_unit" binding:"omitempty,oneof=days months years"`
	CountryOfOrigin *string `json:"country_of_origin"`
}

// CreateProductResponse represents the response after creating a product
type CreateProductResponse struct {
	ID              uuid.UUID  `json:"id"`
	ShopID          uuid.UUID  `json:"shop_id"`
	Name            string     `json:"name"`
	Slug            string     `json:"slug"`
	Description     *string    `json:"description"`
	SKU             *string    `json:"sku"`
	Price           float64    `json:"price"`
	CostPrice       *float64   `json:"cost_price"`
	CompareAtPrice  *float64   `json:"compare_at_price"`
	TaxRate         float64    `json:"tax_rate"`
	StockQuantity   int        `json:"stock_quantity"`
	TrackInventory  bool       `json:"track_inventory"`
	AllowBackorder  bool       `json:"allow_backorder"`
	Category        *string    `json:"category"`
	Brand           *string    `json:"brand"`
	Tags            *string    `json:"tags"`
	PrimaryImage    *string    `json:"primary_image"`
	Status          string     `json:"status"`
	IsActive        bool       `json:"is_active"`
	IsFeatured      bool       `json:"is_featured"`
	Visibility      string     `json:"visibility"`
	Manufacturer    *string    `json:"manufacturer"`
	CountryOfOrigin *string    `json:"country_of_origin"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	PublishedAt     *time.Time `json:"published_at"`
	Message         string     `json:"message"`
}

// UpdateProductRequest represents the payload for updating a product
type UpdateProductRequest struct {
	// Basic Information
	Name             *string `json:"name"`
	Description      *string `json:"description"`
	ShortDescription *string `json:"short_description"`
	SKU              *string `json:"sku"`
	Barcode          *string `json:"barcode"`

	// Pricing
	Price          *float64 `json:"price" binding:"omitempty,gt=0"`
	CostPrice      *float64 `json:"cost_price" binding:"omitempty,gt=0"`
	CompareAtPrice *float64 `json:"compare_at_price" binding:"omitempty,gt=0"`
	TaxRate        *float64 `json:"tax_rate" binding:"omitempty,gte=0,lte=100"`

	// Inventory Management
	StockQuantity     *int  `json:"stock_quantity" binding:"omitempty,gte=0"`
	LowStockThreshold *int  `json:"low_stock_threshold" binding:"omitempty,gte=0"`
	TrackInventory    *bool `json:"track_inventory"`
	AllowBackorder    *bool `json:"allow_backorder"`

	// Categorization
	Category *string `json:"category"`
	Brand    *string `json:"brand"`
	Tags     *string `json:"tags"`

	// Product Specifications
	Weight        *float64 `json:"weight" binding:"omitempty,gt=0"`
	WeightUnit    *string  `json:"weight_unit" binding:"omitempty,oneof=kg lbs g oz"`
	Length        *float64 `json:"length" binding:"omitempty,gt=0"`
	Width         *float64 `json:"width" binding:"omitempty,gt=0"`
	Height        *float64 `json:"height" binding:"omitempty,gt=0"`
	DimensionUnit *string  `json:"dimension_unit" binding:"omitempty,oneof=cm inch mm m"`

	// Media
	PrimaryImage *string `json:"primary_image" binding:"omitempty,url"`
	Images       *string `json:"images"`
	VideoURL     *string `json:"video_url" binding:"omitempty,url"`

	// Status & Visibility
	Status     *string `json:"status" binding:"omitempty,oneof=draft published archived"`
	IsActive   *bool   `json:"is_active"`
	IsFeatured *bool   `json:"is_featured"`
	Visibility *string `json:"visibility" binding:"omitempty,oneof=public private hidden"`

	// SEO
	MetaTitle       *string `json:"meta_title"`
	MetaDescription *string `json:"meta_description"`
	MetaKeywords    *string `json:"meta_keywords"`

	// Additional Information
	Manufacturer    *string `json:"manufacturer"`
	WarrantyPeriod  *int    `json:"warranty_period" binding:"omitempty,gte=0"`
	WarrantyUnit    *string `json:"warranty_unit" binding:"omitempty,oneof=days months years"`
	CountryOfOrigin *string `json:"country_of_origin"`
}

// ProductListResponse represents a single product in list view
type ProductListResponse struct {
	ID             uuid.UUID `json:"id"`
	ShopID         uuid.UUID `json:"shop_id"`
	Name           string    `json:"name"`
	Slug           string    `json:"slug"`
	Price          float64   `json:"price"`
	CompareAtPrice *float64  `json:"compare_at_price"`
	StockQuantity  int       `json:"stock_quantity"`
	PrimaryImage   *string   `json:"primary_image"`
	Category       *string   `json:"category"`
	Brand          *string   `json:"brand"`
	Status         string    `json:"status"`
	IsActive       bool      `json:"is_active"`
	IsFeatured     bool      `json:"is_featured"`
	CreatedAt      time.Time `json:"created_at"`
}

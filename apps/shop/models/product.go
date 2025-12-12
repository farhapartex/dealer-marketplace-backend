package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID     uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ShopID uuid.UUID `json:"shop_id" gorm:"type:uuid;not null"`

	// Basic Information
	Name             string  `json:"name" gorm:"not null"`
	Slug             string  `json:"slug" gorm:"unique;not null"`
	Description      *string `json:"description"`
	ShortDescription *string `json:"short_description"`
	SKU              *string `json:"sku" gorm:"unique"`
	Barcode          *string `json:"barcode"`

	// Pricing
	Price          float64  `json:"price" gorm:"type:decimal(10,2);not null"`
	CostPrice      *float64 `json:"cost_price" gorm:"type:decimal(10,2)"`
	CompareAtPrice *float64 `json:"compare_at_price" gorm:"type:decimal(10,2)"`
	TaxRate        float64  `json:"tax_rate" gorm:"type:decimal(5,2);default:0"`

	// Inventory Management
	StockQuantity     int  `json:"stock_quantity" gorm:"default:0"`
	LowStockThreshold int  `json:"low_stock_threshold" gorm:"default:10"`
	TrackInventory    bool `json:"track_inventory" gorm:"default:true"`
	AllowBackorder    bool `json:"allow_backorder" gorm:"default:false"`

	// Categorization
	Category *string `json:"category"`
	Brand    *string `json:"brand"`
	Tags     *string `json:"tags"`

	// Product Specifications
	Weight        *float64 `json:"weight" gorm:"type:decimal(8,2)"`
	WeightUnit    string   `json:"weight_unit" gorm:"default:'kg'"`
	Length        *float64 `json:"length" gorm:"type:decimal(8,2)"`
	Width         *float64 `json:"width" gorm:"type:decimal(8,2)"`
	Height        *float64 `json:"height" gorm:"type:decimal(8,2)"`
	DimensionUnit string   `json:"dimension_unit" gorm:"default:'cm'"`

	// Media
	PrimaryImage *string `json:"primary_image"`
	Images       *string `json:"images"`
	VideoURL     *string `json:"video_url"`

	// Status & Visibility
	Status     string `json:"status" gorm:"default:'draft'"`
	IsActive   bool   `json:"is_active" gorm:"default:true"`
	IsFeatured bool   `json:"is_featured" gorm:"default:false"`
	Visibility string `json:"visibility" gorm:"default:'public'"`

	// SEO
	MetaTitle       *string `json:"meta_title"`
	MetaDescription *string `json:"meta_description"`
	MetaKeywords    *string `json:"meta_keywords"`

	// Additional Information
	Manufacturer    *string `json:"manufacturer"`
	WarrantyPeriod  *int    `json:"warranty_period"`
	WarrantyUnit    string  `json:"warranty_unit" gorm:"default:'months'"`
	CountryOfOrigin *string `json:"country_of_origin"`

	// Timestamps
	CreatedAt   time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	PublishedAt *time.Time `json:"published_at"`

	// Relationship
	Shop *Shop `json:"shop,omitempty" gorm:"foreignKey:ShopID;constraint:OnDelete:CASCADE"`
}

func (Product) TableName() string {
	return "product"
}

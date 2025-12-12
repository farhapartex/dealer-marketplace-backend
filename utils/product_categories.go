package utils

// Category represents a product category with its subcategories
type Category struct {
	Name          string   `json:"name"`
	Slug          string   `json:"slug"`
	Subcategories []string `json:"subcategories"`
}

// ProductCategories contains all available categories and subcategories
var ProductCategories = []Category{
	{
		Name: "Electronics",
		Slug: "electronics",
		Subcategories: []string{
			"Mobile Phones",
			"Phone Accessories",
			"Laptops",
			"Desktops",
			"Tablets",
			"Cameras",
			"Camera Accessories",
			"Headphones",
			"Earphones",
			"Speakers",
			"Smart Watches",
			"Televisions",
			"Gaming Consoles",
			"Gaming Accessories",
			"Power Banks",
			"Chargers & Cables",
			"Smart Home Devices",
			"Computer Accessories",
			"Networking Devices",
			"Printers & Scanners",
			"Storage Devices",
			"Monitors",
		},
	},
	{
		Name: "Men's Fashion",
		Slug: "mens-fashion",
		Subcategories: []string{
			"T-Shirts",
			"Shirts",
			"Jeans",
			"Trousers",
			"Shorts",
			"Jackets",
			"Sweaters",
			"Suits & Blazers",
			"Ethnic Wear",
			"Innerwear",
			"Sleepwear",
			"Sportswear",
			"Casual Shoes",
			"Formal Shoes",
			"Sports Shoes",
			"Sandals & Slippers",
			"Watches",
			"Belts",
			"Wallets",
			"Sunglasses",
			"Bags & Backpacks",
			"Caps & Hats",
		},
	},
	{
		Name: "Women's Fashion",
		Slug: "womens-fashion",
		Subcategories: []string{
			"Dresses",
			"Tops",
			"T-Shirts",
			"Jeans",
			"Trousers",
			"Skirts",
			"Ethnic Wear",
			"Sarees",
			"Kurtas & Kurtis",
			"Leggings",
			"Jackets",
			"Sweaters",
			"Innerwear",
			"Sleepwear",
			"Sportswear",
			"Flats",
			"Heels",
			"Sandals",
			"Sports Shoes",
			"Handbags",
			"Clutches",
			"Watches",
			"Jewelry",
			"Sunglasses",
			"Scarves",
		},
	},
	{
		Name: "Kids Fashion",
		Slug: "kids-fashion",
		Subcategories: []string{
			"Boys Clothing",
			"Girls Clothing",
			"Baby Boys Clothing",
			"Baby Girls Clothing",
			"Kids Shoes",
			"Kids Accessories",
			"School Uniforms",
			"Kids Bags",
			"Kids Watches",
			"Toys",
			"Infant Care",
		},
	},
	{
		Name: "Home & Living",
		Slug: "home-living",
		Subcategories: []string{
			"Sofas",
			"Beds",
			"Dining Tables",
			"Chairs",
			"Wardrobes",
			"Mattresses",
			"Cushions & Pillows",
			"Bed Sheets",
			"Curtains",
			"Carpets & Rugs",
			"Wall Decor",
			"Lighting",
			"Kitchen Storage",
			"Cookware",
			"Dinnerware",
			"Kitchen Appliances",
			"Bath Towels",
			"Bath Accessories",
			"Home Decor",
			"Clocks",
			"Mirrors",
			"Plants & Planters",
		},
	},
	{
		Name: "Beauty & Personal Care",
		Slug: "beauty-personal-care",
		Subcategories: []string{
			"Skincare",
			"Face Wash & Cleansers",
			"Moisturizers",
			"Sunscreen",
			"Face Masks",
			"Makeup",
			"Lipstick",
			"Foundation",
			"Kajal & Eyeliner",
			"Nail Polish",
			"Shampoo",
			"Conditioner",
			"Hair Oil",
			"Hair Styling",
			"Hair Color",
			"Perfumes",
			"Deodorants",
			"Body Lotion",
			"Shower Gel",
			"Men's Grooming",
			"Shaving",
			"Beard Care",
		},
	},
	{
		Name: "Health & Wellness",
		Slug: "health-wellness",
		Subcategories: []string{
			"Vitamins & Supplements",
			"Protein Supplements",
			"Weight Management",
			"Ayurvedic Products",
			"Health Monitors",
			"First Aid",
			"Medical Devices",
			"Sexual Wellness",
			"Pain Relief",
			"Cold & Cough",
			"Fitness Equipment",
			"Yoga Mats",
			"Dumbbells",
			"Treadmills",
			"Exercise Bikes",
		},
	},
	{
		Name: "Sports & Fitness",
		Slug: "sports-fitness",
		Subcategories: []string{
			"Cricket",
			"Football",
			"Badminton",
			"Tennis",
			"Gym Equipment",
			"Cycling",
			"Swimming",
			"Running",
			"Outdoor Sports",
			"Sports Shoes",
			"Sports Clothing",
			"Sports Accessories",
			"Fitness Trackers",
			"Yoga Equipment",
			"Boxing & MMA",
		},
	},
	{
		Name: "Books & Media",
		Slug: "books-media",
		Subcategories: []string{
			"Fiction Books",
			"Non-Fiction Books",
			"Self-Help Books",
			"Business Books",
			"Children's Books",
			"Comics & Graphic Novels",
			"Textbooks",
			"E-Books",
			"Magazines",
			"Stationery",
			"Office Supplies",
			"Art Supplies",
			"Educational Toys",
		},
	},
	{
		Name: "Appliances",
		Slug: "appliances",
		Subcategories: []string{
			"Refrigerators",
			"Washing Machines",
			"Air Conditioners",
			"Microwave Ovens",
			"Water Purifiers",
			"Vacuum Cleaners",
			"Air Purifiers",
			"Fans",
			"Geysers",
			"Irons",
			"Mixer Grinders",
			"Juicers",
			"Toasters",
			"Coffee Makers",
			"Electric Kettles",
			"Food Processors",
			"Dishwashers",
		},
	},
	{
		Name: "Automotive",
		Slug: "automotive",
		Subcategories: []string{
			"Car Accessories",
			"Car Electronics",
			"Car Care",
			"Car Covers",
			"Car Mats",
			"Bike Accessories",
			"Bike Electronics",
			"Bike Care",
			"Helmets",
			"Riding Gear",
			"Car Parts",
			"Bike Parts",
			"Tires & Wheels",
			"Tools & Equipment",
		},
	},
	{
		Name: "Baby & Kids",
		Slug: "baby-kids",
		Subcategories: []string{
			"Baby Diapers",
			"Baby Food",
			"Baby Care",
			"Baby Bath",
			"Baby Grooming",
			"Feeding & Nursing",
			"Bottles & Nipples",
			"Baby Gear",
			"Strollers",
			"Car Seats",
			"Baby Carriers",
			"Cribs & Cradles",
			"Baby Monitors",
			"Baby Safety",
			"Toys & Games",
			"Educational Toys",
			"Action Figures",
			"Dolls",
			"Remote Control Toys",
		},
	},
	{
		Name: "Food & Beverages",
		Slug: "food-beverages",
		Subcategories: []string{
			"Snacks",
			"Chocolates & Sweets",
			"Biscuits & Cookies",
			"Beverages",
			"Tea & Coffee",
			"Health Drinks",
			"Soft Drinks",
			"Breakfast Cereals",
			"Noodles & Pasta",
			"Cooking Oils",
			"Spices & Masalas",
			"Rice & Flour",
			"Dry Fruits",
			"Organic Foods",
			"Gourmet Foods",
		},
	},
	{
		Name: "Pet Supplies",
		Slug: "pet-supplies",
		Subcategories: []string{
			"Dog Food",
			"Cat Food",
			"Bird Food",
			"Fish Food",
			"Pet Toys",
			"Pet Grooming",
			"Pet Accessories",
			"Pet Beds",
			"Pet Bowls",
			"Pet Collars & Leashes",
			"Pet Health Care",
			"Aquarium Supplies",
			"Pet Training",
		},
	},
	{
		Name: "Jewelry",
		Slug: "jewelry",
		Subcategories: []string{
			"Gold Jewelry",
			"Silver Jewelry",
			"Diamond Jewelry",
			"Platinum Jewelry",
			"Rings",
			"Earrings",
			"Necklaces",
			"Bracelets",
			"Bangles",
			"Pendants",
			"Anklets",
			"Gemstones",
			"Fashion Jewelry",
			"Men's Jewelry",
		},
	},
	{
		Name: "Toys & Games",
		Slug: "toys-games",
		Subcategories: []string{
			"Action Figures",
			"Dolls & Dollhouses",
			"Board Games",
			"Puzzles",
			"Building Blocks",
			"Remote Control Toys",
			"Educational Toys",
			"Outdoor Toys",
			"Soft Toys",
			"Musical Toys",
			"Electronic Toys",
			"Art & Craft Toys",
			"Card Games",
			"Video Games",
		},
	},
	{
		Name: "Bags & Luggage",
		Slug: "bags-luggage",
		Subcategories: []string{
			"Backpacks",
			"Laptop Bags",
			"Handbags",
			"Messenger Bags",
			"Suitcases",
			"Travel Bags",
			"Duffel Bags",
			"Trolley Bags",
			"School Bags",
			"Wallets",
			"Clutches",
			"Waist Bags",
			"Gym Bags",
		},
	},
	{
		Name: "Gift Items",
		Slug: "gift-items",
		Subcategories: []string{
			"Gift Sets",
			"Greeting Cards",
			"Gift Vouchers",
			"Flowers",
			"Cakes",
			"Personalized Gifts",
			"Mugs",
			"Photo Frames",
			"Soft Toys",
			"Chocolates",
			"Gift Hampers",
			"Decorative Items",
		},
	},
}

// CategoryMap maps category names to their subcategories for quick lookup
var CategoryMap = make(map[string][]string)

// SubcategoryMap maps all subcategories to their parent category
var SubcategoryMap = make(map[string]string)

// AllSubcategories contains all subcategories across all categories
var AllSubcategories []string

func init() {
	for _, category := range ProductCategories {
		CategoryMap[category.Name] = category.Subcategories
		for _, subcategory := range category.Subcategories {
			SubcategoryMap[subcategory] = category.Name
			AllSubcategories = append(AllSubcategories, subcategory)
		}
	}
}

// IsValidCategory checks if a given category name is valid
func IsValidCategory(categoryName string) bool {
	_, exists := CategoryMap[categoryName]
	return exists
}

// IsValidSubcategory checks if a given subcategory is valid
func IsValidSubcategory(subcategory string) bool {
	_, exists := SubcategoryMap[subcategory]
	return exists
}

// GetSubcategories returns all subcategories for a given category
func GetSubcategories(categoryName string) []string {
	return CategoryMap[categoryName]
}

// GetParentCategory returns the parent category for a given subcategory
func GetParentCategory(subcategory string) string {
	return SubcategoryMap[subcategory]
}

// GetAllCategories returns all categories
func GetAllCategories() []Category {
	return ProductCategories
}

// GetCategoryNames returns just the category names
func GetCategoryNames() []string {
	names := make([]string, len(ProductCategories))
	for i, category := range ProductCategories {
		names[i] = category.Name
	}
	return names
}

package fahasa

type Product struct {
	EntityID                int64                `json:"entity_id"`
	EntityIDSub             int64                `json:"entity_id_sub"`
	Category                string               `json:"category"`
	CategoryMain            string               `json:"category_main"`
	CategoryMid             string               `json:"category_mid"`
	Category3               string               `json:"category_3"`
	Category4               string               `json:"category_4"`
	CategoryMainID          int64                `json:"category_main_id"`
	CategoryMidID           int64                `json:"category_mid_id"`
	Category3ID             int64                `json:"category_3_id"`
	Category4ID             int64                `json:"category_4_id"`
	TypeID                  string               `json:"type_id"`
	SKU                     string               `json:"sku"`
	HasOptions              int                  `json:"has_options"`
	RequiredOptions         int                  `json:"required_options"`
	CreatedAt               *string              `json:"created_at"`
	UpdatedAt               *string              `json:"updated_at"`
	Name                    string               `json:"name"`
	Episode                 string               `json:"episode"`
	Quantity                int                  `json:"qty"`
	MetaTitle               *string              `json:"meta_title"`
	MetaKeyword             *string              `json:"meta_keyword"`
	MetaDescription         *string              `json:"meta_description"`
	ImagePath               string               `json:"image_path"`
	CountDownReleaseProduct *string              `json:"count_down_release_product"`
	AttrKey                 *string              `json:"attr_key"`
	Image                   string               `json:"image"`
	SmallImage              string               `json:"small_image"`
	IsConfigurable          bool                 `json:"isConfigurable"`
	Childs                  []any                `json:"childs"`
	DisableSelect           bool                 `json:"disable_select"`
	FinalPrice              int64                `json:"final_price"`
	Price                   int64                `json:"price"`
	Visibility              int                  `json:"visibility"`
	Status                  int                  `json:"status"`
	SoonRelease             string               `json:"soon_release"`
	PageCount               int                  `json:"qty_of_page"`
	IsAvailable             bool                 `json:"is_available"`
	DiscountPercent         int                  `json:"discount_percent"`
	StockAvailable          string               `json:"stock_available"`
	HasStock                bool                 `json:"has_stock"`
	ImageLabel              *string              `json:"image_label"`
	SmallImageLabel         *string              `json:"small_image_label"`
	ThumbnailLabel          *string              `json:"thumbnail_label"`
	GiftMessageAvailable    *string              `json:"gift_message_available"`
	PublishYear             int                  `json:"publish_year"`
	Size                    string               `json:"size"`
	Author                  string               `json:"author"`
	Publisher               string               `json:"publisher"`
	Supplier                string               `json:"supplier"`
	SupplierID              string               `json:"supplier_id"`
	Translator              string               `json:"translator"`
	Weight                  int                  `json:"weight"`
	CountryOfManufacture    *string              `json:"country_of_manufacture"`
	TaxClassID              *int                 `json:"tax_class_id"`
	WeightType              *string              `json:"weight_type"`
	Featured                *string              `json:"featured"`
	BookLayout              string               `json:"book_layout"`
	Exclusive               *string              `json:"exclusive"`
	Description             string               `json:"description"`
	ShortDescription        *string              `json:"short_description"`
	Attributes              []ProductAttribute   `json:"attributes"`
	ListComment             *string              `json:"list_comment"`
	NewsFromDate            *string              `json:"news_from_date"`
	NewsToDate              *string              `json:"news_to_date"`
	MediaGallery            ProductMediaGallery  `json:"media_gallery"`
	RatingSummary           ProductRatingSummary `json:"rating_summary"`
	ListRelated             []any                `json:"list_related"`
	ListRelated2            []any                `json:"list_related2"`
	MaxRelated              int                  `json:"maxRelated"`
	ListBundled             []any                `json:"list_bundled"`
	EnableVoteProduct       string               `json:"enableVoteProduct"`
	EnableFlashsale         string               `json:"enableFlashsale"`
	EnableBuffetCombo       string               `json:"enableBuffetCombo"`
	Promotion               ProductPromotion     `json:"promotion"`
	MaxSaleQuantity         *int                 `json:"max_sale_qty"`
	MinimumQuantity         int                  `json:"min_qty"`
	SoldQuantity            string               `json:"sold_qty"`
	FrameImage              *string              `json:"frame_image"`
	Evoucher                int                  `json:"evoucher"`
	NoReturn                int                  `json:"no_return"`
	Label                   string               `json:"label"`
	NoInvoice               int                  `json:"no_invoice"`
	Success                 bool                 `json:"success"`
}

type ProductAttribute struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
}

type ProductGalleryImage struct {
	ValueID  int64  `json:"value_id"`
	File     string `json:"file"`
	Label    string `json:"label"`
	Position int    `json:"position"`
	Type     string `json:"type"`
	GroupID  int64  `json:"group_id"`
	Title    string `json:"title"`
	EntityID *int64 `json:"entity_id"`
}

type ProductMediaGallery struct {
	Images []ProductGalleryImage `json:"images"`
}

type ProductStarCount struct {
	One   int `json:"1"`
	Two   int `json:"2"`
	Three int `json:"3"`
	Four  int `json:"4"`
	Five  int `json:"5"`
}

type ProductRatingSummary struct {
	TotalStar        ProductStarCount `json:"total_star"`
	ReviewsCountFhs  int              `json:"reviews_count_fhs"`
	RatingSummaryFhs int              `json:"rating_summary_fhs"`
	ReviewsCountAmz  int              `json:"reviews_count_amz"`
	RatingSummaryAmz int              `json:"rating_summary_amz"`
	ReviewsCountGr   int              `json:"reviews_count_gr"`
	RatingSummaryGr  int              `json:"rating_summary_gr"`
	ReviewsCountAma  *int             `json:"reviews_count_ama"`
	RatingSummaryAma *int             `json:"rating_summary_ama"`
}

type ProductPromotionRule struct {
	ID            int64   `json:"id"`
	BlockDetailID string  `json:"block_detail_id"`
	RuleContent   string  `json:"rule_content"`
	ButtonLink    *string `json:"btn_link"`
	ButtonTitle   *string `json:"btn_title"`
	Title         string  `json:"title"`
	Title2        string  `json:"title_2"`
	CouponCode    string  `json:"coupon_code"`
	ExpireDate    string  `json:"expire_date"`
	EventType     int     `json:"event_type"`
	Matched       bool    `json:"matched"`
	CloseButton   bool    `json:"close_btn"`
	Priority      int     `json:"priority,omitempty"`
}

type ProductPromotion struct {
	AffectAll      []ProductPromotionRule `json:"affect_all"`
	AffectCoupons  []ProductPromotionRule `json:"affect_coupons"`
	AffectPayments []ProductPromotionRule `json:"affect_payments"`
	AffectCarts    []ProductPromotionRule `json:"affect_carts"`
}

// category

type Category struct {
	ID    int64   `json:"id,string"`
	Name  string  `json:"name"`
	Count *int64  `json:"count,string"` // optional
	Path  string  `json:"path"`
	URL   *string `json:"url"` // optional
}

type CategoryProducts struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	// ParentCategories   []Category        `json:"parent_categories"`
	// Category           Category          `json:"category"`
	// ChildrenCategories []Category        `json:"children_categories"`
	Attributes    []FilterAttribute `json:"attributes"`
	PriceRange    FilterPriceRange  `json:"price_range"`
	TotalProducts int64             `json:"total_products"`
	ProductList   []CategoryProduct `json:"product_list"`
	NoOfPages     int64             `json:"noofpages"`
	Success       bool              `json:"success"`
}

type CategoryProduct struct {
	TypeID            string      `json:"type_id"`
	Type              string      `json:"type"`
	ProductID         int64       `json:"product_id,string"`
	ProductName       string      `json:"product_name"`
	ProductFinalPrice string      `json:"product_finalprice"`
	ProductPrice      string      `json:"product_price"`
	RatingHTML        string      `json:"rating_html"`
	SoonRelease       string      `json:"soon_release"`
	ProductURL        string      `json:"product_url"`
	ImageSrc          string      `json:"image_src"`
	Discount          int         `json:"discount"`
	DiscountLabelHTML string      `json:"discount_label_html"`
	Episode           *string     `json:"episode"` // Pointer to handle nulls
	Label             string      `json:"label"`
	FrameImage        interface{} `json:"frame_image"`
}

// filters

type FilterPriceRange struct {
	PriceRange
	Limits PriceRange `json:"price_range"` // price range limits of current filter
}

type PriceRange struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type FilterAttribute struct {
	ID      int64                   `json:"id,string"`
	Code    string                  `json:"code"`
	Label   string                  `json:"label"`
	Param   string                  `json:"param"`
	Options []FilterAttributeOption `json:"options"`
}

type FilterAttributeOption struct {
	ID       int64  `json:"id,string"`
	Label    string `json:"label"`
	Selected bool   `json:"selected"`
	Param    string `json:"param"`
	Count    int64  `json:"count,string"`
}

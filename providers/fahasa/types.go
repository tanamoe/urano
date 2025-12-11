package fahasa

type Product struct {
	EntityID                int64         `json:"entity_id"`
	EntityIDSub             int64         `json:"entity_id_sub"`
	Category                string        `json:"category"`
	CategoryMain            string        `json:"category_main"`
	CategoryMid             string        `json:"category_mid"`
	Category3               string        `json:"category_3"`
	Category4               string        `json:"category_4"`
	CategoryMainID          int64         `json:"category_main_id"`
	CategoryMidID           int64         `json:"category_mid_id"`
	Category3ID             int64         `json:"category_3_id"`
	Category4ID             int64         `json:"category_4_id"`
	TypeID                  string        `json:"type_id"`
	Sku                     string        `json:"sku"`
	HasOptions              int           `json:"has_options"`
	RequiredOptions         int           `json:"required_options"`
	CreatedAt               *string       `json:"created_at"`
	UpdatedAt               *string       `json:"updated_at"`
	Name                    string        `json:"name"`
	Episode                 string        `json:"episode"`
	Qty                     int           `json:"qty"`
	MetaTitle               *string       `json:"meta_title"`
	MetaKeyword             *string       `json:"meta_keyword"`
	MetaDescription         *string       `json:"meta_description"`
	ImagePath               string        `json:"image_path"`
	CountDownReleaseProduct *string       `json:"count_down_release_product"`
	AttrKey                 *string       `json:"attr_key"`
	MaxSaleQty              *int          `json:"max_sale_qty"`
	Image                   string        `json:"image"`
	SmallImage              string        `json:"small_image"`
	IsConfigurable          bool          `json:"isConfigurable"`
	Childs                  []any         `json:"childs"`
	DisableSelect           bool          `json:"disable_select"`
	FinalPrice              int64         `json:"final_price"`
	Price                   int64         `json:"price"`
	Visibility              int           `json:"visibility"`
	Status                  int           `json:"status"`
	SoonRelease             string        `json:"soon_release"`
	QtyOfPage               int           `json:"qty_of_page"`
	IsAvailable             bool          `json:"is_available"`
	DiscountPercent         int           `json:"discount_percent"`
	StockAvailable          string        `json:"stock_available"`
	HasStock                bool          `json:"has_stock"`
	ImageLabel              *string       `json:"image_label"`
	SmallImageLabel         *string       `json:"small_image_label"`
	ThumbnailLabel          *string       `json:"thumbnail_label"`
	GiftMessageAvailable    *string       `json:"gift_message_available"`
	PublishYear             int           `json:"publish_year"`
	Size                    string        `json:"size"`
	Author                  string        `json:"author"`
	Publisher               string        `json:"publisher"`
	Supplier                string        `json:"supplier"`
	SupplierID              string        `json:"supplier_id"`
	Translator              string        `json:"translator"`
	Weight                  int           `json:"weight"`
	CountryOfManufacture    *string       `json:"country_of_manufacture"`
	TaxClassID              *int          `json:"tax_class_id"`
	WeightType              *string       `json:"weight_type"`
	Featured                *string       `json:"featured"`
	BookLayout              string        `json:"book_layout"`
	Exclusive               *string       `json:"exclusive"`
	Description             string        `json:"description"`
	ShortDescription        *string       `json:"short_description"`
	Attributes              []Attribute   `json:"attributes"`
	ListComment             *string       `json:"list_comment"`
	NewsFromDate            *string       `json:"news_from_date"`
	NewsToDate              *string       `json:"news_to_date"`
	MediaGallery            MediaGallery  `json:"media_gallery"`
	RatingSummary           RatingSummary `json:"rating_summary"`
	ListRelated             []any         `json:"list_related"`
	ListRelated2            []any         `json:"list_related2"`
	MaxRelated              int           `json:"maxRelated"`
	ListBundled             []any         `json:"list_bundled"`
	EnableVoteProduct       string        `json:"enableVoteProduct"`
	EnableFlashsale         string        `json:"enableFlashsale"`
	EnableBuffetCombo       string        `json:"enableBuffetCombo"`
	Promotion               Promotion     `json:"promotion"`
	MinQty                  int           `json:"min_qty"`
	FrameImage              *string       `json:"frame_image"`
	Evoucher                int           `json:"evoucher"`
	NoReturn                int           `json:"no_return"`
	SoldQty                 string        `json:"sold_qty"`
	Label                   string        `json:"label"`
	NoInvoice               int           `json:"no_invoice"`
	Success                 bool          `json:"success"`
}

type Attribute struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
}

type GalleryImage struct {
	ValueID  int64  `json:"value_id"`
	File     string `json:"file"`
	Label    string `json:"label"`
	Position int    `json:"position"`
	Type     string `json:"type"`
	GroupID  int64  `json:"group_id"`
	Title    string `json:"title"`
	EntityID *int64 `json:"entity_id"`
}

type MediaGallery struct {
	Images []GalleryImage `json:"images"`
}

type StarCount struct {
	One   int `json:"1"`
	Two   int `json:"2"`
	Three int `json:"3"`
	Four  int `json:"4"`
	Five  int `json:"5"`
}

type RatingSummary struct {
	TotalStar        StarCount `json:"total_star"`
	ReviewsCountFhs  int       `json:"reviews_count_fhs"`
	RatingSummaryFhs int       `json:"rating_summary_fhs"`
	ReviewsCountAmz  int       `json:"reviews_count_amz"`
	RatingSummaryAmz int       `json:"rating_summary_amz"`
	ReviewsCountGr   int       `json:"reviews_count_gr"`
	RatingSummaryGr  int       `json:"rating_summary_gr"`
	ReviewsCountAma  *int      `json:"reviews_count_ama"`
	RatingSummaryAma *int      `json:"rating_summary_ama"`
}

type PromotionRule struct {
	ID            int64   `json:"id"`
	BlockDetailID string  `json:"block_detail_id"`
	RuleContent   string  `json:"rule_content"`
	BtnLink       *string `json:"btn_link"`
	BtnTitle      *string `json:"btn_title"`
	Title         string  `json:"title"`
	Title2        string  `json:"title_2"`
	CouponCode    string  `json:"coupon_code"`
	ExpireDate    string  `json:"expire_date"`
	EventType     int     `json:"event_type"`
	Matched       bool    `json:"matched"`
	CloseBtn      bool    `json:"close_btn"`
	Priority      int     `json:"priority,omitempty"`
}

type Promotion struct {
	AffectAll      []PromotionRule `json:"affect_all"`
	AffectCoupons  []PromotionRule `json:"affect_coupons"`
	AffectPayments []PromotionRule `json:"affect_payments"`
	AffectCarts    []PromotionRule `json:"affect_carts"`
}

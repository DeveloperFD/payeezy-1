package payeezy

type Authorize struct {
	Amount         string `json:"amount"`
	BillingAddress struct {
		City    string `json:"city,omitempty"`
		Country string `json:"country,omitempty"`
		Email   string `json:"email,omitempty"`
		Phone   struct {
			Number string `json:"number,omitempty"`
			Type   string `json:"type,omitempty"`
		} `json:"phone,omitempty"`
		StateProvince string `json:"state_province,omitempty"`
		Street        string `json:"street,omitempty"`
		ZipPostalCode string `json:"zip_postal_code,omitempty"`
	} `json:"billing_address,omitempty"`
	CreditCard struct {
		CardNumber     string `json:"card_number"`
		CardholderName string `json:"cardholder_name"`
		Cvv            string `json:"cvv"`
		ExpDate        string `json:"exp_date"`
		Type           string `json:"type"`
	} `json:"credit_card,omitempty"`
	CurrencyCode      string `json:"currency_code"`
	MerchantRef       string `json:"merchant_ref,omitempty"`
	Method            string `json:"method,omitempty"`
	PartialRedemption string `json:"partial_redemption,omitempty"`
	SoftDescriptors   struct {
		City                string `json:"city,omitempty"`
		CountryCode         string `json:"country_code,omitempty"`
		DbaName             string `json:"dba_name,omitempty"`
		Mcc                 string `json:"mcc,omitempty"`
		MerchantContactInfo string `json:"merchant_contact_info,omitempty"`
		Mid                 string `json:"mid,omitempty"`
		PostalCode          string `json:"postal_code,omitempty"`
		Region              string `json:"region,omitempty"`
		Street              string `json:"street,omitempty"`
	} `json:"soft_descriptors,omitempty"`
	SplitTenderID   string `json:"split_tender_id,omitempty"`
	TransactionType string `json:"transaction_type"`
}

// LEVEL 2
// Level 2 data fields such as Freight, Duty, etc. - values are for informational purposes only and are not directly reflected in the Total Amount field for the transaction when processed.
// Level2       struct {
// 	CustomerRef string `json:"customer_ref,omitempty"`
// 	Tax1Amount  string `json:"tax1_amount,omitempty"`
// 	Tax1Number  string `json:"tax1_number,omitempty"`
// 	Tax2Amount  string `json:"tax2_amount,omitempty"`
// 	Tax2Number  string `json:"tax2_number,omitempty"`
// } `json:"level2,omitempty"`

// LEVEL 3
// The following properties are used to populate additional information about the transaction, including shipping details and line item information.
//     Visa
//     Mastercard

// Level3 struct {
// 	AltTaxAmount   string `json:"alt_tax_amount,omitempty"`
// 	AltTaxID       string `json:"alt_tax_id,omitempty"`
// 	DiscountAmount string `json:"discount_amount,omitempty"`
// 	DutyAmount     string `json:"duty_amount,omitempty"`
// 	FreightAmount  string `json:"freight_amount,omitempty"`
// 	LineItems      []struct {
// 		CommodityCode     string `json:"commodity_code,omitempty"`
// 		Description       string `json:"description,omitempty"`
// 		DiscountAmount    string `json:"discount_amount,omitempty"`
// 		DiscountIndicator string `json:"discount_indicator,omitempty"`
// 		GrossNetIndicator string `json:"gross_net_indicator,omitempty"`
// 		LineItemTotal     string `json:"line_item_total,omitempty"`
// 		ProductCode       string `json:"product_code,omitempty"`
// 		Quantity          string `json:"quantity,omitempty"`
// 		TaxAmount         string `json:"tax_amount,omitempty"`
// 		TaxRate           string `json:"tax_rate,omitempty"`
// 		TaxType           string `json:"tax_type,omitempty"`
// 		UnitCost          string `json:"unit_cost,omitempty"`
// 		UnitOfMeasure     string `json:"unit_of_measure,omitempty"`
// 	} `json:"line_items,omitempty"`
// 	ShipFromZip   string `json:"ship_from_zip,omitempty"`
// 	ShipToAddress struct {
// 		Address1       string `json:"address_1,omitempty"`
// 		City           string `json:"city,omitempty"`
// 		Country        string `json:"country,omitempty"`
// 		CustomerNumber string `json:"customer_number,omitempty"`
// 		Email          string `json:"email,omitempty"`
// 		Name           string `json:"name,omitempty"`
// 		Phone          string `json:"phone,omitempty"`
// 		State          string `json:"state,omitempty"`
// 		Zip            string `json:"zip,omitempty"`
// 	} `json:"ship_to_address,omitempty"`
// } `json:"level3,omitempty"`

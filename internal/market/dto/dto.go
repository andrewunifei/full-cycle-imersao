package dto

// Recebe do Kafka nesse formato de objeto
type TradeInput struct {
	OrderID			string	`json:"order_id"`
	InvestorID		string	`json:"investor_id"`
	AssetID			string	`json:"asset_id"`
	CurrentShares	int		`json:"current_shares"`
	Shares 			int		`json:"shares"`
	Price 			float64	`json:"price"`
	OrderType 		string	`json:"order_type"`
}

// Retorna para o Kafka nesse formato de objeto
type OrderOutput struct {
	OrderID				string					`json:"order_id"`
	InvestorID			string					`json:"investor_id"`
	AssetID				string					`json:"asset_id"`
	OrderType			string					`json:"order_type"`
	Status				string					`json:"status"`
	Partial				int						`json:"partial"`
	Shares				int						`json:"shares"`
	TransactionOutput	[]*TransactionOutput	`json:"trasactions"`
}

type TransactionOutput struct {
	TransactionID	string	`json:"transaction_id"`
	BuyerID			string	`json:"buyer_id"`
	SellerID		string	`json:"seller_id"`
	AssetID			string	`json:"asset_id"`
	Price			float64	`json:"price_id"`
	Shares			int		`json:"shares_id"`
}
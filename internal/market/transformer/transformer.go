package transformer

import (
	"github.com/andrewunifei/full-cycle-imersao/go/internal/market/dto"
	"github.com/andrewunifei/full-cycle-imersao/go/internal/market/entity"
	"golang.org/x/tools/go/analysis/passes/appends"
)

// Preenche o conteúdo da Order a partir do DTO recebido
// Ou hidratação do objeto
func TransformInput(input dto.TradeInput) *entity.Order {
	asset := entity.NewAsset(input.AssetID, input.AssetID, 1000)
	investor := entity.NewInvestor(input.InvestorID, input.InvestorID)
	order := entity.NewOrder(input.OrderID, investor, asset, input.Shares, input.Price, input.OrderType)

	if input.CurrentShares > 0 {
		assetPosition := entity.NewInvestorAssetPosition(input.AssetID, input.CurrentShares)
		investor.AddAssetPosition(assetPosition)
	}

	return order
}

// Prepara o DTO para ser enviado a partir da entidade Order
func TransformOutput(order *entity.Order) *dto.OrderOutput {
	output := &dto.OrderOutput{
		OrderID: order.ID,
		InvestorID: order.Investor.ID,
		AssetID: order.Asset.ID,
		OrderType: order.OrderType,
		Status: order.Status,
		Partial: order.PendingShares,
		Shares: order.Shares,
	}

	var transactionsOutput []*dto.TransactionOutput

	for _, t := range order.Transactions {
		transactionDto := &dto.TransactionOutput{
			TransactionID: t.ID,
			BuyerID: t.BuyingOrder.Investor.ID,
			SellerID: t.SellingOrder.Investor.ID,
			AssetID: t.SellingOrder.Asset.ID,
			Price: t.Price,
			Shares: t.SellingOrder.Shares - t.SellingOrder.PendingShares,
		}

		transactionsOutput = append(transactionsOutput, transactionDto)
	}

	output.TransactionOutput = transactionsOutput

	return output
}
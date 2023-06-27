package mongo

import (
	"github/mrbryside/pocket-service/internal/domain/eventgen"
	"github/mrbryside/pocket-service/internal/entity"
)

func genPocketTransactionAddedEvent(em []eventModel, eg *eventgen.EventGen) {
	for _, val := range em {
		if val.EventType != "transactionAdded" {
			continue
		}
		for _, v := range val.Transactions {
			t := entity.Transaction{
				Amount:    v.Amount,
				Category:  v.Category,
				CreatedAt: v.CreatedAt,
			}
			eg.AddTransactions(t)
		}
	}
}

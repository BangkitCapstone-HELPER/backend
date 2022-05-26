package services

import (
	"encoding/json"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dto"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/repo"
	"go.uber.org/fx"
)

type transactionServiceParams struct {
	fx.In
	TransactionRepo repo.TransactionRepo
}

type TransactionService interface {
	CreateTransaction(transaction dto.CreateTransactionRequestDTO) (dto.TransactionDTO, error)
	GetAllTransactionByUID(userId uint) ([]dto.TransactionDTO, error)
	UpdateTransaction(trxId dto.UpdateTransactionDTO) (dto.TransactionDTO, error)
}

func NewTransactionService(params transactionServiceParams) TransactionService {
	return &params
}

func (u *transactionServiceParams) CreateTransaction(transaction dto.CreateTransactionRequestDTO) (dto.TransactionDTO, error) {
	tempTransaction := transaction.ToDAO()
	tempTransaction.Remaining = tempTransaction.Count
	newTransaction, err := u.TransactionRepo.CreateTransaction(tempTransaction)
	if err != nil {
		return dto.TransactionDTO{}, err
	}
	return dto.NewTransactionDTO(newTransaction), err
}

func (u *transactionServiceParams) GetAllTransactionByUID(userId uint) ([]dto.TransactionDTO, error) {
	transactions, err := u.TransactionRepo.GetTransactionByUID(userId)

	var result []dto.TransactionDTO
	for _, transaction := range transactions {
		result = append(result, dto.ToTransactionDTO(transaction))
	}

	return result, err
}

func (u *transactionServiceParams) UpdateTransaction(trx dto.UpdateTransactionDTO) (dto.TransactionDTO, error) {
	var updateMap map[string]interface{}
	data, _ := json.Marshal(trx)
	json.Unmarshal(data, &updateMap)
	for k, v := range updateMap {
		if v == "" {
			delete(updateMap, k)
		}
		if v == false {
			delete(updateMap, k)
		}
	}
	transaction, err := u.TransactionRepo.UpdateTransaction(trx.ID, updateMap)

	if err != nil {
		return dto.TransactionDTO{}, err
	}
	return dto.ToTransactionDTO(transaction), err
}

package usecase

import "github.com/brenda-martins/code-bank-golang/domain"

type UseCaseTransaction struct {
	TransactionRepository domain.TransactionRepository
	CreditCardRepository  domain.CreditCardRepository
}

func NewUseCaseTransaction(transactionRepository domain.TransactionRepository, creditCardRepository domain.CreditCardRepository) UseCaseTransaction {
	return UseCaseTransaction{TransactionRepository: transactionRepository, CreditCardRepository: creditCardRepository}
}

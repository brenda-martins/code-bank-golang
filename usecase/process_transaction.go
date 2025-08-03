package usecase

import (
	"time"

	"github.com/brenda-martins/code-bank-golang/domain"
	"github.com/brenda-martins/code-bank-golang/dto"
)

type UseCaseTransaction struct {
	TransactionRepository domain.TransactionRepository
	CreditCardRepository  domain.CreditCardRepository
}

func NewUseCaseTransaction(transactionRepository domain.TransactionRepository, creditCardRepository domain.CreditCardRepository) UseCaseTransaction {
	return UseCaseTransaction{TransactionRepository: transactionRepository, CreditCardRepository: creditCardRepository}
}

func (u UseCaseTransaction) ProcessTransaction(transactionDto dto.Transaction) (domain.Transaction, error) {
	creditCard := hydrateCreditCard(transactionDto)
	ccBalanceAndLimit, err := u.CreditCardRepository.GetCreditCard(*creditCard)
	if err != nil {
		return domain.Transaction{}, err
	}

	creditCard.ID = ccBalanceAndLimit.ID
	creditCard.Limit = ccBalanceAndLimit.Limit
	creditCard.Balance = ccBalanceAndLimit.Balance

	t := hydrateTransaction(transactionDto, ccBalanceAndLimit)
	t.ProcessAndValidate(creditCard)

	err = u.TransactionRepository.SaveTransaction(*t, *creditCard)
	if err != nil {
		return domain.Transaction{}, err
	}

	return *t, nil
}

func hydrateCreditCard(transactionDto dto.Transaction) *domain.CreditCard {
	creditCard := domain.NewCreditCard()
	creditCard.Name = transactionDto.Name
	creditCard.Number = transactionDto.Number
	creditCard.ExpirationMonth = transactionDto.ExpirationMonth
	creditCard.ExpirationYear = transactionDto.ExpirationYear
	creditCard.CVV = transactionDto.CVV
	return creditCard
}

func hydrateTransaction(transaction dto.Transaction, cc domain.CreditCard) *domain.Transaction {
	t := domain.NewTransaction()
	t.CreditCardId = cc.ID
	t.Amount = transaction.Amount
	t.Store = transaction.Store
	t.Description = transaction.Description
	t.CreatedAt = time.Now()
	return t
}

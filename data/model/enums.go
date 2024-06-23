package model

type AccountStatus string
type AccountType string
type CurrencySymbol string
type TransactionType string
type TransactionPurpose string
type TransactionStatus string

const (
	AccountStatusActive AccountStatus = "ACTIVE"
	AccountStatusInactive  AccountStatus = "INACTIVE"
	AccountStatusSuspended AccountStatus = "SUSPENDED"

	AccountTypeChecking AccountType = "CHECKING"
	AccountTypeSavings  AccountType = "SAVINGS"

	CurrencySymbolUSD CurrencySymbol = "USD"
	CurrencySymbolNGN CurrencySymbol = "NGN"

	Debit  TransactionType = "DEBIT"
	Credit TransactionType = "CREDIT"

	DepositInternal    TransactionPurpose = "DEPOSIT_INTERNAL"
	DepositExternal    TransactionPurpose = "DEPOSIT_EXTERNAL"
	WithdrawalExternal TransactionPurpose = "WITHDRAWAL_EXTERNAL"
	WithdrawalInternal TransactionPurpose = "WITHDRAWAL_INTERNAL"

	Pending   TransactionStatus = "PENDING"
	Completed TransactionStatus = "SUCCESS"
	Failed    TransactionStatus = "FAILED"
)
package transaction_status

type TransactionStatus string

const (
	Pending   TransactionStatus = "pending"
	Completed TransactionStatus = "completed"
	Cancelled TransactionStatus = "cancelled"
	OnGoing   TransactionStatus = "ongoing"
)

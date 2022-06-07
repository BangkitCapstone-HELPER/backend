package transaction_status

type TransactionStatus string

const (
	Pending   TransactionStatus = "pending"
	Waiting   TransactionStatus = "waiting"
	Completed TransactionStatus = "completed"
	Cancelled TransactionStatus = "cancelled"
	OnGoing   TransactionStatus = "ongoing"
)

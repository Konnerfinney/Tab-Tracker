package models

// UserBalance represents the balance between two users.
type UserBalance struct {
	BalanceID  int     `json:"balance_id"`
	User1ID    int     `json:"user1_id"`
	User2ID    int     `json:"user2_id"`
	NetBalance float64 `json:"net_balance"`
}

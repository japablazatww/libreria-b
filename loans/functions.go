package loans

type LoanRequest struct {
	Amount   float64 `json:"amount"`
	Term     int     `json:"term"` // months
	UserType string  `json:"user_type"`
}

type LoanResponse struct {
	Approved     bool    `json:"approved"`
	InterestRate float64 `json:"interest_rate"`
	MonthlyPay   float64 `json:"monthly_pay"`
	Message      string  `json:"message"`
}

// CalculateLoan determines if a loan is feasible and calculates payment.
func CalculateLoan(req LoanRequest) (LoanResponse, error) {
	rate := 5.0
	if req.UserType == "PREMIUM" {
		rate = 3.5
	}

	if req.Amount > 50000 && req.UserType != "PREMIUM" {
		return LoanResponse{
			Approved: false,
			Message:  "Amount too high for standard user",
		}, nil
	}

	monthlyRate := rate / 100 / 12
	payment := (req.Amount * monthlyRate) / (1 - (1 / float64(req.Term))) // Simplified formula, careful with 0 term

	return LoanResponse{
		Approved:     true,
		InterestRate: rate,
		MonthlyPay:   payment,
		Message:      "Loan approved",
	}, nil
}

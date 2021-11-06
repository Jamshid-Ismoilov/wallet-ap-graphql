package model

type BalanceBody struct {
	Balance float64 `json:"balance"`
}

type ByCategoryRequestBody struct {
	Token string `json:"token"`
	ID    int    `json:"id"`
}

type DailyRequestBody struct {
	Token string `json:"token"`
	Day   int    `json:"day"`
	Month int    `json:"month"`
	Year  int    `json:"year"`
}

type IsDone struct {
	Content bool `json:"content"`
}

type JWTUser struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type MonthlyRequestBody struct {
	Token string `json:"token"`
	Month int    `json:"month"`
	Year  int    `json:"year"`
}

type NewUser struct {
	Email string `json:"email"`
}

type Payment struct {
	Token      string  `json:"token"`
	Comment    string  `json:"comment"`
	Amount     float64 `json:"amount"`
	CreatedAt  string  `json:"createdAt"`
	CategoryID int     `json:"categoryId"`
}

type RequestBody struct {
	Token string `json:"token"`
}

type SetBalanceBody struct {
	Token  string  `json:"token"`
	Amount float64 `json:"amount"`
}

type StatisticsBody struct {
	Amount       float64 `json:"amount"`
	Time         string  `json:"time"`
	CategoryName string  `json:"categoryName"`
	Comment      *string `json:"comment"`
}

type Status struct {
	Content string `json:"content"`
}

type Token struct {
	Content string `json:"content"`
}

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

package databasepq

import (
	"app/graph/model"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DB struct {
	Conn *sql.DB
}

func (d DB) SelectDBPost(input model.NewUser) (result bool) {
	err := d.Conn.QueryRow(SELECTNEW, input.Email).Scan(&result)

	if err != nil {
		log.Fatalf("failed to execute query SELECTNEW: %v", err)
	}
	return
}

func (d DB) InsertUser(input model.User) (userId int) {
	err := d.Conn.QueryRow(
		INSERTUSER,
		input.Firstname,
		input.Lastname,
		input.Email,
		input.Password,
	).Scan(&userId)

	if err != nil {
		log.Fatalf("failed to insert user: %v", err)
	}
	return
}

func (d DB) UpdateUser(input model.User) (userId int) {
	err := d.Conn.QueryRow(
		UPDATEUSER,
		input.Firstname,
		input.Lastname,
		input.Email,
		input.Password,
	).Scan(&userId)

	if err != nil {
		log.Fatalf("failed to insert user: %v", err)
	}
	return
}

func (d DB) CheckUser(input model.User) (isTrue bool, userId int) {
	fmt.Println(input)
	err := d.Conn.QueryRow(
		CHECKUSER,
		input.Email,
		input.Password,
		input.Firstname,
		input.Lastname,
	).Scan(&isTrue)
	fmt.Println(isTrue)
	if err != nil {
		log.Fatalf("failed to execute CHECKUSER query: %v", err)
	}
	if isTrue == true {
		err := d.Conn.QueryRow(
			SELECTUSERID,
			input.Email,
			input.Password,
			input.Firstname,
			input.Lastname,
		).Scan(&userId)
		if err != nil {
			log.Fatalf("failed to execute SELECTUSERID query: %v", err)
		}
	}
	fmt.Println(userId)

	return
}

func (d DB) DeleteUserById(userId int) bool {
	_ = d.Conn.QueryRow(DELETEUSERBYID, userId)
	log.Println(userId)
	return true
}

func (d DB) AddIncome(userId int, input model.Payment) bool {
	var incomeId int
	err := d.Conn.QueryRow(
		ADDINCOMEDB,
		userId,
		input.Amount,
		input.CategoryID,
	).Scan(&incomeId)

	if err != nil {
		log.Fatalf("failed to execute query ADDINCOMEDB: %v", err)
	}

	if input.CreatedAt != "" {
		_ = d.Conn.QueryRow(
			UPDATECREATEDATINCOME,
			input.CreatedAt,
			incomeId,
		)

		if err != nil {
			log.Fatalf("failed to execute query UPDATECREATEDAT: %v", err)
		}

	}

	if input.Comment != "" {
		_ = d.Conn.QueryRow(
			UPDATECOMMENTINCOME,
			input.Comment,
			incomeId,
		)

		if err != nil {
			log.Fatalf("failed to execute query UPDATECOMMENTINCOME: %v", err)
		}
	}

	return true
}

func (d DB) AddOutgoing(userId int, input model.Payment) bool {
	var outgoingId int
	err := d.Conn.QueryRow(
		ADDOUTGOINGDB,
		userId,
		input.Amount,
		input.CategoryID,
	).Scan(&outgoingId)

	if err != nil {
		log.Fatalf("failed to execute query ADDOUTGOINGDB: %v", err)
	}

	if input.CreatedAt != "" {
		_ = d.Conn.QueryRow(
			UPDATECREATEDATOUTGOING,
			input.CreatedAt,
			outgoingId,
		)

	}

	if input.Comment != "" {
		_ = d.Conn.QueryRow(
			UPDATECOMMENTOUTGOING,
			input.Comment,
			outgoingId,
		)
	}

	return true
}

func (d DB) CheckUserByIdAndEmail(userId int, email string) (result bool) {
	err := d.Conn.QueryRow(
		CHECKUSERBYIDANDEMAIL,
		userId,
		email,
	).Scan(&result)

	if err != nil {
		log.Fatalf("failed to execute query CHECKUSERBYIDANDEMAIL: %v", err)
	}

	return
}

func (d DB) SetBalance(userId int, balance float64) bool {
	_ = d.Conn.QueryRow(
		SETBALANCE,
		balance,
		userId,
	)

	return true
}

func (d DB) GetDailySpendings(input model.DailyRequestBody, userId int) (results []*model.StatisticsBody) {
	rows, err := d.Conn.Query(
		GETDAILYOUTGOINGS,
		input.Day,
		input.Month,
		input.Year,
		userId,
	)

	defer rows.Close()

	if err != nil {
		log.Fatalf("failed to execute query GETDAILYOUTGOINGS: %v", err)
	}

	for rows.Next() {
		var result model.StatisticsBody

		rows.Scan(
			&result.Amount,
			&result.Time,
			&result.CategoryName,
			&result.Comment,
		)
		results = append(results, &result)
	}

	return results
}

func (d DB) GetDailyIncomes(input model.DailyRequestBody, userId int) (results []*model.StatisticsBody) {
	rows, err := d.Conn.Query(
		GETDAILYINCOMES,
		input.Day,
		input.Month,
		input.Year,
		userId,
	)

	defer rows.Close()

	if err != nil {
		log.Fatalf("failed to execute query GETDAILYINCOMES: %v", err)
	}

	for rows.Next() {
		var result model.StatisticsBody

		rows.Scan(
			&result.Amount,
			&result.Time,
			&result.CategoryName,
			&result.Comment,
		)
		results = append(results, &result)
	}

	return results
}

func (d DB) GetMonthlySpendings(input model.MonthlyRequestBody, userId int) (results []*model.StatisticsBody) {
	rows, err := d.Conn.Query(
		GETMONTHLYOUTGOINGS,
		input.Month,
		input.Year,
		userId,
	)

	defer rows.Close()

	if err != nil {
		log.Fatalf("failed to execute query GETMONTHLYOUTGOINGS: %v", err)
	}

	for rows.Next() {
		var result model.StatisticsBody

		rows.Scan(
			&result.Amount,
			&result.Time,
			&result.CategoryName,
			&result.Comment,
		)
		results = append(results, &result)
	}

	return results
}

func (d DB) GetMonthlyIncomes(input model.MonthlyRequestBody, userId int) (results []*model.StatisticsBody) {
	rows, err := d.Conn.Query(
		GETMONTHLYINCOMES,
		input.Month,
		input.Year,
		userId,
	)

	defer rows.Close()

	if err != nil {
		log.Fatalf("failed to execute query GETMONTHLYINCOMES: %v", err)
	}

	for rows.Next() {
		var result model.StatisticsBody

		rows.Scan(
			&result.Amount,
			&result.Time,
			&result.CategoryName,
			&result.Comment,
		)
		results = append(results, &result)
	}

	return results
}
func (d DB) GetSpendingsByCategory(input model.ByCategoryRequestBody, userId int) (results []*model.StatisticsBody) {
	rows, err := d.Conn.Query(
		GETSPENDINGBYCATEGORY,
		userId,
		input.ID,
	)

	defer rows.Close()

	if err != nil {
		log.Fatalf("failed to execute query GETSPENDINGBYCATEGORY: %v", err)
	}

	for rows.Next() {
		var result model.StatisticsBody

		rows.Scan(
			&result.Amount,
			&result.Time,
			&result.CategoryName,
			&result.Comment,
		)
		results = append(results, &result)
	}

	return results
}

func (d DB) GetIncomesByCategory(input model.ByCategoryRequestBody, userId int) (results []*model.StatisticsBody) {
	rows, err := d.Conn.Query(
		GETINCOMESBYCATEGORY,
		userId,
		input.ID,
	)

	defer rows.Close()

	if err != nil {
		log.Fatalf("failed to execute query GETINCOMESBYCATEGORY: %v", err)
	}

	for rows.Next() {
		var result model.StatisticsBody

		rows.Scan(
			&result.Amount,
			&result.Time,
			&result.CategoryName,
			&result.Comment,
		)
		results = append(results, &result)
	}

	return results
}

func (d DB) GetBalanceOfUser(input model.JWTUser) (balance float64) {
	err := d.Conn.QueryRow(
		GETBALANCE,
		input.ID,
		input.Email,
	).Scan(&balance)

	if err != nil {
		log.Fatalf("failed to execute query GETBALANCE: %v", err)
	}
	return
}

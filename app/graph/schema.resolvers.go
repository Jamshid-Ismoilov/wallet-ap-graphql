package graph


import (
	"app/gojwt"
	"app/graph/generated"
	"app/graph/model"
	"app/myerrors"
	"app/redis"
	"app/tools"
	"context"
	"log"
)

func (r *mutationResolver) Registration(ctx context.Context, input model.NewUser) (*model.Status, error) {
	isExists := r.DB.SelectDBPost(input)

	if isExists == true {
		return &model.Status{Content: "You had already signup. You can login."}, nil
	}

	if isExists == false {

		password := tools.SendMail(input.Email)

		redis.SetPasswordByEmail(input, password)
	}

	return &model.Status{Content: "Password has been send to your email"}, nil
}

func (r *mutationResolver) ChangePassword(ctx context.Context, input model.NewUser) (*model.Status, error) {
	isExists := r.DB.SelectDBPost(input)

	if isExists == false {
		return &model.Status{Content: "You had not signed up yet. Signup first"}, nil
	}

	if isExists == true {

		password := tools.SendMail(input.Email)

		redis.SetPasswordByEmail(input, password)
	}

	return &model.Status{Content: "Password has been send to your email"}, nil
}

func (r *mutationResolver) CheckRegistration(ctx context.Context, input model.User) (*model.Token, error) {
	isExists := r.DB.SelectDBPost(model.NewUser{Email: input.Email})

	if isExists == true {
		return &model.Token{Content: ""}, &myerrors.UserExists{}
	}

	password, err := redis.GetPasswordByEmail(input.Email)

	if err != nil {
		log.Fatalf("failed to get password by email in redis: %v", err)
	}

	if password == input.Password {
		userId := r.DB.InsertUser(input)
		redis.DeletePasswordByEmail(input.Email)
		return &model.Token{Content: gojwt.GenerateJWT(model.JWTUser{ID: userId, Email: input.Email})}, nil
	}

	return &model.Token{Content: "404"}, nil
}

func (r *mutationResolver) CheckChangePassword(ctx context.Context, input model.User) (*model.Token, error) {
	password, err := redis.GetPasswordByEmail(input.Email)

	if err != nil {
		log.Fatalf("failed to get password by email in redis: %v", err)
	}

	if password == input.Password {
		userId := r.DB.UpdateUser(input)
		redis.DeletePasswordByEmail(input.Email)
		return &model.Token{Content: gojwt.GenerateJWT(model.JWTUser{ID: userId, Email: input.Email})}, nil
	}

	return &model.Token{Content: "404"}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.User) (*model.Token, error) {
	isTrue, userId := r.DB.CheckUser(input)

	if isTrue == true {
		return &model.Token{Content: gojwt.GenerateJWT(model.JWTUser{ID: userId, Email: input.Email})}, nil
	}

	return &model.Token{Content: ""}, &myerrors.Unauthorized{}
}

func (r *mutationResolver) Logout(ctx context.Context, input model.User) (*model.IsDone, error) {
	isTrue, userId := r.DB.CheckUser(input)

	if isTrue == true {
		return &model.IsDone{Content: r.DB.DeleteUserById(userId)}, nil
	}

	return &model.IsDone{Content: false}, &myerrors.NotExists{}
}

func (r *mutationResolver) AddIncoming(ctx context.Context, input model.Payment) (*model.IsDone, error) {
	user, err := gojwt.ParseJWT(input.Token)
	log.Println("OK")
	if err != nil {
		return &model.IsDone{Content: false}, err
	}

	isExists := r.DB.CheckUserByIdAndEmail(user.ID, user.Email)

	if isExists == false {
		return &model.IsDone{Content: false}, &myerrors.NotExists{}
	}

	if input.CategoryID == 0 || input.Amount == 0 || input.Token == "" {
		return &model.IsDone{Content: false}, &myerrors.Notvalid{}
	}
	log.Println("IOK")
	return &model.IsDone{Content: r.DB.AddIncome(user.ID, input)}, nil
}

func (r *mutationResolver) AddOutgoing(ctx context.Context, input model.Payment) (*model.IsDone, error) {
	user, err := gojwt.ParseJWT(input.Token)

	if err != nil {
		return &model.IsDone{Content: false}, err
	}

	isExists := r.DB.CheckUserByIdAndEmail(user.ID, user.Email)

	if isExists == false {
		return &model.IsDone{Content: false}, &myerrors.NotExists{}
	}

	if input.CategoryID == 0 || input.Amount == 0 || input.Token == "" {
		return &model.IsDone{Content: false}, &myerrors.Notvalid{}
	}

	return &model.IsDone{Content: r.DB.AddOutgoing(user.ID, input)}, nil
}

func (r *mutationResolver) SetBalance(ctx context.Context, input model.SetBalanceBody) (*model.IsDone, error) {
	user, err := gojwt.ParseJWT(input.Token)

	if err != nil {
		return &model.IsDone{Content: false}, err
	}

	isExists := r.DB.CheckUserByIdAndEmail(user.ID, user.Email)

	if isExists == false {
		return &model.IsDone{Content: false}, &myerrors.NotExists{}
	}

	return &model.IsDone{Content: r.DB.SetBalance(user.ID, input.Amount)}, nil
}

func (r *queryResolver) DailySpendings(ctx context.Context, input *model.DailyRequestBody) ([]*model.StatisticsBody, error) {
	user, err := gojwt.ParseJWT(input.Token)

	if err != nil {
		return []*model.StatisticsBody{&model.StatisticsBody{Amount: 0, Time: "", CategoryName: ""}}, &myerrors.NotExists{}
	}

	isExists := r.DB.CheckUserByIdAndEmail(user.ID, user.Email)

	if isExists == false {
		return []*model.StatisticsBody{&model.StatisticsBody{Amount: 0, Time: "", CategoryName: ""}}, &myerrors.NotExists{}
	}

	return r.DB.GetDailySpendings(*input, user.ID), nil
}

func (r *queryResolver) MonthlySpendings(ctx context.Context, input *model.MonthlyRequestBody) ([]*model.StatisticsBody, error) {
	user, err := gojwt.ParseJWT(input.Token)

	if err != nil {
		return []*model.StatisticsBody{&model.StatisticsBody{Amount: 0, Time: "", CategoryName: ""}}, &myerrors.NotExists{}
	}

	isExists := r.DB.CheckUserByIdAndEmail(user.ID, user.Email)

	if isExists == false {
		return []*model.StatisticsBody{&model.StatisticsBody{Amount: 0, Time: "", CategoryName: ""}}, &myerrors.NotExists{}
	}

	return r.DB.GetMonthlySpendings(*input, user.ID), nil
}

func (r *queryResolver) SpendingsByCategory(ctx context.Context, input *model.ByCategoryRequestBody) ([]*model.StatisticsBody, error) {
	user, err := gojwt.ParseJWT(input.Token)

	if err != nil {
		return []*model.StatisticsBody{&model.StatisticsBody{Amount: 0, Time: "", CategoryName: ""}}, &myerrors.NotExists{}
	}

	isExists := r.DB.CheckUserByIdAndEmail(user.ID, user.Email)

	if isExists == false {
		return []*model.StatisticsBody{&model.StatisticsBody{Amount: 0, Time: "", CategoryName: ""}}, &myerrors.NotExists{}
	}

	return r.DB.GetSpendingsByCategory(*input, user.ID), nil
}

func (r *queryResolver) DailyIncomes(ctx context.Context, input *model.DailyRequestBody) ([]*model.StatisticsBody, error) {
	user, err := gojwt.ParseJWT(input.Token)

	if err != nil {
		return []*model.StatisticsBody{&model.StatisticsBody{Amount: 0, Time: "", CategoryName: ""}}, &myerrors.NotExists{}
	}

	isExists := r.DB.CheckUserByIdAndEmail(user.ID, user.Email)

	if isExists == false {
		return []*model.StatisticsBody{&model.StatisticsBody{Amount: 0, Time: "", CategoryName: ""}}, &myerrors.NotExists{}
	}
	return r.DB.GetDailyIncomes(*input, user.ID), nil
}

func (r *queryResolver) MonthlyIncomes(ctx context.Context, input *model.MonthlyRequestBody) ([]*model.StatisticsBody, error) {
	user, err := gojwt.ParseJWT(input.Token)

	if err != nil {
		return []*model.StatisticsBody{&model.StatisticsBody{Amount: 0, Time: "", CategoryName: ""}}, &myerrors.NotExists{}
	}

	isExists := r.DB.CheckUserByIdAndEmail(user.ID, user.Email)

	if isExists == false {
		return []*model.StatisticsBody{&model.StatisticsBody{Amount: 0, Time: "", CategoryName: ""}}, &myerrors.NotExists{}
	}
	return r.DB.GetMonthlyIncomes(*input, user.ID), nil
}

func (r *queryResolver) IncomesByCategory(ctx context.Context, input *model.ByCategoryRequestBody) ([]*model.StatisticsBody, error) {
	user, err := gojwt.ParseJWT(input.Token)

	if err != nil {
		return []*model.StatisticsBody{&model.StatisticsBody{Amount: 0, Time: "", CategoryName: ""}}, &myerrors.NotExists{}
	}

	isExists := r.DB.CheckUserByIdAndEmail(user.ID, user.Email)

	if isExists == false {
		return []*model.StatisticsBody{&model.StatisticsBody{Amount: 0, Time: "", CategoryName: ""}}, &myerrors.NotExists{}
	}

	return r.DB.GetIncomesByCategory(*input, user.ID), nil
}

func (r *queryResolver) GetBalance(ctx context.Context, input *model.RequestBody) (*model.BalanceBody, error) {
	user, err := gojwt.ParseJWT(input.Token)

	if err != nil {
		return &model.BalanceBody{Balance: 0.0}, &myerrors.NotExists{}
	}

	isExists := r.DB.CheckUserByIdAndEmail(user.ID, user.Email)

	if isExists == false {
		return &model.BalanceBody{Balance: 0.0}, &myerrors.NotExists{}
	}

	return &model.BalanceBody{Balance: r.DB.GetBalanceOfUser(user)}, nil
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

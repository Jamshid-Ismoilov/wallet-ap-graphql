package myerrors


type Notvalid struct{}

func (*Notvalid) Error() string {
	return "data is not valid"
}

type Unauthorized struct{}

func (*Unauthorized) Error() string {
	return "unauthorized"
}

type ErrSignatureInvalid struct {}

func (*ErrSignatureInvalid) Error() string {
	return "Signature is invalid"
}

type UserExists struct{}

func (*UserExists) Error() string {
	return "User with this email exists"
}

type NotExists struct{}

func (*NotExists) Error() string {
	return "not exists"
}

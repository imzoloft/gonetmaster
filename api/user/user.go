package user

type User struct {
	Id  string `json:"id" validate:"uuid,required"`
	Key []byte `json:"key" validate:"required"`
}

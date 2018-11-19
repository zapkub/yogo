package models

// UserDocument basic user data
type UserDocument struct {
	ID string
}

// UserModel data of user
// to use in application
type UserModel struct {
	UserDocument
	context context
}

// Save do mutate User data to DB
func (u *UserModel) Save() {

}

// CreateUserModel create new user and save
// into database
func CreateUserModel(ctx context) YogoModel {
	return &UserModel{
		context: ctx,
	}
}

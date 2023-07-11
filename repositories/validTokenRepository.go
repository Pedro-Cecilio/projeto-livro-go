package repositories


type validToken interface{

	Create(User_id, jwt string)
}
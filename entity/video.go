package entity

type Video struct {
	Title       string `json:"title" binding:"min=2,max=10" validate:"is-cool"`
	Description string `json : "description" binding:"max=20"`
	URL         string `json : "url" binding:"required,url"`
}

type Person struct {
	FirstName string `json:"firstname" binding : "required"`
	LastName  string `json:"lastname" binding : "required"`
	Email     string `json:"email" validate:"required,email"`
	Age       int8   `json : "age" binding : "gte=1,lte=130" `
}

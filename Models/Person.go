package PersonModel

type Person struct {
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	Age       int    `bson:"age" binding:"required,gt=10"`
}

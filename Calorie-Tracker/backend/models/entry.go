package models
import(
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Entry struct{
	ID			primitive.ObjectID `bson:"id"`
	Dish		*string				`json:"dish"`
	Carbs		*string				`json:"carbs"`
	Proteins	*string				`json:"proteins"`
	Fats		*string				`json:"fats"`
	Calories 	*string				`json:"calories"`
	Ingredients *string				`json:"ingredients"`
}
package routes

import (
	"Calorie-Tracker/models"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


var validate = validator.New()
var entryCollection *mongo.Collection = openCollection(Client, "calories")

// 1. AddEntry - (https://shorturl.at/BvsMa)t
func AddEntry(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	var entry models.Entry

	// c.BindJSON(&entry) - This line attempts to bind the JSON body of the incoming HTTP request(which has data) to the entry variable, which is of type models.Entry.
	//If the binding fails (e.g., if the JSON is invalid or doesn't match the Entry struct), it returns a 500 Internal Server Error response.
	if err := c.BindJSON(&entry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	validationError := validate.Struct(entry)

	if validationError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": validationError.Error()})
		fmt.Println(validationError)
		return
	}

	entry.ID = primitive.NewObjectID()
	result, insertErr := entryCollection.InsertOne(ctx, entry)

	if insertErr != nil {
		msg := fmt.Sprint("Ordered Item not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		fmt.Println(insertErr)
		return
	}

	defer cancel()
	c.JSON(http.StatusOK, result)
}

// 2. GetEntries function
func GetEntries(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var entries []bson.M
	cursor, err := entryCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &entries); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()
	fmt.Println(entries)
	c.JSON(http.StatusOK, entries)
}

// 3. Get Entry by id Function
func GetEntryById(c *gin.Context) {

	EntryID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(EntryID)

	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	var entry bson.M

	if err := entryCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&entry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()
	fmt.Println(entry)
	c.JSON(http.StatusOK, entry)

}

// 4. GetEntries by Ingredients function
func GetEntriesByIngredients(c *gin.Context) {
	ingredient := c.Params.ByName("id")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var entries []bson.M

	cursor, err := entryCollection.Find(ctx, bson.M{"ingredients": ingredient})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		fmt.Println(err)
		return
	}

	if err := cursor.All(ctx, &entries); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()
	fmt.Println(entries)
	c.JSON(http.StatusOK, entries)
}

// 5. Update INgredient Function
func UpdateIngredients(c *gin.Context) {
	entryID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(entryID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	type Ingredient struct {
		Ingredients *string `json:"ingredients"`
	}

	var ingredient Ingredient

	if err := c.BindJSON(&ingredient); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	result, err := entryCollection.UpdateOne(ctx, bson.M{"_id": docID},
        bson.D{{"$set", bson.D{{"ingredients", ingredient.Ingredients}}}},
    )

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result.ModifiedCount)
}

// 6. Update Entry FUNCTION
func UpdateEntry(c *gin.Context) {

	enrtyID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(enrtyID)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var entry models.Entry
	if err := c.BindJSON(&entry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	validationError := validate.Struct(entry)

	if validationError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": validationError.Error()})
		fmt.Println(validationError)
		return
	}

	// For each field, if a json tag is present (e.g., json:"dish"), the encoder uses the tag's value ("dish") as the key in the JSON object.
	result, err := entryCollection.ReplaceOne(
		ctx,
		bson.M{"_id": docID},
		bson.M{
			"dish":        entry.Dish,
			"fat":         entry.Fats,
			"ingredients": entry.Ingredients,
			"calorie":     entry.Calories,
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()
	c.JSON(http.StatusOK, result.ModifiedCount)

}

// 7. DeleteEntry FUNCTION
func DeleteEntry(c *gin.Context) {
	entryId := c.Params.ByName("id")
	docId, _ := primitive.ObjectIDFromHex(entryId)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	result, err := entryCollection.DeleteOne(ctx, bson.M{"_id": docId})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()
	c.JSON(http.StatusOK, result.DeletedCount)

}

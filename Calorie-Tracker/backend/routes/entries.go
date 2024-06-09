package routes

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var entryCollection *mongo.Collection = openCollection(client, "calories")

func AddEntry(c *gin.Context) {

}

func GetEntries(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var entries []bson.M
	cursor,err := entryCollection.Find(ctx,bson.M{})

	if err !=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx,&entries);
	err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()
	fmt.Println(entries)
	c.JSON(http.StatusInternalServerError,entries)
}

func GetEntryById(c *gin.Context) {

}

func GetEntriesByIngredients(c *gin.Context) {

}
func UpdateIngredients(c *gin.Context) {

}
func UpdateEntry(c *gin.Context) {

}
func DeleteEntry(c *gin.Context) {
	entryId := c.Params.ByName("id")
	docId, _ := primitive.ObjectIdFromHex(entryId)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	result, err := entryCollection.Deleteone(ctx, bson.M{"_id": docId})

	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		fmt.Println(err)
		return 
	}

	defer cancel()
	c.JSON(http.StatusOK,result.DeletedCount)

}

package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/swlee/Ecommerce-sw/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func AddAddress() gin.HandlerFunc {

}

func EditHomeAddress() gin.HandlerFunc {

}

func EditWorkAddress() gin.HandlerFunc {

}

func DeleteAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		// id 값으로 주소를 찾음
		user_id := c.Query("id")

		if user_id == "" {
			c.Header("Context-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid Search Index"})
			c.Abort()
			return
		}

		addresses := make([]models.Address, 0)
		// 16진수객체 ID
		usert_id, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}

		ctx, cancle := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancle()
		filter := bson.D{primitive.E{Key: "_id", Value: usert_id}}
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "address", Value: addresses}}}}
		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(404, "Wrong command")
			return
		}
		defer cancle()
		ctx.Done()
		c.IndentedJSON(200, "Successfully Deleted")
	}

}

package controller

import (
	"context"
	"log"
	"net/http"
	"restaurant-management/database"
	"restaurant-management/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type OrderItemPack struct {
	Table_id    *string
	Order_items []models.OrderItem
}

var orderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderItem")

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		result, err := orderItemCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error ocurred while listing order items"})
			return
		}

		var allOrderItems []bson.M
		if err := result.All(ctx, &allOrderItems); err != nil {
			log.Fatal(err)
			return
		}
		c.JSON(http.StatusOK, allOrderItems)

	}
}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		orderItemId := c.param("order_item_id")
		var orderItem models.OrderItem

		err := orderItemCollection.FindOne(ctx, bson.M{"orderItem_id": orderItemId}).Decode(&orderItem)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing order item"})
			return
		}
		c.JSON(http.StatusOK, orderItem)
	}
}

func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderId := c.Param("order_id")
		allOrderItems, err := ItemsByOrder(orderId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error ocurred while listing order items by order"})
			return
		}
		c.JSON(http.StatusOK, allOrderItems)
	}
}

func ItemsByOrder(order string) (OrderItems []primitive.M, err error) {

}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var orderItemPack OrderItemPack
		var order models.Order

		if err := c.BindJSON(&orderItemPack);err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}

		order.Order_Date, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		orderItemsToBeInserted := []interface{}
		order.Table_id = orderItemPack.Table_id
		order_id := OrderItemOrderCreator(order)
		for _, orderItem := range orderItemPack.Order_items{
			orderItem.Order_id = order_id

			validateErr := validate.Struct(orderItem)

			if validateErr != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error":validateErr.Error()})
				return
			}
		}

	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

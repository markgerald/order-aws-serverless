package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/guregu/dynamo"
	"github.com/markgerald/vw-order/db"
	"github.com/markgerald/vw-order/model"
	"github.com/markgerald/vw-order/service"
	"net/http"
	"os"
)

func GetDb() *dynamo.DB {
	return db.InitDb()
}

func AddOrder(c *gin.Context) {
	var newOrder model.Order
	if err := c.BindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	table := GetDb().Table(os.Getenv("DYNAMODB_TABLE"))
	calc := service.SumOrder(newOrder)
	if err := table.Put(calc).Run(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, calc)
}

func GetOrder(c *gin.Context) {
	var order model.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	table := GetDb().Table("orders")
	if err := table.Get("id", c.Params.ByName("id")).One(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

func GetOrders(c *gin.Context) {
	var orders []model.Order
	table := GetDb().Table("orders")
	if err := table.Scan().All(&orders); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

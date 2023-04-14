package main

import (
    "mymodule/controller"
    _ "mymodule/database"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    api := router.Group("api/v1")
    items := api.Group("items")
    items.GET("", controller.GetItems)
    items.POST("", controller.PostItems)
    items.GET("/:id", controller.GetItemByID)
    items.POST("/:id", controller.PostItemByID)
    items.DELETE("/:id", controller.DeleteItemByID)
    items.GET("/search", controller.SearchItem)

    router.Run(":8000")
}
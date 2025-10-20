package main

import (
	"fmt"
	"log"
	"os"
	"stage1/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main (){
	r := gin.Default()
  

	 if err := godotenv.Load(); err != nil{
		log.Println("Error loading .env", err.Error())
	 }

    
   routes.StringRoutes(r) // string routes

	 PORT := os.Getenv("PORT")

	r.Run(fmt.Sprintf(`:%s`, PORT))
}
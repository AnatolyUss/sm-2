package main

import (
	"fmt"
	"log"
	"net/http"
	"sm-2/main/internal/lib/dto"

	"github.com/gin-gonic/gin"
)

func main() {
	greeting := "Hello, sm-2!"
	fmt.Println(greeting)

	log.SetPrefix("sm-2: ")
	log.SetFlags(0)

	router := gin.Default()
	router.GET("/healthcheck", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.POST("/model", func(c *gin.Context) {
		var modelDto dto.ModelDto

		if err := c.ShouldBindJSON(&modelDto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Model validation failed": err.Error()})
			return
		}

		// const modelService = new ModelService();
		// await modelService.saveModel(modelDto);
		// response.status(201).json(new ResponseDto('Ok'));

		c.JSON(http.StatusOK, gin.H{"status": "Good model!"})
	})

	router.POST("/request", func(c *gin.Context) {
		var requestDto dto.RequestDto

		if err := c.ShouldBindJSON(&requestDto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Request validation failed": err.Error()})
			return
		}

		// const modelService = new ModelService();
		//        const modelDto = await modelService.getModel(requestDto.path, requestDto.method);
		//
		//        if (!modelDto) {
		//          const msg = `Model not found for path:method "${requestDto.path}:${requestDto.method}"`;
		//          response.status(404).json(new ResponseDto(msg));
		//          return;
		//        }
		//
		//        const requestService = new RequestService();
		//        const validationResponse = requestService.validateRequest(requestDto, modelDto);
		//
		//        // Return status 200 instead of 201 as per explicit request at the task definition.
		//        response.status(200).json(new ResponseDto(validationResponse));

		c.JSON(http.StatusOK, gin.H{"status": "Good request!"})
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}

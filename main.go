package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// sample represents data about a record sample.
type sample struct {
	Name  string `json:"name"`
	ID    string `json:"id"`
	Title string `json:"title"`
}

// samples slice to seed record sample data.
var samples = []sample{
	{Name: "Test1", ID: "1", Title: "Blue Train"},
	{Name: "Test2", ID: "2", Title: "Jeru"},
	{Name: "Test3", ID: "3", Title: "Sarah Vaughan and Clifford Brown"},
}

func main() {
	router := gin.Default()
	router.GET("/samples", getsamples)
	router.GET("/samples/:name", getsampleByName)
	router.POST("/samples", postsamples)
	router.PUT("/samples/:name", putsampleByName)
	router.DELETE("/samples/:name", deletesampleByName)
	router.Run(":8080")
}

// getsamples godoc
// @Summary      List samples
// @Description  Get all sample items
// @Tags         samples
// @Produce      json
// @Success      200  {array}  sample
// @Router       /samples [get]
func getsamples(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, samples)
}

// postsamples godoc
// @Summary      Create a new sample
// @Description  Add a new sample from the request body
// @Tags         samples
// @Accept       json
// @Produce      json
// @Param        sample  body      sample  true  "sample to add"
// @Success      201     {object}  sample
// @Failure      400     {object}  gin.H
// @Failure      409     {object}  gin.H
// @Router       /samples [post]
func postsamples(c *gin.Context) {
	var newsample sample
	if err := c.BindJSON(&newsample); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}
	// Check if the sample already exists
	for _, s := range samples {
		if s.Name == newsample.Name {
			c.IndentedJSON(http.StatusConflict, gin.H{"message": "sample with this name already exists"})
			return
		}
	}
	samples = append(samples, newsample)
	c.IndentedJSON(http.StatusCreated, newsample)
}

// getsampleByName godoc
// @Summary      Get sample by Name
// @Description  Get a single sample by its name
// @Tags         samples
// @Produce      json
// @Param        name  path      string  true  "sample Name"
// @Success      200   {object}  sample
// @Failure      404   {object}  gin.H
// @Router       /samples/{name} [get]
func getsampleByName(c *gin.Context) {
	name := c.Param("name")
	// Loop through the list of samples, looking for
	// a sample whose name value matches the parameter.
	for _, a := range samples {
		if a.Name == name {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "sample not found"})
}

// putsampleByName godoc
// @Summary      Update sample by Name
// @Description  Update an existing sample by its name
// @Tags         samples
// @Accept       json
// @Produce      json
// @Param        name    path      string  true  "sample Name"
// @Param        sample  body      sample  true  "Updated sample Data"
// @Success      200     {object}  sample
// @Failure      400     {object}  gin.H
// @Failure      404     {object}  gin.H
// @Router       /samples/{name} [put]
func putsampleByName(c *gin.Context) {
	name := c.Param("name")
	var updatedsample sample
	if err := c.BindJSON(&updatedsample); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}
	for i, a := range samples {
		if a.Name == name {
			samples[i] = updatedsample
			c.IndentedJSON(http.StatusOK, updatedsample)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "sample not found"})
}

// deletesampleByName godoc
// @Summary      Delete sample by Name
// @Description  Delete an existing sample by its name
// @Tags         samples
// @Produce      json
// @Param        name  path      string  true  "sample Name"
// @Success      200   {object}  gin.H  "sample deleted message"
// @Failure      404   {object}  gin.H  "sample not found message"
// @Router       /samples/{name} [delete]
func deletesampleByName(c *gin.Context) {
	name := c.Param("name")
	// Loop through the list of samples, looking for
	// a sample whose name value matches the parameter.
	for i, a := range samples {
		if a.Name == name {
			// Remove the sample from the slice.
			samples = append(samples[:i], samples[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "sample deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "sample not found"})
}

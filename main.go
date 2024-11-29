package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type car struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Model  string `json:"model"`
	Run    int    `json:"run"`
	Owners byte   `json:"owners"`
}

var cars []car

type furniture struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	Height       int    `json:"height"`
	Width        int    `json:"width"`
	Length       int    `json:"length"`
}

var furnitures []furniture

type Flowerbase struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Quantity int       `json:"quantity"`
	Price    int       `json:"price"`
	Date     time.Time `json:"date"`
}

var Flowerbases []Flowerbase

func loadCarsFromFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &cars)
}

func loadfurnitureFromFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &furnitures)
}

func loadFlowerbaseFromFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &Flowerbases)
}

// сохраняет  в  JSON файл.
func saveCarsToFile(filename string) error {
	data, err := json.MarshalIndent(cars, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}

func savefurnitureToFile(filename string) error {
	data, err := json.MarshalIndent(furnitures, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}

func saveFlowerbaseToFile(filename string) error {
	data, err := json.MarshalIndent(Flowerbases, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}

// добавляет из JSON, полученного в теле запроса.
func postCars(c *gin.Context) {
	var newCar car

	// Вызываем BindJSON, чтобы привязать полученный JSON
	if err := c.BindJSON(&newCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Добавляем в срез.
	cars = append(cars, newCar)

	// Сохраняет  в файл
	if err := saveCarsToFile("cars.json"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newCar)
}

func postfurnitures(c *gin.Context) {
	var newfurniture furniture

	// Вызываем BindJSON, чтобы привязать полученный JSON
	if err := c.BindJSON(&newfurniture); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Добавляем в срез.
	furnitures = append(furnitures, newfurniture)

	// Сохраняет  в файл
	if err := savefurnitureToFile("furniture.json"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newfurniture)
}

func postFlowerbase(c *gin.Context) {
	var newFlowerbase Flowerbase

	// Вызываем BindJSON, чтобы привязать полученный JSON
	if err := c.BindJSON(&newFlowerbase); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Добавляем в срез.
	Flowerbases = append(Flowerbases, newFlowerbase)

	// Сохраняет в файл
	if err := saveFlowerbaseToFile("Flowerbase.json"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newFlowerbase)
}

//	находим значение id который совпадает с параметром id
//
// , отправленным клиентом, и возвращает в качестве ответа.
func getCarByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range cars {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "car not found"})
}

func getfurnitureByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range furnitures {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "furniture not found"})
}

func getFlowerbaseByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range Flowerbases {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Flowerbase not found"})
}

// обновляем поля заданные клиентом id.
func updateCarByID(c *gin.Context) {
	id := c.Param("id")

	for i, a := range cars {
		if a.ID == id {
			var updatedCar car
			if err := c.BindJSON(&updatedCar); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			cars[i] = updatedCar

			if err := saveCarsToFile("cars.json"); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.IndentedJSON(http.StatusOK, cars[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "car not found"})
}

func updatefurnitureByID(c *gin.Context) {
	id := c.Param("id")

	for i, a := range furnitures {
		if a.ID == id {
			var updatedfurniture furniture
			if err := c.BindJSON(&updatedfurniture); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			furnitures[i] = updatedfurniture

			if err := savefurnitureToFile("furniture.json"); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.IndentedJSON(http.StatusOK, furnitures[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "furniture not found"})
}

func updateFlowerbaseByID(c *gin.Context) {
	id := c.Param("id")

	for i, a := range Flowerbases {
		if a.ID == id {
			var updatedFlowerbase Flowerbase
			if err := c.BindJSON(&updatedFlowerbase); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			Flowerbases[i] = updatedFlowerbase

			if err := saveFlowerbaseToFile("Flowerbase.json"); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.IndentedJSON(http.StatusOK, Flowerbases[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Flowerbase not found"})
}

// частично обновляем поля с заданным ID.
func patchCarByID(c *gin.Context) {
	id := c.Param("id")

	for i, a := range cars {
		if a.ID == id {

			var patchCar car
			if err := c.BindJSON(&patchCar); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if patchCar.Name != "" {
				cars[i].Name = patchCar.Name
			}
			if patchCar.Model != "" {
				cars[i].Model = patchCar.Model
			}
			if patchCar.Run != 0 {
				cars[i].Run = patchCar.Run
			}
			if patchCar.Owners != 0 {
				cars[i].Owners = patchCar.Owners
			}

			if err := saveCarsToFile("cars.json"); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.IndentedJSON(http.StatusOK, cars[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "car not found"})
}

func patchfurnitureByID(c *gin.Context) {
	id := c.Param("id")

	for i, a := range furnitures {
		if a.ID == id {

			var patchfurniture furniture
			if err := c.BindJSON(&patchfurniture); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if patchfurniture.Name != "" {
				furnitures[i].Name = patchfurniture.Name
			}
			if patchfurniture.Manufacturer != "" {
				furnitures[i].Manufacturer = patchfurniture.Manufacturer
			}
			if patchfurniture.Height != 0 {
				furnitures[i].Height = patchfurniture.Height
			}
			if patchfurniture.Width != 0 {
				furnitures[i].Width = patchfurniture.Width
			}
			if patchfurniture.Length != 0 {
				furnitures[i].Length = patchfurniture.Length
			}

			if err := savefurnitureToFile("furniture.json"); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.IndentedJSON(http.StatusOK, furnitures[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "furniture not found"})
}

func patchFlowerbaseByID(c *gin.Context) {
	id := c.Param("id")

	for i, a := range Flowerbases {
		if a.ID == id {

			var patchFlowerbase Flowerbase
			if err := c.BindJSON(&patchFlowerbase); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if patchFlowerbase.Name != "" {
				Flowerbases[i].Name = patchFlowerbase.Name
			}
			if patchFlowerbase.Quantity != 0 {
				Flowerbases[i].Quantity = patchFlowerbase.Quantity
			}
			if patchFlowerbase.Price != 0 {
				Flowerbases[i].Price = patchFlowerbase.Price
			}
			Flowerbases[i].Date = patchFlowerbase.Date

			if err := saveFlowerbaseToFile("Flowerbase.json"); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.IndentedJSON(http.StatusOK, cars[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Flowerbase not found"})
}

// удаляем с заданным ID.
func deleteCarByID(c *gin.Context) {

	id := c.Param("id")
	for i, a := range cars {
		if a.ID == id {

			cars = append(cars[:i], cars[i+1:]...)

			if err := saveCarsToFile("cars.json"); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.Status(http.StatusNoContent) // Return 204 No Content
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "car not found"})
}

func deletefurnitureByID(c *gin.Context) {

	id := c.Param("id")
	for i, a := range furnitures {
		if a.ID == id {

			furnitures = append(furnitures[:i], furnitures[i+1:]...)

			if err := savefurnitureToFile("furniture.json"); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.Status(http.StatusNoContent) // Return 204 No Content
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "furniture not found"})
}

func deleteFlowerbaseByID(c *gin.Context) {

	id := c.Param("id")
	for i, a := range Flowerbases {
		if a.ID == id {

			Flowerbases = append(Flowerbases[:i], Flowerbases[i+1:]...)

			if err := saveFlowerbaseToFile("Flowerbase.json"); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.Status(http.StatusNoContent) // Return 204 No Content
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Flowerbase not found"})
}

func main() {
	// Загружаем из файла при запуске.
	if err := loadCarsFromFile("cars.json"); err != nil {
		// Если файл не существует, создаем новый по умолчанию
		if os.IsNotExist(err) {
			fmt.Println("File not found, creating a new one with default models.")
			cars = []car{
				{ID: "1", Name: "JAECOO", Model: "J8", Run: 10000, Owners: 0},
				{ID: "2", Name: "Lada", Model: "Granta", Run: 96000, Owners: 1},
				{ID: "3", Name: "Chery", Model: "Tiggo 4 Pro", Run: 117000, Owners: 2}}
		}
		if err := saveCarsToFile("cars.json"); err != nil {
			fmt.Println("Error saving cars to file:", err)
		}
	} else {
		fmt.Println("Error loading cars from file:", err)
	}

	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/cars", getCars)
	router.POST("/cars", postCars)
	router.GET("/cars/:id", getCarByID)
	router.PUT("/cars/:id", updateCarByID)
	router.PATCH("/cars/:id", patchCarByID)
	router.DELETE("/cars/:id", deleteCarByID)

	if err := loadfurnitureFromFile("furniture.json"); err != nil {
		// Если файл не существует, создаем новый по умолчанию
		if os.IsNotExist(err) {
			fmt.Println("File not found, creating a new one with default models.")
			furnitures = []furniture{
				{ID: "1", Name: "Стул", Manufacturer: "МебельТорг", Height: 90, Width: 50, Length: 50},
				{ID: "2", Name: "Стол", Manufacturer: "МебельТорг", Height: 90, Width: 100, Length: 100},
				{ID: "3", Name: "Кровать", Manufacturer: "МебельТорг", Height: 60, Width: 180, Length: 210}}
		}
		if err := savefurnitureToFile("furniture.json"); err != nil {
			fmt.Println("Error saving furniture to file:", err)
		}
	} else {
		fmt.Println("Error loading furniture from file:", err)
	}

	router.GET("/furniture", getfurniture)
	router.POST("/furniture", postfurnitures)
	router.GET("/furniture/:id", getfurnitureByID)
	router.PUT("/furniture/:id", updatefurnitureByID)
	router.PATCH("/furniture/:id", patchfurnitureByID)
	router.DELETE("/furniture/:id", deletefurnitureByID)

	if err := loadFlowerbaseFromFile("Flowerbase.json"); err != nil {
		// Если файл не существует, создаем новый по умолчанию
		if os.IsNotExist(err) {
			fmt.Println("File not found, creating a new one with default models.")
			Flowerbases = []Flowerbase{
				{ID: "1", Name: "Розы", Quantity: 100, Price: 10000, Date: time.Now()},
				{ID: "2", Name: "Хризантемы", Quantity: 233, Price: 500, Date: time.Now()},
				{ID: "3", Name: "Ромашки", Quantity: 444, Price: 1033, Date: time.Now()}}
		}
		if err := saveFlowerbaseToFile("Flowerbase.json"); err != nil {
			fmt.Println("Error saving Flowerbase to file:", err)
		}
	} else {
		fmt.Println("Error loading Flowerbase from file:", err)
	}

	router.GET("/Flowerbase", getFlowerbase)
	router.POST("/Flowerbase", postFlowerbase)
	router.GET("/Flowerbase/:id", getFlowerbaseByID)
	router.PUT("/Flowerbase/:id", updateFlowerbaseByID)
	router.PATCH("/Flowerbase/:id", patchFlowerbaseByID)
	router.DELETE("/Flowerbase/:id", deleteFlowerbaseByID)
	router.Run("localhost:8080")
}

// выдает список в формате JSON.
func getCars(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cars)
}

func getfurniture(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, furnitures)
}

func getFlowerbase(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Flowerbases)
}

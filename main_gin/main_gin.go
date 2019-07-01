package main_gin

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"personal/guitar_collection/domain"
)

var db *gorm.DB

func main() {
	router := gin.Default()

	v1 := router.Group("/v1/guitar")
	{
		v1.POST("/", createGuitar)
		v1.GET("/", fetchAllGuitars)
		v1.GET("/:id", fetchSingleGuitar)
		v1.PUT("/:id", updateGuitar)
		v1.DELETE("/:id", deleteGuitar)
	}

	err := router.Run()
	if err != nil {
		panic("failed to run the router")
	}
}

func init() {
	// Open DB connection
	var err error
	db, err = gorm.Open("mysql", "root:12345elsen@/guitar_collection?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	db.AutoMigrate(&domain.Guitar{})
}

func createGuitar(c *gin.Context) {
	var newGuitar domain.Guitar
	err := c.BindJSON(&newGuitar)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	} else {
		result := db.Save(&newGuitar)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": result.Error.Error(),
				"data":    nil,
			})
		} else {
			c.JSON(http.StatusCreated, gin.H{
				"status":  "success",
				"message": "guitar is successfully created!",
				"data":    nil,
			})
		}
	}
}

func fetchAllGuitars(c *gin.Context) {
	var guitars [] domain.Guitar
	var viewGuitars [] domain.ViewGuitar

	result := db.Find(&guitars)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": result.Error.Error(),
			"data":    nil,
		})
	} else {
		if len(guitars) <= 0 {
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "no guitar found",
				"data":    viewGuitars,
			})
		} else {
			for _, item := range guitars {
				viewGuitars = append(viewGuitars, domain.ViewGuitar{
					Id:    item.ID,
					Type:  item.Type,
					Brand: item.Brand,
					Price: item.Price,
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "success in fetching all guitars",
				"data":    viewGuitars,
			})
		}
	}
}

func fetchSingleGuitar(c *gin.Context) {
	var savedGuitar domain.Guitar
	var viewGuitar domain.ViewGuitar
	guitarId := c.Param("id")
	result := db.First(&savedGuitar, guitarId)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": result.Error.Error(),
			"data":    nil,
		})
	} else {
		if savedGuitar.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  "fail",
				"message": "no guitar found with that id",
				"data":    nil,
			})
		} else {
			viewGuitar = domain.ViewGuitar{
				Id:    savedGuitar.ID,
				Type:  savedGuitar.Type,
				Brand: savedGuitar.Brand,
				Price: savedGuitar.Price,
			}
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "success in fetching all guitars",
				"data":    viewGuitar,
			})
		}
	}
}

func updateGuitar(c *gin.Context) {
	var savedGuitar domain.Guitar
	guitarId := c.Param("id")
	result := db.First(&savedGuitar, guitarId)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": result.Error.Error(),
			"data":    nil,
		})
	} else {
		if savedGuitar.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  "fail",
				"message": "no guitar found with that id",
				"data":    nil,
			})
		} else {
			var newGuitar domain.Guitar
			err := c.BindJSON(&newGuitar)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  "error",
					"message": err.Error(),
					"data":    nil,
				})
			} else {
				result := db.Model(&savedGuitar).Update(
					"brand", newGuitar.Brand,
					"type", newGuitar.Type,
					"price", newGuitar.Price,
				)
				if result.Error != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"status":  "error",
						"message": result.Error.Error(),
						"data":    nil,
					})
				} else {
					var viewGuitar domain.ViewGuitar
					viewGuitar = domain.ViewGuitar{
						Id:    savedGuitar.ID,
						Type:  savedGuitar.Type,
						Brand: savedGuitar.Brand,
						Price: savedGuitar.Price,
					}
					c.JSON(http.StatusOK, gin.H{
						"status":  "success",
						"message": "success in updating guitar",
						"data":    viewGuitar,
					})
				}
			}
		}
	}
}

func deleteGuitar(c *gin.Context) {
	var savedGuitar domain.Guitar
	guitarId := c.Param("id")

	result := db.First(&savedGuitar, guitarId)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": result.Error.Error(),
			"data":    nil,
		})
	} else {
		if savedGuitar.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  "fail",
				"message": "no guitar found with that id",
				"data":    nil,
			})
		} else {
			result := db.Delete(&savedGuitar)
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  "error",
					"message": result.Error.Error(),
					"data":    nil,
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status":  "success",
					"message": "success in deleting guitar",
					"data":    nil,
				})
			}
		}
	}
}

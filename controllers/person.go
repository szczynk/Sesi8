package controllers

import (
	"net/http"

	"github.com/szczynk/Sesi8/structs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InDB struct {
	db *gorm.DB
}

func NewPersonInDB(db *gorm.DB) *InDB {
	return &InDB{db}
}

// get all persons
func (idb *InDB) GetPersons(c *gin.Context) {
	var (
		persons []structs.Person
		result  gin.H
	)

	idb.db.Find(&persons)

	if len(persons) <= 0 {
		result = gin.H{
			"result": "nil",
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": persons,
			"count":  len(persons),
		}
	}

	c.JSON(http.StatusOK, result)
}

// create person
func (idb *InDB) CreatePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	person.First_Name = first_name
	person.Last_Name = last_name

	idb.db.Create(&person)
	result = gin.H{
		"result": person,
	}

	c.JSON(http.StatusOK, result)
}

// get one person by id
func (idb *InDB) GetPerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	id := c.Param("id")
	err := idb.db.Where("id = ?", id).First(&person).Error

	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// update person by id as query
func (idb *InDB) UpdatePerson(c *gin.Context) {
	id := c.Query("id")
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	var (
		person    structs.Person
		newPerson structs.Person
		result    gin.H
	)

	err := idb.db.Where("id = ?", id).First(&person).Error

	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	newPerson.First_Name = first_name
	newPerson.Last_Name = last_name
	err = idb.db.Model(&person).Updates(newPerson).Error
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": "succesfully updated data",
		}
	}

	c.JSON(http.StatusOK, result)
}

// delete person by id
func (idb *InDB) DeletePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	id := c.Param("id")
	err := idb.db.Where("id = ?", id).First(&person).Error

	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	err = idb.db.Delete(&person).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "data deleted succesfully",
		}
	}

	c.JSON(http.StatusOK, result)
}

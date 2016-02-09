package handler

import (
	"fmt"
	"net/http"
	"squarecode/dynastymasra/rhino/model"
	"squarecode/dynastymasra/rhino/thrift/service"
	"squarecode/dynastymasra/rhino/util"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/gin-gonic/gin"
)

/**
 * Created by Dynastymasra
 * Name     : Dimas Ragil T
 * Email    : dynastymasra@gmail.com
 * LinkedIn : http://www.linkedin.com/in/dynastymasra
 * Github   : https://github.com/dynastymasra
 * Mobile and Backend Developer
 */

func GetAll(c *gin.Context) {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	result, err := service.GetAll(transportFactory, protocolFactory)

	if err != nil {
		m := fmt.Sprintf("%v", err)
		c.JSON(http.StatusExpectationFailed, util.FailResponse(m))
	} else {
		parse := make([]interface{}, len(result))
		for i := range result {
			parse[i] = model.PersonResponse(result[i])
		}
		c.JSON(http.StatusCreated, util.ListResponse("success", "", parse))
	}
}

func GetById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data":   "GetById",
		"status": "success",
	})
}

func Create(c *gin.Context) {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())

	var json model.Person
	c.BindJSON(&json)
	person := model.PersonMapping(json.Identity, json.Firstname, json.Middlename, json.Lastname, json.Address, json.Age, json.Active)
	result, err := service.Create(transportFactory, protocolFactory, person)
	if err != nil {
		m := fmt.Sprintf("%v", err)
		c.JSON(http.StatusExpectationFailed, util.FailResponse(m))
	} else {
		parse := map[string]interface{}{
			"person": model.PersonResponse(result),
		}
		c.JSON(http.StatusCreated, util.ObjectResponse("success", "", parse))
	}
}

func Update(c *gin.Context) {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())

	var json model.Person
	c.BindJSON(&json)
	person := model.PersonUpdate(json.ID, json.Identity, json.Firstname, json.Middlename, json.Lastname, json.Address, json.Age, json.Active)
	result, err := service.Update(transportFactory, protocolFactory, person)
	if err != nil {
		m := fmt.Sprintf("%v", err)
		c.JSON(http.StatusExpectationFailed, util.FailResponse(m))
	} else {
		parse := map[string]interface{}{
			"person": model.PersonResponse(result),
		}
		c.JSON(http.StatusCreated, util.ObjectResponse("success", "", parse))
	}
}

func Delete(c *gin.Context) {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())

	var json model.Person
	c.BindJSON(&json)
	err := service.Delete(transportFactory, protocolFactory, json.ID)
	if err != nil {
		m := fmt.Sprintf("%v", err)
		c.JSON(http.StatusExpectationFailed, util.FailResponse(m))
	} else {
		parse := map[string]interface{}{
			"person": json,
		}
		c.JSON(http.StatusCreated, util.ObjectResponse("success", "", parse))
	}
}

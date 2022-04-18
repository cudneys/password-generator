package api

import (
	"fmt"
	config "github.com/cudneys/password-generator/config"
	models "github.com/cudneys/password-generator/models"
	"github.com/gin-gonic/gin"
	"github.com/sethvargo/go-password/password"
	"net/http"
	"strconv"
)

func getIntQP(c *gin.Context, key string, defaultValue string) (int, error) {
	length := c.DefaultQuery(key, defaultValue)
	lNum, err := strconv.Atoi(length)
	if err != nil {
		return 0, err
	}
	return lNum, nil
}

func getLength(c *gin.Context) (int, error) {
	return getIntQP(c, "length", fmt.Sprint(config.GetMinLen()))
}

func getDigits(c *gin.Context) (int, error) {
	return getIntQP(c, "digits", "4")
}

func getSymbols(c *gin.Context) (int, error) {
	return getIntQP(c, "symbols", "4")
}

// @BasePath /v1
// passwordGenerator godoc
// @Summary Password Generator
// @Schemes
// @Description Generates somewhat secure-ish passwords
// @Param        length   query      int  false  "Password Length"
// @Param        digits   query      int  false  "Num Digits"
// @Param        symbols   query      int  false  "Num Symbols"
// @Produce json
// @Success 200 {object} models.Response
// @Router /generate [get]
func PasswordGenerator(c *gin.Context) {
	var reqParams models.RequestParams

	// Bind the request params
	err := c.ShouldBind(&reqParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Request: reqParams, Password: "", Error: fmt.Sprintf("%s", err)})
		return
	}

	// Validate the request params and set default values
	if err = reqParams.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Request: reqParams, Password: "", Error: fmt.Sprintf("%s", err)})
		return
	}

	// Generate the password
	password, err := password.Generate(reqParams.Length, reqParams.Digits, reqParams.Symbols, false, true)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Request: reqParams, Password: "", Error: fmt.Sprintf("%s", err)})
		return
	}

	// Build the response and send it
	resp := models.Response{Request: reqParams, Password: password}
	c.JSON(http.StatusOK, resp)
}

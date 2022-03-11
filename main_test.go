package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/tjaime/go-gin-api-rest/controller"
)

func SetupRotasTest() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestVerificaSaudacaoStatusOk(t *testing.T) {
	r := SetupRotasTest()
	r.GET("/:nome", controller.Saudacao)

	req, _ := http.NewRequest("GET", "/thiago", nil)
	rsp := httptest.NewRecorder()

	r.ServeHTTP(rsp, req)
	if rsp.Code != http.StatusOK {
		t.Fatalf("Actual: %d - Expected: %d",
			rsp.Code, http.StatusOK)
	}

}

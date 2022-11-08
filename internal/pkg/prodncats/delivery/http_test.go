package delivery

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"mb_api/internal/pkg/models"
	"mb_api/internal/pkg/network"
	"mb_api/internal/pkg/prodncats"
	"mb_api/internal/pkg/prodncats/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	testProducts = models.Products{
		ProductName: "name",
		ProductCategory: "category",
	}
	testCategories = models.Categories{
		CategoryName: "name",
		CategoryProduct: "product",
	}
	testPairs = models.Pairs{
		ProductName: "pname",
		CategoryName: "cname",
	}
)

func getTestDelivery(mockUC *mocks.MockUseCase) prodncats.Delivery {
	return &prodncatsDelivery{UseCase:mockUC}
}

func TestProdncatsDelivery_GetProducts_Incorrect(t *testing.T) {
	// Create mock
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUC := mocks.NewMockUseCase(mockCtrl)
	ed := getTestDelivery(mockUC)

	body, _ := json.Marshal("")
	req, err := http.NewRequest("GET", "/api/getprods", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
		return
	}
	rr := httptest.NewRecorder()
	var ps map[string]string

	mockUC.EXPECT().InitProducts(new([]models.Products)).Return(http.StatusInternalServerError, errors.New("error"))
	ed.GetProducts(rr, req, ps)

	msg, err := network.DecodeToMsg(rr.Body)
	if err != nil {
		t.Fail()
		return
	}

	assert.Equal(t, http.StatusInternalServerError, msg.Status)
	assert.Equal(t, "error", msg.Message)
}

func TestProdncatsDelivery_GetProducts_Correct(t *testing.T) {
	// Create mock
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUC := mocks.NewMockUseCase(mockCtrl)
	ed := getTestDelivery(mockUC)

	body, _ := json.Marshal("")
	req, err := http.NewRequest("GET", "/api/getprods", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
		return
	}
	rr := httptest.NewRecorder()
	var ps map[string]string

	mockUC.EXPECT().InitProducts(new([]models.Products))
	ed.GetProducts(rr, req, ps)

	msg, err := network.DecodeToMsg(rr.Body)
	if err != nil {
		t.Fail()
		return
	}

	assert.Equal(t, 0, msg.Status)
}

func TestProdncatsDelivery_GetCategories_Incorrect(t *testing.T) {
	// Create mock
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUC := mocks.NewMockUseCase(mockCtrl)
	ed := getTestDelivery(mockUC)

	body, _ := json.Marshal("")
	req, err := http.NewRequest("GET", "/api/getcats", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
		return
	}
	rr := httptest.NewRecorder()
	var ps map[string]string

	mockUC.EXPECT().InitCategories(new([]models.Categories)).Return(http.StatusInternalServerError, errors.New("error"))
	ed.GetCategories(rr, req, ps)

	msg, err := network.DecodeToMsg(rr.Body)
	if err != nil {
		t.Fail()
		return
	}

	assert.Equal(t, http.StatusInternalServerError, msg.Status)
	assert.Equal(t, "error", msg.Message)
}

func TestProdncatsDelivery_GetCategories_Correct(t *testing.T) {
	// Create mock
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUC := mocks.NewMockUseCase(mockCtrl)
	ed := getTestDelivery(mockUC)

	body, _ := json.Marshal("")
	req, err := http.NewRequest("GET", "/api/getcats", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
		return
	}
	rr := httptest.NewRecorder()
	var ps map[string]string

	mockUC.EXPECT().InitCategories(new([]models.Categories))
	ed.GetCategories(rr, req, ps)

	msg, err := network.DecodeToMsg(rr.Body)
	if err != nil {
		t.Fail()
		return
	}

	assert.Equal(t, 0, msg.Status)
}

func TestProdncatsDelivery_GetPairs_Incorrect(t *testing.T) {
	// Create mock
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUC := mocks.NewMockUseCase(mockCtrl)
	ed := getTestDelivery(mockUC)

	body, _ := json.Marshal("")
	req, err := http.NewRequest("GET", "/api/getpairs", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
		return
	}
	rr := httptest.NewRecorder()
	var ps map[string]string

	mockUC.EXPECT().InitPairs(new([]models.Pairs)).Return(http.StatusInternalServerError, errors.New("error"))
	ed.GetPairs(rr, req, ps)

	msg, err := network.DecodeToMsg(rr.Body)
	if err != nil {
		t.Fail()
		return
	}

	assert.Equal(t, http.StatusInternalServerError, msg.Status)
	assert.Equal(t, "error", msg.Message)
}

func TestProdncatsDelivery_GetPairs_Correct(t *testing.T) {
	// Create mock
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUC := mocks.NewMockUseCase(mockCtrl)
	ed := getTestDelivery(mockUC)

	body, _ := json.Marshal("")
	req, err := http.NewRequest("GET", "/api/getpairs", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
		return
	}
	rr := httptest.NewRecorder()
	var ps map[string]string

	mockUC.EXPECT().InitPairs(new([]models.Pairs))
	ed.GetPairs(rr, req, ps)

	msg, err := network.DecodeToMsg(rr.Body)
	if err != nil {
		t.Fail()
		return
	}

	assert.Equal(t, 0, msg.Status)
}
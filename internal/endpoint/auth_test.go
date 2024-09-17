package endpoint

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/lavatee/shop_api_gateway/internal/service"
	mock_service "github.com/lavatee/shop_api_gateway/internal/service/mocks"
	"github.com/magiconair/properties/assert"
)

type TestUser struct {
	Name     string
	Email    string
	Password string
}

func TestEndpoint_SignUp(t *testing.T) {
	type MockBehavior func(s *mock_service.MockAuth, name string, email string, password string)
	testTable := []struct {
		Name               string
		InputBody          string
		InputUser          TestUser
		ExpectedStatusCode int
		ExpectedResponse   string
		MockBehavior       MockBehavior
	}{
		{
			Name:               "ok",
			InputBody:          `{"name": "sasha", "email": "sasha@gmail.com", "password": "qwerty"}`,
			InputUser:          TestUser{Name: "sasha", Email: "sasha@gmail.com", Password: "qwerty"},
			ExpectedStatusCode: 200,
			ExpectedResponse:   `{"id":1}`,
			MockBehavior: func(s *mock_service.MockAuth, name, email, password string) {
				s.EXPECT().SignUp(name, email, password).Return(1, nil)
			},
		},
	}
	for _, test := range testTable {
		t.Run(test.Name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			auth := mock_service.NewMockAuth(c)
			test.MockBehavior(auth, test.InputUser.Name, test.InputUser.Email, test.InputUser.Password)
			svc := &service.Service{Auth: auth}
			end := &Endpoint{Services: svc}
			router := gin.New()
			router.POST("/signup", end.SignUp)
			recorder := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/signup", bytes.NewBufferString(test.InputBody))
			router.ServeHTTP(recorder, req)
			assert.Equal(t, recorder.Code, test.ExpectedStatusCode)
			assert.Equal(t, recorder.Body.String(), test.ExpectedResponse)
		})
	}
}

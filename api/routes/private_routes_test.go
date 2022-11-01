package routes

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/mahadevans87/short-url/helpers"
	"github.com/stretchr/testify/assert"
)

func TestPrivateRoutes(t *testing.T) {
	// Load .env.test file from the root folder.
	if err := godotenv.Load("../.env"); err != nil {
		panic(err)
	}

	// Create a sample data string.
	dataString := `{"url":"https://www.google.com","short":"","expiry":24}`

	// Create access token.
	token, err := helpers.GenerateNewAccessToken()
	if err != nil {
		panic(err)
	}

	// Define a structure for specifying input and output data of a single test case.
	tests := []struct {
		description   string
		route         string // input route
		method        string // input method
		tokenString   string // input token
		body          io.Reader
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "post without JWT and body",
			route:         "/api/v1/shorten",
			method:        "POST",
			tokenString:   "",
			body:          nil,
			expectedError: false,
			expectedCode:  400,
		},
		{
			description:   "post without right credentials",
			route:         "/api/v1/shorten",
			method:        "POST",
			tokenString:   "Bearer " + token,
			body:          strings.NewReader(dataString),
			expectedError: false,
			expectedCode:  403,
		},
		{
			description:   "post correctly",
			route:         "/api/v1/shorten",
			method:        "POST",
			tokenString:   "Bearer " + token,
			body:          strings.NewReader(dataString),
			expectedError: false,
			expectedCode:  200,
		},
	}

	// Define a new Fiber app.
	app := fiber.New()

	// Define routes.
	PrivateRoutes(app)

	// Iterate through test single test cases
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			// Create a new http request with the route from the test case.
			req := httptest.NewRequest(test.method, test.route, test.body)
			req.Header.Set("Authorization", test.tokenString)
			req.Header.Set("Content-Type", "application/json")

			// Perform the request plain with the app.
			resp, err := app.Test(req, -1) // the -1 disables request latency

			// Verify, that no error occurred, that is not expected
			assert.Equalf(t, test.expectedError, err != nil, test.description)

			// Verify, if the status code is as expected.
			assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
		})
	}
}

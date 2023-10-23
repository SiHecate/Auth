package controllers

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	type wanted struct {
		StatusCode   int
		expectedKeys []string
	}

	app := fiber.New()

	app.Post("/login", func(c *fiber.Ctx) error {
		var data map[string]string

		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request data",
			})
		}
		if _, ok := data["password"]; !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errorPassword": "Missing 'password' field",
			})
		}
		if _, ok := data["email"]; !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errorEmail": "Missing 'email' field",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Login successful",
		})
	})

	tests := []struct {
		name           string
		requestPayload string
		expected       wanted
	}{
		{
			name:           "Valid Registration (Login)",
			requestPayload: `{"email": "user123@user.com", "password": "user123"}`,
			expected: wanted{
				StatusCode: 200,
				expectedKeys: []string{
					"message",
				},
			},
		},
		{
			name:           "Invalid Registration (Missing password)",
			requestPayload: `{"email": "user123@user.com"}`,
			expected: wanted{
				StatusCode: 400,
				expectedKeys: []string{
					"errorPassword",
				},
			},
		},
		{
			name:           "Invalid Registration (Missing email)",
			requestPayload: `{"password": "user123"}`,
			expected: wanted{
				StatusCode: 400,
				expectedKeys: []string{
					"errorEmail",
				},
			},
		},
		{
			name:           "Invalid Registration (Login)",
			requestPayload: `{"",}`,
			expected: wanted{
				StatusCode: 400,
				expectedKeys: []string{
					"error",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/login", strings.NewReader(test.requestPayload))
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()

			assert.Equal(t, test.expected.StatusCode, resp.StatusCode)

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}
			var response map[string]interface{}
			err = json.Unmarshal(body, &response)
			if err != nil {
				t.Fatal(err)
			}

			for _, key := range test.expected.expectedKeys {
				assert.Contains(t, response, key, "Expected JSON key '%s' not found in the response", key)
			}
		})
	}
}

func TestRegister(t *testing.T) {

	type wanted struct {
		StatusCode   int
		expectedKeys []string
	}

	app := fiber.New()

	app.Post("/signup", func(c *fiber.Ctx) error {
		var data map[string]string

		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request data",
			})
		}
		if _, ok := data["password"]; !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errorPassword": "Missing 'password' field",
			})
		}
		if _, ok := data["password_confirm"]; !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errorPassword_confirm": "Missing 'password_confirm' field",
			})
		}
		if _, ok := data["name"]; !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errorName": "Missing 'name' field",
			})
		}
		if _, ok := data["email"]; !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errorEmail": "Missing 'email' field",
			})
		}
		if data["password"] != data["password_confirm"] {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errorMatch": "Password and Password_confirm does not match",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Registration successful",
		})
	})

	tests := []struct {
		name           string
		requestPayload string
		expected       wanted
	}{
		{
			name:           "Valid Registration (Register)",
			requestPayload: `{"name": "user123", "email": "user123@user.com", "password": "user123", "password_confirm": "user123"}`,
			expected: wanted{
				StatusCode: 200,
				expectedKeys: []string{
					"message",
				},
			},
		},
		{
			name:           "Unvalid Registration (Register)",
			requestPayload: `{""}`,
			expected: wanted{
				StatusCode: 400,
				expectedKeys: []string{
					"error",
				},
			},
		},
		{
			name:           "Invalid Registration (Missing Password)",
			requestPayload: `{"name": "user123", "email": "user123@user.com","password_confirm": "user123"}`,
			expected: wanted{
				StatusCode: 400,
				expectedKeys: []string{
					"errorPassword",
				},
			},
		},
		{
			name:           "Invalid Registration (Missing Confirm Password)",
			requestPayload: `{"name": "user123", "email": "user123@user.com", "password": "user123"}`,
			expected: wanted{
				StatusCode: 400,
				expectedKeys: []string{
					"errorPassword_confirm",
				},
			},
		},
		{
			name:           "Invalid Registration (Match Password)",
			requestPayload: `{"name": "user123", "email": "user123@user.com", "password": "user123","password_confirm": "user321"}`,
			expected: wanted{
				StatusCode: 400,
				expectedKeys: []string{
					"errorMatch",
				},
			},
		},
		{
			name:           "Invalid Registration (Missing Username)",
			requestPayload: `{"email": "user123@user.com", "password": "user123", "password_confirm": "user123"}`,
			expected: wanted{
				StatusCode: 400,
				expectedKeys: []string{
					"errorName",
				},
			},
		},
		{
			name:           "Invalid Registration (Missing Email)",
			requestPayload: `{"name": "user123", "password": "user123", "password_confirm": "user123"}`,
			expected: wanted{
				StatusCode: 400,
				expectedKeys: []string{
					"errorEmail",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/signup", strings.NewReader(test.requestPayload))
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()

			assert.Equal(t, test.expected.StatusCode, resp.StatusCode)

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}
			var response map[string]interface{}
			err = json.Unmarshal(body, &response)
			if err != nil {
				t.Fatal(err)
			}

			for _, key := range test.expected.expectedKeys {
				assert.Contains(t, response, key, "Expected JSON key '%s' not found in the response", key)
			}
		})
	}
}

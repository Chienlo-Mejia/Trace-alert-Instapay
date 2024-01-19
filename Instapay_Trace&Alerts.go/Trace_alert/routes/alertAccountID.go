package routes

import (
	"instapay/db"
	"instapay/model"

	"github.com/gofiber/fiber/v2"
)

var requestCount int

const allowedRate = 5

func checkRateLimits() bool {
	requestCount++
	if requestCount > allowedRate {
		requestCount = 0
		return true
	}
	return false
}
func GetAccInfo(c *fiber.Ctx) error {

	if checkRateLimits() {
		return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Gateway",
						"ReasonCode":  "RATE_LIMIT_EXCEEDED",
						"Description": "You have exceeded the service rate limit. Maximum allowed: ${rate_limit.output} TPS",
						"Recoverable": true,
						"Details":     nil,
					},
				},
			},
		})
	} else if c.Method() != fiber.MethodPost {
		return c.Status(fiber.StatusMethodNotAllowed).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "TRACE_FINANCIAL_CRIME",
						"ReasonCode":  "METHOD_NOT_ALLOWED",
						"Description": "Only POST method allowed",
						"Recoverable": false,
						"Details":     nil,
					},
				},
			},
		})
	}
	// Parse the request body
	info := &model.AccountAlertRequest1{}
	resp := &model.ResponseInfo1{}

	if err := c.BodyParser(&info); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "ALERT_FINANCIAL_CRIME",
						"ReasonCode":  "BAD_REQUEST",
						"Description": "We could not handle your request",
						"Recoverable": false,
						"Details":     "The request contains a bad payload",
					},
				},
			},
		})
	} else if fetchErr := db.DB.Debug().Raw(`SELECT * FROM response_info WHERE account_alert_id = ?`, info.AccountAlertID).Scan(&resp).Error; fetchErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "ALERT_FINANCIAL_CRIME",
						"ReasonCode":  "BAD_REQUEST",
						"Description": "We could not handle your request",
						"Recoverable": false,
						"Details":     "The request contains a bad payload",
					},
				},
			},
		})
	}

	// NetworkAlertID := resp.NetworkAlertID
	// AccountID := resp.AccountID

	// logMessage := fmt.Sprintf("%s: %s  %s ", resp.ID, NetworkAlertID, AccountID)
	// loggers.GetAccInfoLoggs(c.Path(), "folderName", logMessage, NetworkAlertID, AccountID)

	return c.JSON(resp)
}

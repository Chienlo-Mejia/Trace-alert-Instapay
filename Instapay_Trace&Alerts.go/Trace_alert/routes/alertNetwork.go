package routes

import (
	"fmt"
	"instapay/db"
	"instapay/model"

	"github.com/gofiber/fiber/v2"
)

// func Alertnetwork(c *fiber.Ctx) error {

// 	network := &model.NetworkBody{}

// 	if parsErr := c.BodyParser(&network); parsErr != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"Errors": fiber.Map{
// 				"Error": []fiber.Map{
// 					{
// 						"Source":      "ALERT_FINANCIAL_CRIME",
// 						"ReasonCode":  "BAD_REQUEST",
// 						"Description": "We could not handle your request",
// 						"Recoverable": false,
// 						"Details":     "The request contains a bad payload",
// 					},
// 				},
// 			},
// 		})
// 	} else if network == nil || (network.Since == "" && network.Limit == 0 && network.PaginationToken == "" && network.Filter == "" && !network.Include_all_alerts) {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"Errors": fiber.Map{
// 				"Error": []fiber.Map{
// 					{
// 						"Source":      "TRACE_FINANCIAL_CRIME",
// 						"ReasonCode":  "BAD_REQUEST",
// 						"Description": "We could not handle your request",
// 						"Recoverable": false,
// 						"Details":     "The request body is empty",
// 					},
// 				},
// 			},
// 		})
// 	}

// 	account := model.Transaction_Response{}
// 	transaction := &model.Transaction_Response{}
// 	networks := &model.NetworkResponse{}
// 	if err := db.DB.Debug().Raw(`SELECT * FROM public.trace_alert`).Scan(&account).Error; err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error()})
// 	}
// 	if err := db.DB.Debug().Raw(`SELECT * FROM public.trace_visualisation`).Scan(transaction).Error; err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error()})
// 	}
// 	if err := db.DB.Debug().Raw(`SELECT * FROM public.trace_accountalert`).Scan(networks).Error; err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error()})
// 	}
// 	// Initialize an empty slice to store the retrieved alerts
// 	var transactions []model.NetworkResponse

// 	// Construct the SQL query with date filtering and order by id
// 	query := `SELECT * FROM public.trace_accountalert WHERE 1 = 1`
// 	countQuery := "SELECT COUNT(*) FROM public.trace_accountalert WHERE 1=1"
// 	countFilteredQuery := "SELECT COUNT(*) FROM public.trace_accountalert WHERE 1=1"

// 	// Add conditions based on the request body
// 	if network.Since != "" {
// 		query += " AND DATE_TRUNC('day', decisiondate) = $1"
// 		countFilteredQuery += " AND DATE_TRUNC('day', decisiondate) = $1"
// 	}

// 	if network.Filter != "" {
// 		query += " AND " + network.Filter
// 		countFilteredQuery += " AND " + network.Filter
// 	}

// 	query += ` AND EXISTS (
// 		SELECT 1
// FROM public.trace_alert AS ad
// LEFT JOIN public.trace_accountalert AS ta ON ad.decisiondate = ta.decisiondate
// LEFT JOIN public.alert_data AS tv ON ad.decisiondate = tv.decisiondate
// WHERE ad.decisiondate = ta.decisiondate
// 	)`

// 	var totalCount int64
// 	if err := db.DB.Raw(countQuery).Count(&totalCount).Error; err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Error counting all transactions",
// 		})
// 	}

// 	// Execute the query and retrieve the count for filtered transactions
// 	var filteredCount int64
// 	if err := db.DB.Raw(countFilteredQuery, network.Since).Count(&filteredCount).Error; err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Error counting filtered transactions",
// 		})
// 	}
// 	err := db.DB.Debug().Raw(query, network.Since).Scan(&transactions).Error

// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"Errors": map[string]interface{}{
// 				"Error": []map[string]interface{}{
// 					{
// 						"Source":      "ALERT_FINANCIAL_CRIME",
// 						"ReasonCode":  "BAD_REQUEST",
// 						"Description": "We could not handle your request",
// 						"Recoverable": false,
// 						"Details":     "The request contains a bad payload",
// 					},
// 				},
// 			},
// 		})
// 	}

// 	// Check if no data was found for the given date
// 	if len(transactions) == 0 {
// 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
// 			"message": "No data found for the specified date",
// 		})
// 	}

// 	alert := fiber.Map{
// 		"id":        transaction.ID,
// 		"Time":      account.Time,
// 		"networkID": account.Networkalertid,
// 	}

// 	return c.JSON(fiber.Map{
// 		"alerts":            alert,
// 		"accountAlerts":     account,
// 		"transactionAlerts": transaction,
// 		"network":           networks,
// 		"totalCount":        totalCount,
// 		"filteredCount":     filteredCount,
// 	})

// }

func Alertnetwork(c *fiber.Ctx) error {
	network := &model.NetworkBody{}

	if parsErr := c.BodyParser(&network); parsErr != nil {
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
	} else if network == nil || (network.Since == "" && network.Limit == 0 && network.PaginationToken == "" && network.Filter == "" && !network.Include_all_alerts) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "TRACE_FINANCIAL_CRIME",
						"ReasonCode":  "BAD_REQUEST",
						"Description": "We could not handle your request",
						"Recoverable": false,
						"Details":     "The request body is empty",
					},
				},
			},
		})
	}

	// Initialize models
	account := model.TransactionAlert{}
	transaction := &model.Alertnetwork{}
	networks := &model.NetworkResponse{}

	// Fetch data from the database
	if err := db.DB.Debug().Raw(`SELECT * FROM public.transactionAlerts`).Scan(&account).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error()})
	}
	if err := db.DB.Debug().Raw(`SELECT * FROM public.trace_alerts`).Scan(transaction).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error()})
	}
	if err := db.DB.Debug().Raw(`SELECT * FROM public.trace_alert`).Scan(networks).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error()})
	}

	// Initialize an empty slice to store the retrieved alerts
	var transactions []model.NetworkResponse

	// Construct the SQL query with date filtering and order by id
	query := `SELECT * FROM public.transactionAlerts WHERE 1 = 1`
	countQuery := "SELECT COUNT(*) FROM public.transactionAlerts WHERE 1=1"
	countFilteredQuery := "SELECT COUNT(*) FROM public.transactionAlerts WHERE 1=1"

	// Add conditions based on the request body

	if network.Since != "" {
		query += " AND DATE_TRUNC('day', decisiondate) = DATE_TRUNC('day', $1::date)"
		countFilteredQuery += " AND DATE_TRUNC('day', decisiondate) = DATE_TRUNC('day', $1::date)"
	}

	if network.Filter != "" {
		query += " AND " + network.Filter
		countFilteredQuery += " AND " + network.Filter
	}
	// Include the EXISTS condition

	// Execute the query and retrieve the count for all transactions
	var totalCount int64
	if err := db.DB.Raw(countQuery).Count(&totalCount).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error counting all transactions",
		})
	}

	// Execute the query and retrieve the count for filtered transactions

	var filteredCount int64

	if network.Since != "" {
		if err := db.DB.Raw(countFilteredQuery, network.Since).Count(&filteredCount).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error counting filtered transactions"})
		}
	} else {
		if err := db.DB.Raw(countFilteredQuery).Count(&filteredCount).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error counting filtered transactions"})
		}
	}

	//Limit

	if network.Limit > 0 {
		if int64(network.Limit) > filteredCount {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Limit exceeds the total number of filtered transactions",
			})
		}
		query += fmt.Sprintf(" LIMIT %d", network.Limit)
	}

	// Execute the main query
	err := db.DB.Debug().Raw(query, network.Since).Scan(&transactions).Error

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Errors": map[string]interface{}{
				"Error": []map[string]interface{}{
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
	// Check if no data was found for the given date
	if len(transactions) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No data found for the specified date",
		})
	}

	// Prepare response
	alert := fiber.Map{
		"id":        transaction.ID,
		"Time":      account.Time,
		"networkID": account.Networkalertid,
	}

	return c.JSON(fiber.Map{
		"alerts":            alert,
		"accountAlerts":     account,
		"transactionAlerts": transaction,
		"network":           networks,
		"totalCount":        totalCount,
		"filteredCount":     filteredCount,
	})
}

//404
//200
//400
//405
//429

//401
//403

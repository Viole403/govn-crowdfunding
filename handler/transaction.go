package handler

import (
	"govn-crowdfunding/helper"
	"govn-crowdfunding/transaction"
	"govn-crowdfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignsTransaction(c *gin.Context) {
	var input transaction.GetCampaignTransactionsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil) // Update to use the correct response
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := h.service.GetTransactionByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil) // Update to use the correct response
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of campaign's transactions", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions)) // Update to use the correct response
	c.JSON(http.StatusOK, response)
}

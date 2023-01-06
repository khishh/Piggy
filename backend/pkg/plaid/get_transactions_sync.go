package plaid

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/plaid/plaid-go/v3/plaid"
)

type GetTransactionsSyncRequestBody struct {
	AccessToken string
}

func GetTransactionsSync(c *gin.Context) {

	fmt.Println("===== GetTransactionsSync ======")

	ctx := context.Background()
	client := c.Value("PlaidClient").(*plaid.APIClient)
	hasMore := true
	var cursor string

	var getTransactionsSyncRequestBody GetTransactionsSyncRequestBody

	if err := c.BindJSON(&getTransactionsSyncRequestBody); err != nil {
		c.JSON(http.StatusNotFound, "GetTransactionsSync failed with ln 22")
		return
	}

	accessToken := getTransactionsSyncRequestBody.AccessToken
	fmt.Println(accessToken)

	for hasMore {
		request := plaid.NewTransactionsSyncRequest(accessToken)
		request.SetCount(10) // for test

		res, _, err := client.PlaidApi.TransactionsSync(ctx).TransactionsSyncRequest(*request).Execute()

		if err != nil {
			c.JSON(http.StatusNotFound, "GetTransactionsSync failed during client.PlaidApi.TransactionsSync(ctx).TransactionsSyncRequest(*request).Execute()")
			return
		}

		// hasMore = res.GetHasMore()
		hasMore = false
		cursor = res.GetNextCursor()
		fmt.Println(res.HasMore)
		c.JSON(http.StatusNotFound, res)

	}

	fmt.Println(cursor)
}

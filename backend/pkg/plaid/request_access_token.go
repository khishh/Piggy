package plaid

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/plaid/plaid-go/v3/plaid"
)

type CreateAccessTokenRequestBody struct {
	PublicToken string
}

// obtain access token in exchange for public token
func CreateAccessToken(c *gin.Context) {

	client := c.Value("PlaidClient").(*plaid.APIClient)

	var createAccessTokenRequestBody CreateAccessTokenRequestBody
	ctx := context.Background()

	if err := c.BindJSON(&createAccessTokenRequestBody); err != nil {
		c.JSON(http.StatusNotFound, "CreateAccessToken Error: c.BindJSON(&createAccessTokenRequestBody) failed")
		return
	}

	publicToken := createAccessTokenRequestBody.PublicToken
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Printf("==== Public Token :%+v\n", string(body))
	fmt.Printf("==== Public Token :%+v\n", createAccessTokenRequestBody)
	fmt.Printf("==== Public Token : %s\n", publicToken)

	exchangePublicTokenResponse, _, err := client.PlaidApi.ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(
		*plaid.NewItemPublicTokenExchangeRequest(publicToken),
	).Execute()

	if err != nil {
		RenderError(c, err)
		return
	}

	accessToken := exchangePublicTokenResponse.GetAccessToken()
	itemID := exchangePublicTokenResponse.GetItemId()

	fmt.Println("public token: " + publicToken)
	fmt.Println("access token: " + accessToken)
	fmt.Println("item ID: " + itemID)

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
		"item_id":      itemID,
	})

}

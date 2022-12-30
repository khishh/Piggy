package plaid

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/plaid/plaid-go/v3/plaid"
)

type CreateLinkRequestBody struct {
	SubId string
}

func CreateLinkToken(c *gin.Context) {

	fmt.Println("==== CreateLinkToken ====")

	client := c.Value("PlaidClient").(*plaid.APIClient)
	countryCodes := c.Value("CountryCodes").(string)
	redirectUri := c.Value("RedirectUri").(string)
	plaidProducts := c.Value("PlaidProducts").(string)

	fmt.Printf("%+v\n", client)
	fmt.Printf("%+v\n", countryCodes)
	fmt.Printf("%+v\n", redirectUri)
	fmt.Printf("%+v\n", plaidProducts)
	fmt.Printf("%+v\n", c.Request)

	var createLinkRequestBody CreateLinkRequestBody

	if err := c.BindJSON(&createLinkRequestBody); err != nil {
		c.JSON(http.StatusNotFound, "This User is not registered for the application.")
		return
	}

	clientUserId := createLinkRequestBody.SubId

	fmt.Println(clientUserId)

	linkToken, err := linkTokenCreate(client, countryCodes, redirectUri, plaidProducts, clientUserId)

	if err != nil {
		RenderError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"link_token": linkToken})
}

func linkTokenCreate(
	client *plaid.APIClient,
	countryCodes string,
	redirectUri string,
	plaidProducts string,
	clientUserId string,
) (string, error) {

	ctx := context.Background()

	user := plaid.LinkTokenCreateRequestUser{
		ClientUserId: clientUserId,
	}

	request := plaid.NewLinkTokenCreateRequest(
		"personal-finance-app",
		"en",
		ConvertCountryCodes(strings.Split(countryCodes, ",")),
		user,
	)

	products := ConvertProducts(strings.Split(plaidProducts, ","))
	request.SetProducts(products)

	if redirectUri != "" {
		request.SetRedirectUri(redirectUri)
	}

	linkTokenCreateResp, _, err := client.PlaidApi.LinkTokenCreate(ctx).LinkTokenCreateRequest(*request).Execute()

	if err != nil {
		return "", err
	}

	return linkTokenCreateResp.GetLinkToken(), nil
}

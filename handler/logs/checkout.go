package logs

import (
	"fmt"
)

// func logCheckoutBackend(checkout *CheckoutLogRequest) {
// 	checkoutClient, _ := hclient.NewClient()
// 	checkout.AllowPublic = loading.Data.Settings.Settings.AllowPublicWebhook

// 	_, err := checkoutClient.NewRequest().
// 		SetURL("https://api.hellasaio.com/api/checkout").
// 		SetMethod("POST").
// 		SetHeader("Content-Type", "application/json").
// 		SetHeader("Accept", "*/*").
// 		SetHeader("Authorization", "Bearer "+auth.AuthToken).
// 		SetJSONBody(checkout).
// 		Do()

// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// }

// func logCheckoutDiscord(checkout *CheckoutLogRequest, discordWebhook string) {
// 	var title string
// 	var color int

// 	if checkout.Status == "success" {
// 		title = "**Successful Checkout!**"
// 		color = 2524623
// 	} else if checkout.Status == "denied" {
// 		title = "**Checkout Failed!**"
// 		color = 16711680
// 	} else {
// 		return // invalid status
// 	}

// 	checkoutClient, _ := hclient.NewClient()

// 	requestData := fmt.Sprintf(`{"content":null,"embeds":[{"title":"%s","description":"%s","color":%d,"fields":[{"name":"MSKU","value":"%s","inline":true},{"name":"Mode","value":"%s","inline":true},{"name":"Size","value":"[%s](https://quicktask.hellasaio.com/quicktask?product_id=%s&siteId=%s&size=%s)","inline":true},{"name":"Checkout Time","value":"%sms","inline":true},{"name":"Price","value":"€%.2f","inline":true},{"name":"Store","value":"%s","inline":true},{"name":"Quicktask Link","value":"[Link](https://quicktask.hellasaio.com/quicktask?product_id=%s&siteId=%s&size=random)","inline":true}],"thumbnail":{"url":"%s"}}],"attachments":[]}`,
// 		title, checkout.ProductName, color, checkout.ProductMSKU, checkout.Mode, checkout.Size, checkout.ProductMSKU, ConvertSiteToID[strings.ToLower(checkout.Website)], url.QueryEscape(checkout.Size), strconv.Itoa(checkout.CheckoutMs), checkout.Price, checkout.Website, checkout.ProductMSKU, ConvertSiteToID[strings.ToLower(checkout.Website)], checkout.ImageUrl)

// 	_, err := checkoutClient.NewRequest().
// 		SetURL(discordWebhook).
// 		SetMethod("POST").
// 		SetHeader("Content-Type", "application/json").
// 		SetHeader("Accept", "*/*").
// 		SetBody(requestData).
// 		Do()

// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	}
// }

func LogCheckout(checkout *CheckoutLogRequest, discordWebhook string) {
	// go logCheckoutBackend(checkout)
	// go logCheckoutDiscord(checkout, discordWebhook)
	fmt.Println(checkout)
}
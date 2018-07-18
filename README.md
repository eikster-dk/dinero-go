# dinero-go
The unofficial Golang SDK for Dinero Regnskab (dinero.dk)


# How to use
The SDK should be fairly easy to use, all it requires are a client key and client secret that you can get from https://api.dinero.dk/apply

When you have the above key and secret you can use the SDK like this
```
package main

import (
	"fmt"

	"github.com/eikc/dinero-go"
	"github.com/eikc/dinero-go/contacts"
)

func main() {
	clientKey, secret, := "testKey", "topSecretKeyHere"
	organizationApiKey, organizationId := "apikey from user here", 1233321

	c := dinero.NewClient(clientKey, secret)
	c.Authorize(organizationApiKey, organizationId)

	contactParams := contacts.ContactParams{
		Name:                         "Hello awesome",
		ExternalReference:            "external",
		AttPerson:                    "",
		City:                         "city",
		EanNumber:                    "",
		Email:                        "test@test.dk",
		PaymentConditionType:         contacts.NettoCash,
		Phone:                        "88 88 88 88",
		Street:                       "street",
		VatNumber:                    "",
		Webpage:                      "http://awesome.dk",
		ZipCode:                      "2700",
		CountryKey:                   "DK",
		IsPerson:                     true,
		PaymentConditionNumberOfDays: 10,
	}

	resp, err := contacts.Add(c, contactParams)
	if err != nil {
		// do something with the err, look at https://api.dinero.dk/docs/errorcodes for codes and why
	}

	// do something with contact created resp
	fmt.Println(resp)
}
```

every api endpoint is packed into it's own package that takes the dinero.Client interface and a set of parameters if needed.

The api is heavy inspired by stripe-go and feedback are welcome :-)


Enjoy!

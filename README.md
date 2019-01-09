# go-awardwallet-api
Libreria de Golang para integrar AwardWallet API como servicio.

## Instalacion
```sh
go get github.com/arthurnavah/go-awardwallet-api
```

## Ejemplo de uso
```go
package main

import (
	"fmt"

    "github.com/arthurnavah/go-awardwallet-api"
)

func main() {
    clientAward := awardwallet.NewClient("https://business.awardwallet.com/api/export/v1/account/1", "1010")

    responseAward, errAward, err := clientAward.GetAccountDetails()

    fmt.Println(responseAward)
    fmt.Println(errAward)
    fmt.Println(err)
}
```

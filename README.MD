SH5 WebApi Client for StoreHouse v5

Example:

package main

import (
"fmt"

    sh5apiclient "github.com/alexandr-andreyev/sh5-apiclient"

)

const (
BaseUrl = "http://127.0.0.1"
Sh5Port = 9797
)

func main() {
Sh5Client := sh5apiclient.NewClient(BaseUrl, Sh5Port, "Admin", "Password")
resp, err := Sh5Client.Sh5Info()
if err != nil {
fmt.Print("error:", err)
}
fmt.Printf("%s", resp)
}

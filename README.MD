SH5 WebApi Client for StoreHouse v5

Example:
// GET SH5 INFO

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

// SH5 PROC EXEC

    package main

    import (
        "encoding/json"
        "fmt"
        sh5apiclient "sh5test/sh5-apiclient"
    )

    const (
        BaseUrl = "http://127.0.0.1"
        Sh5Port = 9797
    )

    func main() {
        Sh5Client := sh5apiclient.NewClient(BaseUrl, Sh5Port, "Admin", "Password")

        resp, err := Sh5Client.Sh5Exec("Departs", nil)
        if err != nil {
            fmt.Print("error sh5exec:", err)
        }

        result, err := Sh5Client.Sh5ExecWithParse(resp, true)
        if err != nil {
            fmt.Println("error sh5execParse: ", err)
        }

        b, err := json.Marshal(result)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println(string(b))
    }

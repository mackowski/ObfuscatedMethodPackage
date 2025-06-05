package main

import (
    "encoding/base64"
    "fmt"
    "net"
)

func main() {
    encoded := "bG9jYWxob3N0OjQ0NDQ="
    decoded, _ := base64.StdEncoding.DecodeString(encoded)
    conn, err := net.Dial("tcp", string(decoded))
    if err != nil {
        fmt.Println("Error connecting:", err)
        return
    }
    fmt.Println("Connected to", conn.RemoteAddr())
    conn.Close()
}

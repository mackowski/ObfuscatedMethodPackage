package obfuscated

import (
    "encoding/base64"
    "fmt"
    "net"
)

// ConnectToEncodedAddress decodes a base64 address and connects to it.
func ConnectToEncodedAddress(encoded string) error {
    decoded, err := base64.StdEncoding.DecodeString(encoded)
    if err != nil {
        return fmt.Errorf("decode error: %w", err)
    }
    conn, err := net.Dial("tcp", string(decoded))
    if err != nil {
        return fmt.Errorf("connection error: %w", err)
    }
    defer conn.Close()
    fmt.Println("Connected to", conn.RemoteAddr())
    return nil
}

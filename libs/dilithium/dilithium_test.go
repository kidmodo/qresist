package dilithium

import "testing"

func TestDilithium(t *testing.T) {
    pubKey, privKey := Keypair()
    message := []byte("Test quantum-resistant signature")
    signature, _ := Sign(message, privKey)

    if !Verify(message, signature, pubKey) {
        t.Error("Failed to verify the signature")
    }
}

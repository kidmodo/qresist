package dilithium

/*
#cgo CFLAGS: -I/Users/kidmodo/qresist/libs/dilithium
#cgo LDFLAGS: -L/Users/kidmodo/qresist/libs/dilithium  -ldilithium
#include "api.h"
*/
import "C"
import (
    "unsafe"
    "fmt"
)

const PublicKeySize = 1312
const SecretKeySize = 2560
const SignatureSize = 2420

func Keypair() ([]byte, []byte) {
    publicKey := make([]byte, PublicKeySize)
    secretKey := make([]byte, SecretKeySize)

    result := C.pqcrystals_dilithium2_ref_keypair(
        (*C.uint8_t)(unsafe.Pointer(&publicKey[0])),
        (*C.uint8_t)(unsafe.Pointer(&secretKey[0])),
    )
    if result != 0 {
        panic("Keypair generation failed")
    }

    return publicKey, secretKey
}

func Sign(message []byte, secretKey []byte) ([]byte, error) {
    signature := make([]byte, SignatureSize)
    var siglen C.size_t

    result := C.pqcrystals_dilithium2_ref_signature(
        (*C.uint8_t)(unsafe.Pointer(&signature[0])),
        &siglen,
        (*C.uint8_t)(unsafe.Pointer(&message[0])),
        C.size_t(len(message)),
        nil, // No context
        0,
        (*C.uint8_t)(unsafe.Pointer(&secretKey[0])),
    )
    if result != 0 {
        return nil, fmt.Errorf("Signing failed")
    }

    return signature[:siglen], nil
}

func Verify(message []byte, signature []byte, publicKey []byte) bool {
    result := C.pqcrystals_dilithium2_ref_verify(
        (*C.uint8_t)(unsafe.Pointer(&signature[0])),
        C.size_t(len(signature)),
        (*C.uint8_t)(unsafe.Pointer(&message[0])),
        C.size_t(len(message)),
        nil, // No context
        0,
        (*C.uint8_t)(unsafe.Pointer(&publicKey[0])),
    )

    return result == 0
}

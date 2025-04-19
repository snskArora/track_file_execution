package main

import (
    "fmt"
    "log"
    "github.com/zeebo/xxh3"
)

type u128 = xxh3.Uint128
type hash struct {
    ui64 *uint64
    ui128 *u128
}

func show(h hash) string {
    if h.ui64 != nil {
        return fmt.Sprintf("%d", *h.ui64)
    }
    if h.ui128 != nil {
        return fmt.Sprintf("%d%d", h.ui128.Hi, h.ui128.Lo)
    }
    return ""
}

func hashString(hash_len int, data string) hash {
    switch hash_len {
        case 64:
            v := xxh3.HashString(data)
            return hash{ui64: &v}
        case 128:
            h := xxh3.HashString128(data)
            return hash{ui128: &h}
        default:
            log.Printf("Invalid hash size: %d, Acceptable are 64 and 128", hash_len)
            return hash{}
    }
}

// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls12

import (
	"bytes"
	"crypto/internal/fips"
	"crypto/internal/fips/sha256"
	"errors"
)

func init() {
	fips.CAST("TLSv1.2-SHA2-256", func() error {
		input := []byte{
			0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
			0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10,
		}
		transcript := []byte{
			0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18,
			0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f, 0x20,
		}
		want := []byte{
			0x8c, 0x3e, 0xed, 0xa7, 0x1c, 0x1b, 0x4c, 0xc0,
			0xa0, 0x44, 0x90, 0x75, 0xa8, 0x8e, 0xbc, 0x7c,
			0x5e, 0x1c, 0x4b, 0x1e, 0x4f, 0xe3, 0xc1, 0x06,
			0xeb, 0xdc, 0xc0, 0x5d, 0xc0, 0xc8, 0xec, 0xf3,
			0xe2, 0xb9, 0xd1, 0x03, 0x5e, 0xb2, 0x60, 0x5d,
			0x12, 0x68, 0x4f, 0x49, 0xdf, 0xa9, 0x9d, 0xcc,
		}
		if got := MasterSecret(sha256.New, input, transcript); !bytes.Equal(got, want) {
			return errors.New("unexpected result")
		}
		return nil
	})
}

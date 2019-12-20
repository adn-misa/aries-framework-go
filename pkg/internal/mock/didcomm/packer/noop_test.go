/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package packer

import (
	"encoding/base64"
	"testing"

	"github.com/btcsuite/btcutil/base58"
	"github.com/stretchr/testify/require"
)

// note: does not replicate correct packing
// when msg needs to be escaped.
func testPack(msg, senderKey, recKey []byte) []byte {
	headerValue := base64.URLEncoding.EncodeToString([]byte(`{"typ":"NOOP"}`))

	return []byte(`{"protected":"` + headerValue +
		`","spk":"` + base58.Encode(senderKey) +
		`","kid":"` + base58.Encode(recKey) +
		`","msg":"` + string(msg) + `"}`)
}

func TestPacker(t *testing.T) {
	p := New(nil)
	require.NotNil(t, p)
	require.Equal(t, encodingType, p.EncodingType())

	t.Run("no rec keys", func(t *testing.T) {
		_, err := p.Pack(nil, nil, nil)
		require.Error(t, err)
		require.Contains(t, err.Error(), "no recipients")
	})

	t.Run("pack, compare against correct data", func(t *testing.T) {
		msgin := []byte("hello my name is zoop")
		key := []byte("senderkey")
		rec := []byte("recipient")

		msgout, err := p.Pack(msgin, key, [][]byte{rec})
		require.NoError(t, err)

		correct := testPack(msgin, key, rec)
		require.Equal(t, correct, msgout)
	})

	t.Run("unpack fixed value, confirm data", func(t *testing.T) {
		correct := []byte("this is not a test message")
		key := []byte("testKey")
		rec := []byte("key2")
		msgin := testPack(correct, key, rec)

		envOut, err := p.Unpack(msgin)
		require.NoError(t, err)

		require.Equal(t, correct, envOut.Message)
		require.Equal(t, key, envOut.FromVerKey)
		require.Equal(t, rec, envOut.ToVerKey)
	})

	t.Run("multiple pack/unpacks", func(t *testing.T) {
		cleartext := []byte("this is not a test message")
		key1 := []byte("testKey")
		rec1 := []byte("rec1")
		key2 := []byte("wrapperKey")
		rec2 := []byte("rec2")

		correct1 := testPack(cleartext, key1, rec1)

		msg1, err := p.Pack(cleartext, key1, [][]byte{rec1})
		require.NoError(t, err)
		require.Equal(t, correct1, msg1)

		msg2, err := p.Pack(msg1, key2, [][]byte{rec2})
		require.NoError(t, err)

		env1, err := p.Unpack(msg2)
		require.NoError(t, err)
		require.Equal(t, key2, env1.FromVerKey)
		require.Equal(t, rec2, env1.ToVerKey)
		require.Equal(t, correct1, env1.Message)

		env2, err := p.Unpack(env1.Message)
		require.NoError(t, err)
		require.Equal(t, key1, env2.FromVerKey)
		require.Equal(t, rec1, env2.ToVerKey)
		require.Equal(t, cleartext, env2.Message)
	})

	t.Run("unpack errors", func(t *testing.T) {
		_, err := p.Unpack(nil)
		require.Error(t, err)
		require.Contains(t, err.Error(), "end of JSON input")

		_, err = p.Unpack([]byte("{}"))
		require.Error(t, err)
		require.Contains(t, err.Error(), "end of JSON input")

		_, err = p.Unpack([]byte("{\"protected\":\"$$$$$$$$$$$$\"}"))
		require.Error(t, err)
		require.Contains(t, err.Error(), "illegal base64 data")

		_, err = p.Unpack([]byte("{\"protected\":\"e3t7\"}"))
		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid character")
	})
}

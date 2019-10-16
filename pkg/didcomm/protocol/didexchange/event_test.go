/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package didexchange

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDIDExchangeEvent(t *testing.T) {
	ev := didExchangeEvent{connectionID: "abc", invitationID: "xyz"}
	require.Equal(t, ev.ConnectionID(), "abc")
	require.Equal(t, ev.InvitationID(), "xyz")

	err := errors.New("processing error")
	evErr := didExchangeEventError{err: err}
	require.Equal(t, err.Error(), evErr.Error())

	evErr = didExchangeEventError{}
	require.Equal(t, "", evErr.Error())
}

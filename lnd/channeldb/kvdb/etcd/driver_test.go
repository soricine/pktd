// +build kvdb_etcd

package etcd

import (
	"testing"

	"github.com/pkt-cash/pktd/pktwallet/walletdb"
	"github.com/stretchr/testify/require"
)

func TestOpenCreateFailure(t *testing.T) {
	t.Parallel()

	db, err := walletdb.Open(dbType)
	require.Error(t, err)
	require.Nil(t, db)

	db, err = walletdb.Open(dbType, "wrong")
	require.Error(t, err)
	require.Nil(t, db)

	db, err = walletdb.Create(dbType)
	require.Error(t, err)
	require.Nil(t, db)

	db, err = walletdb.Create(dbType, "wrong")
	require.Error(t, err)
	require.Nil(t, db)
}
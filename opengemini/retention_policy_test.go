package opengemini

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClientCreateRetentionPolicy(t *testing.T) {
	c := testDefaultClient(t)
	databaseName := randomDatabaseName()
	err := c.CreateDatabase(databaseName)
	require.Nil(t, err)
	err = c.CreateRetentionPolicy(databaseName, RpConfig{Name: "test_rp1", Duration: "3d"}, false)
	require.Nil(t, err)
	err = c.CreateRetentionPolicy(databaseName, RpConfig{Name: "test_rp2", Duration: "3d", ShardGroupDuration: "1h"}, false)
	require.Nil(t, err)
	err = c.CreateRetentionPolicy(databaseName, RpConfig{Name: "test_rp3", Duration: "3d", ShardGroupDuration: "1h", IndexDuration: "7h"}, false)
	require.Nil(t, err)
	err = c.CreateRetentionPolicy(databaseName, RpConfig{Name: "test_rp4", Duration: "3d"}, true)
	require.Nil(t, err)
	err = c.DropRetentionPolicy(databaseName, "test_rp4")
	require.Nil(t, err)
	err = c.DropRetentionPolicy(databaseName, "test_rp3")
	require.Nil(t, err)
	err = c.DropRetentionPolicy(databaseName, "test_rp2")
	require.Nil(t, err)
	err = c.DropRetentionPolicy(databaseName, "test_rp1")
	require.Nil(t, err)
	err = c.DropDatabase(databaseName)
	require.Nil(t, err)
}

func TestClientCreateRetentionPolicyNotExistDatabase(t *testing.T) {
	c := testDefaultClient(t)
	databaseName := randomDatabaseName()
	err := c.CreateRetentionPolicy(databaseName, RpConfig{Name: "test_rp1", Duration: "3d"}, false)
	require.NotNil(t, err)
	err = c.DropDatabase(databaseName)
	require.Nil(t, err)
}

func TestClientCreateRetentionPolicyEmptyDatabase(t *testing.T) {
	c := testDefaultClient(t)
	err := c.CreateRetentionPolicy("", RpConfig{Name: "test_rp1", Duration: "3d"}, false)
	require.NotNil(t, err)
}

func TestClientDropRetentionPolicy(t *testing.T) {
	c := testDefaultClient(t)
	err := c.DropRetentionPolicy("test_rp1", "test_database")
	require.Nil(t, err)
}

func TestClientShowRetentionPolicy(t *testing.T) {
	c := testDefaultClient(t)
	databaseName := randomDatabaseName()
	err := c.CreateDatabase(databaseName)
	require.Nil(t, err)
	rpResult, err := c.ShowRetentionPolicy(databaseName)
	require.Nil(t, err)
	require.NotEqual(t, len(rpResult), 0)
	err = c.DropDatabase(databaseName)
	require.Nil(t, err)
}

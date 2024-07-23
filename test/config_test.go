package test

import (
	"github.com/mrizkisaputra/golang-database/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConnection(t *testing.T) {
	connection := config.GetConnection()
	assert.NotNil(t, connection, "Should have returned a valid connection")
}

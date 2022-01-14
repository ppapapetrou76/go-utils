package types

import (
	"math/rand"
	"testing"
	"time"

	"github.com/ppapapetrou76/go-testing/assert"
)

func TestNewBool(t *testing.T) {
	value := rand.Intn(2) == 0
	expected := &value

	assert.That(t, NewBool(value)).IsEqualTo(expected)
}

func TestNewInt(t *testing.T) {
	value := rand.Intn(100)
	expected := &value

	assert.That(t, NewInt(value)).IsEqualTo(expected)
}

func TestNewInt64(t *testing.T) {
	value := rand.Int63()
	expected := &value

	assert.That(t, NewInt64(value)).IsEqualTo(expected)
}

func TestNewInt32(t *testing.T) {
	value := rand.Int31()
	expected := &value

	assert.That(t, NewInt32(value)).IsEqualTo(expected)
}

func TestNewUint32(t *testing.T) {
	value := rand.Uint32()
	expected := &value

	assert.That(t, NewUint32(value)).IsEqualTo(expected)
}

func TestNewUint64(t *testing.T) {
	value := rand.Uint64()
	expected := &value

	assert.That(t, NewUint64(value)).IsEqualTo(expected)
}

func TestNewUint(t *testing.T) {
	value := uint(rand.Uint64())
	expected := &value

	assert.That(t, NewUint(value)).IsEqualTo(expected)
}

func TestNewString(t *testing.T) {
	value := "random_string." + time.Now().GoString()
	expected := &value

	assert.That(t, NewString(value)).IsEqualTo(expected)
}

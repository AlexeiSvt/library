package tests

import (
	numbers"Library/internal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result, err := numbers.Add(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, 3, result)

	resultFloat, err := numbers.Add(1.5, 2.5)
	assert.NoError(t, err)
	assert.Equal(t, 4.0, resultFloat)
}

func TestSub(t *testing.T) {
	result, err := numbers.Sub(5, 3)
	assert.NoError(t, err)
	assert.Equal(t, 2, result)

	resultFloat, err := numbers.Sub(5.5, 2.5)
	assert.NoError(t, err)
	assert.Equal(t, 3.0, resultFloat)
}

func TestMult(t *testing.T) {
	result, err := numbers.Mult(2, 3)
	assert.NoError(t, err)
	assert.Equal(t, 6, result)

	resultFloat, err := numbers.Mult(2.5, 2.0)
	assert.NoError(t, err)
	assert.Equal(t, 5.0, resultFloat)
}

func TestDiv(t *testing.T) {
	result, err := numbers.Div(6, 3)
	assert.NoError(t, err)
	assert.Equal(t, 2, result)

	resultFloat, err := numbers.Div(5.0, 2.0)
	assert.NoError(t, err)
	assert.Equal(t, 2.5, resultFloat)

	_, err = numbers.Div(1, 0)
	assert.Error(t, err)
}

func TestMin(t *testing.T) {
	result, err := numbers.Min(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, 1, result)

	resultFloat, err := numbers.Min(1.5, 2.5)
	assert.NoError(t, err)
	assert.Equal(t, 1.5, resultFloat)
}

func TestMax(t *testing.T) {
	result, err := numbers.Max(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, 2, result)

	resultFloat, err := numbers.Max(1.5, 2.5)
	assert.NoError(t, err)
	assert.Equal(t, 2.5, resultFloat)
}

func TestAbs(t *testing.T) {
	result, err := numbers.Abs(-1)
	assert.NoError(t, err)
	assert.Equal(t, 1, result)

	resultFloat, err := numbers.Abs(-1.5)
	assert.NoError(t, err)
	assert.Equal(t, 1.5, resultFloat)
}

func TestFact(t *testing.T) {
	result, err := numbers.Fact(5)
	assert.NoError(t, err)
	assert.Equal(t, 120, result)
}

func TestIsEven(t *testing.T) {
	result, err := numbers.IsEven(4)
	assert.NoError(t, err)
	assert.True(t, result)

	result, err = numbers.IsEven(5)
	assert.NoError(t, err)
	assert.False(t, result)
}

func TestIsOdd(t *testing.T) {
	result, err := numbers.IsOdd(4)
	assert.NoError(t, err)
	assert.False(t, result)

	result, err = numbers.IsOdd(5)
	assert.NoError(t, err)
	assert.True(t, result)
}

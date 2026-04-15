package tests

import (
	strings"Library/internal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcat(t *testing.T) {
    result, err := strings.Concat("Hello, ", "World!")
    assert.NoError(t, err)
    assert.Equal(t, "Hello, World!", result)

    result, err = strings.Concat("", "World!")
    assert.Error(t, err)
    assert.Equal(t, "", result)
}

func TestContains(t *testing.T) {
    contains, err := strings.Contains("Hello, World!", "World")
    assert.NoError(t, err)
    assert.True(t, contains)

    contains, err = strings.Contains("Hello, World!", "")
    assert.Error(t, err)
    assert.False(t, contains)
}

func TestCount(t *testing.T) {
    count, err := strings.Count("Hello, World!", "o")
    assert.NoError(t, err)
    assert.Equal(t, 2, count)

    count, err = strings.Count("Hello, World!", "")
    assert.Error(t, err)
    assert.Equal(t, 0, count)
}

func TestReverse(t *testing.T) {
    reversed, err := strings.Reverse("Hello, World!")
    assert.NoError(t, err)
    assert.Equal(t, "!dlroW ,olleH", reversed)

    reversed, err = strings.Reverse("")
    assert.Error(t, err)
    assert.Equal(t, "", reversed)
}

func TestToUpper(t *testing.T) {
    upper, err := strings.ToUpper("Hello, World!")
    assert.NoError(t, err)
    assert.Equal(t, "HELLO, WORLD!", upper)

    upper, err = strings.ToUpper("")
    assert.Error(t, err)
    assert.Equal(t, "", upper)
}

func TestToLower(t *testing.T) {
    lower, err := strings.ToLower("Hello, World!")
    assert.NoError(t, err)
    assert.Equal(t, "hello, world!", lower)

    lower, err = strings.ToLower("")
    assert.Error(t, err)
    assert.Equal(t, "", lower)
}

func TestIsPalindrome(t *testing.T) {
    isPalindrome, err := strings.IsPalindrome("madam")
    assert.NoError(t, err)
    assert.True(t, isPalindrome)

    isPalindrome, err = strings.IsPalindrome("hello")
    assert.NoError(t, err)
    assert.False(t, isPalindrome)

    isPalindrome, err = strings.IsPalindrome("")
    assert.Error(t, err)
    assert.False(t, isPalindrome)
}
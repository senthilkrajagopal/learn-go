package senthil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//I forgot my comments.. i will update in place... BIG NOOO
func TestUserNameChange(t *testing.T) {
	u := user{
		name: "mukund",
	}
	u.setName("mukund change")
	assert.Equal(t, u.getName(), "mukund change")
}

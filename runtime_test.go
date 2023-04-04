package toolkit

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestRuntimeInjectData(t *testing.T) {
	type Book struct {
		Name string
	}

	runtime := NewRuntime()
	runtime.Log().Info("set book")
	runtime.Set("book", &Book{Name: "book 1"})
	runtime.Log().Infof("set price with value %v", 6000)
	runtime.Set("normalPrice", 6000)

	runtime.Log().Warningf("posilble error : %v", true)
	book := Get[*Book](runtime, "book")
	assert.Equal(t, "book 1", book.Name)
	price := Get[int](runtime, "normalPrice")
	assert.Equal(t, 6000, price)
	p := Get[int64](runtime, "normalPrice")
	assert.Equal(t, int64(0), p)
}

func TestRuntimeConfig(t *testing.T) {
	path := "./test/app.yaml"

	runtime := NewRuntime()
	runtime.Config().SetConfigType("yaml")
	runtime.ReadConfig(&path)

	s := runtime.Config().GetString("secret")
	runtime.log.Infof("%v = %v", "super secret", s)
	assert.Equal(t, "super secret", s)

}

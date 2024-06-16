package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygolang/fc-rest-api/internal/entity"
)

func TestNewProduct(t *testing.T) {
	p, err := entity.NewProduct("Product 1", 10.0)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "Product 1", p.Name)
	assert.Equal(t, 10.0, p.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := entity.NewProduct("", 10.0)
	assert.Nil(t, p)
	assert.Equal(t, entity.ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := entity.NewProduct("Product 1", 0)
	assert.Nil(t, p)
	assert.Equal(t, entity.ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := entity.NewProduct("Product 1", -10)
	assert.Nil(t, p)
	assert.Equal(t, entity.ErrInvalidPrice, err)
}

func TestProductValidate(t *testing.T) {
	p, err := entity.NewProduct("Product 1", 10.0)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}

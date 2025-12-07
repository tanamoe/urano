package fahasa

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	assert := assert.New(t)

	rawJSON, err := os.ReadFile("../../test/data/fahasa_product.json")
	assert.NoError(err)

	var expected Product
	err = json.Unmarshal(rawJSON, &expected)
	assert.NoError(err)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(rawJSON)
		assert.NoError(err)
	}))

	client := NewClient(WithDomain(srv.URL))

	product, err := client.Product(t.Context(), 1)
	assert.NoError(err)
	assert.NotNil(product)
	assert.Equal(&expected, product)
}

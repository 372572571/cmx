package definition

import (
	"embed"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed test.yml
var api_yaml embed.FS

var api_definition = []byte(`
api_definition:
  oss:
    - name: get
      http:
        method: get
        path: v1/oss/get
        summary: "get"
        body: "*"
      request: api.oss.oss.get_request
      response: api.oss.oss.get_response
`)

func TestApiParseFile(t *testing.T) {
	e := new(Apidefinition)
	_ = e
	test_data := []Api{
		{
			Name: "get",
			Http: struct {
				IsOpenApi bool   "json:\"is_open_api\" yaml:\"is_open_api\""
				Method    string "json:\"method\" yaml:\"method\""
				Path      string "json:\"path\" yaml:\"path\""
				Body      string "json:\"body\" yaml:\"body\""
				Summary   string "json:\"summary\" yaml:\"summary\""
			}{
				IsOpenApi: false,
				Method:    "get",
				Path:      "v1/oss/get",
				Body:      "*",
				Summary:   "get",
			},
			Request:  "api.oss.oss.get_request",
			Response: "api.oss.oss.get_response",
		},
	}
	tf, err := os.CreateTemp("", "test.yml")
	assert.NoError(t, err)
	defer os.Remove(tf.Name())
	os.WriteFile(tf.Name(), api_definition, 0666)
	em := e.ParseFile(tf.Name())
	assert.NotNil(t, em.Definition["oss"])
	_ = em
	for k, v := range em.Definition["oss"] {
		assert.Equal(t, test_data[k].Name, v.Name)
		assert.Equal(t, test_data[k].Http, v.Http)
		assert.Equal(t, test_data[k].Request, v.Request)
		assert.Equal(t, test_data[k].Response, v.Response)
	}
}

func TestApiParse(t *testing.T) {
	e := new(Apidefinition)
	_ = e
	test_data := []Api{
		{
			Name: "get",
			Http: struct {
				IsOpenApi bool   "json:\"is_open_api\" yaml:\"is_open_api\""
				Method    string "json:\"method\" yaml:\"method\""
				Path      string "json:\"path\" yaml:\"path\""
				Body      string "json:\"body\" yaml:\"body\""
				Summary   string "json:\"summary\" yaml:\"summary\""
			}{
				Method:  "get",
				Path:    "v1/oss/get",
				Body:    "*",
				Summary: "get",
			},
			Request:  "api.oss.oss.get_request",
			Response: "api.oss.oss.get_response",
		},
	}

	em := e.Parse(api_definition)
	assert.NotNil(t, em.Definition["oss"])
	_ = em
	for k, v := range em.Definition["oss"] {
		assert.Equal(t, test_data[k].Name, v.Name)
		assert.Equal(t, test_data[k].Http, v.Http)
		assert.Equal(t, test_data[k].Request, v.Request)
		assert.Equal(t, test_data[k].Response, v.Response)
	}
}

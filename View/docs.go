package View

import "github.com/swaggo/swag"

const docTemplate = `{
  "openapi": "3.0.3",
  "info": {
    "title": "Golang MVC Structure",
    "version": "1.0.11"
  },
  "tags": [
    {
      "name": "Product"
    }
  ],
  "paths": {
    "/getAllProducts": {
      "get": {
        "tags": [
          "Product"
        ],
        "summary": "get Product",
        "responses": {
          "201": {
            "description": "get Successful",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/Product"
                }
              }
            }
          }
        }
      }
    },
    "/getProductById": {
      "post": {
        "tags": [
          "Product"
        ],
        "summary": "get Product",
        "requestBody": {
          "description": "get Product",
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/Product"
              }
            }
          }
        },
        "responses": {
          "201": {
            "description": "get Successful",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/Product"
                }
              }
            }
          }
        }
      }
    },
    "/getUpdateById": {
      "post": {
        "tags": [
          "Product"
        ],
        "summary": "get Product",
        "requestBody": {
          "description": "get Product",
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/ProductOperation"
              }
            }
          }
        },
        "responses": {
          "201": {
            "description": "get Successful",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/Product"
                }
              }
            }
          }
        }
      }
    }
  },
  "components": {
    "schemas": {
      "Product": {
        "type": "object",
        "properties": {
          "sequence": {
            "type": "integer",
            "format": "int64"
          }, 
          "pid": {
            "type": "String",
            "example": "Ticket-01"
          }, 
          "pname": {
            "type": "String",
            "example": "123"
          }, 
          "price": {
            "type": "number",
            "format": "float"
          }, 
          "qty": {
            "type": "integer",
            "format": "int64",
            "example": "10"
          }
        }
      },
      "ProductOperation": {
        "type": "object",
        "properties": { 
          "pid": {
            "type": "String",
            "example": "Ticket-01"
          }, 
          "qty": {
            "type": "integer",
            "format": "int64",
            "example": "10"
          }, 
          "operator": {
            "type": "String",
            "example": "increase"
          }
        }
      }
    }
  }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Golang MVC",
	Description:      "This is a sample server todo server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

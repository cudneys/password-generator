// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://github.com/cudneys",
            "email": "password-generator@cudneys.net"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/generate": {
            "get": {
                "description": "Generates somewhat secure-ish passwords",
                "produces": [
                    "application/json"
                ],
                "summary": "Password Generator",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Password Length",
                        "name": "length",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Num Digits",
                        "name": "digits",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Num Symbols",
                        "name": "symbols",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.RequestParams": {
            "type": "object",
            "properties": {
                "digits": {
                    "type": "integer"
                },
                "length": {
                    "type": "integer"
                },
                "symbols": {
                    "type": "integer"
                }
            }
        },
        "models.Response": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "request": {
                    "$ref": "#/definitions/models.RequestParams"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Password Generator API",
	Description:      "This is a simple password generator",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

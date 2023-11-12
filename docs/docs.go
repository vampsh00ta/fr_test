// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/newsletter/create": {
            "post": {
                "description": "Создает рассылку",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Newsletter"
                ],
                "summary": "CreateNewsletter",
                "parameters": [
                    {
                        "description": "username",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_transport_http_v1.createNewsletterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/internal_transport_http_v1.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_transport_http_v1.response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/internal_transport_http_v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_transport_http_v1.response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "fr_internal_repository_models.Status": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "text": {
                    "type": "string",
                    "example": "status"
                },
                "time": {
                    "type": "string"
                }
            }
        },
        "internal_transport_http_v1.createNewsletterRequest": {
            "type": "object"
        },
        "internal_transport_http_v1.response": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "petstore.swagger.io",
	BasePath:         "/v2",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server Petstore server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
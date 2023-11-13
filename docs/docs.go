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
        "/client": {
            "put": {
                "description": "Изменяет  запись клиента",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Client"
                ],
                "summary": "Update",
                "parameters": [
                    {
                        "description": "body",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_transport_http_v1.updateClientRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
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
            },
            "delete": {
                "description": "Удаляет  запись клиента",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Client"
                ],
                "summary": "Delete",
                "parameters": [
                    {
                        "description": "body",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_transport_http_v1.deleteClientRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
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
        },
        "/client/add": {
            "post": {
                "description": "Создает запись клиента",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Client"
                ],
                "summary": "Add",
                "parameters": [
                    {
                        "description": "username",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_transport_http_v1.addClientRequest"
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
        },
        "/newsletter": {
            "delete": {
                "description": "Удаляет рассылку",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Newsletter"
                ],
                "summary": "Delete",
                "parameters": [
                    {
                        "description": "body",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/fr_internal_transport_http_v1.deleteNewsletterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/fr_internal_transport_http_v1.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/fr_internal_transport_http_v1.response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/fr_internal_transport_http_v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/fr_internal_transport_http_v1.response"
                        }
                    }
                }
            },
            "patch": {
                "description": "Изменяет рассылку",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Newsletter"
                ],
                "summary": "Update",
                "parameters": [
                    {
                        "description": "body",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/fr_internal_transport_http_v1.updateNewsletterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/fr_internal_transport_http_v1.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/fr_internal_transport_http_v1.response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/fr_internal_transport_http_v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/fr_internal_transport_http_v1.response"
                        }
                    }
                }
            }
        },
        "/newsletter/create": {
            "post": {
                "description": "Создает рассылку.Если убрать фильтр, то добавит в рассылку всех клиентов",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Newsletter"
                ],
                "summary": "Create",
                "parameters": [
                    {
                        "description": "body",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/fr_internal_transport_http_v1.createNewsletterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/fr_internal_transport_http_v1.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/fr_internal_transport_http_v1.response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/fr_internal_transport_http_v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/fr_internal_transport_http_v1.response"
                        }
                    }
                }
            }
        },
        "/newsletter/{id}": {
            "get": {
                "description": "Выводит информацию о рассылке",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Newsletter"
                ],
                "summary": "GetNewsletter",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Newsletter ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/fr_internal_transport_http_v1.getNewsletter"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/fr_internal_transport_http_v1.response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/fr_internal_transport_http_v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/fr_internal_transport_http_v1.response"
                        }
                    }
                }
            }
        },
        "/newsletter/{id}/messages": {
            "get": {
                "description": "Выводит последние статусы сообщения по id рассылки",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Newsletter"
                ],
                "summary": "GetLastMessageStatuses",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Newsletter ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "status",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/fr_internal_transport_http_v1.getLastMessageStatusesResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/fr_internal_transport_http_v1.response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/fr_internal_transport_http_v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/fr_internal_transport_http_v1.response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "fr_internal_transport_http_v1.addClientRequest": {
            "type": "object",
            "required": [
                "mobile_code",
                "tag",
                "tel_number",
                "timezone"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "mobile_code": {
                    "type": "string",
                    "example": "7"
                },
                "tag": {
                    "type": "string",
                    "example": "tag"
                },
                "tel_number": {
                    "type": "integer",
                    "example": 9999999999
                },
                "timezone": {
                    "type": "string",
                    "example": "UTC"
                }
            }
        },
        "fr_internal_transport_http_v1.createNewsletterRequest": {
            "type": "object",
            "required": [
                "start_time",
                "text"
            ],
            "properties": {
                "filter": {
                    "$ref": "#/definitions/models.Filter"
                },
                "start_time": {
                    "type": "string",
                    "example": "2023-11-12T16:45:00Z"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "fr_internal_transport_http_v1.deleteClientRequest": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "fr_internal_transport_http_v1.deleteNewsletterRequest": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "fr_internal_transport_http_v1.getLastMessageStatusesResponse": {
            "type": "object",
            "required": [
                "messages"
            ],
            "properties": {
                "messages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.MessageStatus"
                    }
                }
            }
        },
        "fr_internal_transport_http_v1.getNewsletter": {
            "type": "object",
            "required": [
                "messages",
                "start_time",
                "text"
            ],
            "properties": {
                "filter": {
                    "$ref": "#/definitions/models.Filter"
                },
                "messages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.MessageStatus"
                    }
                },
                "start_time": {
                    "type": "string",
                    "example": "2023-11-12T16:45:00Z"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "fr_internal_transport_http_v1.response": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "fr_internal_transport_http_v1.updateClientRequest": {
            "type": "object",
            "required": [
                "mobile_code",
                "tag",
                "tel_number",
                "timezone"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "mobile_code": {
                    "type": "string",
                    "example": "7"
                },
                "tag": {
                    "type": "string",
                    "example": "tag"
                },
                "tel_number": {
                    "type": "integer",
                    "example": 9999999999
                },
                "timezone": {
                    "type": "string",
                    "example": "UTC"
                }
            }
        },
        "fr_internal_transport_http_v1.updateNewsletterRequest": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "text": {
                    "type": "string",
                    "example": "1"
                },
                "time": {
                    "type": "string",
                    "example": "1"
                }
            }
        },
        "internal_transport_http_v1.addClientRequest": {
            "type": "object",
            "required": [
                "mobile_code",
                "tag",
                "tel_number",
                "timezone"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "mobile_code": {
                    "type": "string",
                    "example": "7"
                },
                "tag": {
                    "type": "string",
                    "example": "tag"
                },
                "tel_number": {
                    "type": "integer",
                    "example": 9999999999
                },
                "timezone": {
                    "type": "string",
                    "example": "UTC"
                }
            }
        },
        "internal_transport_http_v1.createNewsletterRequest": {
            "type": "object",
            "required": [
                "start_time",
                "text"
            ],
            "properties": {
                "filter": {
                    "$ref": "#/definitions/models.Filter"
                },
                "start_time": {
                    "type": "string",
                    "example": "2023-11-12T16:45:00Z"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "internal_transport_http_v1.deleteClientRequest": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "internal_transport_http_v1.deleteNewsletterRequest": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "internal_transport_http_v1.getLastMessageStatusesResponse": {
            "type": "object",
            "required": [
                "messages"
            ],
            "properties": {
                "messages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.MessageStatus"
                    }
                }
            }
        },
        "internal_transport_http_v1.getNewsletter": {
            "type": "object",
            "required": [
                "messages",
                "start_time",
                "text"
            ],
            "properties": {
                "filter": {
                    "$ref": "#/definitions/models.Filter"
                },
                "messages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.MessageStatus"
                    }
                },
                "start_time": {
                    "type": "string",
                    "example": "2023-11-12T16:45:00Z"
                },
                "text": {
                    "type": "string"
                }
            }
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
        },
        "internal_transport_http_v1.updateClientRequest": {
            "type": "object",
            "required": [
                "mobile_code",
                "tag",
                "tel_number",
                "timezone"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "mobile_code": {
                    "type": "string",
                    "example": "7"
                },
                "tag": {
                    "type": "string",
                    "example": "tag"
                },
                "tel_number": {
                    "type": "integer",
                    "example": 9999999999
                },
                "timezone": {
                    "type": "string",
                    "example": "UTC"
                }
            }
        },
        "internal_transport_http_v1.updateNewsletterRequest": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "text": {
                    "type": "string",
                    "example": "1"
                },
                "time": {
                    "type": "string",
                    "example": "1"
                }
            }
        },
        "models.Filter": {
            "type": "object",
            "properties": {
                "mobile_code": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "test"
                    ]
                },
                "tag": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "tag"
                    ]
                }
            }
        },
        "models.MessageStatus": {
            "type": "object",
            "properties": {
                "client_id": {
                    "type": "integer"
                },
                "creation_time": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "newsletter_id": {
                    "type": "integer"
                },
                "send_time": {
                    "type": "string"
                },
                "status": {
                    "type": "string",
                    "example": "status"
                },
                "time": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8000",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

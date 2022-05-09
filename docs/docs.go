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
        "contact": {
            "name": "Dimas Restu Hidayanto",
            "url": "https://github.com/dimaskiddo",
            "email": "drh.dimasrestu@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/whatsapp": {
            "get": {
                "description": "Get The Server Status",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Root"
                ],
                "summary": "Show The Status of The Server",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/v1/whatsapp/auth": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Get Authentication Token",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Root"
                ],
                "summary": "Generate Authentication Token",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/v1/whatsapp/login": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get QR Code for WhatsApp Multi-Device Login",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json",
                    "text/html"
                ],
                "tags": [
                    "WhatsApp"
                ],
                "summary": "Generate QR Code for WhatsApp Multi-Device Login",
                "parameters": [
                    {
                        "enum": [
                            "html",
                            "json"
                        ],
                        "type": "string",
                        "default": "html",
                        "description": "Change Output Format in HTML or JSON",
                        "name": "output",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/v1/whatsapp/logout": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Make Device Logout from WhatsApp Multi-Device",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp"
                ],
                "summary": "Logout Device from WhatsApp Multi-Device",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/v1/whatsapp/send/text": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Make Device Logout from WhatsApp Multi-Device",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp"
                ],
                "summary": "Logout Device from WhatsApp Multi-Device",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Destination Phone Number",
                        "name": "msisdn",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Text Message Content",
                        "name": "message",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        },
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.x",
	Host:             "127.0.0.1:3000",
	BasePath:         "",
	Schemes:          []string{"http"},
	Title:            "Go WhatsApp Multi-Device REST API",
	Description:      "This is WhatsApp Multi-Device Implementation in Go REST API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
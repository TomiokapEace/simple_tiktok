// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/feed": {
            "get": {
                "tags": [
                    "基础接口"
                ],
                "summary": "视频流",
                "parameters": [
                    {
                        "type": "string",
                        "description": "latest_time",
                        "name": "latest_time",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/hello": {
            "get": {
                "tags": [
                    "公共接口"
                ],
                "summary": "首页",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/publish/action/": {
            "post": {
                "tags": [
                    "基础接口"
                ],
                "summary": "视频投稿",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "文件",
                        "name": "data",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "标题",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

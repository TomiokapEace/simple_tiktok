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
                            "$ref": "#/definitions/controller.UploadRespStruct"
                        }
                    }
                }
            }
        },
        "/publish/list/": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "视频发布列表接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录用户的token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "查找目标用户的id",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "status_msg为成功",
                        "schema": {
                            "$ref": "#/definitions/controller.GetPublishListResponse"
                        }
                    }
                }
            }
        },
        "/relation/action/": {
            "post": {
                "tags": [
                    "社交接口"
                ],
                "summary": "关注操作",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户id",
                        "name": "to_user_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "关注操作",
                        "name": "action_type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.FollowRespStruct"
                        }
                    }
                }
            }
        },
        "/relation/follow/list/": {
            "get": {
                "tags": [
                    "社交接口"
                ],
                "summary": "关注列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户id",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.FollowListRespStruct"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.FollowListRespStruct": {
            "type": "object",
            "properties": {
                "status_code": {
                    "type": "integer"
                },
                "status_msg": {
                    "type": "string"
                },
                "user_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/service.User"
                    }
                }
            }
        },
        "controller.FollowRespStruct": {
            "type": "object",
            "properties": {
                "status_code": {
                    "type": "integer"
                },
                "status_msg": {
                    "type": "string"
                }
            }
        },
        "controller.GetPublishListResponse": {
            "type": "object",
            "properties": {
                "status_code": {
                    "type": "integer"
                },
                "status_msg": {
                    "type": "string"
                },
                "video_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/service.Video"
                    }
                }
            }
        },
        "controller.UploadRespStruct": {
            "type": "object",
            "properties": {
                "status_code": {
                    "type": "integer"
                },
                "status_msg": {
                    "type": "string"
                }
            }
        },
        "service.Author": {
            "type": "object",
            "properties": {
                "follow_count": {
                    "description": "default",
                    "type": "integer"
                },
                "follower_count": {
                    "description": "default",
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "is_follow": {
                    "description": "defalut",
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "service.User": {
            "type": "object",
            "properties": {
                "follow_count": {
                    "type": "integer"
                },
                "follower_count": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "is_follow": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "service.Video": {
            "type": "object",
            "properties": {
                "author": {
                    "$ref": "#/definitions/service.Author"
                },
                "comment_count": {
                    "type": "integer"
                },
                "cover_url": {
                    "type": "string"
                },
                "favorite_count": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "is_favorite": {
                    "type": "boolean"
                },
                "play_url": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
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

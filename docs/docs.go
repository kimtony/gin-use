// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "MIT",
            "url": "https://www.baidu.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/health": {
            "get": {
                "description": "服务是否启动正常检查",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "监测服务"
                ],
                "summary": "健康检查接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "name",
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
        "/api/v1/account/info": {
            "post": {
                "description": "用户个人信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account服务"
                ],
                "summary": "request.WechatAccount",
                "parameters": [
                    {
                        "maxLength": 10,
                        "minLength": 5,
                        "type": "string",
                        "description": "string valid",
                        "name": "string",
                        "in": "query"
                    },
                    {
                        "description": "微信参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Account"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Account"
                        }
                    }
                }
            }
        },
        "/api/wechat": {
            "post": {
                "security": [
                    {
                        "OAuth2Application": [
                            "account"
                        ]
                    }
                ],
                "description": "微信服务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "微信服务"
                ],
                "summary": "微信服务",
                "parameters": [
                    {
                        "description": "微信参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.WechatAccount"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "token",
                        "in": "header",
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
    "definitions": {
        "model.Account": {
            "type": "object",
            "required": [
                "id",
                "name"
            ],
            "properties": {
                "id": {
                    "type": "string"
                },
                "mobile": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "request.WechatAccount": {
            "type": "object",
            "required": [
                "mobile",
                "nickname",
                "openid"
            ],
            "properties": {
                "mobile": {
                    "type": "integer"
                },
                "nickname": {
                    "type": "string"
                },
                "openid": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "2.0",
	Host:        "192.168.1.163:8081",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "swagger 接口文档",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}

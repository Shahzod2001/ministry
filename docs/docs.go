// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/admin/teachers": {
            "get": {
                "description": "Get all teachers from all universities",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "teachers"
                ],
                "summary": "Get all teachers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.response"
                        }
                    }
                }
            }
        },
        "/teacher/all": {
            "get": {
                "description": "Get all teachers associated with the university",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "teachers"
                ],
                "summary": "Get teachers of the university",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.response"
                        }
                    }
                }
            }
        },
        "/teacher/create": {
            "post": {
                "description": "Add a new teacher to the university",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "teachers"
                ],
                "summary": "Add a new teacher",
                "parameters": [
                    {
                        "description": "Add Teacher",
                        "name": "teacher",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.addTeacherInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.addTeacherInput": {
            "type": "object",
            "required": [
                "academic_degree_id",
                "academic_position_id",
                "birth_date",
                "birth_place",
                "direction_spec_id",
                "first_name",
                "from_year",
                "gender",
                "job_title",
                "last_name",
                "spec_id",
                "to_year",
                "type_id"
            ],
            "properties": {
                "academic_degree_id": {
                    "type": "integer"
                },
                "academic_position_id": {
                    "type": "integer"
                },
                "birth_date": {
                    "type": "string"
                },
                "birth_place": {
                    "type": "string"
                },
                "direction_spec_id": {
                    "type": "integer"
                },
                "first_name": {
                    "type": "string"
                },
                "from_year": {
                    "type": "integer"
                },
                "gender": {
                    "type": "integer"
                },
                "job_title": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "middle_name": {
                    "type": "string"
                },
                "other_job": {
                    "type": "string"
                },
                "spec_id": {
                    "type": "integer"
                },
                "to_year": {
                    "type": "integer"
                },
                "type_id": {
                    "type": "integer"
                }
            }
        },
        "handler.response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
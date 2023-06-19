// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "https://github.com/stefanymelda",
            "email": "1214026@std.ulbi.ac.id"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/dlt/{id}": {
            "delete": {
                "description": "Hapus data kuesioner.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kuesioner"
                ],
                "summary": "Delete data kuesioner.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/ins1": {
            "post": {
                "description": "Input data kuesioner.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kuesioner"
                ],
                "summary": "Insert data kuesioner.",
                "parameters": [
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.Kuesioner"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Kuesioner"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/kuesioner": {
            "get": {
                "description": "Mengambil semua data kuesioner.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kuesioner"
                ],
                "summary": "Get All Data Kuesioner.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Kuesioner"
                        }
                    }
                }
            }
        },
        "/kuesioner/{id}": {
            "get": {
                "description": "Ambil per ID data kuesioner.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kuesioner"
                ],
                "summary": "Get By ID Data Kuesioner.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Kuesioner"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/update/{id}": {
            "put": {
                "description": "Ubah data kuesioner.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kuesioner"
                ],
                "summary": "Update data kuesioner.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.Kuesioner"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Kuesioner"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.Kuesioner": {
            "type": "object",
            "properties": {
                "biodata": {
                    "$ref": "#/definitions/controller.Responden"
                },
                "email": {
                    "type": "string",
                    "example": "jasmine@gmail.com"
                },
                "latitude": {
                    "type": "number",
                    "example": 123.11
                },
                "location": {
                    "type": "string",
                    "example": "Bandung"
                },
                "longitude": {
                    "description": "ID           primitive.ObjectID ` + "`" + `bson:\"_id,omitempty\" json:\"_id,omitempty\"` + "`" + `",
                    "type": "number",
                    "example": 123.11
                },
                "status": {
                    "description": "Datetime     primitive.DateTime ` + "`" + `bson:\"datetime,omitempty\" json:\"datetime,omitempty\"` + "`" + `",
                    "type": "string",
                    "example": "Done"
                }
            }
        },
        "controller.Responden": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "jasmine@gmail.com"
                },
                "jenis_kelamin": {
                    "type": "string",
                    "example": "Prempuan"
                },
                "nama": {
                    "description": "ID               primitive.ObjectID ` + "`" + `bson:\"_id,omitempty\" json:\"_id,omitempty\"` + "`" + `",
                    "type": "string",
                    "example": "Jasmine"
                },
                "phone_number": {
                    "type": "string",
                    "example": "081243284212"
                },
                "usia": {
                    "type": "integer",
                    "example": 18
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "jaehyun.herokuapp.com",
	BasePath:         "/",
	Schemes:          []string{"https", "http"},
	Title:            "TES SWAG",
	Description:      "This is a sample server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

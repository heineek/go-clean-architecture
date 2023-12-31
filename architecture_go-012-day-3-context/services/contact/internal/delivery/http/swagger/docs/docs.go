// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
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
            "name": "API Support",
            "email": "kolyadkons@gmail.com"
        },
        "license": {
            "name": "kolyadkons"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/contacts": {
            "get": {
                "description": "Метод позволяет получить список контактов.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "contacts"
                ],
                "summary": "Получить список контактов.",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Количество записей",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "Смещение при получении записей",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "name",
                        "description": "Сортировка по полю",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Список контактов",
                        "schema": {
                            "$ref": "#/definitions/contact.ListContact"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "404 Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Метод позволяет создать контакт.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "contacts"
                ],
                "summary": "Метод позволяет создать контакт.",
                "parameters": [
                    {
                        "description": "Данные по контакту",
                        "name": "contact",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/contact.ShortContact"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "201": {
                        "description": "Структура контакта",
                        "schema": {
                            "$ref": "#/definitions/contact.ContactResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "404 Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/contacts/{id}": {
            "get": {
                "description": "Метод позволяет получить контакт по мдентификатору контакта.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "contacts"
                ],
                "summary": "Получить контакт.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор контакта",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Структура контакта",
                        "schema": {
                            "$ref": "#/definitions/contact.ContactResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "404 Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Метод позволяет обновить данные контакта.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "contacts"
                ],
                "summary": "Метод позволяет обновить данные контакта.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор контакта",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Данные по контакту",
                        "name": "contact",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/contact.ShortContact"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Структура контакта",
                        "schema": {
                            "$ref": "#/definitions/contact.ContactResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "404 Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Метод позволяет удалить контакт.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "contacts"
                ],
                "summary": "Метод позволяет удалить контакт.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор контакта",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "404 Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/groups/": {
            "get": {
                "description": "Метод позволяет получить список групп.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Метод позволяет получить список групп.",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Количество записей",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "Смещение при получении записей",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "name",
                        "description": "Сортировка по полю",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/group.GroupList"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "404 Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Метод позволяет создать группу контактов.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Метод позволяет создать группу контактов.",
                "parameters": [
                    {
                        "description": "Данные по группе",
                        "name": "group",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/group.ShortGroup"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/group.GroupResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "404 Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/groups/{id}": {
            "get": {
                "description": "Метод позволяет получить данные по группе.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Метод позволяет получить данные по группе.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор группы контактов",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/group.GroupResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "404 Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Метод позволяет обновить данные группы.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Метод позволяет обновить данные группы.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор группы",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Данные по группе",
                        "name": "group",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/group.ShortGroup"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/group.GroupResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "404 Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Метод позволяет удалить группу.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Метод позволяет удалить группу.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор группы",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "404 Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/groups/{id}/contacts/": {
            "post": {
                "security": [
                    {
                        "Cookies": []
                    }
                ],
                "description": "Создание контакта и добавление его в существующую группу.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Создание контакта и добавление его в существующую группу.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор группы контактов",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Данные по контакту",
                        "name": "contact",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/contact.ShortContact"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "404 Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/groups/{id}/contacts/{contactId}": {
            "post": {
                "description": "Метод позволяет добавить контакты в группу.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Метод позволяет добавить контакты в группу.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор группы",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Идентификатор контакта",
                        "name": "contactId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "404 Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Метод позволяет удалить контакт из группы.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Метод позволяет удалить контакт из группы.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор группы",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Идентификатор контакта",
                        "name": "contactId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "404 Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "contact.ContactResponse": {
            "type": "object",
            "required": [
                "createdAt",
                "id",
                "modifiedAt",
                "phoneNumber"
            ],
            "properties": {
                "age": {
                    "description": "Возраст",
                    "type": "integer",
                    "default": 0,
                    "maximum": 200,
                    "minimum": 0,
                    "example": 42
                },
                "createdAt": {
                    "description": "Дата создания контакта",
                    "type": "string"
                },
                "email": {
                    "description": "Электронная почта",
                    "type": "string",
                    "format": "email",
                    "maxLength": 250,
                    "example": "example@gmail.com"
                },
                "gender": {
                    "description": "Пол",
                    "type": "integer",
                    "enum": [
                        1,
                        2
                    ],
                    "example": 1
                },
                "id": {
                    "description": "Идетификатор записи",
                    "type": "string",
                    "format": "uuid",
                    "example": "00000000-0000-0000-0000-000000000000"
                },
                "modifiedAt": {
                    "description": "Дата последнего изменения контакта",
                    "type": "string"
                },
                "name": {
                    "description": "Имя клиента",
                    "type": "string",
                    "maxLength": 50,
                    "example": "Иван"
                },
                "patronymic": {
                    "description": "Отчество клиента",
                    "type": "string",
                    "maxLength": 100,
                    "example": "Иванович"
                },
                "phoneNumber": {
                    "description": "Мобильный телефон",
                    "type": "string",
                    "maxLength": 50,
                    "example": "78002002020"
                },
                "surname": {
                    "description": "Фамилия клиента",
                    "type": "string",
                    "maxLength": 100,
                    "example": "Иванов"
                }
            }
        },
        "contact.ListContact": {
            "type": "object",
            "properties": {
                "limit": {
                    "description": "Количество записей",
                    "type": "integer",
                    "default": 10,
                    "minimum": 0,
                    "example": 10
                },
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/contact.ContactResponse"
                    }
                },
                "offset": {
                    "description": "Смещение при получении записей",
                    "type": "integer",
                    "default": 0,
                    "minimum": 0,
                    "example": 20
                },
                "total": {
                    "description": "Всего",
                    "type": "integer",
                    "default": 0,
                    "minimum": 0,
                    "example": 10
                }
            }
        },
        "contact.ShortContact": {
            "type": "object",
            "required": [
                "phoneNumber"
            ],
            "properties": {
                "age": {
                    "description": "Возраст",
                    "type": "integer",
                    "default": 0,
                    "maximum": 200,
                    "minimum": 0,
                    "example": 42
                },
                "email": {
                    "description": "Электронная почта",
                    "type": "string",
                    "format": "email",
                    "maxLength": 250,
                    "example": "example@gmail.com"
                },
                "gender": {
                    "description": "Пол",
                    "type": "integer",
                    "enum": [
                        1,
                        2
                    ],
                    "example": 1
                },
                "name": {
                    "description": "Имя клиента",
                    "type": "string",
                    "maxLength": 50,
                    "example": "Иван"
                },
                "patronymic": {
                    "description": "Отчество клиента",
                    "type": "string",
                    "maxLength": 100,
                    "example": "Иванович"
                },
                "phoneNumber": {
                    "description": "Мобильный телефон",
                    "type": "string",
                    "maxLength": 50,
                    "example": "78002002020"
                },
                "surname": {
                    "description": "Фамилия клиента",
                    "type": "string",
                    "maxLength": 100,
                    "example": "Иванов"
                }
            }
        },
        "group.GroupList": {
            "type": "object",
            "properties": {
                "limit": {
                    "description": "Количество записей",
                    "type": "integer",
                    "default": 10,
                    "minimum": 0,
                    "example": 10
                },
                "list": {
                    "description": "Список групп",
                    "type": "array",
                    "minItems": 0,
                    "items": {
                        "$ref": "#/definitions/group.GroupResponse"
                    }
                },
                "offset": {
                    "description": "Смещение при получении записей",
                    "type": "integer",
                    "default": 0,
                    "minimum": 0,
                    "example": 20
                },
                "total": {
                    "description": "Всего",
                    "type": "integer",
                    "default": 0,
                    "minimum": 0,
                    "example": 10
                }
            }
        },
        "group.GroupResponse": {
            "type": "object",
            "required": [
                "createdAt",
                "id",
                "modifiedAt",
                "name"
            ],
            "properties": {
                "contactsAmount": {
                    "description": "Кол-во контактов в группе",
                    "type": "integer",
                    "default": 10,
                    "minimum": 0
                },
                "createdAt": {
                    "description": "Дата создания группы",
                    "type": "string"
                },
                "description": {
                    "description": "Описание",
                    "type": "string",
                    "maxLength": 1000,
                    "example": "Описание группы"
                },
                "id": {
                    "description": "Идентификатор группы",
                    "type": "string",
                    "format": "uuid",
                    "example": "00000000-0000-0000-0000-000000000000"
                },
                "modifiedAt": {
                    "description": "Дата последнего изменения группы",
                    "type": "string"
                },
                "name": {
                    "description": "Название группы",
                    "type": "string",
                    "maxLength": 100,
                    "example": "Название группы"
                }
            }
        },
        "group.ShortGroup": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "description": {
                    "description": "Описание",
                    "type": "string",
                    "maxLength": 1000,
                    "example": "Описание группы"
                },
                "name": {
                    "description": "Название группы",
                    "type": "string",
                    "maxLength": 100,
                    "example": "Название группы"
                }
            }
        },
        "http.ErrorResponse": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "id": {
                    "type": "string"
                },
                "info": {},
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "slurm contact service on clean architecture",
	Description:      "contact service on clean architecture",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

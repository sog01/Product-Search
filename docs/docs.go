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
        "/products/bulk": {
            "post": {
                "description": "BulkInsert bulk insert products",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products API"
                ],
                "summary": "BulkInsert bulk insert products",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.BulkInsertReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BulkInsertResp"
                        }
                    }
                }
            }
        },
        "/products/bulk/update": {
            "post": {
                "description": "BulkUpdate bulk update products",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products API"
                ],
                "summary": "BulkUpdate bulk update products",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.BulkUpdateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BulkUpdateResp"
                        }
                    }
                }
            }
        },
        "/search": {
            "get": {
                "description": "Search product from given q",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Search API"
                ],
                "summary": "Search product from given q",
                "parameters": [
                    {
                        "type": "string",
                        "description": "q",
                        "name": "q",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "catalog",
                        "name": "catalog",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "size",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "sort_by",
                        "name": "sort_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "next_cursor",
                        "name": "next_cursor",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SearchResponse"
                        }
                    }
                }
            }
        },
        "/search/autocomplete": {
            "get": {
                "description": "Search autocomplete from given q",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Search API"
                ],
                "summary": "Search autocomplete from given q",
                "parameters": [
                    {
                        "type": "string",
                        "description": "q",
                        "name": "q",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.AutocompleteResp"
                        }
                    }
                }
            }
        },
        "/search/catalogs": {
            "get": {
                "description": "Search product catalogs from given q",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Search API"
                ],
                "summary": "Search product catalogs from given q",
                "parameters": [
                    {
                        "type": "string",
                        "description": "q",
                        "name": "q",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SearchCatalogsResp"
                        }
                    }
                }
            }
        },
        "/search/total": {
            "get": {
                "description": "Search product total from given q",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Search API"
                ],
                "summary": "Search product total from given q",
                "parameters": [
                    {
                        "type": "string",
                        "description": "q",
                        "name": "q",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SearchTotalResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Autocomplete": {
            "type": "object",
            "properties": {
                "highlight": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.AutocompleteResp": {
            "type": "object",
            "properties": {
                "autocompletes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Autocomplete"
                    }
                }
            }
        },
        "model.BulkInsertReq": {
            "type": "object",
            "properties": {
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ProductSearchInsert"
                    }
                }
            }
        },
        "model.BulkInsertResp": {
            "type": "object"
        },
        "model.BulkUpdateReq": {
            "type": "object",
            "properties": {
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ProductSearchUpdate"
                    }
                }
            }
        },
        "model.BulkUpdateResp": {
            "type": "object"
        },
        "model.ProductSearchCatalogs": {
            "type": "object",
            "properties": {
                "catalog": {
                    "type": "string"
                },
                "count": {
                    "type": "integer"
                }
            }
        },
        "model.ProductSearchInsert": {
            "type": "object",
            "properties": {
                "catalog": {
                    "type": "string"
                },
                "cta_url": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.ProductSearchResponse": {
            "type": "object",
            "properties": {
                "catalog": {
                    "type": "string"
                },
                "cta_url": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.ProductSearchUpdate": {
            "type": "object",
            "properties": {
                "catalog": {
                    "type": "string"
                },
                "cta_url": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.SearchCatalogsResp": {
            "type": "object",
            "properties": {
                "catalogs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ProductSearchCatalogs"
                    }
                }
            }
        },
        "model.SearchResponse": {
            "type": "object",
            "properties": {
                "next_cursor": {
                    "type": "string"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ProductSearchResponse"
                    }
                }
            }
        },
        "model.SearchTotalResp": {
            "type": "object",
            "properties": {
                "total": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Product Search API",
	Description:      "This is a product search API swagger documentation.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
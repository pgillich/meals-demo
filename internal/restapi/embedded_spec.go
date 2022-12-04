// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is demo for a foodstore (meals) service",
    "title": "OpenAPI Foodstore",
    "license": {
      "name": "Apache-2.0",
      "url": "https://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "foodstore.kind-01.company.com",
  "basePath": "/v1",
  "paths": {
    "/ingredients": {
      "get": {
        "description": "ll ingredients are stored",
        "produces": [
          "application/json"
        ],
        "tags": [
          "meal"
        ],
        "summary": "Get all ingredients",
        "operationId": "getIngredients",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Ingredient"
              }
            }
          },
          "500": {
            "description": "Invalid tag value",
            "schema": {
              "$ref": "#/definitions/ApiError"
            }
          }
        }
      }
    },
    "/livez": {
      "get": {
        "description": "Returns OK",
        "tags": [
          "info"
        ],
        "summary": "Liveness status for orchestrator",
        "operationId": "getLivez",
        "responses": {
          "200": {
            "description": "successful operation"
          }
        }
      }
    },
    "/login": {
      "post": {
        "description": "Returns token for authorized User",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "operationId": "login",
        "parameters": [
          {
            "description": "Login Payload",
            "name": "login",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/LoginInfo"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful login",
            "schema": {
              "$ref": "#/definitions/LoginSuccess"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "User not found",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/meal/findByTag": {
      "get": {
        "description": "One tag ID can be provided",
        "produces": [
          "application/json"
        ],
        "tags": [
          "meal"
        ],
        "summary": "Finds Meals by tag",
        "operationId": "findMealsByTag",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "Tag to filter by",
            "name": "tag",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Meal"
              }
            }
          },
          "500": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ApiError"
            }
          }
        }
      }
    },
    "/meal/{id}": {
      "get": {
        "description": "Returns a single meal",
        "produces": [
          "application/json"
        ],
        "tags": [
          "meal"
        ],
        "summary": "Find meal by ID",
        "operationId": "getMealById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of meal to return",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Meal"
            }
          },
          "500": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ApiError"
            }
          }
        }
      },
      "put": {
        "security": [
          {
            "JWT": []
          }
        ],
        "description": "the ID at the end of path is needed, but skipped",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "meal"
        ],
        "summary": "Update an existing meal",
        "operationId": "updateMeal",
        "parameters": [
          {
            "description": "Mea object that needs to be updated",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Meal"
            }
          },
          {
            "type": "integer",
            "format": "int64",
            "default": -1,
            "description": "ID for generator workaround",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Meal"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/ApiError"
            }
          },
          "500": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ApiError"
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "JWT": []
          }
        ],
        "description": "the ID at the end of path is needed, but skipped",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "meal"
        ],
        "summary": "Create a new meal",
        "operationId": "createMeal",
        "parameters": [
          {
            "description": "Meal object that needs to be created",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Meal"
            }
          },
          {
            "type": "integer",
            "format": "int64",
            "default": -1,
            "description": "ID for generator workaround",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Meal"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/ApiError"
            }
          },
          "500": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ApiError"
            }
          }
        }
      },
      "delete": {
        "security": [
          {
            "JWT": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "meal"
        ],
        "summary": "Deletes a meal",
        "operationId": "deleteMeal",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "Meal id to delete",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation"
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/ApiError"
            }
          },
          "500": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ApiError"
            }
          }
        }
      }
    },
    "/tags": {
      "get": {
        "description": "All tags are stored",
        "produces": [
          "application/json"
        ],
        "tags": [
          "meal"
        ],
        "summary": "Get all tags",
        "operationId": "getTags",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Tag"
              }
            }
          },
          "500": {
            "description": "Invalid tag value",
            "schema": {
              "$ref": "#/definitions/ApiError"
            }
          }
        }
      }
    },
    "/version": {
      "get": {
        "description": "Version anf build info",
        "produces": [
          "application/json"
        ],
        "tags": [
          "info"
        ],
        "summary": "Version",
        "operationId": "getVersion",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Version"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ApiError": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "Ingredient": {
      "description": "An ingredient for a meal",
      "type": "object",
      "title": "Meal Ingredient",
      "properties": {
        "description": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "x-go-custom-tag": "gorm:\"primaryKey\""
        },
        "name": {
          "type": "string"
        }
      }
    },
    "LoginInfo": {
      "type": "object",
      "required": [
        "email",
        "password"
      ],
      "properties": {
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "LoginSuccess": {
      "type": "object",
      "properties": {
        "success": {
          "type": "boolean"
        },
        "token": {
          "type": "string"
        }
      }
    },
    "Meal": {
      "description": "A meal",
      "type": "object",
      "title": "Meal",
      "required": [
        "name"
      ],
      "properties": {
        "description": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "x-go-custom-tag": "gorm:\"primaryKey\""
        },
        "ingredients": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Ingredient"
          },
          "x-go-custom-tag": "gorm:\"many2many:meal_ingredients\""
        },
        "kcal": {
          "type": "number"
        },
        "name": {
          "type": "string",
          "x-go-custom-tag": "valid:\"customNameValidator\""
        },
        "pictureUrl": {
          "type": "string"
        },
        "price": {
          "type": "number"
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Tag"
          },
          "x-go-custom-tag": "gorm:\"many2many:meal_tags\""
        }
      }
    },
    "Tag": {
      "description": "A tag for a meal",
      "type": "object",
      "title": "Meal Tag",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "x-go-custom-tag": "gorm:\"primaryKey\""
        },
        "name": {
          "type": "string"
        }
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "fullName": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "Version": {
      "description": "Version and build info",
      "type": "object",
      "title": "Version",
      "properties": {
        "appName": {
          "type": "string"
        },
        "buildTime": {
          "type": "string"
        },
        "goMod": {
          "type": "string"
        },
        "version": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "JWT": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "Everything about your Meals",
      "name": "meal"
    },
    {
      "description": "Providing info about service",
      "name": "info"
    },
    {
      "description": "Operations about user",
      "name": "user"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is demo for a foodstore (meals) service",
    "title": "OpenAPI Foodstore",
    "license": {
      "name": "Apache-2.0",
      "url": "https://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "foodstore.kind-01.company.com",
  "basePath": "/v1",
  "paths": {
    "/ingredients": {
      "get": {
        "description": "ll ingredients are stored",
        "produces": [
          "application/json"
        ],
        "tags": [
          "meal"
        ],
        "summary": "Get all ingredients",
        "operationId": "getIngredients",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Ingredient"
              }
            }
          },
          "500": {
            "description": "Invalid tag value",
            "schema": {
              "$ref": "#/definitions/ApiError"
            }
          }
        }
      }
    },
    "/livez": {
      "get": {
        "description": "Returns OK",
        "tags": [
          "info"
        ],
        "summary": "Liveness status for orchestrator",
        "operationId": "getLivez",
        "responses": {
          "200": {
            "description": "successful operation"
          }
        }
      }
    },
    "/login": {
      "post": {
        "description": "Returns token for authorized User",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "operationId": "login",
        "parameters": [
          {
            "description": "Login Payload",
            "name": "login",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/LoginInfo"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful login",
            "schema": {
              "$ref": "#/definitions/LoginSuccess"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "User not found",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/meal/findByTag": {
      "get": {
        "description": "One tag ID can be provided",
        "produces": [
          "application/json"
        ],
        "tags": [
          "meal"
        ],
        "summary": "Finds Meals by tag",
        "operationId": "findMealsByTag",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "Tag to filter by",
            "name": "tag",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Meal"
              }
            }
          },
          "500": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ApiError"
            }
          }
        }
      }
    },
    "/meal/{id}": {
      "get": {
        "description": "Returns a single meal",
        "produces": [
          "application/json"
        ],
        "tags": [
          "meal"
        ],
        "summary": "Find meal by ID",
        "operationId": "getMealById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of meal to return",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Meal"
            }
          },
          "500": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ApiError"
            }
          }
        }
      },
      "put": {
        "security": [
          {
            "JWT": []
          }
        ],
        "description": "the ID at the end of path is needed, but skipped",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "meal"
        ],
        "summary": "Update an existing meal",
        "operationId": "updateMeal",
        "parameters": [
          {
            "description": "Mea object that needs to be updated",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Meal"
            }
          },
          {
            "type": "integer",
            "format": "int64",
            "default": -1,
            "description": "ID for generator workaround",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Meal"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/ApiError"
            }
          },
          "500": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ApiError"
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "JWT": []
          }
        ],
        "description": "the ID at the end of path is needed, but skipped",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "meal"
        ],
        "summary": "Create a new meal",
        "operationId": "createMeal",
        "parameters": [
          {
            "description": "Meal object that needs to be created",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Meal"
            }
          },
          {
            "type": "integer",
            "format": "int64",
            "default": -1,
            "description": "ID for generator workaround",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Meal"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/ApiError"
            }
          },
          "500": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ApiError"
            }
          }
        }
      },
      "delete": {
        "security": [
          {
            "JWT": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "meal"
        ],
        "summary": "Deletes a meal",
        "operationId": "deleteMeal",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "Meal id to delete",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation"
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/ApiError"
            }
          },
          "500": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ApiError"
            }
          }
        }
      }
    },
    "/tags": {
      "get": {
        "description": "All tags are stored",
        "produces": [
          "application/json"
        ],
        "tags": [
          "meal"
        ],
        "summary": "Get all tags",
        "operationId": "getTags",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Tag"
              }
            }
          },
          "500": {
            "description": "Invalid tag value",
            "schema": {
              "$ref": "#/definitions/ApiError"
            }
          }
        }
      }
    },
    "/version": {
      "get": {
        "description": "Version anf build info",
        "produces": [
          "application/json"
        ],
        "tags": [
          "info"
        ],
        "summary": "Version",
        "operationId": "getVersion",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Version"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ApiError": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "Ingredient": {
      "description": "An ingredient for a meal",
      "type": "object",
      "title": "Meal Ingredient",
      "properties": {
        "description": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "x-go-custom-tag": "gorm:\"primaryKey\""
        },
        "name": {
          "type": "string"
        }
      }
    },
    "LoginInfo": {
      "type": "object",
      "required": [
        "email",
        "password"
      ],
      "properties": {
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "LoginSuccess": {
      "type": "object",
      "properties": {
        "success": {
          "type": "boolean"
        },
        "token": {
          "type": "string"
        }
      }
    },
    "Meal": {
      "description": "A meal",
      "type": "object",
      "title": "Meal",
      "required": [
        "name"
      ],
      "properties": {
        "description": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "x-go-custom-tag": "gorm:\"primaryKey\""
        },
        "ingredients": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Ingredient"
          },
          "x-go-custom-tag": "gorm:\"many2many:meal_ingredients\""
        },
        "kcal": {
          "type": "number"
        },
        "name": {
          "type": "string",
          "x-go-custom-tag": "valid:\"customNameValidator\""
        },
        "pictureUrl": {
          "type": "string"
        },
        "price": {
          "type": "number"
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Tag"
          },
          "x-go-custom-tag": "gorm:\"many2many:meal_tags\""
        }
      }
    },
    "Tag": {
      "description": "A tag for a meal",
      "type": "object",
      "title": "Meal Tag",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "x-go-custom-tag": "gorm:\"primaryKey\""
        },
        "name": {
          "type": "string"
        }
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "fullName": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "Version": {
      "description": "Version and build info",
      "type": "object",
      "title": "Version",
      "properties": {
        "appName": {
          "type": "string"
        },
        "buildTime": {
          "type": "string"
        },
        "goMod": {
          "type": "string"
        },
        "version": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "JWT": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "Everything about your Meals",
      "name": "meal"
    },
    {
      "description": "Providing info about service",
      "name": "info"
    },
    {
      "description": "Operations about user",
      "name": "user"
    }
  ]
}`))
}

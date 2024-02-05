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
        "/api/v1/create/user": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new user with username, password, and email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create new user",
                "parameters": [
                    {
                        "description": "Create User Command",
                        "name": "createUserCommand",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.CreateUserCommand"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User created successfully",
                        "schema": {
                            "$ref": "#/definitions/domain.CreateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request - invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized - invalid credentials",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/delete/user/{id}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Deletes a user with the specified ID from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/domain.DeleteUserResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid user ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Error deleting user",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/login": {
            "post": {
                "description": "Handles user login and returns a JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "User Login",
                "parameters": [
                    {
                        "description": "Login Credentials",
                        "name": "loginBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.LoginCommand"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JWT Token returned on successful authentication",
                        "schema": {
                            "$ref": "#/definitions/domain.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/user/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieve a single user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get User by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful retrieval of user information",
                        "schema": {
                            "$ref": "#/definitions/TradingBot_domain_getUser.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/users": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieve a list of all users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.GetUserAllResponse"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/users/{id}/email": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update user email by user ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update user email",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Email Object",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Email"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Email updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "TradingBot_domain_getAllUsers.UserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "TradingBot_domain_getUser.UserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "domain.CreateResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "domain.CreateUserCommand": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "domain.DeleteUserResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "domain.Email": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "user@example.com"
                }
            }
        },
        "domain.GetUserAllResponse": {
            "type": "object",
            "properties": {
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/TradingBot_domain_getAllUsers.UserResponse"
                    }
                }
            }
        },
        "domain.LoginCommand": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "domain.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "TradingBot API",
	Description:      "This is a sample server for TradingBot.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/login": {
            "post": {
                "description": "Generate JWT is user is authenticated",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Authenticate user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg_auth_handlers.ErrResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/pb.LoginResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/pb.LoginResponse"
                        }
                    }
                }
            }
        },
        "/project": {
            "get": {
                "description": "view all project",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "View Projects",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/pb.ProjectsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg_task_handlers.ErrResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/pb.ProjectsResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pb.ProjectsResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new project",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "Create Project",
                "parameters": [
                    {
                        "description": "ProjectReq",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.ProjectReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/pb.ProjectResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg_task_handlers.ErrResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/pb.ProjectResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pb.ProjectResponse"
                        }
                    }
                }
            }
        },
        "/project/{ID}": {
            "get": {
                "description": "view a project by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "View Project",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/pb.ProjectResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg_task_handlers.ErrResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/pb.ProjectResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pb.ProjectResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a project",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "Update Project",
                "parameters": [
                    {
                        "description": "ProjectReq",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.ProjectReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/pb.ProjectResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg_task_handlers.ErrResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/pb.ProjectResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pb.ProjectResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "remove a project",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "Remove Project",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/pb.RemovedProjectResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg_task_handlers.ErrResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/pb.RemovedProjectResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pb.RemovedProjectResponse"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "RegisterRequest",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/pb.RegisterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg_auth_handlers.ErrResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/pb.RegisterResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pb.RegisterResponse"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "description": "Get user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "View user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.ViewUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg_auth_handlers.ErrResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/pkg_auth_handlers.ErrResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/pb.ViewUserResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pb.ViewUserResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.ProjectReq": {
            "type": "object"
        },
        "handlers.RegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "firstname",
                "lastname",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "johnsmith@example.com"
                },
                "firstname": {
                    "type": "string",
                    "example": "john"
                },
                "lastname": {
                    "type": "string",
                    "example": "smith"
                },
                "password": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 8,
                    "example": "John123!"
                }
            }
        },
        "pb.CommonResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "pb.LoginResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "pb.Project": {
            "type": "object",
            "properties": {
                "assignees": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "created_on": {
                    "type": "string"
                },
                "creator": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "tasks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/pb.Task"
                    }
                },
                "updated_on": {
                    "type": "string"
                }
            }
        },
        "pb.ProjectResponse": {
            "type": "object",
            "properties": {
                "common_response": {
                    "$ref": "#/definitions/pb.CommonResponse"
                },
                "project": {
                    "$ref": "#/definitions/pb.Project"
                }
            }
        },
        "pb.ProjectsResponse": {
            "type": "object",
            "properties": {
                "common_response": {
                    "$ref": "#/definitions/pb.CommonResponse"
                },
                "project": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/pb.Project"
                    }
                }
            }
        },
        "pb.RegisterResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "pb.RemovedProjectResponse": {
            "type": "object",
            "properties": {
                "common_response": {
                    "$ref": "#/definitions/pb.CommonResponse"
                }
            }
        },
        "pb.Status": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3
            ],
            "x-enum-varnames": [
                "Status_TODO",
                "Status_INPROGRESS",
                "Status_READY",
                "Status_COMPLETED"
            ]
        },
        "pb.Task": {
            "type": "object",
            "properties": {
                "assignee": {
                    "type": "integer"
                },
                "creator": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "task_status": {
                    "$ref": "#/definitions/pb.Status"
                }
            }
        },
        "pb.ViewUserResponse": {
            "type": "object",
            "properties": {
                "activitystatus": {
                    "type": "boolean"
                },
                "email": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "pkg_auth_handlers.ErrResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "pkg_task_handlers.ErrResponse": {
            "type": "object",
            "properties": {
                "error": {
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

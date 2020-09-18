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
  "consumes": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Secen Pick Server",
    "version": "1.0.0"
  },
  "paths": {
    "/dialog": {
      "get": {
        "operationId": "getDialog",
        "parameters": [
          {
            "type": "string",
            "description": "““, “anime“, “manga“, “book“のうちのどれか",
            "name": "genre",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "オフセット。初めは0",
            "name": "offset",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "取得する記録の数",
            "name": "limit",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "デフォで新着順。””, “like“, “comment“, “combined“",
            "name": "sort",
            "in": "query"
          },
          {
            "type": "string",
            "description": "#スタートの場合はタグ, そうでない場合は文字列と判断",
            "name": "q",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "return dialogs",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                },
                "schema": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/dialog"
                  }
                }
              }
            }
          },
          "400": {
            "description": "request error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "internal serever error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "operationId": "postDialog",
        "parameters": [
          {
            "type": "string",
            "name": "token",
            "in": "header",
            "required": true
          },
          {
            "description": "セリフ投稿時に必要なパラメータ",
            "name": "content",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "author": {
                  "type": "string"
                },
                "comment": {
                  "type": "string"
                },
                "content": {
                  "type": "string"
                },
                "link": {
                  "type": "string"
                },
                "style": {
                  "type": "string"
                },
                "title": {
                  "type": "string"
                },
                "user_id": {
                  "type": "integer"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "投稿成功",
            "schema": {
              "type": "object",
              "properties": {
                "id": {
                  "type": "integer"
                },
                "message": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "request error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "internal serever error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/dialog/{id}/comment": {
      "get": {
        "operationId": "getCommentById",
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "オフセット。初めは0",
            "name": "offset",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "取得する記録の数",
            "name": "limit",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "取得成功",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                },
                "schema": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/comment"
                  }
                }
              }
            }
          },
          "400": {
            "description": "request error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "internal serever error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "operationId": "postCommentById",
        "parameters": [
          {
            "type": "string",
            "name": "token",
            "in": "header",
            "required": true
          },
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "name": "comment",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "comment": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "投稿成功",
            "schema": {
              "type": "object",
              "properties": {
                "id": {
                  "type": "integer"
                },
                "message": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "request error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "internal serever error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/tag": {
      "get": {
        "operationId": "getTag",
        "parameters": [
          {
            "type": "string",
            "description": "“title“, “author“, “speaker“, “other“のうちのどれか。typeによる絞り込み。空欄だと全部。",
            "name": "genre",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "オフセット。初めは0",
            "name": "offset",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "取得する記録の数",
            "name": "limit",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "デフォで新着順。",
            "name": "sort",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "取得成功",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                },
                "schema": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/tag"
                  }
                }
              }
            }
          },
          "400": {
            "description": "request error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "internal serever error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "operationId": "postTag",
        "parameters": [
          {
            "type": "string",
            "name": "token",
            "in": "header",
            "required": true
          },
          {
            "name": "tag",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "name": {
                  "type": "string"
                },
                "type": {
                  "description": "“title“, “author“, “speaker“ “other“のうちのどれか。typeによる絞り込み。空欄だと全部。",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "投稿成功",
            "schema": {
              "type": "object",
              "properties": {
                "id": {
                  "type": "integer"
                },
                "message": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "request error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "internal serever error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "comment": {
      "type": "object",
      "properties": {
        "content": {
          "type": "string"
        },
        "user": {
          "$ref": "#/definitions/user"
        }
      }
    },
    "dialog": {
      "type": "object",
      "properties": {
        "author": {
          "type": "string"
        },
        "content": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "link": {
          "type": "string"
        },
        "source": {
          "type": "string"
        },
        "style": {
          "type": "string"
        },
        "title": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "string"
    },
    "principal": {
      "type": "string"
    },
    "tag": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "user": {
      "type": "object",
      "properties": {
        "display_name": {
          "description": "ユーザ名",
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "photo_url": {
          "description": "画像の保存先URL",
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "tokenAuth": {
      "type": "apiKey",
      "name": "X-API-Key",
      "in": "header"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Secen Pick Server",
    "version": "1.0.0"
  },
  "paths": {
    "/dialog": {
      "get": {
        "operationId": "getDialog",
        "parameters": [
          {
            "type": "string",
            "description": "““, “anime“, “manga“, “book“のうちのどれか",
            "name": "genre",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "オフセット。初めは0",
            "name": "offset",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "取得する記録の数",
            "name": "limit",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "デフォで新着順。””, “like“, “comment“, “combined“",
            "name": "sort",
            "in": "query"
          },
          {
            "type": "string",
            "description": "#スタートの場合はタグ, そうでない場合は文字列と判断",
            "name": "q",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "return dialogs",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                },
                "schema": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/dialog"
                  }
                }
              }
            }
          },
          "400": {
            "description": "request error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "internal serever error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "operationId": "postDialog",
        "parameters": [
          {
            "type": "string",
            "name": "token",
            "in": "header",
            "required": true
          },
          {
            "description": "セリフ投稿時に必要なパラメータ",
            "name": "content",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "author": {
                  "type": "string"
                },
                "comment": {
                  "type": "string"
                },
                "content": {
                  "type": "string"
                },
                "link": {
                  "type": "string"
                },
                "style": {
                  "type": "string"
                },
                "title": {
                  "type": "string"
                },
                "user_id": {
                  "type": "integer"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "投稿成功",
            "schema": {
              "type": "object",
              "properties": {
                "id": {
                  "type": "integer"
                },
                "message": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "request error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "internal serever error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/dialog/{id}/comment": {
      "get": {
        "operationId": "getCommentById",
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "オフセット。初めは0",
            "name": "offset",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "取得する記録の数",
            "name": "limit",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "取得成功",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                },
                "schema": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/comment"
                  }
                }
              }
            }
          },
          "400": {
            "description": "request error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "internal serever error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "operationId": "postCommentById",
        "parameters": [
          {
            "type": "string",
            "name": "token",
            "in": "header",
            "required": true
          },
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "name": "comment",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "comment": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "投稿成功",
            "schema": {
              "type": "object",
              "properties": {
                "id": {
                  "type": "integer"
                },
                "message": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "request error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "internal serever error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/tag": {
      "get": {
        "operationId": "getTag",
        "parameters": [
          {
            "type": "string",
            "description": "“title“, “author“, “speaker“, “other“のうちのどれか。typeによる絞り込み。空欄だと全部。",
            "name": "genre",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "オフセット。初めは0",
            "name": "offset",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "取得する記録の数",
            "name": "limit",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "デフォで新着順。",
            "name": "sort",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "取得成功",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                },
                "schema": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/tag"
                  }
                }
              }
            }
          },
          "400": {
            "description": "request error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "internal serever error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "operationId": "postTag",
        "parameters": [
          {
            "type": "string",
            "name": "token",
            "in": "header",
            "required": true
          },
          {
            "name": "tag",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "name": {
                  "type": "string"
                },
                "type": {
                  "description": "“title“, “author“, “speaker“ “other“のうちのどれか。typeによる絞り込み。空欄だと全部。",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "投稿成功",
            "schema": {
              "type": "object",
              "properties": {
                "id": {
                  "type": "integer"
                },
                "message": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "request error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "internal serever error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "comment": {
      "type": "object",
      "properties": {
        "content": {
          "type": "string"
        },
        "user": {
          "$ref": "#/definitions/user"
        }
      }
    },
    "dialog": {
      "type": "object",
      "properties": {
        "author": {
          "type": "string"
        },
        "content": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "link": {
          "type": "string"
        },
        "source": {
          "type": "string"
        },
        "style": {
          "type": "string"
        },
        "title": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "string"
    },
    "principal": {
      "type": "string"
    },
    "tag": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "user": {
      "type": "object",
      "properties": {
        "display_name": {
          "description": "ユーザ名",
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "photo_url": {
          "description": "画像の保存先URL",
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "tokenAuth": {
      "type": "apiKey",
      "name": "X-API-Key",
      "in": "header"
    }
  }
}`))
}

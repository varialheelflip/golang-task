## 项目说明
```bash
// 运行环境 go 1.23
// 依赖安装
// clone项目后进入到项目根目录
go mod tidy
// 配置启动参数(config.yaml)
// 项目启动
go run main.go
```

## 启动参数
```yaml
app:
  port: 8080 //启动端口


database:
  host: "192.168.200.130" // 数据库相关配置
  port: 3306
  dbname: "gorm"
  user: "root"
  password: "root"
  debug: true

jwt:
  secret_key: "your_secret_key" // jwt令牌密钥
```

## 相关接口(postman)
### 用户接口
```json
{
  "info": {
    "_postman_id": "7420b83d-609d-42df-b9f5-ba808c211860",
    "name": "用户",
    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
    "_exporter_id": "9523421"
  },
  "item": [
    {
      "name": "注册",
      "request": {
        "method": "POST",
        "header": [],
        "body": {
          "mode": "raw",
          "raw": "{\r\n    \"username\": \"123456789\",\r\n    \"password\": \"pwd1234\",\r\n    \"email\": \"2345678@qq.com\"\r\n}",
          "options": {
            "raw": {
              "language": "json"
            }
          }
        },
        "url": {
          "raw": "localhost:8080/api/v1/users/register",
          "host": [
            "localhost"
          ],
          "port": "8080",
          "path": [
            "api",
            "v1",
            "users",
            "register"
          ]
        }
      },
      "response": []
    },
    {
      "name": "登录",
      "request": {
        "method": "POST",
        "header": [],
        "body": {
          "mode": "raw",
          "raw": "{\r\n    \"username\": \"abcde\",\r\n    \"password\": \"pwd1234\"\r\n}",
          "options": {
            "raw": {
              "language": "json"
            }
          }
        },
        "url": {
          "raw": "localhost:8080/api/v1/users/login",
          "host": [
            "localhost"
          ],
          "port": "8080",
          "path": [
            "api",
            "v1",
            "users",
            "login"
          ]
        }
      },
      "response": []
    }
  ]
}
```

### 文章接口
```json
{
	"info": {
		"_postman_id": "676fdbf5-e2ad-48fb-84bd-c9c191c0bedf",
		"name": "文章",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "9523421"
	},
	"item": [
		{
			"name": "新建文章",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\r\n    \"title\": \"文章标题\",\r\n    \"content\": \"文章内容\"\r\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:8080/api/v1/posts",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"api",
						"v1",
						"posts"
					]
				}
			},
			"response": []
		},
		{
			"name": "更新文章",
			"request": {
				"method": "PUT",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\r\n    \"id\": 14,\r\n    \"title\": \"文章标题2\",\r\n    \"content\": \"文章内容2\"\r\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:8080/api/v1/posts",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"api",
						"v1",
						"posts"
					]
				}
			},
			"response": []
		},
		{
			"name": "删除文章",
			"request": {
				"method": "DELETE",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\r\n    \"title\": \"文章标题\",\r\n    \"content\": \"文章内容\"\r\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:8080/api/v1/posts/13",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"api",
						"v1",
						"posts",
						"13"
					]
				}
			},
			"response": []
		},
		{
			"name": "查询文章",
			"protocolProfileBehavior": {
				"disableBodyPruning": true
			},
			"request": {
				"method": "GET",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\r\n    \"title\": \"文章标题\",\r\n    \"content\": \"文章内容\"\r\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:8080/api/v1/posts/14",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"api",
						"v1",
						"posts",
						"14"
					]
				}
			},
			"response": []
		},
		{
			"name": "分页查询文章",
			"protocolProfileBehavior": {
				"disableBodyPruning": true
			},
			"request": {
				"method": "GET",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\r\n    \"title\": \"文章标题\",\r\n    \"content\": \"文章内容\"\r\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:8080/api/v1/posts/page?pageNo=1&pageSize=5&userId=4",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"api",
						"v1",
						"posts",
						"page"
					],
					"query": [
						{
							"key": "pageNo",
							"value": "1"
						},
						{
							"key": "pageSize",
							"value": "5"
						},
						{
							"key": "userId",
							"value": "4"
						}
					]
				}
			},
			"response": []
		}
	],
	"event": [
		{
			"listen": "prerequest",
			"script": {
				"type": "text/javascript",
				"packages": {},
				"exec": [
					"pm.request.headers.add({key: 'token', value: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTE4ODIxMzEsImlkIjo0LCJ1c2VybmFtZSI6ImFiY2RlIn0.iYTkeIN4PBD5UbRrOsJY4Wi2KnTVRE_wj7Q7ov4Mx4c' });"
				]
			}
		},
		{
			"listen": "test",
			"script": {
				"type": "text/javascript",
				"packages": {},
				"exec": [
					""
				]
			}
		}
	]
}
```

### 评论接口
```json
{
	"info": {
		"_postman_id": "f626f2dd-d492-4d3f-ac70-c270583a7173",
		"name": "评论",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "9523421"
	},
	"item": [
		{
			"name": "新建评论",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\r\n    \"postId\": 5,\r\n    \"content\": \"评论1\"\r\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:8080/api/v1/comments",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"api",
						"v1",
						"comments"
					]
				}
			},
			"response": []
		},
		{
			"name": "查询评论",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:8080/api/v1/comments/5",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"api",
						"v1",
						"comments",
						"5"
					]
				}
			},
			"response": []
		}
	],
	"event": [
		{
			"listen": "prerequest",
			"script": {
				"type": "text/javascript",
				"packages": {},
				"exec": [
					"pm.request.headers.add({key: 'token', value: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTE4MjE3MzAsImlkIjo0LCJ1c2VybmFtZSI6ImFiY2RlIn0.fqXcJKTrXBQxm9A2ymyveA1yC5Jc8r8egW2I_iT2-B0' });"
				]
			}
		},
		{
			"listen": "test",
			"script": {
				"type": "text/javascript",
				"packages": {},
				"exec": [
					""
				]
			}
		}
	]
}
```
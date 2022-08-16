# bluebell

- 项目结构



- 目录结构

  ```
  .
  ├── bluebell.log
  ├── config.yaml
  ├── controller
  │   ├── code.go
  │   ├── community.go
  │   ├── docs_models.go
  │   ├── post.go
  │   ├── request.go
  │   ├── response.go
  │   ├── user.go
  │   ├── validator.go
  │   └── vote.go
  ├── dao
  │   ├── mysql
  │   │   ├── community.go
  │   │   ├── error_code.go
  │   │   ├── mysql.go
  │   │   ├── post.go
  │   │   └── user.go
  │   └── redis
  │       ├── key.go
  │       ├── post.go
  │       ├── redis.go
  │       └── vote.go
  ├── docs
  │   ├── docs.go
  │   ├── swagger.json
  │   └── swagger.yaml
  ├── go.mod
  ├── go.sum
  ├── logger
  │   └── logger.go
  ├── logic
  │   ├── community.go
  │   ├── post.go
  │   ├── user.go
  │   └── vote.go
  ├── main.go
  ├── middlewares
  │   └── auth.go
  ├── models
  │   ├── community.go
  │   ├── create_table.sql
  │   ├── params.go
  │   ├── post.go
  │   └── user.go
  ├── pkg
  │   ├── jwt
  │   │   └── jwt.go
  │   └── snowflake
  │       └── snowflake.go
  ├── routes
  │   └── routes.go
  └── settings
      └── settings.go
  
  14 directories, 41 files
  
  ```

  

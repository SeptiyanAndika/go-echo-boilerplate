# go-echo-starter
Simple boilerplate modular using Echo Framework (Golang)

### HOW-TO

Before you start, please make sure that you have install `glide` and start `mysql`. Afterward, please follow below steps.

1. Clone this project.
2. In current directory, execute `glide install`.
3. Update mysql config in `config.toml`.
3. Last, `go run main.go`

### Middleware Authorizer

Middleware Authorizer can support multiple roles, like in  in `modules/user/main.go` 

1. e.GET("/restricted", restricted, utils.Authorizer()) // all roles can acess
2. e.GET("/restricted-user", restricted, utils.Authorizer("user")) // only roles  user can access
3. e.GET("/restricted-admin", restricted, utils.Authorizer("admin")) // only roles admin can access
3. e.GET("/restricted-user-admin", restricted, utils.Authorizer("admin", "user")) // roles admin and rols user can access

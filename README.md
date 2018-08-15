# QSC Recruit Open Platform (*rop*) in Go

<img align="right" width="159px" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png">

ROP-go is a web project written in Go (Golang). It features a lot of APIs with robust behavior, high performance, and easily read code.
If you are interested in this project, any contricution is welcomed.

## Tech Stack
Gin, Gorm, Vendor, JWT, Viper.

## Routers
```go
auth := g.Group("/v1/auth")
{
    auth.POST("/login", user.Login)
}

userGroup := g.Group("/v1/user")
userGroup.Use(middleware.AuthMiddleware())
{
    userGroup.GET("/info", user.Info)
}

insGroup := g.Group("/v1/instance")
insGroup.Use(middleware.AuthMiddleware())
{
    insGroup.POST("", instance.Create)
    insGroup.GET("", instance.List)
    insGroup.PUT("/:id", instance.Update)
}

formGroup := g.Group("/v1/form")
insGroup.Use(middleware.AuthMiddleware())
{
    formGroup.POST("", form.Create)
    formGroup.GET("", form.List)
    formGroup.PUT("/:id", form.Update)
}

freGroup := g.Group("/v1/freshman")
freGroup.Use(middleware.AuthMiddleware())
{
    freGroup.POST("/submit/:instanceId", freshman.Submit)
}

svcd := g.Group("/sd")
{
    svcd.GET("/health", sd.HealthCheck)
    svcd.GET("/disk", sd.DiskCheck)
    svcd.GET("/cpu", sd.CPUCheck)
    svcd.GET("/ram", sd.RAMCheck)
}
```

## Code Amount
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              34            210             63           1220
XML                              6              0              0            872
YAML                             2              0              0             78
-------------------------------------------------------------------------------
SUM:                            42            210             63           2170
-------------------------------------------------------------------------------


## Directory Tree
.
├── README.md
├── conf
│   ├── config.yaml
│   └── config_sample.yaml
├── config
│   └── config.go
├── handler
│   ├── form
│   │   ├── create.go
│   │   ├── form.go
│   │   ├── list.go
│   │   └── update.go
│   ├── freshman
│   │   ├── freshman.go
│   │   └── submit.go
│   ├── handler.go
│   ├── instance
│   │   ├── create.go
│   │   ├── instance.go
│   │   ├── list.go
│   │   └── update.go
│   ├── sd
│   │   └── check.go
│   └── user
│       ├── info.go
│       ├── login.go
│       └── user.go
├── log
├── main.go
├── model
│   ├── form.go
│   ├── freshman.go
│   ├── init.go
│   ├── instance.go
│   ├── intent.go
│   ├── interview.go
│   └── user.go
├── pkg
│   ├── auth
│   ├── errno
│   │   ├── code.go
│   │   └── errno.go
│   ├── timerange
│   │   └── time.go
│   └── token
│       └── token.go
├── router
│   ├── middleware
│   │   ├── auth.go
│   │   ├── header.go
│   │   ├── logging.go
│   │   └── requestid.go
│   └── router.go
└── service
    └── SMS.go

18 directories, 37 files
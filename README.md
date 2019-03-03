# QSC Recruit Open Platform (*rop*) in Go

<img align="right" width="159px" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png">

ROP-go is a web project written in Go (Golang). It features a lot of APIs with robust behavior, high performance, and easily read code.
If you are interested in this project, any contricution is welcomed.

## Tech Stack
Gin, Gorm, Vendor, JWT, Viper.

## Constant
substage 1 => INIT, 2 <= ASSIGNED
AutoJoinable 1 => OK， -1 <= BAN

## Routers
```go
auth := g.Group("/v1/auth")
{
    auth.GET("/association", association.ListAssociaton)
    auth.POST("/login/qsc", user.LoginByQSC)
    auth.POST("/login/passwd", user.LoginByPassword)
}

userGroup := g.Group("/v1/user")
userGroup.Use(middleware.AuthMiddleware())
{
    userGroup.GET("/info", user.Info)
    userGroup.POST("/avatar", user.Avatar)
    userGroup.GET("/sms", user.GetSMSAccount)
}

insGroup := g.Group("/v1/instance")
insGroup.Use(middleware.AuthMiddleware())
{
    insGroup.POST("", instance.Create)
    insGroup.GET("", instance.List)
    insGroup.PUT("/:id", instance.Update)
}

formGroup := g.Group("/v1/form")
formGroup.Use(middleware.AuthMiddleware())
{
    formGroup.POST("", form.Create)
    formGroup.GET("", form.List)
    formGroup.PUT("/:id", form.Update)
}

freGroup := g.Group("/v1/freshman")
//freGroup.Use(middleware.AuthMiddleware())
{
    freGroup.POST("/submit", freshman.Submit)
}

intentGroup := g.Group("/v1/intent")
intentGroup.Use(middleware.AuthMiddleware())
{
    intentGroup.POST("/assign", intent.Assign)
    intentGroup.POST("/reject/:id", intent.Reject)
    intentGroup.GET("", intent.List)
}

interviewGroup := g.Group("/v1/interview")
interviewGroup.Use(middleware.AuthMiddleware())
{
    interviewGroup.POST("", interview.Create)
    interviewGroup.PUT("/:id", interview.Update)
    interviewGroup.GET("", interview.List)
    interviewGroup.POST("/enable/:id", interview.Enable)
    interviewGroup.POST("/disable/:id", interview.Disable)
}

associationGroup := g.Group("/v1/association")
associationGroup.Use(middleware.AuthMiddleware())
{
    associationGroup.POST("", association.Create)
    associationGroup.GET("", association.ListAssociaton)
    associationGroup.GET("/name/:name", association.Get)
    associationGroup.GET("/user", association.GetUser)
    associationGroup.POST("/user", association.AddUser)
    associationGroup.POST("/notice", association.SendNotice)
}

ssrGroup := g.Group("/v1/ssr")
{
    ssrGroup.GET("/schedule", ssr.Schedule)
    ssrGroup.POST("/join/:id", interview.Join)
    ssrGroup.GET("/form", ssr.GetFormByIns)
    ssrGroup.POST("/reject/:id", interview.Reject)
    ssrGroup.GET("/instance", instance.List)
    ssrGroup.GET("/register", ssr.Register)
    ssrGroup.POST("/register", user.Register)
}

fileGroup := g.Group("/v1/file")
{
    fileGroup.POST("/upload/img", file.UploadImage)
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

```
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              70            545            265           3195
XML                              5              0              0            896
YAML                             4              6             19            150
Markdown                         1             18              0            150
Dockerfile                       1              5              0              6
-------------------------------------------------------------------------------
SUM:                            81            574            284           4397
-------------------------------------------------------------------------------
```

## Directory Tree
```
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
```
# gin-dj
This is a Golang [Gin web frameworks](https://github.com/gin-gonic/gin) template that provides a Django-like folder structure and command line tools for initializing projects and creating apps. It also includes some common and useful middleware that are set as default.

## Prerequisites
This project are made depeneds on following packages.

- [Gin](https://github.com/gin-gonic/gin)

- [Gin-Contrib-Cors](https://github.com/gin-contrib/cors)

- [Gorm](https://github.com/go-gorm/gorm)

- [Gorilla-csrf](https://github.com/gorilla/csrf)

- [Gin-adapter](https://github.com/gwatts/gin-adapter)

- [GormPostgresDriver](https://github.com/go-gorm/postgres) (Optional)

## Installation
Run following command to install the Golang Gin web framework template.

```bash
go install github.com/Zncl2222/gin-dj@latest
```

Or clone this project and run the following command in project's root directory

```bash
go install .
```

## Quick Start
Use the command line tools to initialize project and create app like what django-admin and manage.py does.

Create project
```bash
gin-dj init example-project
```
Change directory to your project directory
```bash
cd example-project
```

Create custom app
```bash
gin-dj createapp example-app
```

## Folder Structure
```
example-project/
├── example-app/
|   ├── models.go
|   ├── serializers.go
|   ├── urls.go
|   ├── views.go
|   ├── ...
├── core/
│   ├── settings.go
├── urls.go
└── main.go
```

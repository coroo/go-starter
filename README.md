# Go Limonilo

:bulb: Create a backend using GoLang for User Profile CRUD (UserID, Email, Address, Password) and Login Process (Username & Password) Process which follow SOLID Principles.

> Please read [how the app work section](#mag-how-the-app-works-) for overall flow or view this video below:

<a href="https://www.loom.com/share/bf4b9e7e4c944158a5179c1837301bad" target="_blank"> <img style="max-width:300px;" src="https://cdn.loom.com/sessions/thumbnails/bf4b9e7e4c944158a5179c1837301bad-with-play.gif"> <p>Create your own dynamic form - Watch Video</p> </a>

### App Information
This app use [gin](https://github.com/gin-gonic/gin) as a framework, [gorm](https://gorm.io/index.html) as an ORM and [go-swagger](https://github.com/go-swagger/go-swagger) as API documentation.

# :page_facing_up: Installation

1. Go installation: 
- Download: https://golang.org/dl/
- Instruction: https://golang.org/doc/install#install

2. You need to have database and set it later in .env:
```
Create your own form-generator database, lets call it `enigmaschool`
```

## Install Swagger Library
```bash
go get -u github.com/swaggo/swag/cmd/swag
```

## Easy Setup & Run Go Limonilo
For make it easier to do configuration in Go Limonilo, I create shell script that you can use by type:
```bash
./run.sh
```

> If you need to run manual, please see [manual handling section](#wheelchair-manual-handling)

![Go Limonilo Configuration](storage/assets/img/form-generator-run-sh.png)

You need to use this step-by-step for running the app at the first time:
```go
1. Set Up Environment (.env) 
// Used for set-up app environment. For the first setup, you need to change your environment detail. For more information about environment that you need to add, please contact developer.

2. Do Unit Test 
// You can always running unit test by this feature

3. Update Swagger Documentation 
// You use swagger for maintain API restful documentation. You can check it later after running app (4. Run Go Limonilo) and redirect to your app_link/swagger/index.html

4. Run Go Limonilo 
// This feature is used for running the app
```

## API Documentation
> This feature only can be used after running the app
```bash
redirect to this link --> /swagger/index.html
```
e.g. http://localhost:3000/swagger/index.html

![Go Limonilo Swagger Viewer](storage/assets/img/form-generator-swagger-view.png)

---

# :mag: How The App Works ?

![Go Limonilo Diagram](storage/assets/img/form-generator-diagram.png)

---

# :wheelchair: Manual Handling

## Set Environment
```bash
cp .env.example .env
```
> Fulfill your environment setup

## Unit Test
```
go test -coverprofile cp.out -v ./... && go tool cover -html=cp.out
```

## Update Swagger Documentation
```bash
swag init
```
or
```bash
{GOLANG_PATH}/go/bin/swag init 
//e.g. /UserProfiles/kuncoro.barot/go/bin/swag init
```

## Running App
```bash
go run main.go
```


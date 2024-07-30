# GOLANG HTTP SERVER

This is a simple Go server using the Gin framework. It serves two HTML files: `home.html` and `about.html`.

## Prerequisites

- [Go](https://golang.org/doc/install)
- [Git](https://git-scm.com/downloads)

## Project Structure

myapp/
├── main.go
├── templates/
│ ├── home.html
│ └── about.html

## Getting Started

### Step 1: Initialize the project

```bash
mkdir myapp
cd myapp

go mod init myapp
```
### Step 2: install dependencies:
```
go get -u github.com/gin-gonic/gin
go mod tidy
```

### Step 3: run the server:
```
go run main.go
```

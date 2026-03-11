# Go Naming Conventions

Go follows simple and consistent naming conventions. These conventions improve readability and are widely used in the Go ecosystem.

---

## 1. Package Names

**Rules**

- Use short, lowercase names
- Avoid underscores and camelCase
- Should describe the package purpose
- Usually a single word

### Good

```go
package math
package utils
package handler
package auth
```

### Avoid

```go
package MathUtils
package user_service
package userService
```

---

## 2. Variable Names

Variables follow **camelCase** naming.

### Example

```go
userName := "Kalyan"
totalAmount := 100
isLoggedIn := true
```

Short variable names are acceptable in small scopes.

```go
i := 0
x := 10
```

But for larger scopes use descriptive names.

```go
userCount := 10
orderTotal := 500
```

---

## 3. Function Names

Functions use **camelCase**.

### Example

```go
func getUser() {}
func calculateTotal() {}
func sendEmail() {}
```

For exported functions use **PascalCase**.

```go
func GetUser() {}
func CalculateTotal() {}
```

---

## 4. Exported vs Unexported

In Go, **capitalization controls visibility**.

### Exported (Public)

Accessible from other packages.

```go
func GetUser() {}
type User struct {}
var AppConfig string
```

### Unexported (Private)

Accessible only inside the same package.

```go
func getUser() {}
type user struct {}
var appConfig string
```

Example:

```go
package user

func GetUser() {}       // public
func validateUser() {}  // private
```

---

## 5. Struct Names

Struct names use **PascalCase**.

### Example

```go
type User struct {
    ID   int
    Name string
}
```

Avoid:

```go
type user_data struct{}
```

---

## 6. Interface Names

Interfaces often end with **-er**.

### Example

```go
type Reader interface {}
type Writer interface {}
type Logger interface {}
type UserService interface {}
```

Example with method:

```go
type Storage interface {
    Save(data string) error
}
```

---

## 7. Constants

Constants usually use **PascalCase**.

### Example

```go
const MaxUsers = 100
const DefaultTimeout = 30
```

Grouped constants:

```go
const (
    StatusActive  = "active"
    StatusPending = "pending"
)
```

---

## 8. Acronyms

Acronyms should remain **fully capitalized**.

### Correct

```go
userID
HTTPServer
APIClient
```

### Incorrect

```go
userId
HttpServer
ApiClient
```

---

## 9. File Names

File names are usually **lowercase**.

Underscores may be used if needed.

### Example

```go
user_service.go
auth_handler.go
db_connection.go
```

---

## 10. Receiver Names

Receiver variables should be **short abbreviations** of the struct name.

### Example

```go
type User struct {
    Name string
}

func (u User) GetName() string {
    return u.Name
}
```

Avoid:

```go
func (this User) GetName() {}
func (self User) GetName() {}
```

Correct:

```go
func (u User) GetName() {}
```

---

# Example Go Code

```go
package user

type User struct {
    ID   int
    Name string
}

func GetUser(id int) User {
    return User{
        ID:   id,
        Name: "Kalyan",
    }
}

func validateUser(u User) bool {
    return u.Name != ""
}
```

---

# Quick Cheat Sheet

| Item | Convention | Example |
|-----|------------|---------|
| Package | lowercase | `package auth` |
| Variables | camelCase | `userName` |
| Functions | camelCase | `getUser()` |
| Exported | PascalCase | `GetUser()` |
| Structs | PascalCase | `UserService` |
| Interfaces | -er pattern | `Reader` |
| Constants | PascalCase | `MaxSize` |
| Files | lowercase | `user_service.go` |
| Receiver | short name | `(u User)` |

---

# Summary

Go naming conventions focus on:

- Simplicity
- Readability
- Consistency
- Capitalization for visibility control

Following these conventions ensures your code aligns with the Go community and makes it easier for others to understand and maintain.


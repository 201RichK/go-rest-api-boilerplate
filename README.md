# ARO Go (Golang) REST 

- Endpoint (REST API Swagger): http://localhost:8080/swagger/index.html

**Used libraries:**
- [gin](https://github.com/gin-gonic)
- [gin-swagger](https://github.com/swaggo/gin-swagger)
- [gorm](https://gorm.io/docs/)
- [godotenv](https://pkg.go.dev/github.com/joho/godotenv?tab=doc)
- [i18n](https://github.com/gin-contrib/i18n)
<!-- - [testify](https://github.com/stretchr/testify) -->
<!-- - [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) -->


### Features

- [ ] User Auth functionality (Signup, Login, Forgot Password, Reset Password)
- [ ] JWT Authentication
- [x] REST API
- [x] Gorm (Golang SQL DB ORM) with Postgres implementation and auto migration
- [x] Configs via environmental variables
- [ ] Email notification (Welcome email, Reset password email)
- [x] Swagger REST API documentation
- [ ] Unit tests
- [ ] Dependency injection
- [x] I18n 
- [x] Sms service

---

### Run locally

```sh
cp .env.exemple .env
docker-compose up --build
```

See Swagger Doc ```http://localhost:3000/swagger/index.html```


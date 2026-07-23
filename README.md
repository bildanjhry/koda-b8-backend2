# Program Autentikasi

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin_gonic-008ECF?style=for-the-badge&logo=gin&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-000000?style=for-the-badge&logo=jsonwebtokens&logoColor=white)
![Swagger](https://img.shields.io/badge/Swagger-85EA2D?style=for-the-badge&logo=swagger&logoColor=black)
![JQuery](https://img.shields.io/badge/jquery-0769AD?style=for-the-badge&logo=jquery&logoColor=white)

Aplikasi Fullstack dengan Backend menggunakan Gin-Gonic sebagai frameworknya, HTML & CSS dan Jquery untuk Frontend-nya serta Jason Web Token(JWT) untuk engine autentikasi-nya.

### Tech Stack:
- Go v1.25.4
- github.com/bildanjhry/go_shared-lib v1.0.1
- github.com/gin-gonic/gin v1.12.0
- github.com/jackc/pgx/v5 v5.10.0
- github.com/golang-jwt/jwt/v5 v5.3.1
- github.com/swaggo/files v1.0.1
- github.com/swaggo/gin-swagger v1.6.1
- github.com/swaggo/swag v1.16.6


### ERD:

```mermaid
erDiagram
    USERS {
        int64 id PK
        string name
        string email
        string password
        string picture
        date created_at
        date updated_at
    }

```

Untuk dokumentasi API, aplikasi ini menggunakan swaggo (swagger-go) dan berikut merupakan endpoint-endpoint dari program ini
### Preview:
![alt text](swagger.png)

<p align="center">

<h3 align="center">Bank Merchant API</h3>
</p>

### Built With

- [Go v1.22.3](https://go.dev/)
- [GORM v1.25.12](https://gorm.io/)
- [Gin v1.10.0](https://gin-gonic.com/)
- [PostgreSQL](https://www.postgresql.org/docs/)
- [Google UUID](https://github.com/google/uuid)
- [Golang JWT](https://golang-jwt.github.io/jwt/)

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.

- Text Editor
- PgAdmin or other PostgreSQL DBMS
- Postman or other API Testing Apps<br/>
- Golang Packages v1.22.3

### Installation

1. Clone the repo
   ```sh
    https://github.com/maulanadityaa/bank-merchant-api.git
   ```
2. Rename .env.example to .env and change few config

   ```env
    PORT=8080
    DB_HOST=YOUR DB HOST
    DB_PORT=YOUR DB PORT
    DB_USER=YOUR DB USER
    DB_PASSWORD=YOUR DB PASSWORD
    DB_NAME=YOUR DB NAME
    DB_SSLMODE=disable
    JWT_SECRET=YOUR JWT SECRET
    APP_ENV=development
    LOG_LEVEL=debug
    APP_NAME=Bank Merchant API
   ```

3. Install all libraries

   ```sh
   go mod tidy
   ```

4. Then run the project

   ```sh
   go run main.go
   ```

## API Documentation

Postman : https://documenter.getpostman.com/view/32332849/2sAXxLCa2E

### Example Request

- Endpoint : `/api/v1/payment/pay`
- Method : POST
- Header : Authorization : Bearer {token}
- Content-Type: application/json
- Accept: application/json
- Body :

```json
{
  "to": "97dda43f-6678-4135-91b0-e30f3aa2205b",
  "amount": 150000
}
```

### Example Response

```json
{
  "statusCode": 200,
  "message": "OK",
  "data": {
    "from": {
      "id": "c2df7ced-d05b-489c-9163-205981134f92",
      "name": "Lynn Schmeler",
      "balance": 20000,
      "created_at": "2024-10-03 21:28:12.772083 +0700 WIB",
      "updated_at": "2024-10-03 23:14:41.892194 +0700 WIB"
    },
    "to": {
      "id": "97dda43f-6678-4135-91b0-e30f3aa2205b",
      "name": "Emmett Little",
      "balance": 3480000,
      "created_at": "2024-10-03 21:28:34.010495 +0700 WIB",
      "updated_at": "2024-10-03 23:14:41.892593 +0700 WIB"
    },
    "amount": 10000,
    "created_at": "2024-10-03 23:14:41.8938546 +0700 WIB m=+1332.260772601"
  }
}
```

<!-- CONTACT -->

## Contact

M Maulana Z Aditya -
Instagram - [@maulanadityaa](https://instagram.com/maulanadityaa)

Project Link: [https://github.com/maulanadityaa/bank-merchant-api](https://github.com/maulanadityaa/bank-merchant-api)

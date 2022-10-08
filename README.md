# Assignment 2 Sesi 8

Assignment ini bertujuan untuk mempelajari use case untuk package Gin, Gorm, dan gin-swagger. Dengan menggunakan package tersebut, kita diharapkan dapat membuat API yang lebih baik dan lebih mudah.

## How to run

1. Clone this repository,
2. Run this command in your terminal,
   ```bash
   > go install && swag i -g ../../cmd/assignment-2/main.go -d ./api/handlers/
   > docker compose up --build
   ```
3. Wait until the building process is done, 
4. open your browser and go to `localhost:8000`

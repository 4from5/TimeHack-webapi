migrate -path ../schema/ -database 'postgres://localhost:5432/timehack_db?sslmode=disable' down
migrate -path ../schema/ -database 'postgres://localhost:5432/timehack_db?sslmode=disable' up
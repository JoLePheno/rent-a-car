psql -c "CREATE DATABASE stoik"

go run cmd/rentalctl/*.go init

go run cmd/rentalctl/*.go version

go run cmd/rentalctl/*.go up
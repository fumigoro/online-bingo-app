cd frontend/
npm run generate
echo "Nuxt.js build completed."
cd ../backend/
go build
go run .
cd ../


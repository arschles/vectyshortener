.PHONY: frontend
frontend:
	cd ./frontend && GOOS=js GOARCH=wasm go build -o ../public/frontend.wasm

.PHONY: backend
backend:
	go run .

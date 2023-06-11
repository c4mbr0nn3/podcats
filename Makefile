.PHONY: run, golang, vuejs

golang:
	bash -c "cd backend && go run main.go"

vuejs:
	bash -c "cd frontend && npm start"

dev:
	make golang & make vuejs


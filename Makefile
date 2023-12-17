test:
	go list -f '{{.Dir}}/...' -m | xargs go test

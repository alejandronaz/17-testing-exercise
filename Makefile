.PHONY: test
test:
	go test ./...

.PHONY: report
report:
	go test -coverprofile=coverage.out ./...

.PHONY: html_report
html_report: report
	go tool cover -html=coverage.out

.PHONY: func_report
func_report: report
	go tool cover -func=coverage.out

.PHONY: linter
linter:
	staticcheck ./...
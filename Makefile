# Makefile para tests y coverage

.PHONY: test coverage coverage-html clean ci

# Corre todos los tests
test:
	go test ./... -v

# Genera reporte de coverage en consola
coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out	

# Genera reporte de coverage en HTML (se abre en navegador)
coverage-html: coverage
	go tool cover -html=coverage.out -o coverage.html
	@echo "Reporte HTML generado en coverage.html"

# Limpia archivos temporales
clean:
	rm -f coverage.out coverage.html

# Regla para CI/CD - falla si coverage < 80%
ci:
	go test ./... -coverprofile=coverage.out
	@total=$$(go tool cover -func=coverage.out | grep total | awk '{print substr($$3, 1, length($$3)-1)}'); \
	req=80.0; \
	echo "Coverage: $$total% (mínimo requerido: $$req%)"; \
	if awk -v t=$$total -v r=$$req 'BEGIN { exit (t < r) }'; then \
		echo "✅ Coverage suficiente: $$total%"; \
	else \
		echo "❌ Coverage insuficiente: $$total% (mínimo $$req%)"; \
		go tool cover -html=coverage.out -o coverage.html; \
		echo "👉 Revisa el reporte en coverage.html"; \
		exit 1; \
	fi

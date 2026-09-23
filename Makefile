default:
	@echo 'Targets:'
	@echo '  run     Build and run main app.'
	@echo '  test    Run unit tests.'

run:
	go run .

test:
	go test -v practicalastro/tests

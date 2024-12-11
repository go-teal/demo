run:
	rm -f store/db_file.duckdb
	teal gen
	go run ./cmd/demo
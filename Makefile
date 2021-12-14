.PHONY = proto
SERVICES = activity auth

proto:
	@for srv in $(SERVICES); do \
		if [ -d "./"$$srv ]; then \
			echo "generating proto for "$$srv; \
			protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./$$srv/*.proto; \
		fi; \
	done
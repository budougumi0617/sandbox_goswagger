.PHONY: swagger

swagger:
	docker run --rm -it -v ${PWD}:/local -w /local quay.io/goswagger/swagger \
	generate server -f ./swagger.yaml --exclude-main -t ./gen





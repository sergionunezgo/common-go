format:
	./script/format.sh

lint:
	./script/lint.sh

test:
	./script/test.sh

protoc:
	./script/protoc.sh

precommit: format lint test
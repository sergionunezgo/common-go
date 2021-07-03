format:
	./script/format.sh

lint:
	./script/lint.sh

test:
	./script/test.sh

precommit: format lint test
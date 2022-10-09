bin=wechat_work

.PHONY: clean help ALL

help:
	@echo "usage: make <option>"
	@echo "options and effects:"
	@echo "    help   : Show help"
	@echo "    all    : Build multiple binary of this project"
	@echo "    darwin : Build the darwin binary of this project"

ALL: ${bin}

wechat_work:
	go build -o wechat_work

darwin:
	go build -o ${bin}-darwin
	@echo "build darwin done"

clean:
	rm -rf wechat_work

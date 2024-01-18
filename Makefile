.PHONY: clean fstabfmt install

PREFIX ?= /usr
MANDIR ?= $(PREFIX)/share/man/man1
DESTDIR ?=
SRCFILES := $(wildcard *.go)
GOBUILD ?= go build -mod=vendor -v

fstabfmt: $(SRCFILES)
	$(GOBUILD)

install: fstabfmt
	mkdir -p "$(DESTDIR)$(PREFIX)/bin"
	install -m755 fstabfmt "$(DESTDIR)$(PREFIX)/bin/fstabfmt"

clean:
	rm -f fstabfmt

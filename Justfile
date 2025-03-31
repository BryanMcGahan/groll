
default: build

build:
    go build -o ./bin/groll ./cmd/groll

run ARGS='': build
    @echo {{ARGS}}
    ./bin/groll {{ARGS}}

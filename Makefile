
all: build create_user read_user update_user delete_user

build:
	mkdir -p build

create_user: build
	go build -o build/create_user cmd/myapp/create_user.go

read_user: build
	go build -o build/read_user cmd/myapp/read_user.go

update_user: build
	go build -o build/update_user cmd/myapp/update_user.go

delete_user: build
	go build -o build/delete_user cmd/myapp/delete_user.go

clean:
	rm -rf build

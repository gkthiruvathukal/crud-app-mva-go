
all: create_user read_user update_user delete_user

create_user:
	go build -o create_user create_user.go

read_user:
	go build -o read_user read_user.go

update_user:
	go build -o update_user update_user.go

delete_user:
	go build -o delete_user delete_user.go

clean:
	rm -f create_user read_user update_user delete_user

export GOPROXY=https://goproxy.cn,direct
export GO111MODULE=on

OBJ = udpsend

all: $(OBJ)

$(OBJ):
	go mod tidy && go build -gcflags "-N -l" -o $@ .
	#cp -af $(OBJ) ./docker/

clean:
	rm -fr $(OBJ)

-include .deps
dep:
	echo -n "$(OBJ):" > .deps
	find src -name '*.go' | awk '{print $$0 " \\"}' >> .deps
	echo "" >> .deps
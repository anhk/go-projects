export GOPROXY=https://goproxy.cn,direct
export GO111MODULE=on

OBJ = udpsend

all: $(OBJ)

$(OBJ):
	go mod tidy && go build -gcflags "-N -l" -o $@ .

clean:
	rm -fr $(OBJ)

-include .deps
dep:
	echo '$(OBJ): \\' > .deps
	find . -name '*.go' | awk '{print $$0 " \\"}' >> .deps
	echo "" >> .deps
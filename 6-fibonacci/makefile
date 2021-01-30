all: build

build:
	go build -o fibonacci.out .

run: build
	./fibonacci.out hello

clean:
	rm *.out -rf

rebuild: clean build
	

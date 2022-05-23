build-c:
	gcc -c -Wall -S -o ./internal/${BIN_NAME}.asm  ./internal/elo.c
	gcc -c -Wall -o ./internal/${BIN_NAME}.out ./internal/elo.c
	ar rsc ./internal/libelo.a ./internal/${BIN_NAME}.out
benchmark:
	go test ./pkg/elo_test.go -bench=. -v
# commands: 
# https://borretti.me/article/compiling-llvm-ir-binary
# https://stackoverflow.com/questions/9148890/how-to-make-clang-compile-to-llvm-ir

# don't delete intermediate files
# source: https://stackoverflow.com/questions/9638275/how-to-undo-intermediate-file-deletion
.PRECIOUS: %.ll %.obj

all: fib.bin

# gen.ll: ../ast/llvm_test.go
# 	cd ../ast && go test -timeout 30s -run ^TestAstTestSuite$$ -testify.m "^(TestBasic)$$" github.com/klvnptr/k/ast

%.ll: %.c
	clang -S -emit-llvm -o $@ $<

%.obj: %.ll
	llc -filetype=obj -o $@ $<

%.bin: %.obj
	clang -g -v -o $@ $<

%.s: %.ll
	llc -filetype=asm -o $@ $<

clean:
	rm -f *.ll *.obj *.bin *.s
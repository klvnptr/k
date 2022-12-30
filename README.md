# K-lang

## Intro

K (K-lang) is a strongly-typed, C-like programming language that I am writing to learn how to write a compiler. It is written in Go and uses LLVM as a backend.

All it can do right now is to run bunch of tests. See ```ast/*_test.go``` files.

If you want to run the tests, run.
```bash
go test ./...
```

One of the goals it to take and adopt as many test cases possible from the standard [c-testsuite](https://github.com/c-testsuite/c-testsuite) and have K pass them. See the runner in ```ast/k_testsuite_test.go``` and the adapted test cases in ```testsuite``` folder.

K (or K-lang) is named after the inital of [Konstruktor](https://konstruktor.online/), a fantastic product engineering team in Budapest, Hungary that I am a part of.

## Examples

K can link to standard C libraries and call functions from them. Here is an example of a program that calls the ```printf``` function from the standard C library.

```c
i64 printf(i8 *fmt, ...);

i64 fib(i64 n) {
	if (n < 2) {
		return n;
	}

	return fib(n - 1) + fib(n - 2);
}

i64 main() {
	i64 i = 0;
	
	while (i++ < 6)
		printf("fib(%d) = %d\n", i, fib(i));
	
	return 0;
}
```
And it's output:
```
fib(1) = 1
fib(2) = 1
fib(3) = 2
fib(4) = 3
fib(5) = 5
fib(6) = 8
```
K already has a fine pointer arithmetic implementation. Here is an example of a program that uses it to reverse a string.
```c
i64 printf(i8 *fmt, ...);

i8* strrev(i8 *s) {
    i8 *p = s;
    i8 *q = s;
    i8 tmp;

    while (*q != (i8)0) {
        q++;
    }

    q--;

    while (p < q) {
        tmp = *p;
        *p = *q;
        *q = tmp;
        p++;
        q--;
    }

    return s;
}

i64 main() {
    i8 *s = "hello";
    printf("%s", strrev(s));
    
    return 0;
}
```
And it's output:
```
olleh
```

## Architecture

1. [alecthomas/participle](https://github.com/alecthomas/participle) is used in the ```parser``` package to parse the source code into an AST.
2. Transform() called on the AST and returns a new AST that is easier to work with in the next (code generation step) step.
3. In the ```ast``` package [llir/llvm](https://github.com/llir/llvm) is used for code generation. It is a Go package that generates easy to read, plain text LLVM IR.
4. We call ```llc``` to compile the LLVM IR into machine code.

## Todo

Many todos. Here is a list of some of them:

- [ ] implement more operators from C
- [ ] convert operators to head + tail
- [ ] implement arrays
  - [X] partially implemented, needs more testing
- [ ] implement globals
- [ ] sizeof should accept expressions
- [ ] struct assign {.a = 1, .b = 2}
- [ ] continue, break
- [ ] for loop
- [ ] +=, -=, *=, /=

## Useful Resources

### Code Generation
- [LLVM IR and Go](https://blog.gopheracademy.com/advent-2018/llvm-ir-and-go/) This is a good introduction to LLVM IR and how to use the llir/llvm package to generate LLVM IR in Go. This inspired me to start writing this compiler.

### C syntax
- [The syntax of C in Backus-Naur Form](https://cs.wmich.edu/~gupta/teaching/cs4850/sumII06/The%20syntax%20of%20C%20in%20Backus-Naur%20form.htm) This helped me a lot to understand the syntax of C.
- [A Small-C language definition for teaching compiler design](https://medium.com/@efutch/a-small-c-language-definition-for-teaching-compiler-design-b70198531a2f) This is a good example of how to use EBNF to define a basic C-like language.
- [Participle's MicroC](https://github.com/alecthomas/participle/tree/master/_examples/microc) This is a good example of how to use participle to parse C-like syntax.

### LLVM IR
- [Mapping High Level Constructs to LLVM IR](https://mapping-high-level-constructs-to-llvm-ir.readthedocs.io/en/latest/README.html) This is a good introduction to LLVM IR and how to map C constructs to LLVM IR.
- [LLVM IR reference](https://llvm.org/docs/LangRef.html) This is the official LLVM IR reference.

### EBNF
- [General Extended BNF syntax](https://www.cs.nmsu.edu/~rth/cs/cs471/Syntax%20Module/EBNF.html) This is a good introduction to EBNF. Can be useful to understand participle's EBNF-like syntax in the Go structs.
# AUSPL : Another Useless Programming Language

AUSPL is an interpreted language. It is an implementation of ["Monkey"](https://monkeylang.org/) programming language from the book ["Writing an Interpreter in Go"](https://interpreterbook.com/).

The syntax of AUSPL is same as that of Monkey.

Here's how AUSPL looks like:

```
//Integer Variable
let var = 5;

//String
let name = "This is AUSPL!";

//Boolean
let isTrue = true;

//Arrays
let numbers = [1, 2, 3, 4, 5];

//Hash Map
let maps = {"Cat": 0, "Dog": 1};
```

AUSPL supports various functionality such as:
 - Functions
   ```
   let add = fn(a, b) {
     return a + b; 
   };
   ```
 - Conditional Expressions
   ```
   let useless = true;
   if (useless) {
      puts("Yup, Its useless!");
   } else {
      puts("Not useless!");
   }
   ```
 - While Loop:
   ```
   let i = 0;
   while (i<5) {
     puts(i);
     i = i + 1;
   }
   ```

// #include <stdio.h>
i8 printf(i8* f, ...);

i64 main() 
{
   i64 a;
   a = 42;
   printf("%d\n", a);

   i64 b = 64;
   printf("%d\n", b);

   // int c = 12, d = 34;
   i64 c = 12;
   i64 d = 34;
   printf("%d, %d\n", c, d);

   return 0;
}

// vim: set expandtab ts=4 sw=3 sts=3 tw=80 :

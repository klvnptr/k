// #include <stdio.h>
// #include <string.h>
i8 printf(i8* f, ...);
i8* strcpy(i8* d, i8* s);
i8* strncpy(i8* d, i8* s, i64 n);
i64 strcmp(i8* s1, i8* s2);
i64 strlen(i8* s);
i8* strcat(i8* d, i8* s);
i64 strncmp(i8* s1, i8* s2, i64 n);
i8* strchr(i8* s, i8 c);
i8* strrchr(i8* s, i8 c);
i8* memset(i8* d, i8 c, i64 n);
i8* memcpy(i8* d, i8* s, i64 n);
i64 memcmp(i8* s1, i8* s2, i64 n);

i64 main()
{
   // i8 a[10];
   [10]i8 aa;
   i8* a = &aa[0];

   strcpy(a, "hello");
   printf("%s\n", a);

   strncpy(a, "gosh", 2);
   printf("%s\n", a);

   printf("%d\n", strcmp(a, "apple") > 0);
   printf("%d\n", strcmp(a, "goere") > 0);
   printf("%d\n", strcmp(a, "zebra") < 0);

   printf("%d\n", strlen(a));

   strcat(a, "!");
   printf("%s\n", a);

   printf("%d\n", strncmp(a, "apple", 2) > 0);
   printf("%d\n", strncmp(a, "goere", 2) == 0);
   printf("%d\n", strncmp(a, "goerg", 2) == 0);
   printf("%d\n", strncmp(a, "zebra", 2) < 0);

   printf("%s\n", strchr(a, 'o'));
   printf("%s\n", strrchr(a, 'l'));
   // printf("%d\n", strrchr(a, 'x') == NULL);
   printf("%d\n", strrchr(a, 'x') == 0);

   memset(&a[1], 'r', 4);
   printf("%s\n", a);

   memcpy(&a[2], a, 2);
   printf("%s\n", a);

   printf("%d\n", memcmp(a, "apple", 4) > 0);
   printf("%d\n", memcmp(a, "grgr", 4) == 0);
   printf("%d\n", memcmp(a, "zebra", 4) < 0);

   return 0;
}

// /* vim: set expandtab ts=4 sw=3 sts=3 tw=80 :*/

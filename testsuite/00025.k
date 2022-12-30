i64 strlen(i8* s);

i64
main()
{
	i8 *p;
	
	p = "hello";
	return strlen(p) - 5;
}

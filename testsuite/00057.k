i64
main()
{
	// char a[16], b[16];
	[16]i8 a;
	[16]i8 b;
	
	if(sizeof(a) != sizeof(b))
		return 1;
	return 0;
}

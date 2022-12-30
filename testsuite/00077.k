i64
foo([100]i64 x)
{
	[100]i64 y;
	i64 *p;
	
	y[0] = 2000;
	
	if(x[0] != 1000)
	{
		return 1;
	}
	
	p = (i64*)x;
	
	if(p[0] != 1000)
	{
		return 2;
	}
	
	p = (i64*)y;
	
	if(p[0] != 2000)
	{
		return 3;
	}
	
	// sizeof an array in K is the number of bytes allocated for it
	// if(sizeof(x) != sizeof(void*))
	// {
	// 	return 4;
	// }

	// sizeof an array in K is the number of bytes allocated for it	
	// if(sizeof(y) <= sizeof(x))
	// {
	// 	return 5;
	// }
	
	return 0;
}

i64
main()
{
	[100]i64 x;
	x[0] = 1000;
	
	return foo(x);
}

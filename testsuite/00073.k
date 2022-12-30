i64
main()
{
	[2]i64 arr;
	i64 *p;
	
	p = &arr[1];
	p = p - 1;
	*p = 123;
	
	if(arr[0] != 123)
		return 1;
	return 0;
}

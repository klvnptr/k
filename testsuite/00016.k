i64
main()
{
	[2]i64 arr;
	i64 *p;
	
	p = &arr[1];
	*p = 0;
	return arr[1];
}

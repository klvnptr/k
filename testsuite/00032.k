i64
main()
{
	[2]i64 arr;
	i64 *p;
	
	arr[0] = 2;
	arr[1] = 3;
	p = &arr[0];
	if(*(p++) != 2)
		return 1;
	if(*(p++) != 3)
		return 2;
	
	p = &arr[1];
	if(*(p--) != 3)
		return 1;
	if(*(p--) != 2)
		return 2;
		
	p = &arr[0];
	if(*(++p) != 3)
		return 1;
	
	p = &arr[1];
	if(*(--p) != 2)
		return 1;

	return 0;
}

i64
main()
{
	i8* p;
	i64 x;
	
	x = 2;
	p = (i8*)&x;
	
	if(*((i64*)p) != 2)
		return 1;
	return 0;
}

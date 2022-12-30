i64
main()
{
	i64 x;
	i64 *p;
	
	x = 4;
	p = &x;
	*p = 0;

	return *p;
}

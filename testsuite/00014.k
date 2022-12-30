i64
main()
{
	i64 x;
	i64 *p;
	
	x = 1;
	p = &x;
	p[0] = 0;
	return x;
}

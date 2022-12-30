i64
main()
{
	// int x, *p, **pp;
	i64 x;
	i64 *p;
	i64 **pp;
	
	x = 0;
	p = &x;
	pp = &p;
	return **pp;
}

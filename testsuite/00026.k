i8
main()
{
	i8 *p;
	
	p = "hello";
	// cannot implicitly cast i64 to i8
	return p[0] - (i8)104;
}

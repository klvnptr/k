i64
main()
{
	i64 x;
	i64 *p;
	i64 **pp;

	x = 0;
	p = &x;
	pp = &p;

	// if condition expressions require boolean values
	if(*p != 0)
		return 1;
	if(**pp != 0)
		return 1;
	else
		**pp = 1;

	if(x != 0)
		return 0;
	else
		return 1;

	// return was missing
	return 1;
}

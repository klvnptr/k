i64
main()
{
	i16 x;
	i16 *p;

	if (sizeof(0) < 2)
		return 1;
	if (sizeof 0 < 2)
		return 1;
	if (sizeof(i8) < 1)
		return 1;
	if (sizeof(i16) - 2 < 0)
		return 1;
	if (sizeof(&x) != sizeof p)
		return 1;
	return 0;
}

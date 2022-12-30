i64
zero()
{
	return 0;
}

i64
one()
{
	return 1;
}

i64
main()
{
	i64 x;
	i64 y;
	
	x = zero();
	y = ++x;
	if (x != 1)
		return 1;
	if (y != 1)
		return 1;
	
	x = one();	
	y = --x;
	if (x != 0)
		return 1;
	if (y != 0)
		return 1;
	
	x = zero();
	y = x++;
	if (x != 1)
		return 1;
	if (y != 0)
		return 1;
	
	x = one();
	y = x--;
	if (x != 0)
		return 1;
	if (y != 1)
		return 1;
	
	return 0;
}

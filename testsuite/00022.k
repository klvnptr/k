type i64 x;

i64
main()
{
	x v;
	// K is strongly-typed, so implicit casting of i64 to x is not allowed
	v = (x)0;
	return v;
}


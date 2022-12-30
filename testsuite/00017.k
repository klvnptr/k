i64
main()
{
	struct { i64 x, i64 y, } s;
	
	s.x = 3;
	s.y = 5;
	return s.y - s.x - 2; 
}

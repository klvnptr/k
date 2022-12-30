// types must be defined globally
type struct { i64 x, i64 y, } S;

i64
main()
{
	// struct S { int x; int y; } s;
	S s;
	// struct S *p;
	S *p;

	p = &s;	
	s.x = 1;
	p->y = 2;
	return p->y + p->x - 3; 
}


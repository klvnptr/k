type struct { S *p, i64 x, } S;

i64
main()
{
	// struct S { struct S *p; int x; } s;
	S s;
	
	s.x = 0;
	s.p = &s;
	return s.p->p->p->p->p->x;
}


type struct {
    i64 x,
    struct {
        i64 y,
        i64 z,
    } nest,
} s;

i64
main() {
    s v;
    v.x = 1;
    v.nest.y = 2;
    v.nest.z = 3;
    if (v.x + v.nest.y + v.nest.z != 6)
        return 1;
    return 0;
}


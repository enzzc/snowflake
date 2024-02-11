package snowflake

import "testing"

func TestSnowflakeSucc(t *testing.T) {
	g := NewSnowflakeGen()
	id1 := g.NewID()
	id2 := g.NewID()
	if id1 >= id2 {
		t.Errorf("id1=%d should not be greater or equal to id2=%d", id1, id2)
	}
}

func BenchmarkSnowflakeGen(b *testing.B) {
	g := NewSnowflakeGen()
	for n := 0; n < b.N; n++ {
		g.NewID()
	}
}

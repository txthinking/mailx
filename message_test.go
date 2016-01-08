package xmail

import(
    "testing"
)

func TestMakeBoundary(t *testing.T){
    s := MakeBoundary()
    t.Log(s)
}

func TestChunkSplit(t *testing.T){
    s := `
aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb
ccccccccccccccccccccccccccccccccccccccc
`
    s, err := ChunkSplit(s)
    if err != nil{
        t.Fatal(err)
    }
    t.Log(s)
}


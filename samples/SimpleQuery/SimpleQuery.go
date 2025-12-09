package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "fakegosqldriver"
)

/*
cd fakegosqldriver/samples/SimpleQuery
go run SimpleQuery.go

fakeDriver.Open sConParams=good
Executing first query
fakeConnection.QueryContext sRequestText=good
Completed first query
Executing second query
fakeConnection.QueryContext sRequestText=bad
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [sync.RWMutex.Lock]:
sync.runtime_SemacquireRWMutex(0x7ff64c966ec0?, 0xe0?, 0x7ff64c7daeb9?)
        runtime/sema.go:105 +0x25
sync.(*RWMutex).Lock(0x7ff64c966ec0?)
        sync/rwmutex.go:155 +0x6b
database/sql.(*Conn).close(0x7ff64c879cf8?, {0x7ff64c879cf8, 0xc00009e080})
        database/sql/sql.go:2138 +0x75
database/sql.(*Conn).closemuRUnlockCondReleaseConn(0xc0000c4000, {0x7ff64c879cf8, 0xc00009e080})
        database/sql/sql.go:2123 +0x78
database/sql.(*DB).queryDC(0xc0000b0080?, {0x7ff64c87a010, 0x7ff64c966ec0}, {0x0, 0x0}, 0xc0000c2000, 0xc00009e050, {0x7ff64c84b790, 0x3}, {0x0, ...})
        database/sql/sql.go:1790 +0x596
database/sql.(*Conn).QueryContext(0xc0000c4000, {0x7ff64c87a010, 0x7ff64c966ec0}, {0x7ff64c84b790, 0x3}, {0x0, 0x0, 0x0})
        database/sql/sql.go:2037 +0xc5
main.main()
        fakegosqldriver/samples/SimpleQuery/SimpleQuery.go:78 +0x305

goroutine 20 [select]:
database/sql.(*DB).connectionOpener(0xc000092270, {0x7ff64c87a080, 0xc0000b6000})
        database/sql/sql.go:1261 +0x87
created by database/sql.OpenDB in goroutine 1
        database/sql/sql.go:841 +0x130
exit status 2
*/

func main () {

	pool, err := sql.Open ("fake", "good")
	if err != nil {
		fmt.Printf ("sql.Open failed: %v\n", err)
		return
	}

	defer pool.Close ()

	con, err := pool.Conn (context.Background ())
	if err != nil {
		fmt.Printf ("pool.Conn failed: %v\n", err)
		return
	}

	defer con.Close ()

	fmt.Printf ("Executing first query\n")

	rows1, err := con.QueryContext (context.Background (), "good")
	if err != nil {
		fmt.Printf ("con.QueryContext failed: %v\n", err)
		return
	}

	fmt.Printf ("Completed first query\n")

	defer rows1.Close ()

	fmt.Printf ("Executing second query\n")

	rows2, err := con.QueryContext (context.Background (), "bad")
	if err != nil {
		fmt.Printf ("con.QueryContext failed: %v\n", err)
		return
	}

	fmt.Printf ("Completed second query\n")

	defer rows2.Close ()

} // end main

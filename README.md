## Fake Go SQL Driver and Sample Program

This fake Go SQL driver and sample program demonstrates Go `sql` package connection pool deadlock when `QueryContext` returns `driver.ErrBadConn`.

A `database/sql.Conn` connection object has a `closemu` field of type `sync.RWMutex`.
* Before a `Conn` object is closed, a write lock is acquired on its `closemu`.
* Each outstanding child `Rows` object holds a read lock on its parent's `closemu` to prevent its parent connection from being closed until all child `Rows` objects are closed.

When a driver returns `driver.ErrBadConn` from the `QueryContext` method, as part of the post-processing of the result from the driver, before returning to the application, the Go `sql` package connection pool tries to close the parent connection and remove the connection from its connection pool.

Before removing the connection from the pool, a write lock must be obtained on the connection's `closemu`. If there are any outstanding child `Rows` objects, the attempt to acquire a write lock on `closemu` will deadlock. The deadlock situation will cause a hang if undetected by the Go runtime, or cause a panic if detected by the Go runtime.

### How to run
```
cd fakegosqldriver/samples/SimpleQuery
go run SimpleQuery.go
```

### Example output
```
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
```

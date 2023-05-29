# scylla

### 封裝項目
- github.com/syhlion/gocql
- pika.rdtech.vip/eden-lib/gocqlx/v2
- pika.rdtech.vip/eden-lib/gocqlx/v2/qb


### use Unsafe

```golang
scylla.Unsafe()
```

### New Session

```golang
   cqlSession, err := scylla.NewSession([]string{"127.0.0.1"},
        scylla.WithPort(9042),
        scylla.WithConsistency(gocql.Quorum),
        scylla.WithTimeout(time.Duration(60)*time.Second),
        scylla.WithConnectTimeout(time.Duration(60)*time.Second),
        scylla.WithMaxPreparedStmts(1),
        scylla.WithNumConns(5),
    )

    defer cqlSession.Close()
```

### New Table
```golang
type Test struct {
        ID                int64             `json:"id"`
        Timezone          string            `json:"timezone"`
        AvailableCurrency []string          `json:"available_currency"`
        CommonSetting     map[string]string `json:"common_setting"`
        PlatformDomain    map[int64]string  `json:"platform_domain"`
    }

    var testMetadata = scylla.Metadata{
        Name:    "testbatch.test",
        Columns: []string{"id", "available_currency", "common_setting", "platform_domain", "timezone"},
        PartKey: []string{"id"},
    }

    var personTable = scylla.NewTable(testMetadata)

```

### Insert Table
```golang
    p := Test{
        99,
        "-10",
        []string{"test"},
        map[string]string{"aa": "bb"},
        map[int64]string{111: "bb"},
    }
    q1 := cqlSession.InsertBuilderQuery(personTable.InsertBuilder())
    exec1 := q1.BindStruct(p)

    if err := exec1.ExecRelease(); err != nil {
        log.Fatal(err)
    }

    stmt1, names1 := personTable.InsertBuilder().ToCql()

    fmt.Println(`[insert SQL] query:`, stmt1, `,names:`, names1)
    // insert SQL] query: INSERT INTO testbatch.test (id,available_currency,common_setting,platform_domain,timezone) VALUES (?,?,?,?,?)  ,names: [id available_currency common_setting platform_domain timezone]

```

### Select Table
```golang
    var testdata []Test

    query1 := cqlSession.SelectBuilderQuery(personTable.SelectBuilder(`id`, `timezone`))
    exec2 := query1.BindMap(map[string]interface{}{
        "id": 1,
    })

    if err := exec2.SelectRelease(&testdata); err != nil {
        log.Fatal(err)
    }

    stmt2, names2 := personTable.SelectBuilder(`id`, `timezone`).ToCql()

    fmt.Println(`[select SQL] query:`, stmt2, `,names:`, names2, `,testdata:`, testdata)
    // [select SQL] query: SELECT id,timezone FROM testbatch.test WHERE id=?  ,names: [id] ,testdata: [{1 1212 [] map[] map[]}]

```

### Update Table
```golang
    // fmt.Println(`PrimaryKeyCmp():`, personTable.PrimaryKeyCmp())
    query := cqlSession.UpdateBuilderQuery(personTable.UpdateBuilder(`timezone`))
    exec := query.BindMap(map[string]interface{}{
        `timezone`: `+9`,
        `id`:       10,
    })

    if err := exec.ExecRelease(); err != nil {
        log.Fatal(err)
    }

    stmt, names := personTable.UpdateBuilder(`timezone`).ToCql()

    fmt.Println(`[update] query:`, stmt, `,names:`, names)
    // [update] query: UPDATE testbatch.test SET timezone=? WHERE id=?  ,names: [timezone id]
```
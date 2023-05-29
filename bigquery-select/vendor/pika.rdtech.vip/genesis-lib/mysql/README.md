# mysql
-
### Open

```golang
    db, err := mysql.Open("root:1234@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local",
        mysql.WithSetMaxIdleConns(10),
        mysql.WithSetMaxOpenConns(100),
        mysql.WithSetConnMaxLifetime(30),
        mysql.WithSingularTable(true),
        mysql.WithLogMode(true),
    )

    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.fatal(err)
    }
```

### AutoMigrate

```golang
    type Accounts struct {
        mysql.Model
        Account string `gorm:"column:account;" json:"account,omitempty"`
        Price   int    `gorm:"column:price;" json:"price,omitempty"`
    }

    err := db.AutoMigrate(Accounts{}).Error()
    if err != nil {
        log.fatal(err)
    }
```

### Create

```golang
    if err := db.Create(&Accounts{Account: "hello"}).Error(); err != nil {
        return
    }
```

### Find

```golang
    accs := make([]*Accounts, 0)
    err = db.Where("id = ?", 2).Find(&accs).Error()
    if err !=nil {
        return
    }
```

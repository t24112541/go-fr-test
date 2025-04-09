# run mode (!windows)
``` bash
APP_ENV=dev go run main.go
```

structure: 
``` 
go-fr-test/
├── cmd/
│   └── migration/
│       ├── sub/
│       └── main.go
├── datasource/
├── db/                         # for docker init database
├── handlers/
├── migrations/
├── models/
├── routers/
├── configs/
│   ├── .local.env
│   ├── .dev.env
│   ├── .staging.env
│   ├── .prod.env
|   └── .env (default)
├── main.go
└── ...
```

# create user on db(mysql)
``` bash
CREATE USER 'go_fr_db_user'@'%' IDENTIFIED BY 'password';
GRANT ALL ON go_fr_test.* TO 'go_fr_db_user'@'%';
FLUSH PRIVILEGES;
```



# migrate command
``` bash
 go run cmd\migration\main.go [command]
```

- `init` initial migration table
- `up` migration
- `rollback` rollback migration
- `make-sql` new migration file(.sql) for sql language
    - `-name="[name]"` name of migration file
- `make-go` new migration file(.go) for go language
    - `-name="[name]"` name of migration file
- `status` check migration
- `fake` fake migration(not apply to db)
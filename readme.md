# run mode (!windows)
``` console
APP_ENV=dev go run main.go
```

note: 
``` 
go-fr-test/
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
```console
CREATE USER 'go_fr_db_user'@'%' IDENTIFIED BY 'password';
GRANT ALL ON go_fr_test.* TO 'go_fr_db_user'@'%';
FLUSH PRIVILEGES;
```


```
go install gofr.dev/cli/gofr@latest
```


# lapi-hotel-system
## run dev server
```
export AUTH0_DOMAIN=lapi.us.auth0.com
DEBUG=true go run .
```
## migration
reference：https://qiita.com/k-kurikuri/items/946e2bf8c79176ef3ff0

install sql-migrate
```
go install github.com/rubenv/sql-migrate/...@latest
```

see status
```
sql-migrate status
```

dev migrate
```
sql-migrate up -env="test"
sql-migrate up -env="development"
```

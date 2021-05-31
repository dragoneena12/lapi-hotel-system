# lapi-hotel-system
## install dependencies
```
go install
```
## migration
reference：https://qiita.com/k-kurikuri/items/946e2bf8c79176ef3ff0

install sql-migrate
```
go install github.com/rubenv/sql-migrate/...
```

see status
```
sql-migrate status
```

dev migrate
```
sql-migrate up -env="development"
```

prod migrate
```
kubectl port-forward -n lapi-hotel-system <pod-name> 3306:3306
sql-migrate up -env="production"
```
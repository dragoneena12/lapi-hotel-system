# lapi-hotel-system
## install dependencies
```
go install
```
## migration
reference：https://qiita.com/k-kurikuri/items/946e2bf8c79176ef3ff0

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
kubectl port-forward -n lapi-hotel-system mariadb-69bdf58589-zn6jf 3307:3306
sql-migrate up -env="production"
```
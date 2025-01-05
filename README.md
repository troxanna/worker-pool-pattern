Logger implemented using worker-pool pattern

# How to run
```
cd worker-pool-pattern
mkdir users
go run main.go
```

# Result tests
| Count    | Time without worker-pool    | Time with worker-pool    |
|----------|-----------------------------|--------------------------|
|    10    |     11.04 seconds           |        0.91 seconds      |
|    100   |     111.56 seconds          |        9.98 seconds      |
|    1000  |     1120.09 seconds         |        101.42 seconds    |

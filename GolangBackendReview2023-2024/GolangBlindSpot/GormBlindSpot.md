## Gorm Blind Spot

Record the bug.

### Where in syntax error

err code

```go
err := repo.fmsDb.Table("third_party_order_tab").Where("third_party_tracking_num IN ?", thirdPartyTrackingNums).Find(&result).GetError()
```

<span style=color:red>error</span>

```shell
Error 1064: You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near '?,?)' at line 1 
```

<span style=color:green>fix</span>  (add "<span style=color:red>**()**</span>" beside the "?")

```go
err := repo.fmsDb.Table("third_party_order_tab").Where("third_party_tracking_num IN (?)", thirdPartyTrackingNums).Find(&result).GetError()
```

#### How did gorm build SQL

build result

```go
err := repo.fmsDb.Table("third_party_order_tab").Where("third_party_tracking_num IN ?", thirdPartyTrackingNums).Find(&result).GetError()
```

convert into

```sql
SELECT * FROM `third_party_order_tab`  WHERE (third_party_tracking_num IN ?,?,?)
```

**core func**

 `scope.prepareQuerySQL()` in 

`github.com/jinzhu/gorm@v1.9.16/callback_query.go`

core code

```go
scope.Raw(fmt.Sprintf("SELECT %v FROM %v %v", scope.selectSQL(), scope.QuotedTableName(), scope.CombinedConditionSql()))
```


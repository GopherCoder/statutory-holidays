# statutory-holidays

> 基于 Redis 构建的节假日 API


### 1. 设计（redis 数据结构）

table One

| key | value | type | 说明|
| --- |---|---|---|
|ch_name_list| 中文节假日列表| List | 	"元旦","春节","清明节","劳动节","端午节","中秋节","国庆节"|


table two 

| key | value | type | 说明|
|---|---|---|---|
|en_name_list| 英文节假日列表| List | 	"New Year\\'s Day","Spring Festival","Tomb-sweeping Day","Labour Day","Dragon Boat Festival","Mid-autumn Festival","National Day"|


table three

| key | value | type | 说明 |
| ---| ---|---|---|
|history_holidays_map| 节假日集合| hash | key 2019:0  value "2018/12/30~2019/01/01"|

### 2. 如何操作？

根据设计，如何通过年份获取到当年的节假日信息？

- 根据设计，将数据导入到 redis 中
- 根据 history_holidays_map， 可以获取到 keys 或者 values
- 根据 keys (2019:0, 2019 表示年，0 表示当年的第一个法定节假日)
- 根据 values(2018/12/30~2019/01/01, 2018/12/30 表示起始时间，2019/01/01 表示结束时间)

即可获取到当年的节假日安排

### 3. 如何访问？

> http://localhost:8080/v1/api/holiday/holidays?year=2019

### 4. 其他？

暂无

CREATE DICTIONARY my_dictionary (
  id UInt64,
  name String
)
PRIMARY KEY id
SOURCE (
  CLICKHOUSE(
    HOST 'localhost'
    PORT tcpPort()
    TABLE 'my_source_table'
  )
)
LAYOUT(FLAT())
LIFETIME(
  MIN 0
  MAX 1000
)
COMMENT 'My dictionary';

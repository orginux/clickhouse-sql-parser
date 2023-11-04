package parser

const (
	KeywordAdd          = "ADD"
	KeywordAdmin        = "ADMIN"
	KeywordAfter        = "AFTER"
	KeywordAlias        = "ALIAS"
	KeywordAll          = "ALL"
	KeywordAlter        = "ALTER"
	KeywordAnd          = "AND"
	KeywordAnti         = "ANTI"
	KeywordAny          = "ANY"
	KeywordArray        = "ARRAY"
	KeywordAs           = "AS"
	KeywordAsc          = "ASC"
	KeywordAscending    = "ASCENDING"
	KeywordAsof         = "ASOF"
	KeywordAst          = "AST"
	KeywordAsync        = "ASYNC"
	KeywordAttach       = "ATTACH"
	KeywordBetween      = "BETWEEN"
	KeywordBoth         = "BOTH"
	KeywordBy           = "BY"
	KeywordCache        = "CACHE"
	KeywordCase         = "CASE"
	KeywordCast         = "CAST"
	KeywordCheck        = "CHECK"
	KeywordClear        = "CLEAR"
	KeywordClickhouse   = "CLICKHOUSE"
	KeywordCluster      = "CLUSTER"
	KeywordCodec        = "CODEC"
	KeywordCollate      = "COLLATE"
	KeywordColumn       = "COLUMN"
	KeywordColumns      = "COLUMNS"
	KeywordComment      = "COMMENT"
	KeywordCompiled     = "COMPILED"
	KeywordConfig       = "CONFIG"
	KeywordConstraint   = "CONSTRAINT"
	KeywordCreate       = "CREATE"
	KeywordCross        = "CROSS"
	KeywordCube         = "CUBE"
	KeywordCurrent      = "CURRENT"
	KeywordDatabase     = "DATABASE"
	KeywordDatabases    = "DATABASES"
	KeywordDate         = "DATE"
	KeywordDay          = "DAY"
	KeywordDeduplicate  = "DEDUPLICATE"
	KeywordDefault      = "DEFAULT"
	KeywordDelay        = "DELAY"
	KeywordDelete       = "DELETE"
	KeywordDesc         = "DESC"
	KeywordDescending   = "DESCENDING"
	KeywordDescribe     = "DESCRIBE"
	KeywordDetach       = "DETACH"
	KeywordDetached     = "DETACHED"
	KeywordDictionaries = "DICTIONARIES"
	KeywordDictionary   = "DICTIONARY"
	KeywordDisk         = "DISK"
	KeywordDistinct     = "DISTINCT"
	KeywordDistributed  = "DISTRIBUTED"
	KeywordDrop         = "DROP"
	KeywordDNS          = "DNS"
	KeywordElse         = "ELSE"
	KeywordEnd          = "END"
	KeywordEngine       = "ENGINE"
	KeywordEstimate     = "ESTIMATE"
	KeywordEvents       = "EVENTS"
	KeywordExcept       = "EXCEPT"
	KeywordExists       = "EXISTS"
	KeywordExplain      = "EXPLAIN"
	KeywordExpression   = "EXPRESSION"
	KeywordExtract      = "EXTRACT"
	KeywordFetches      = "FETCHES"
	KeywordFileSystem   = "FILESYSTEM"
	KeywordFinal        = "FINAL"
	KeywordFirst        = "FIRST"
	KeywordFile         = "FILE"
	KeywordFlush        = "FLUSH"
	KeywordFollowing    = "FOLLOWING"
	KeywordFor          = "FOR"
	KeywordFormat       = "FORMAT"
	KeywordFreeze       = "FREEZE"
	KeywordFrom         = "FROM"
	KeywordFull         = "FULL"
	KeywordFunction     = "FUNCTION"
	KeywordFunctions    = "FUNCTIONS"
	KeywordGlobal       = "GLOBAL"
	KeywordGrant        = "GRANT"
	KeywordGranularity  = "GRANULARITY"
	KeywordGroup        = "GROUP"
	KeywordHaving       = "HAVING"
	KeywordHierarchical = "HIERARCHICAL"
	KeywordHour         = "HOUR"
	KeywordHttp         = "HTTP"
	KeywordId           = "ID"
	KeywordIf           = "IF"
	KeywordIlike        = "ILIKE"
	KeywordIn           = "IN"
	KeywordIndex        = "INDEX"
	KeywordInf          = "INF"
	KeywordInjective    = "INJECTIVE"
	KeywordInner        = "INNER"
	KeywordInsert       = "INSERT"
	KeywordInterval     = "INTERVAL"
	KeywordInto         = "INTO"
	KeywordIs           = "IS"
	KeywordIs_object_id = "IS_OBJECT_ID"
	KeywordJoin         = "JOIN"
	KeywordKey          = "KEY"
	KeywordKill         = "KILL"
	KeywordLast         = "LAST"
	KeywordLayout       = "LAYOUT"
	KeywordLeading      = "LEADING"
	KeywordLeft         = "LEFT"
	KeywordLifetime     = "LIFETIME"
	KeywordLike         = "LIKE"
	KeywordLimit        = "LIMIT"
	KeywordLive         = "LIVE"
	KeywordLocal        = "LOCAL"
	KeywordLogs         = "LOGS"
	KeywordMark         = "MARK"
	KeywordMaterialize  = "MATERIALIZE"
	KeywordMaterialized = "MATERIALIZED"
	KeywordMax          = "MAX"
	KeywordMerges       = "MERGES"
	KeywordMin          = "MIN"
	KeywordMinute       = "MINUTE"
	KeywordModify       = "MODIFY"
	KeywordMonth        = "MONTH"
	KeywordMongodb      = "MONGODB"
	KeywordMove         = "MOVE"
	KeywordMoves        = "MOVES"
	KeywordMutation     = "MUTATION"
	KeywordMysql        = "MYSQL"
	KeywordNan_sql      = "NAN_SQL"
	KeywordNo           = "NO"
	KeywordNot          = "NOT"
	KeywordNull         = "NULL"
	KeywordNulls        = "NULLS"
	KeywordOdbc         = "ODBC"
	KeywordOffset       = "OFFSET"
	KeywordOn           = "ON"
	KeywordOptimize     = "OPTIMIZE"
	KeywordOption       = "OPTION"
	KeywordOr           = "OR"
	KeywordOrder        = "ORDER"
	KeywordOuter        = "OUTER"
	KeywordOutfile      = "OUTFILE"
	KeywordOver         = "OVER"
	KeywordPartition    = "PARTITION"
	KeywordPipeline     = "PIPELINE"
	KeywordPolicy       = "POLICY"
	KeywordPopulate     = "POPULATE"
	KeywordPostgresql   = "POSTGRESQL"
	KeywordPreceding    = "PRECEDING"
	KeywordPrewhere     = "PREWHERE"
	KeywordPrimary      = "PRIMARY"
	KeywordProjection   = "PROJECTION"
	KeywordQuarter      = "QUARTER"
	KeywordQuery        = "QUERY"
	KeywordQueues       = "QUEUES"
	KeywordQuota        = "QUOTA"
	KeywordRange        = "RANGE"
	KeywordRedis        = "REDIS"
	KeywordRefresh      = "REFRESH"
	KeywordReload       = "RELOAD"
	KeywordRemove       = "REMOVE"
	KeywordRename       = "RENAME"
	KeywordReplace      = "REPLACE"
	KeywordReplica      = "REPLICA"
	KeywordReplicated   = "REPLICATED"
	KeywordReplication  = "REPLICATION"
	KeywordRestart      = "RESTART"
	KeywordRight        = "RIGHT"
	KeywordRole         = "ROLE"
	KeywordRollup       = "ROLLUP"
	KeywordRow          = "ROW"
	KeywordRows         = "ROWS"
	KeywordSample       = "SAMPLE"
	KeywordSecond       = "SECOND"
	KeywordSelect       = "SELECT"
	KeywordSemi         = "SEMI"
	KeywordSends        = "SENDS"
	KeywordSet          = "SET"
	KeywordSettings     = "SETTINGS"
	KeywordShow         = "SHOW"
	KeywordShutdown     = "SHUTDOWN"
	KeywordSource       = "SOURCE"
	KeywordStart        = "START"
	KeywordStop         = "STOP"
	KeywordSubstring    = "SUBSTRING"
	KeywordSync         = "SYNC"
	KeywordSyntax       = "SYNTAX"
	KeywordSystem       = "SYSTEM"
	KeywordTable        = "TABLE"
	KeywordTables       = "TABLES"
	KeywordTemporary    = "TEMPORARY"
	KeywordTest         = "TEST"
	KeywordThen         = "THEN"
	KeywordTies         = "TIES"
	KeywordTimeout      = "TIMEOUT"
	KeywordTimestamp    = "TIMESTAMP"
	KeywordTo           = "TO"
	KeywordTop          = "TOP"
	KeywordTotals       = "TOTALS"
	KeywordTrailing     = "TRAILING"
	KeywordTrim         = "TRIM"
	KeywordTruncate     = "TRUNCATE"
	KeywordTtl          = "TTL"
	KeywordType         = "TYPE"
	KeywordUnbounded    = "UNBOUNDED"
	KeywordUncompressed = "UNCOMPRESSED"
	KeywordUnion        = "UNION"
	KeywordUpdate       = "UPDATE"
	KeywordUse          = "USE"
	KeywordUser         = "USER"
	KeywordUsing        = "USING"
	KeywordUuid         = "UUID"
	KeywordValues       = "VALUES"
	KeywordView         = "VIEW"
	KeywordVolume       = "VOLUME"
	KeywordWatch        = "WATCH"
	KeywordWeek         = "WEEK"
	KeywordWhen         = "WHEN"
	KeywordWhere        = "WHERE"
	KeywordWindow       = "WINDOW"
	KeywordWith         = "WITH"
	KeywordYear         = "YEAR"
)

var keywords = NewSet(
	KeywordAdd,
	KeywordAdmin,
	KeywordAfter,
	KeywordAlias,
	KeywordAll,
	KeywordAlter,
	KeywordAnd,
	KeywordAnti,
	KeywordAny,
	KeywordArray,
	KeywordAs,
	KeywordAsc,
	KeywordAscending,
	KeywordAsof,
	KeywordAst,
	KeywordAsync,
	KeywordAttach,
	KeywordBetween,
	KeywordBoth,
	KeywordBy,
	KeywordCache,
	KeywordCase,
	KeywordCast,
	KeywordCheck,
	KeywordClear,
	KeywordClickhouse,
	KeywordCluster,
	KeywordCodec,
	KeywordCollate,
	KeywordColumn,
	KeywordColumns,
	KeywordComment,
	KeywordCompiled,
	KeywordConfig,
	KeywordConstraint,
	KeywordCreate,
	KeywordCross,
	KeywordCube,
	KeywordCurrent,
	KeywordDatabase,
	KeywordDatabases,
	KeywordDate,
	KeywordDay,
	KeywordDeduplicate,
	KeywordDefault,
	KeywordDelay,
	KeywordDelete,
	KeywordDesc,
	KeywordDescending,
	KeywordDescribe,
	KeywordDetach,
	KeywordDetached,
	KeywordDictionaries,
	KeywordDictionary,
	KeywordDisk,
	KeywordDistinct,
	KeywordDistributed,
	KeywordDrop,
	KeywordDNS,
	KeywordElse,
	KeywordEnd,
	KeywordEngine,
	KeywordEstimate,
	KeywordEvents,
	KeywordExcept,
	KeywordExists,
	KeywordExplain,
	KeywordExpression,
	KeywordExtract,
	KeywordFetches,
	KeywordFileSystem,
	KeywordFinal,
	KeywordFirst,
	KeywordFile,
	KeywordFlush,
	KeywordFollowing,
	KeywordFor,
	KeywordFormat,
	KeywordFreeze,
	KeywordFrom,
	KeywordFull,
	KeywordFunction,
	KeywordFunctions,
	KeywordGlobal,
	KeywordGrant,
	KeywordGranularity,
	KeywordGroup,
	KeywordHaving,
	KeywordHierarchical,
	KeywordHour,
	KeywordHttp,
	KeywordId,
	KeywordIf,
	KeywordIlike,
	KeywordIn,
	KeywordIndex,
	KeywordInf,
	KeywordInjective,
	KeywordInner,
	KeywordInsert,
	KeywordInterval,
	KeywordInto,
	KeywordIs,
	KeywordIs_object_id,
	KeywordJoin,
	KeywordKey,
	KeywordKill,
	KeywordLast,
	KeywordLayout,
	KeywordLeading,
	KeywordLeft,
	KeywordLifetime,
	KeywordLike,
	KeywordLimit,
	KeywordLive,
	KeywordLocal,
	KeywordLogs,
	KeywordMark,
	KeywordMaterialize,
	KeywordMaterialized,
	KeywordMax,
	KeywordMerges,
	KeywordMin,
	KeywordMinute,
	KeywordModify,
	KeywordMonth,
	KeywordMongodb,
	KeywordMove,
	KeywordMoves,
	KeywordMutation,
	KeywordMysql,
	KeywordNan_sql,
	KeywordNo,
	KeywordNot,
	KeywordNull,
	KeywordNulls,
	KeywordOdbc,
	KeywordOffset,
	KeywordOn,
	KeywordOptimize,
	KeywordOption,
	KeywordOr,
	KeywordOrder,
	KeywordOuter,
	KeywordOutfile,
	KeywordOver,
	KeywordPartition,
	KeywordPipeline,
	KeywordPolicy,
	KeywordPopulate,
	KeywordPostgresql,
	KeywordPreceding,
	KeywordPrewhere,
	KeywordPrimary,
	KeywordProjection,
	KeywordQuarter,
	KeywordQuery,
	KeywordQueues,
	KeywordQuota,
	KeywordRange,
	KeywordRedis,
	KeywordRefresh,
	KeywordReload,
	KeywordRemove,
	KeywordRename,
	KeywordReplace,
	KeywordReplica,
	KeywordReplicated,
	KeywordReplication,
	KeywordRestart,
	KeywordRight,
	KeywordRole,
	KeywordRollup,
	KeywordRow,
	KeywordRows,
	KeywordSample,
	KeywordSecond,
	KeywordSelect,
	KeywordSemi,
	KeywordSends,
	KeywordSet,
	KeywordSettings,
	KeywordShow,
	KeywordShutdown,
	KeywordSource,
	KeywordStart,
	KeywordStop,
	KeywordSubstring,
	KeywordSync,
	KeywordSyntax,
	KeywordSystem,
	KeywordTable,
	KeywordTables,
	KeywordTemporary,
	KeywordTest,
	KeywordThen,
	KeywordTies,
	KeywordTimeout,
	KeywordTimestamp,
	KeywordTo,
	KeywordTop,
	KeywordTotals,
	KeywordTrailing,
	KeywordTrim,
	KeywordTruncate,
	KeywordTtl,
	KeywordType,
	KeywordUnbounded,
	KeywordUncompressed,
	KeywordUnion,
	KeywordUpdate,
	KeywordUse,
	KeywordUser,
	KeywordUsing,
	KeywordUuid,
	KeywordValues,
	KeywordView,
	KeywordVolume,
	KeywordWatch,
	KeywordWeek,
	KeywordWhen,
	KeywordWhere,
	KeywordWindow,
	KeywordWith,
	KeywordYear,
)

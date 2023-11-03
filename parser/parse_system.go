package parser

import (
	"fmt"
	"strings"
)

func (p *Parser) parseSetExpr(pos Pos) (*SetExpr, error) {
	if err := p.consumeKeyword(KeywordSet); err != nil {
		return nil, err
	}
	settings, err := p.parseSettingsExprList(p.Pos())
	if err != nil {
		return nil, err
	}
	return &SetExpr{
		SetPos:   pos,
		Settings: settings,
	}, nil
}

func (p *Parser) parseSystemFlushExpr(pos Pos) (*SystemFlushExpr, error) {
	if err := p.consumeKeyword(KeywordFlush); err != nil {
		return nil, err
	}

	switch {
	case p.matchKeyword(KeywordLogs):
		lastToken := p.last()
		_ = p.lexer.consumeToken()
		return &SystemFlushExpr{
			FlushPos:     pos,
			StatementEnd: lastToken.End,
			Logs:         true,
		}, nil
	case p.tryConsumeKeyword(KeywordDistributed) != nil:
		distributed, err := p.parseTableIdentifier(p.Pos())
		if err != nil {
			return nil, err
		}
		return &SystemFlushExpr{
			FlushPos:     pos,
			StatementEnd: distributed.End(),
			Distributed:  distributed,
		}, nil
	default:
		return nil, fmt.Errorf("expected LOGS|DISTRIBUTED")
	}
}

func (p *Parser) parseSystemReloadExpr(pos Pos) (*SystemReloadExpr, error) {
	if err := p.consumeKeyword(KeywordReload); err != nil {
		return nil, err
	}

	switch {
	case p.matchKeyword(KeywordDictionaries):
		lastToken := p.last()
		_ = p.lexer.consumeToken()
		return &SystemReloadExpr{
			ReloadPos:    pos,
			StatementEnd: lastToken.End,
			Type:         KeywordDictionaries,
		}, nil
	case p.tryConsumeKeyword(KeywordDictionary) != nil:
		dictionary, err := p.parseTableIdentifier(p.Pos())
		if err != nil {
			return nil, err
		}
		return &SystemReloadExpr{
			ReloadPos:    pos,
			StatementEnd: dictionary.End(),
			Type:         KeywordDictionary,
			Dictionary:   dictionary,
		}, nil
	case p.tryConsumeKeyword("EMBEDDED") != nil:
		lastToken := p.last()
		if err := p.consumeKeyword(KeywordDictionaries); err != nil {
			return nil, err
		}
		return &SystemReloadExpr{
			ReloadPos:    pos,
			StatementEnd: lastToken.End,
			Type:         "EMBEDDED DICTIONARIES",
		}, nil
	default:
		return nil, fmt.Errorf("expected DICTIONARIES|CONFIG")
	}
}

func (p *Parser) parseSystemSyncExpr(pos Pos) (*SystemSyncExpr, error) {
	if err := p.consumeKeyword(KeywordSync); err != nil {
		return nil, err
	}
	if err := p.consumeKeyword(KeywordReplica); err != nil {
		return nil, err
	}
	cluster, err := p.parseTableIdentifier(p.Pos())
	if err != nil {
		return nil, err
	}
	return &SystemSyncExpr{
		SyncPos: pos,
		Cluster: cluster,
	}, nil
}

func (p *Parser) parseSystemCtrlExpr(pos Pos) (*SystemCtrlExpr, error) {
	if !p.matchKeyword(KeywordStart) && !p.matchKeyword(KeywordStop) {
		return nil, fmt.Errorf("expected START|STOP")
	}
	command := strings.ToUpper(p.last().String)
	_ = p.lexer.consumeToken()

	var typ string
	switch {
	case p.tryConsumeKeyword(KeywordDistributed) != nil:
		switch {
		case p.matchKeyword(KeywordSends):
			typ = "DISTRIBUTED SENDS"
		case p.matchKeyword(KeywordFetches):
			typ = "FETCHES"
		case p.matchKeyword(KeywordMerges):
			typ = "MERGES"
		case p.matchKeyword(KeywordTtl):
			typ = "TTL MERGES"
			if err := p.consumeKeyword(KeywordMerges); err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("expected SENDS|FETCHES|MERGES|TTL")
		}
		cluster, err := p.parseTableIdentifier(p.Pos())
		if err != nil {
			return nil, err
		}
		return &SystemCtrlExpr{
			CtrlPos:      pos,
			StatementEnd: cluster.End(),
			Command:      command,
			Type:         typ,
			Cluster:      cluster,
		}, nil
	case p.tryConsumeKeyword(KeywordReplicated) != nil:
		lastToken := p.last()
		if err := p.consumeKeyword(KeywordSends); err != nil {
			return nil, err
		}
		typ = "REPLICATED SENDS"
		return &SystemCtrlExpr{
			CtrlPos:      pos,
			StatementEnd: lastToken.End,
			Command:      command,
			Type:         typ,
		}, nil
	default:
		return nil, fmt.Errorf("expected DISTRIBUTED|REPLICATED")
	}
}

func (p *Parser) parseSystemDropExpr(pos Pos) (*SystemDropExpr, error) {
	if err := p.consumeKeyword(KeywordDrop); err != nil {
		return nil, err
	}
	switch {
	case p.matchKeyword(KeywordDNS),
		p.matchKeyword(KeywordMark),
		p.matchKeyword(KeywordUncompressed),
		p.matchKeyword(KeywordFileSystem),
		p.matchKeyword(KeywordQuery):
		prefixToken := p.last()
		_ = p.lexer.consumeToken()
		lastToken := p.last()
		if err := p.consumeKeyword(KeywordCache); err != nil {
			return nil, err
		}
		return &SystemDropExpr{
			DropPos:      pos,
			StatementEnd: lastToken.End,
			Type:         prefixToken.String + " CACHE",
		}, nil
	case p.matchKeyword(KeywordCompiled):
		_ = p.lexer.consumeToken()
		if err := p.consumeKeyword(KeywordExpression); err != nil {
			return nil, err
		}
		lastToken := p.last()
		if err := p.consumeKeyword(KeywordCache); err != nil {
			return nil, err
		}
		return &SystemDropExpr{
			DropPos:      pos,
			StatementEnd: lastToken.End,
			Type:         "COMPILED EXPRESSION CACHE",
		}, nil
	default:
		return nil, fmt.Errorf("expected DNS|MARK|REPLICA|DATABASE|UNCOMPRESSION|COMPILED|QUERY")
	}
}

func (p *Parser) tryParseDeduplicateExpr(pos Pos) (*DeduplicateExpr, error) {
	if !p.matchKeyword(KeywordDeduplicate) {
		return nil, nil
	}
	return p.parseDeduplicateExpr(pos)
}

func (p *Parser) parseDeduplicateExpr(pos Pos) (*DeduplicateExpr, error) {
	if err := p.consumeKeyword(KeywordDeduplicate); err != nil {
		return nil, err
	}
	if p.tryConsumeKeyword(KeywordBy) == nil {
		return &DeduplicateExpr{
			DeduplicatePos: pos,
		}, nil
	}

	by, err := p.parseColumnExprList(p.Pos())
	if err != nil {
		return nil, err
	}
	var except *ColumnExprList
	if p.tryConsumeKeyword(KeywordExcept) != nil {
		except, err = p.parseColumnExprList(p.Pos())
		if err != nil {
			return nil, err
		}
	}
	return &DeduplicateExpr{
		DeduplicatePos: pos,
		By:             by,
		Except:         except,
	}, nil
}

func (p *Parser) parseOptimizeExpr(pos Pos) (*OptimizeExpr, error) {
	if err := p.consumeKeyword(KeywordOptimize); err != nil {
		return nil, err
	}
	if err := p.consumeKeyword(KeywordTable); err != nil {
		return nil, err
	}

	table, err := p.parseTableIdentifier(p.Pos())
	if err != nil {
		return nil, err
	}
	statmentEnd := table.End()

	onCluster, err := p.tryParseOnCluster(p.Pos())
	if err != nil {
		return nil, err
	}
	if onCluster != nil {
		statmentEnd = onCluster.End()
	}

	partitionExpr, err := p.tryParsePartitionExpr(p.Pos())
	if err != nil {
		return nil, err
	}
	if partitionExpr != nil {
		statmentEnd = partitionExpr.End()
	}

	hasFinal := false
	lastPos := p.Pos()
	if p.tryConsumeKeyword(KeywordFinal) != nil {
		hasFinal = true
		statmentEnd = lastPos
	}

	deduplicate, err := p.tryParseDeduplicateExpr(p.Pos())
	if err != nil {
		return nil, err
	}
	if deduplicate != nil {
		statmentEnd = deduplicate.End()
	}

	return &OptimizeExpr{
		OptimizePos:  pos,
		StatementEnd: statmentEnd,
		Table:        table,
		OnCluster:    onCluster,
		Partition:    partitionExpr,
		HasFinal:     hasFinal,
		Deduplicate:  deduplicate,
	}, nil
}

func (p *Parser) parseSystemExpr(pos Pos) (*SystemExpr, error) {
	if err := p.consumeKeyword(KeywordSystem); err != nil {
		return nil, err
	}

	var err error
	var expr Expr
	switch {
	case p.matchKeyword(KeywordFlush):
		expr, err = p.parseSystemFlushExpr(p.Pos())
	case p.matchKeyword(KeywordReload):
		expr, err = p.parseSystemReloadExpr(p.Pos())
	case p.matchKeyword(KeywordSync):
		expr, err = p.parseSystemSyncExpr(p.Pos())
	case p.matchKeyword(KeywordStart), p.matchKeyword(KeywordStop):
		expr, err = p.parseSystemCtrlExpr(p.Pos())
	case p.matchKeyword(KeywordDrop):
		expr, err = p.parseSystemDropExpr(p.Pos())
	default:
		return nil, fmt.Errorf("expected FLUSH|RELOAD|SYNC|START|STOP")
	}
	if err != nil {
		return nil, err
	}
	return &SystemExpr{
		SystemPos: pos,
		Expr:      expr,
	}, nil
}

func (p *Parser) parseCheckExpr(pos Pos) (*CheckExpr, error) {
	if err := p.consumeKeyword(KeywordCheck); err != nil {
		return nil, err
	}
	if err := p.consumeKeyword(KeywordTable); err != nil {
		return nil, err
	}
	table, err := p.parseTableIdentifier(p.Pos())
	if err != nil {
		return nil, err
	}
	partition, err := p.tryParsePartitionExpr(p.Pos())
	if err != nil {
		return nil, err
	}
	return &CheckExpr{
		CheckPos:  pos,
		Table:     table,
		Partition: partition,
	}, nil
}

func (p *Parser) parseRoleName(_ Pos) (*RoleName, error) {
	switch {
	case p.matchTokenKind(TokenIdent):
		name, err := p.parseIdent()
		if err != nil {
			return nil, err
		}
		var scope *StringLiteral
		if p.tryConsumeTokenKind("@") != nil {
			scope, err = p.parseString(p.Pos())
			if err != nil {
				return nil, err
			}
		}
		onCluster, err := p.tryParseOnCluster(p.Pos())
		if err != nil {
			return nil, err
		}
		return &RoleName{
			Name:      name,
			Scope:     scope,
			OnCluster: onCluster,
		}, nil
	case p.matchTokenKind(TokenString):
		name, err := p.parseString(p.Pos())
		if err != nil {
			return nil, err
		}
		onCluster, err := p.tryParseOnCluster(p.Pos())
		if err != nil {
			return nil, err
		}
		return &RoleName{
			Name:      name,
			OnCluster: onCluster,
		}, nil
	default:
		return nil, fmt.Errorf("expected <ident> or <string>")
	}
}

func (p *Parser) tryParseRoleSettings(pos Pos) ([]*RoleSetting, error) {
	if p.tryConsumeKeyword(KeywordSettings) == nil {
		return nil, nil
	}
	return p.parseRoleSettings(pos)
}

func (p *Parser) parseRoleSetting(_ Pos) (*RoleSetting, error) {
	pairs := make([]*SettingPair, 0)
	for p.matchTokenKind(TokenIdent) {
		name, err := p.parseIdent()
		if err != nil {
			return nil, err
		}
		switch name.Name {
		case "NONE", "READABLE", "WRITABLE", "CONST", "CHANGEABLE_IN_READONLY":
			return &RoleSetting{
				Modifier:     name,
				SettingPairs: pairs,
			}, nil
		}
		switch {
		case p.matchTokenKind("="),
			p.matchTokenKind(TokenInt),
			p.matchTokenKind(TokenFloat),
			p.matchTokenKind(TokenString):
			_ = p.tryConsumeTokenKind("=")
			value, err := p.parseLiteral(p.Pos())
			if err != nil {
				return nil, err
			}
			pairs = append(pairs, &SettingPair{
				Name:  name,
				Value: value,
			})
		default:
			pairs = append(pairs, &SettingPair{
				Name: name,
			})
		}

	}
	return &RoleSetting{
		SettingPairs: pairs,
	}, nil
}

func (p *Parser) parseRoleSettings(_ Pos) ([]*RoleSetting, error) {
	settings := make([]*RoleSetting, 0)
	for {
		setting, err := p.parseRoleSetting(p.Pos())
		if err != nil {
			return nil, err
		}
		settings = append(settings, setting)
		if p.tryConsumeTokenKind(",") == nil {
			break
		}
	}
	return settings, nil
}

func (p *Parser) parseCreateRole(pos Pos) (*CreateRole, error) {
	if err := p.consumeKeyword(KeywordRole); err != nil {
		return nil, err
	}

	ifNotExists := false
	orReplace := false
	switch {
	case p.matchKeyword(KeywordIf):
		_ = p.lexer.consumeToken()
		if err := p.consumeKeyword(KeywordNot); err != nil {
			return nil, err
		}
		if err := p.consumeKeyword(KeywordExists); err != nil {
			return nil, err
		}
		ifNotExists = true
	case p.matchKeyword(KeywordOr):
		_ = p.lexer.consumeToken()
		if err := p.consumeKeyword(KeywordReplace); err != nil {
			return nil, err
		}
		orReplace = true
	}

	roleNames := make([]*RoleName, 0)
	roleName, err := p.parseRoleName(p.Pos())
	if err != nil {
		return nil, err
	}
	roleNames = append(roleNames, roleName)
	for p.tryConsumeTokenKind(",") != nil {
		roleName, err := p.parseRoleName(p.Pos())
		if err != nil {
			return nil, err
		}
		roleNames = append(roleNames, roleName)
	}
	statementEnd := roleNames[len(roleNames)-1].End()

	var accessStorageType *Ident
	if p.tryConsumeKeyword(KeywordIn) != nil {
		accessStorageType, err = p.parseIdent()
		if err != nil {
			return nil, err
		}
		statementEnd = accessStorageType.NameEnd
	}

	settings, err := p.tryParseRoleSettings(p.Pos())
	if err != nil {
		return nil, err
	}
	if settings != nil {
		statementEnd = settings[len(settings)-1].End()
	}

	return &CreateRole{
		CreatePos:         pos,
		StatementEnd:      statementEnd,
		IfNotExists:       ifNotExists,
		OrReplace:         orReplace,
		RoleNames:         roleNames,
		AccessStorageType: accessStorageType,
		Settings:          settings,
	}, nil
}

func (p *Parser) parserDropUserOrRole(pos Pos) (*DropUserOrRole, error) {
	var target string
	switch {
	case p.matchKeyword(KeywordUser), p.matchKeyword(KeywordRole):
		target = p.last().String
		_ = p.lexer.consumeToken()
	default:
		return nil, fmt.Errorf("expected USER|ROLE")
	}

	ifExists, err := p.tryParseIfExists()
	if err != nil {
		return nil, err
	}

	names := make([]*RoleName, 0)
	name, err := p.parseRoleName(p.Pos())
	if err != nil {
		return nil, err
	}
	names = append(names, name)
	for p.tryConsumeTokenKind(",") != nil {
		name, err := p.parseRoleName(p.Pos())
		if err != nil {
			return nil, err
		}
		names = append(names, name)
	}
	statementEnd := names[len(names)-1].End()

	onCluster, err := p.tryParseOnCluster(p.Pos())
	if err != nil {
		return nil, err
	}
	if onCluster != nil {
		statementEnd = onCluster.End()
	}

	var from *Ident
	if p.tryConsumeKeyword(KeywordFrom) != nil {
		from, err = p.parseIdent()
		if err != nil {
			return nil, err
		}
	}

	modifier, err := p.tryParseModifier()
	if err != nil {
		return nil, err
	}

	return &DropUserOrRole{
		DropPos:      pos,
		StatementEnd: statementEnd,
		Target:       target,
		IfExists:     ifExists,
		Names:        names,
		From:         from,
		Modifier:     modifier,
	}, nil
}

func (p *Parser) parsePrivilegeSelectOrInsert(pos Pos) (*PrivilegeExpr, error) {
	keyword := p.last().String
	_ = p.lexer.consumeToken()
	params, err := p.parseFunctionParams(p.Pos())
	if err != nil {
		return nil, err
	}
	return &PrivilegeExpr{
		PrivilegePos: pos,
		Keywords:     []string{keyword},
		Params:       params,
	}, nil
}

func (p *Parser) parsePrivilegeAlter(pos Pos) (*PrivilegeExpr, error) {
	keywords := []string{KeywordAlter}
	switch {
	case p.tryConsumeKeyword(KeywordIndex) != nil:
		keywords = append(keywords, KeywordIndex)
	case p.matchKeyword(KeywordUpdate), p.matchKeyword(KeywordDelete),
		p.matchKeyword(KeywordUser), p.matchKeyword(KeywordRole), p.matchKeyword(KeywordQuota):
		keyword := p.last().String
		_ = p.lexer.consumeToken()
		keywords = append(keywords, keyword)
	case p.matchKeyword(KeywordAdd), p.matchKeyword(KeywordDrop),
		p.matchKeyword(KeywordModify), p.matchKeyword(KeywordClear),
		p.matchKeyword(KeywordComment), p.matchKeyword(KeywordRename),
		p.matchKeyword(KeywordMaterialized):
		keyword := p.last().String
		_ = p.lexer.consumeToken()
		keywords = append(keywords, keyword)
		switch {
		case p.tryConsumeKeyword(KeywordColumn) != nil:
			keywords = append(keywords, KeywordColumn)
		case p.tryConsumeKeyword(KeywordIndex) != nil:
			keywords = append(keywords, KeywordIndex)
		case p.tryConsumeKeyword(KeywordConstraint) != nil:
			keywords = append(keywords, KeywordConstraint)
		case p.tryConsumeKeyword(KeywordTtl) != nil:
			keywords = append(keywords, KeywordTtl)
		default:
			return nil, fmt.Errorf("expected COLUMN|INDEX")
		}
	case p.tryConsumeKeyword(KeywordOrder) != nil:
		if err := p.consumeKeyword(KeywordBy); err != nil {
			return nil, err
		}
		keywords = append(keywords, KeywordOrder, KeywordBy)
	case p.tryConsumeKeyword(KeywordSample) != nil:
		if err := p.consumeKeyword(KeywordBy); err != nil {
			return nil, err
		}
		keywords = append(keywords, KeywordSample, KeywordBy)
	case p.tryConsumeKeyword(KeywordSettings) != nil:
		keywords = append(keywords, KeywordSettings)
	case p.tryConsumeKeyword(KeywordView) != nil:
		keywords = append(keywords, KeywordView)
		switch {
		case p.tryConsumeKeyword(KeywordModify) != nil:
			keywords = append(keywords, KeywordModify)
		case p.tryConsumeKeyword(KeywordRefresh) != nil:
			keywords = append(keywords, KeywordRefresh)
		default:
			return nil, fmt.Errorf("expected MODIFY|REFRESH")
		}
	case p.matchKeyword(KeywordMove), p.matchKeyword(KeywordFreeze):
		keyword := p.last().String
		_ = p.lexer.consumeToken()
		keywords = append(keywords, keyword)
		if err := p.consumeKeyword(KeywordPartition); err != nil {
			return nil, err
		}
		keywords = append(keywords, KeywordPartition)
	default:
		return nil, fmt.Errorf("expected UPDATE|DELETE|ADD|DROP|MODIFY|CLEAR|COMMENT|RENAME|MATERIALIZED|ORDER|SAMPLE|SETTINGS|VIEW|MOVE|FREEZE")
	}
	return &PrivilegeExpr{
		PrivilegePos: pos,
		Keywords:     keywords,
	}, nil
}

func (p *Parser) parsePrivilegeCreate(pos Pos) (*PrivilegeExpr, error) {
	keywords := []string{KeywordCreate}
	switch {
	case p.matchKeyword(KeywordDatabase), p.matchKeyword(KeywordDictionary),
		p.matchKeyword(KeywordTable), p.matchKeyword(KeywordFunction), p.matchKeyword(KeywordView),
		p.matchKeyword(KeywordUser), p.matchKeyword(KeywordRole), p.matchKeyword(KeywordQuota):
		keyword := p.last().String
		_ = p.lexer.consumeToken()
		keywords = append(keywords, keyword)
	case p.tryConsumeKeyword(KeywordTemporary) != nil:
		if err := p.consumeKeyword(KeywordTable); err != nil {
			return nil, err
		}
		keywords = append(keywords, KeywordTemporary, KeywordTable)
	case p.tryConsumeKeyword(KeywordRows) != nil:
		if err := p.consumeKeyword(KeywordPolicy); err != nil {
			return nil, err
		}
		keywords = append(keywords, KeywordRows, KeywordPolicy)
	default:
		return nil, fmt.Errorf("expected DATABASE|DICTIONARY|TABLE|FUNCTION|VIEW|USER|ROLE|ROWS")
	}
	return &PrivilegeExpr{
		PrivilegePos: pos,
		Keywords:     keywords,
	}, nil
}

func (p *Parser) parsePrivilegeDrop(pos Pos) (*PrivilegeExpr, error) {
	keywords := []string{KeywordDrop}
	switch {
	case p.matchKeyword(KeywordDatabase), p.matchKeyword(KeywordDictionary),
		p.matchKeyword(KeywordUser), p.matchKeyword(KeywordRole), p.matchKeyword(KeywordQuota),
		p.matchKeyword(KeywordTable), p.matchKeyword(KeywordFunction), p.matchKeyword(KeywordView):
		keyword := p.last().String
		_ = p.lexer.consumeToken()
		keywords = append(keywords, keyword)
	default:
		return nil, fmt.Errorf("expected DATABASE|DICTIONARY|TABLE|FUNCTION|VIEW")
	}
	return &PrivilegeExpr{
		PrivilegePos: pos,
		Keywords:     keywords,
	}, nil
}

func (p *Parser) parsePrivilegeShow(pos Pos) (*PrivilegeExpr, error) {
	keywords := []string{KeywordShow}
	switch {
	case p.matchKeyword(KeywordDatabases), p.matchKeyword(KeywordDictionaries),
		p.matchKeyword(KeywordTables), p.matchKeyword(KeywordColumns):
		keyword := p.last().String
		_ = p.lexer.consumeToken()
		keywords = append(keywords, keyword)
	default:
		return nil, fmt.Errorf("expected DATABASES|DICTIONARIES|TABLES|COLUMNS")
	}
	return &PrivilegeExpr{
		PrivilegePos: pos,
		Keywords:     keywords,
	}, nil
}

func (p *Parser) parsePrivilegeSystem(pos Pos) (*PrivilegeExpr, error) {
	keywords := []string{KeywordShow}
	switch {
	case p.matchKeyword(KeywordShutdown), p.matchKeyword(KeywordMerges), p.matchKeyword(KeywordFetches),
		p.matchKeyword(KeywordSends), p.matchKeyword(KeywordMoves), p.matchKeyword(KeywordCluster):
		keyword := p.last().String
		_ = p.lexer.consumeToken()
		keywords = append(keywords, keyword)
	case p.tryConsumeKeyword(KeywordDrop) != nil:
		keywords = append(keywords, KeywordDrop)
		switch {
		case p.tryConsumeKeyword(KeywordCache) != nil:
			keywords = append(keywords, KeywordCache)
		case p.matchKeyword(KeywordMark), p.matchKeyword(KeywordDNS), p.matchKeyword(KeywordUncompressed):
			keyword := p.last().String
			_ = p.lexer.consumeToken()
			keywords = append(keywords, keyword)
			if err := p.consumeKeyword(KeywordCache); err != nil {
				return nil, err
			}
			keywords = append(keywords, KeywordCache)
		default:
			return nil, fmt.Errorf("expected CACHE|MARK|DNS|UNCOMPRESSED")
		}
	case p.tryConsumeKeyword(KeywordReload) != nil:
		keywords = append(keywords, KeywordReload)
		switch {
		case p.matchKeyword(KeywordDictionary), p.matchKeyword(KeywordFunction),
			p.matchKeyword(KeywordFunctions), p.matchKeyword(KeywordConfig):
			keyword := p.last().String
			_ = p.lexer.consumeToken()
			keywords = append(keywords, keyword)
		default:
			return nil, fmt.Errorf("expected DICTIONARY|FUNCTION|FUNCTIONS|CONFIG")
		}
	case p.tryConsumeKeyword(KeywordFlush) != nil:
		keywords = append(keywords, KeywordFlush)
		switch {
		case p.matchKeyword(KeywordLogs), p.matchKeyword(KeywordDistributed):
			keyword := p.last().String
			_ = p.lexer.consumeToken()
			keywords = append(keywords, keyword)
		default:
			return nil, fmt.Errorf("expected LOGS|DISTRIBUTED")
		}
	case p.tryConsumeKeyword(KeywordTtl) != nil:
		keywords = append(keywords, KeywordTtl)
		if err := p.consumeKeyword(KeywordMerges); err != nil {
			return nil, err
		}
		keywords = append(keywords, KeywordMerges)
	case p.matchKeyword(KeywordSync), p.matchKeyword(KeywordRestart):
		keyword := p.last().String
		_ = p.lexer.consumeToken()
		keywords = append(keywords, keyword)
		if err := p.consumeKeyword(KeywordReplica); err != nil {
			return nil, err
		}
		keywords = append(keywords, KeywordReplica)
	case p.tryConsumeKeyword(KeywordReplication) != nil:
		keywords = append(keywords, KeywordReplication)
		if err := p.consumeKeyword(KeywordQueues); err != nil {
			return nil, err
		}
		keywords = append(keywords, KeywordQueues)
	default:
		return nil, fmt.Errorf("expected QUEUES|SHUTDOWN|MERGES|FETCHES|SENDS|MOVES|CLUSTER|DROP|RELOAD|FLUSH|TTL|SYNC|RESTART|REPLICATION")
	}
	return &PrivilegeExpr{
		PrivilegePos: pos,
		Keywords:     keywords,
	}, nil
}

func (p *Parser) parsePrivilege(pos Pos) (*PrivilegeExpr, error) {
	switch {
	case p.matchKeyword(KeywordSelect), p.matchKeyword(KeywordInsert):
		return p.parsePrivilegeSelectOrInsert(pos)
	case p.tryConsumeKeyword(KeywordAlter) != nil:
		return p.parsePrivilegeAlter(pos)
	case p.tryConsumeKeyword(KeywordCreate) != nil:
		return p.parsePrivilegeCreate(pos)
	case p.tryConsumeKeyword(KeywordDrop) != nil:
		return p.parsePrivilegeDrop(pos)
	case p.tryConsumeKeyword(KeywordShow) != nil:
		return p.parsePrivilegeShow(pos)
	case p.tryConsumeKeyword(KeywordAll) != nil:
		return &PrivilegeExpr{
			PrivilegePos: pos,
			Keywords:     []string{KeywordAll},
		}, nil
	case p.tryConsumeKeyword(KeywordKill) != nil:
		if err := p.consumeKeyword(KeywordQuery); err != nil {
			return nil, err
		}
		return &PrivilegeExpr{
			PrivilegePos: pos,
			Keywords:     []string{KeywordKill, KeywordQuery},
		}, nil
	case p.tryConsumeKeyword(KeywordSystem) != nil:
		return p.parsePrivilegeSystem(pos)
	case p.matchKeyword(KeywordOptimize), p.matchKeyword(KeywordTruncate):
		keyword := p.last().String
		_ = p.lexer.consumeToken()
		return &PrivilegeExpr{
			PrivilegePos: pos,
			Keywords:     []string{keyword},
		}, nil
	case p.tryConsumeKeyword(KeywordRole) != nil:
		if err := p.consumeKeyword(KeywordAdmin); err != nil {
			return nil, err
		}
		return &PrivilegeExpr{
			PrivilegePos: pos,
			Keywords:     []string{KeywordRole, KeywordAdmin},
		}, nil
	}
	return nil, fmt.Errorf("expected SELECT|INSERT|ALTER|CREATE|DROP|SHOW|KILL|SYSTEM|OPTIMIZE|TRUNCATE")
}

func (p *Parser) parsePrivilegeRoles(_ Pos) ([]*Ident, error) {
	roles := make([]*Ident, 0)
	role, err := p.parseIdent()
	if err != nil {
		return nil, err
	}
	roles = append(roles, role)
	for p.tryConsumeTokenKind(",") != nil {
		role, err := p.parseIdent()
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

func (p *Parser) parseGrantOptions(_ Pos) ([]string, error) {
	options := make([]string, 0)
	for p.matchKeyword(KeywordWith) {
		option, err := p.parseGrantOption(p.Pos())
		if err != nil {
			return nil, err
		}
		options = append(options, option)
	}
	return options, nil
}

func (p *Parser) parseGrantOption(_ Pos) (string, error) {
	if err := p.consumeKeyword(KeywordWith); err != nil {
		return "", err
	}
	ident, err := p.parseIdent()
	if err != nil {
		return "", err
	}
	if err := p.consumeKeyword(KeywordOption); err != nil {
		return "", err
	}
	return ident.Name, nil
}

func (p *Parser) parseGrantSource(_ Pos) (*TableIdentifier, error) {
	ident, err := p.parseIdentOrStar()
	if err != nil {
		return nil, err
	}

	if p.tryConsumeTokenKind(".") == nil {
		return &TableIdentifier{
			Table: ident,
		}, nil
	}
	dotIdent, err := p.parseIdentOrStar()
	if err != nil {
		return nil, err
	}
	return &TableIdentifier{
		Database: ident,
		Table:    dotIdent,
	}, nil
}

func (p *Parser) parseGrantPrivilege(pos Pos) (*GrantPrivilegeExpr, error) {
	if err := p.consumeKeyword(KeywordGrant); err != nil {
		return nil, err
	}
	onCluster, err := p.tryParseOnCluster(p.Pos())
	if err != nil {
		return nil, err
	}
	var privileges []*PrivilegeExpr
	privilege, err := p.parsePrivilege(p.Pos())
	if err != nil {
		return nil, err
	}
	privileges = append(privileges, privilege)
	for p.tryConsumeTokenKind(",") != nil {
		privilege, err := p.parsePrivilege(p.Pos())
		if err != nil {
			return nil, err
		}
		privileges = append(privileges, privilege)
	}
	statementEnd := privileges[len(privileges)-1].End()

	if err := p.consumeKeyword(KeywordOn); err != nil {
		return nil, err
	}
	on, err := p.parseGrantSource(p.Pos())
	if err != nil {
		return nil, err
	}

	if err := p.consumeKeyword(KeywordTo); err != nil {
		return nil, err
	}
	toRoles, err := p.parsePrivilegeRoles(p.Pos())
	if err != nil {
		return nil, err
	}
	if len(toRoles) != 0 {
		statementEnd = toRoles[len(toRoles)-1].NameEnd
	}
	options, err := p.parseGrantOptions(p.Pos())
	if err != nil {
		return nil, err
	}
	if len(options) != 0 {
		statementEnd = p.last().End
	}

	return &GrantPrivilegeExpr{
		GrantPos:     pos,
		StatementEnd: statementEnd,
		OnCluster:    onCluster,
		Privileges:   privileges,
		On:           on,
		To:           toRoles,
		WithOptions:  options,
	}, nil
}

func (p *Parser) parseAlterRole(pos Pos) (*AlterRole, error) {
	if err := p.consumeKeyword(KeywordRole); err != nil {
		return nil, err
	}

	ifExists, err := p.tryParseIfExists()
	if err != nil {
		return nil, err
	}

	roleRenamePairs := make([]*RoleRenamePair, 0)
	roleRenamePair, err := p.parseRoleRenamePair(p.Pos())
	if err != nil {
		return nil, err
	}
	roleRenamePairs = append(roleRenamePairs, roleRenamePair)
	for p.tryConsumeTokenKind(",") != nil {
		roleRenamePair, err := p.parseRoleRenamePair(p.Pos())
		if err != nil {
			return nil, err
		}
		roleRenamePairs = append(roleRenamePairs, roleRenamePair)
	}
	statementEnd := roleRenamePairs[len(roleRenamePairs)-1].End()

	settings, err := p.tryParseRoleSettings(p.Pos())
	if err != nil {
		return nil, err
	}
	if settings != nil {
		statementEnd = settings[len(settings)-1].End()
	}

	return &AlterRole{
		AlterPos:        pos,
		StatementEnd:    statementEnd,
		IfExists:        ifExists,
		RoleRenamePairs: roleRenamePairs,
		Settings:        settings,
	}, nil
}

func (p *Parser) parseRoleRenamePair(_ Pos) (*RoleRenamePair, error) {
	roleName, err := p.parseRoleName(p.Pos())
	if err != nil {
		return nil, err
	}
	roleRenamePair := &RoleRenamePair{
		RoleName:     roleName,
		StatementEnd: roleName.End(),
	}
	if p.tryConsumeKeyword(KeywordRename) != nil {
		if err := p.consumeKeyword(KeywordTo); err != nil {
			return nil, err
		}
		newName, err := p.parseIdent()
		if err != nil {
			return nil, err
		}
		roleRenamePair.NewName = newName
		roleRenamePair.StatementEnd = newName.NameEnd
	}
	return roleRenamePair, nil
}

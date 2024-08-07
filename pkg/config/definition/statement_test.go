package definition

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var statement_definition = []byte(`
statement_definition:
  partner:
    statement: CREATE TABLE partner (
      id bigint(20) UNSIGNED NOT NULL COMMENT '代理id',
      applicant_id bigint(20) NOT NULL DEFAULT '0' COMMENT '代理商上级',
      directly_id bigint(20) NOT NULL DEFAULT '0' COMMENT '直属负责人',
      partner_log_id bigint(20) NOT NULL DEFAULT '0' COMMENT '代理商日志id',
      partner_change_directly_id bigint(20) NOT NULL DEFAULT '0' COMMENT '代理商变更直属负责人申请记录id',
      main VARCHAR(50) NOT NULL DEFAULT '' COMMENT '经营主体',
      position VARCHAR(50) NOT NULL DEFAULT '' COMMENT '职位',
      region VARCHAR(204) NOT NULL DEFAULT '' COMMENT '区域',
      source VARCHAR(50) NOT NULL DEFAULT '0' COMMENT '来源',
      evaluation text COMMENT '代理商评价',
      price decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '代理商价格',
      partner_contract_id bigint(20) NOT NULL DEFAULT '0' COMMENT '代理商合约id',
      expiration_at bigint(20) NOT NULL DEFAULT 0 COMMENT '到期时间',
      id_card varchar(18) NOT NULL DEFAULT '' COMMENT '身份证',
      mobile varchar(14) NOT NULL DEFAULT '' COMMENT '联系电话',
      uniform_code varchar(18) NOT NULL DEFAULT '' COMMENT '社会统一信用代码',
      real_name varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '真实姓名',
      level char(4) NOT NULL DEFAULT '0' COMMENT '代理类型',
      status char(4) NOT NULL DEFAULT '0' COMMENT '代理商状态',
      created_at datetime NOT NULL COMMENT '创建时间 ISO 8601格式',
      updated_at datetime NOT NULL COMMENT '更新时间 ISO 8601格式',
      deleted_at bigint(20) NOT NULL DEFAULT '0' COMMENT '',
      UNIQUE KEY uq_partner_uc_cart (uniform_code, id_card),
      Key idx_partner_applicant_id (applicant_id),
      Index idx_partner_id_card (id_card),
      PRIMARY KEY (id)
      ) ENGINE = InnoDB ROW_FORMAT = DYNAMIC COMMENT = '代理商基础信息'
`)

func TestStatementParseFile(t *testing.T) {
	e := new(Statementdefinition)
	tf, err := os.CreateTemp("", "test.yml")
	assert.NoError(t, err)
	defer os.Remove(tf.Name())
	os.WriteFile(tf.Name(), statement_definition, 0666)
	em := e.ParseFile(tf.Name())
	assert.NotNil(t, em.Definition)
	assert.NotNil(t, em.Definition["partner"])
	// t.Log(em.Definition["partner"])
}

func TestStatementParse(t *testing.T) {
	e := new(Statementdefinition)
	em := e.Parse(statement_definition)
	assert.NotNil(t, em.Definition)
	assert.NotNil(t, em.Definition["partner"])
	// t.Log(em.Definition["partner"])
}

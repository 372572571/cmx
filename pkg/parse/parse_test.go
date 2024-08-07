package parse

import (
	"fmt"
	"testing"
)

var test_sql_1 = `
CREATE TABLE oauth2_access_tokens (
	id bigint unsigned NOT NULL AUTO_INCREMENT DEFAULT '0',
	client_id varchar(100) NOT NULL DEFAULT '0',
	user_id bigint unsigned NOT NULL DEFAULT 0,
	access_token varchar(512) NOT NULL DEFAULT 'x'  COMMENT '访问token',
	expires_in bigint unsigned NOT NULL,
	expires_x decimal(12,6) NOT NULL DEFAULT 1.000000,
	created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (id) USING BTREE,
	UNIQUE KEY oauth2_access_tokens_access_token_unique (access_token),
	UNIQUE KEY oauth2_access_tokens_client_id_user_id_unique (client_id,user_id),
	KEY user_id (user_id),
	CONSTRAINT oauth2_access_tokens_ibfk_1 FOREIGN KEY (user_id) REFERENCES user (id),
	CONSTRAINT oauth2_access_tokens_ibfk_2 FOREIGN KEY (client_id) REFERENCES oauth2_clients (client_id)
  ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='访问令牌表，用于存储访问令牌信息'
`

// CONSTRAINT oauth2_access_tokens_ibfk_1 FOREIGN KEY (user_id) REFERENCES user (id),
// CONSTRAINT oauth2_access_tokens_ibfk_2 FOREIGN KEY (client_id) REFERENCES oauth2_clients (client_id)

func TestParse(t *testing.T) {
	m := ParseSqlToModel(test_sql_1, Options{})
	fmt.Println(m)
}

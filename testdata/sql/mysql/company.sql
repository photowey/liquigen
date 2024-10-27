/* comments */
/* comments */

-- comments
-- comments

create table if not exists company.employee
(
	id          bigint                                   not null comment 'AAAA' primary key,
	create_by   bigint                                   not null comment 'BBBB',
	update_by   bigint                                   not null comment 'CCCC',
	create_time timestamp      default CURRENT_TIMESTAMP not null comment 'DDDD',
	update_time timestamp      default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment 'EEEE',
	deleted     tinyint(2)                               not null comment 'FFFF',
	employee_no varchar(32)                              not null comment 'GGGG',
	balance     decimal(16, 2) default 0                 not null comment 'HHHH',
	org_id      bigint                                   not null comment 'IIII',
	org_name    varchar(64)                              null comment 'JJJJ',
	sorted      int            default 0                 not null comment 'KKKK',
	states      tinyint        default 0                 not null comment 'LLLL',
	remark      text                                     null comment 'MMMM'
) COMMENT = 'EMPLOYEE' Engine = Innodb;

CREATE TABLE IF NOT EXISTS company.organization
(
	id          BIGINT                               NOT NULL COMMENT 'AAAA' PRIMARY KEY,
	create_by   BIGINT                               NOT NULL COMMENT 'BBBB',
	update_by   BIGINT                               NOT NULL COMMENT 'CCCC',
	create_time TIMESTAMP  DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT 'DDDD',
	update_time TIMESTAMP  DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'EEEE',
	deleted     TINYINT(2)                           NOT NULL COMMENT 'FFFF',
	org_no      VARCHAR(32)                          NOT NULL COMMENT 'GGGG',
	sorted      INT        DEFAULT 0                 NOT NULL COMMENT 'HHHH',
	states      TINYINT(2) DEFAULT 0                 NOT NULL COMMENT 'IIII',
	remark      text                                 NULL COMMENT 'JJJJ'
) COMMENT = 'ORGANIZATION' ENGINE = Innodb;

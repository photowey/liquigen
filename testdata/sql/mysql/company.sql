/*
 * Copyright © 2024 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/* comments */

/*
 1.111
 2.222
 3.333
 */

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

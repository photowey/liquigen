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

package postgresd

const (
	DsnTemplate = "postgresql://%s:%s@%s:%d/%s?TimeZone=%s"

	TableInfoSQL = "SELECT tableName FROM pg_tables WHERE tableName NOT LIKE 'pg%' AND tableName NOT LIKE 'sql_%' ORDER BY tableName"

	ColumnInfoSQL = `SELECT CONCAT('', cast(obj_description(relfilenode, 'pg_class') as varchar))     AS tableComment,
       a.attname                                                                 AS namez,
       pg_type.typname                                                           AS typname,
       col_description(a.attrelid, a.attnum)                                     AS commentz,
       pg_type.typlen                                                            AS typlen,
       CONCAT('', SUBSTRING(format_type(a.atttypid, a.atttypmod) from '\(.*\)')) AS formatlength,
       a.attnotnull                                                              AS notnullz,
       (CASE
            WHEN (SELECT COUNT(*)
                  FROM pg_constraint
                  WHERE conrelid = a.attrelid
                    AND conkey[1] = attnum
                    AND contype = 'p') > 0 THEN
                TRUE
            ELSE FALSE
           END)                                                                  AS primarykey
FROM pg_class AS c,
     pg_attribute AS a
         INNER JOIN pg_type ON pg_type.oid = a.atttypid
WHERE c.relname = '%s'
  AND a.attrelid = c.oid
  AND a.attnum > 0`
)

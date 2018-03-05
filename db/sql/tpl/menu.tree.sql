SELECT
	m1.id AS id,
	(
		CASE
		WHEN (m2.id = 0 OR m2.id IS NULL) THEN
			0
		ELSE
			m2.id
		END
	) AS pId,
	m1.name AS name,
	(
		CASE
		WHEN (m2.id = 0 OR m2.id IS NULL) THEN
			TRUE
		ELSE
			FALSE
		END
	) AS open
FROM
	sys_menu m1
LEFT JOIN sys_menu m2 ON m1.pcode = m2.code
ORDER BY
	m1.id ASC
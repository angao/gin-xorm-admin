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
			true
		ELSE
			false
		END
	) AS open,
	(
		CASE
		WHEN (m3.ID = 0 OR m3.ID IS NULL) THEN
			false
		ELSE
			true
		END
	) AS checked
FROM
	sys_menu m1
LEFT JOIN sys_menu m2 ON m1.pcode = m2.code
LEFT JOIN (
	SELECT
		ID
	FROM
		sys_menu
	WHERE
		ID IN (
        {{ range $index, $value := .menuIDs }}
            {{ if ne $index $.length }}
                {{$value}},
            {{ else }}
                {{$value}}
            {{ end }}
        {{ end }}
        )
) m3 ON m1.id = m3.id
ORDER BY
	m1.id ASC
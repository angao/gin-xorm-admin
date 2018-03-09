SELECT
	id,
	title,
	type,
	content,
	createtime,
	creater,
	(
		SELECT
			NAME
		FROM
			sys_user
		WHERE
			id = creater
	) AS create_name
FROM
	sys_notice
WHERE 1 = 1
{{ if ne .Name ""}}
    AND (title like '%{{ .Name }}%' OR content like '%{{ .Name }}%')
{{ end }}
ORDER BY createtime DESC
{{ if ne .Limit 0}}
	LIMIT {{ .Limit }} OFFSET {{ .Offset }}
{{ end }}
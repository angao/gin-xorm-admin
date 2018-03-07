SELECT
	s.id,
	s.num,
	s.pid,
	(
		SELECT
			fullname
		FROM
			sys_dept
		WHERE
			id = s.pid
	) AS pname,
	s.pids,
	s.simplename,
	s.fullname,
	s.tips,
	s.version
FROM
	sys_dept s
WHERE 1 = 1
{{if ne .name ""}}
    AND s.simplename like '%{{ .name }}%' OR s.fullname like '%{{ .name }}%'
{{ end}}
ORDER BY s.num ASC
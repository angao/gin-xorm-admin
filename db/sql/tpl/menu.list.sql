SELECT
	m.id,
	m.code,
	m.pcode,
	m.pcodes,
	m.name,
	m.icon,
	m.url,
	m.num,
	m.levels,
	( CASE WHEN m.ismenu = 1 THEN '是' ELSE '否' END ) AS is_menu_name,
	m.tips,
	( CASE WHEN m.STATUS = 1 THEN '启用' ELSE '不启用' END ) AS status_name,
	m.isopen 
FROM
	`sys_menu` m
WHERE 1 = 1 AND m.status = 1
    {{ if ne .Name ""}}
        AND (m.name like '%{{ .Name }}%')
    {{ end }}
    {{ if eq .Order "desc" }}
        ORDER BY m.id DESC
    {{ else }}
        ORDER BY m.id ASC
    {{ end }}
SELECT
	m.id,
	m.CODE,
	m.pcode,
	m.pcodes,
	m.NAME,
	m.icon,
	m.url,
	m.num,
	m.levels,
	( CASE WHEN m.ismenu = 1 THEN '是' ELSE '否' END ) AS isMenuName,
	m.tips,
	( CASE WHEN m.STATUS = 1 THEN '启用' ELSE '不启用' END ) AS statusName,
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
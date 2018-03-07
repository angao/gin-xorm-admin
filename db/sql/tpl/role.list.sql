SELECT
    r.id,
    r.`name`,
    r.pid AS pId,
    (
        SELECT
            NAME
        FROM
            sys_role
        WHERE
            id = r.pid
    ) AS p_name,
    r.deptid,
    r.tips,
    d.fullname AS dept_name
FROM
    sys_role r
LEFT JOIN sys_dept d ON r.deptid = d.id
WHERE 1 = 1
{{ if ne .Id 0}}
    AND r.id = {{ .Id }}
{{ end }}
{{ if ne .Name ""}}
    AND (r.name like '%{{ .Name }}%')
{{ end }}
{{ if eq .Order "desc" }}
    ORDER BY r.id DESC
{{ else }}
    ORDER BY r.id ASC
{{ end }}
{{ if ne .Limit 0}}
LIMIT {{ .Limit }} OFFSET {{ .Offset }}
{{ end }}
SELECT
    u.id,
    u.avatar,
    u.account,
    u.`name`,
    u.birthday,
    (
        CASE
        WHEN u.sex = 1 THEN
            '男'
        WHEN u.sex = 2 THEN
            '女'
        ELSE
            '未知'
        END
    ) sex,
    u.email,
    u.phone,
    u.roleid,
    u.deptid,
    (
        CASE
        WHEN u.`status` = 1 THEN
            '启用'
        WHEN u.`status` = 2 THEN
            '冻结'
        WHEN u.`status` = 3 THEN
            '删除'
        END
    ) status,
    date_format(u.createtime, '%Y-%m-%d %T') AS createtime,
    r.name AS role_name,
    d.fullname AS dept_name
FROM
    sys_user u
LEFT JOIN sys_role r ON u.roleid = r.id
LEFT JOIN sys_dept d ON u.deptid = d.id
WHERE u.status != 3
{{ if ne .DeptID 0 }}
    AND u.deptid = {{ .DeptID }}
{{ end }}
{{ if eq .Order "desc" }}
    ORDER BY id DESC
{{ else }}
    ORDER BY id ASC
{{ end }}
LIMIT {{ .Limit }} OFFSET {{ .Offset }}
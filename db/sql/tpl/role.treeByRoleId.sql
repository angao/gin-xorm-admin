SELECT
    r.id AS "id",
    pId AS "pId",
    NAME AS "name",
    (
        CASE
        WHEN (pId = 0 OR pId IS NULL) THEN
            TRUE
        ELSE
            FALSE
        END
    ) "open",
    (
        CASE
        WHEN (r1.ID = 0 OR r1.ID IS NULL) THEN
            FALSE
        ELSE
            TRUE
        END
    ) "checked"
FROM
    sys_role r
LEFT JOIN (
    SELECT
        ID
    FROM
        sys_role
    WHERE
        ID in (
    {{ range $index, $value := .roleIds }}
        {{ if ne $index $.length }}
            {{$value}},
        {{ else }}
            {{$value}}
        {{ end }}
    {{ end }}
        )
) r1 ON r.ID = r1.ID
ORDER BY
    pId,
    num ASC
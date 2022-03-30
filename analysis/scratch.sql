-- Number of provide operations
SELECT count(*)
FROM provides;

-- Number of dials for provide
WITH cte AS (
    SELECT count(d.id) dial_count
    FROM dials d
             INNER JOIN provides_x_dials pxd on d.id = pxd.dial_id
             INNER JOIN provides p on pxd.provide_id = p.id
    WHERE p.provider_id = 1
    GROUP BY p.id
)
SELECT percentile_cont(0.50) WITHIN GROUP ( ORDER BY cte.dial_count ASC) AS p50,
       percentile_cont(0.90) WITHIN GROUP ( ORDER BY cte.dial_count ASC) AS p90,
       percentile_cont(0.95) WITHIN GROUP ( ORDER BY cte.dial_count ASC) AS p95
FROM cte;

-- Number of connections for provide
WITH cte AS (
    SELECT count(c.id) conn_count
    FROM connections c
             INNER JOIN provides_x_connections pxc on c.id = pxc.connection_id
             INNER JOIN provides p on pxc.provide_id = p.id
    WHERE p.provider_id = 1
    GROUP BY p.id
)
SELECT percentile_cont(0.50) WITHIN GROUP ( ORDER BY cte.conn_count ASC) AS p50,
       percentile_cont(0.90) WITHIN GROUP ( ORDER BY cte.conn_count ASC) AS p90,
       percentile_cont(0.95) WITHIN GROUP ( ORDER BY cte.conn_count ASC) AS p95
FROM cte;

FROM percona/percona-server:5.7

COPY ./initdb.d/*.sql /docker-entrypoint-initdb.d/
COPY ./mysql.conf.d/*.cnf /etc/my.cnf.d

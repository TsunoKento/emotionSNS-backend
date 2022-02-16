FROM --platform=linux/amd64 mysql:8.0.28

EXPOSE 3306

ADD ./my.cnf /etc/mysql/cnf.d/my.cnf

CMD ["mysqld"]
docker run \
-v ~/mariadb_data:/var/lib/mysql  \
-e MARIADB_ROOT_PASSWORD=armchair \
-v ~/mariadb_data:/var/lib/mysql  \
-p 3306:3306 \
-d mariadb:latest

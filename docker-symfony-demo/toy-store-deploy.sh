#!/bin/bash
# 假设已经安装git、php、composer
#
if [ ! -f "toy-store" ]; then
    git clone https://github.com/mousezheng/toy-store nginx/www/toy-store;
fi

cd nginx/www/toy-store;

git pull origin master;
composer install;

php bin/console doctrine:database:create -n;
php bin/console doctrine:migrations:migrate -n;
./bin/phpunit;


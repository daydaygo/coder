# laravel

```sh
composer global require laravel/installer
laravel new blog

composer create-project --prefer-dist laravel/laravel blog

php artisan serve
php artisan config:cache # prod
php artisan route:cache
php artisan view:cache

```

- auth: breeze(minimal) jetstram(functional tailwindCss livewire/inertia.js)
- deploy: forge vapor(serverless)

- mix: compile css+js
- fullstack: livewire inertia.js

## scout: full-text search to model

```sh
composer require laravel/scout
php artisan vendor:publish --provider="Laravel\Scout\ScoutServiceProvider"

php artisan make:migration create_posts_table
php artisan migrate
php artisan make:model Post
php artisan make:controller PostController
php artisan make:command ImportPosts
```

- https://laravel.com/docs
- https://learnku.com/laravel

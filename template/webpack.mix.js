const mix = require('laravel-mix')
const env = process.env.NODE_ENV || 'development';

mix.disableNotifications()
    .setPublicPath(__dirname)

mix.js('app/app.js', 'public/js')
    .sass('app/assets/scss/app.scss', 'public/css')

if (env === 'production') {
    mix.version()
}
# url-shortener
Url shortener. Written with golang + fasthttp + redis.

Requires running redis-server (port 6379 by default, but you can change it in configuration file) + golang >= 1.18 + dependencies.

Accepts POST requests on /encode/; argument "url" should contain url you want to encode;
Accepts GET requests on /<shorturl>; service will redirect you to decoded url (if such url exists).

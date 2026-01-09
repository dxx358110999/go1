go get github.com/gin-contrib/cors

git config --global http.proxy http://127.0.0.1:7897
git config --global https.proxy https://127.0.0.1:7897
git config --global core.autocrlf input

git rm --cached -r .idea
git rm --cached -r web_app.log
git rm --cached -r conf
git commit -m "rm .idea"

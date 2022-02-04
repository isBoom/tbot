tar -zxcf ../a.tar ../tbot
scp -r  ../a.tar sky:
ssk sky "rm -rf tbot && tar -zxvf a.tar && rm -rf a.tar"
del ../a.tar
git add .
git commit -m"fix"
git push https://github.com/isboom/tbot master
scp -r --exclude='*.out'  ../tbot sky:
git add .
git commit -m"fix"
git push https://github.com/isboom/tbot master
echo "-----------------------------"
echo "Stash内容の説明を入力"
read STASH_MESSAGE

eval git stash save "$STASH_MESSAGE"
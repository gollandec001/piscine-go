export aloka=$(head -n 179 streets/Buckingham_Place | tail -n 1 | cut -d "#" -f2)
echo $aloka
cat interviews/interview-"$aloka"
echo $MAIN_SUSPECT
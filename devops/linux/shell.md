## Shell
### 常用 Job
日期遍历

```bash
start_date=20200801
end_date=`date +"%Y%m%d" -d "-1 days"`
start_sec=`date -d "$start_date" "+%s"`
end_sec=`date -d "$end_date" "+%s"`
for((i=start_sec;i<=end_sec;i+=86400)); do
    day=$(date -d "@$i" "+%Y-%m-%d")
    echo $day
    ./media-marketing book-data sum -d=${day}
done

date=2021-03-01
end_date=2021-03-20
while [ "$date" != "$end_date" ]; do
  echo $date
	# your command -d=${date}
  date=$(date -I -d "$date + 1 day")

  # mac option for d decl (the +1d is equivalent to + 1 day)
  # d=$(date -j -v +1d -f "%Y-%m-%d" "2020-12-12" +%Y-%m-%d)
done
```

批量日期执行脚本

```bash
cd $(dirname "$0")/..

SCRIPT=$1
BEGIN=$2
END=$3

cronlog="`pwd`/logs/batch-run-`date +"%Y-%m-%d"| tr -d '-'`.log"

while [ "$BEGIN" -le "$END" ];
do
    DATA=$(date -d "$BEGIN" +%Y-%m-%d)

    echo `date +"%Y-%m-%d %H:%M:%S"`"--${SCRIPT} -d=${DATA} start" | tee -a ${cronlog}
    ${SCRIPT} -d=${DATA} | tee -a ${cronlog}
    echo `date +"%Y-%m-%d %H:%M:%S"`"--${SCRIPT} -d=${DATA} end" | tee -a ${cronlog}
    BEGIN=`echo $(date -d "$BEGIN+1days" +%Y-%m-%d) | tr -d '-'`
done
```

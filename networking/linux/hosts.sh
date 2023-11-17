#! /bin/bash
# Identify all hosts in your subnet

ipaddr=$(ip -4 a show eth0 | grep -w inet | awk -F" " {' print $2 '})
ip3=$(echo "$ipaddr" | awk -F"." {' print $1"."$2"."$3 '})

echo -n "Start isp: $ip3."
read startip
echo -n "End ip: $ip3."
read endip
#startip=1
#endip=20

for (( i=startip; i<=endip; i++ )) do
  ping -c 1 "$ip3.$i"
done

arp -e | grep -v incomplete > "arp_$(date +%F-%R).txt"

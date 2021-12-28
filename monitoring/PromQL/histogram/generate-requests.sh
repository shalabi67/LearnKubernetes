while [ 1 -eq 1 ]
do
  x=$(( $RANDOM % 50 + 1 ))
  i=0
  while [ $i -lt $x ]
  do
    curl -s http://promo:8180 >/dev/null
    i=$(( $i + 1 ))
  done
  curl -s http://promo:8180?slow >/dev/null

  x=$(( $RANDOM % 50 + 1 ))
  i=0
  while [ $i -lt $x ]
  do
    curl -s http://promo:8181 >/dev/null
    i=$(( $i + 1 ))
  done
  curl -s http://promo:8181?slow >/dev/null

done

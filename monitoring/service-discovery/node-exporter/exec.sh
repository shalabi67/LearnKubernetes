# run node_exporter in background
node_exporter &

# now you can run other commands
while [ 1 -eq 1 ]
do
 ls -R /
done

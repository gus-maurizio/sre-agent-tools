# sre-agent-tools
Tools to simulate load, memory, and other goodies
```
tail -2 /tmp/sre-agent.cpu.measure.log|jq '. | {ts: .timestamp, plugin: .plugin, at: .at, cpuavg: .measure.cpu, eachcpu: .measure.cpupercent, mhz: .measure.throughput}'

tail -2 /tmp/sre-agent.cpu.measure.log|jq '. | {ts: .timestamp, plugin: .plugin, at: .at, cpuavg: .measure.cpu, cpumax: .measure.cpumax, mhz: .measure.throughput}' -c -a|jq --slurp '.'| jq -r '(.[0] | keys_unsorted) as $keys | ([$keys] + map([.[ $keys[] ]])) [] | @csv'
```

```
tail -1 /tmp/sre-agent.cpu.measure.log|jq '.at| split(" ")[0,1,2,4]'

tail -2 /tmp/sre-agent.cpu.measure.log|jq '. | {ts: .timestamp, plugin: .plugin, at: .at, cpuavg: .measure.cpu, cpumax: .measure.cpumax, mhz: .measure.throughput}' -c -a | \
jq --slurp '.'| jq -r '(.[0] | keys_unsorted) as $keys | ([$keys] + map([.[ $keys[] ]])) [] | @csv'

tail -5 /tmp/sre-agent.cpu.measure.log | \
cat /tmp/sre-agent.cpu.measure.log | \
jq '. | {
ts:     .timestamp,
plugin: .plugin, 
cpuavg: .measure.cpu,
cpumax: .measure.cpumax, 
mhz:    .measure.throughput, 
at_date: (.at| split(" ")[0]),
at_time: (.at| split(" ")[1]),
at_zone: (.at| split(" ")[2]),
at_tz:   (.at| split(" ")[3]),
at_relt: (.at| split(" ")[4])
}' -c -a | \
jq --slurp '.'| jq -r '(.[0] | keys_unsorted) as $keys | ([$keys] + map([.[ $keys[] ]])) [] | @csv'
```

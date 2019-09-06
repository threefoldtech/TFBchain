# Explorer

A block explorer for Threefold Bonus

## Run it yourself

### Prerequisites
* Caddyserver
* Threefold Bonus daemon (`tfbchaind`)


Make sure you have `tfbchaind` (the Threefold Bonus daemon) running with the explorer module enabled:
`tfbchaind -M cgte`

Now start caddy from the `caddy` folder of this repository:
`caddy -conf Caddyfile.local`
and browse to http://localhost:2015

#GPractice

###Installation
1. from `go install -v ...`
2. place configuration file `config.yaml` to `~/.gpractice/` directory:
   - `cd ~`
   - `mkdir .gpractice`
   - `cp config.yaml ./.gpractice` 
3. modify config.yaml if needed
4. - `gpractice` to run CLI
   - `gpractice-web` to run web app, default URL: `http://localhost:3000/app`
5. to see available arguments, run needed command with flag `-help`: `gpractice -help`


###TODO
- [x] import/export csv files
- [x] html/css styling
- [ ] mongodb support
- [x] unit tests
- [x] external configs
- [x] duration web format: 01h23m45s
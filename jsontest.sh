curl -s https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json' |\ 
> jq '.id'
"70"
jq --raw-output '.name'
